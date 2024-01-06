package esi

import (
	"cmp"
	"crypto/sha256"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"slices"
)

// Merge devices together if they have identical RxPDO/TxPDO/Object
// configurations.  This should be a fairly conservative way to catch
// mostly-identical devices that can share the same driver.
//
// This also *dramatically* reduces the side of the output, from 73M
// to 24M.
//
// For example, this decides that the following are mostly identical:
//
// - EJ2008
// - EJ2128
// - EL2008
// - EL2008-0015
// - EL2088
// - EL2788
// - EL2798
// - EL2808
// - EL2828
// - ELX2008
// - EP2008-0001
// - EP2008-0002
// - EP2008-0022
// - EP2028-0001
// - EP2028-0002
// - EP2028-0032
// - EPP2008-0001
// - EPP2008-0002
// - EPP2008-0022
// - EPP2028-0001
// - EPP2028-0002
// - EQ2008-0002
//
// Unfortunately, this is still a bit odd about different revisions of
// devices.  For example, according to the current ESI files from
// Beckhoff, the EL7031 comes in multiple revisions:
//
// - 0x00100000 "rev0"
// - 0x0010001e "rev0e" -- actually differnet HW?
// - 0x00110000 "rev1" -- identical to rev0
// - 0x00120000 "rev2" -- adds a bunch of 0x6020 and 0x7020 PDOs which aren't in rev0
// - 0x00130000 "rev3" -- similar to rev2, but names reformatted
// - 0x00140000 "rev4" -- identical to rev3
// - 0x00150000 "rev5" -- identical to rev3
// - 0x00160000 "rev6" -- identical to rev3
// - 0x00170000 "rev7" -- similar to rev2, but deletes a bunch of PDOs from 'objects'
// - 0x00180000 "rev8" -- similar to rev7, but adds 0x7021 PDOs
// - 0x00190000 "rev9" -- identical to rev8
// - 0x001a0000 "rev10" -- similar to rev8, but adds 0x6020:0x23
//
// These are currently broken into 7 different blocks by this code,
// because (to some extent) these actually *are* several different
// devices.  Attempting to access 0x6020 PDOs with a rev0 or rev1
// device will fail, probably keeping LCEC from starting up at all.
//
// Looking at our code, the el70x1 driver supports the EL7031 and the
// EL7041-0052.  Amazingly, this code actually groups the "rev10"
// EL7031 together with all of the revs of the EL7041-0052.
func MergeDevices(devices []*ESIDevice) ([]*ESIDevice, error) {
	pdomap := make(map[string]*ESIDevice)

	for _, d := range devices {
		profile, err := yaml.Marshal(d.Profile)
		if err != nil {
			return nil, fmt.Errorf("can't marshal profile: %v", err)
		}

		txpdos, err := yaml.Marshal(d.TxPDOs)
		if err != nil {
			return nil, fmt.Errorf("can't marshal txpdos: %v", err)
		}

		rxpdos, err := yaml.Marshal(d.RxPDOs)
		if err != nil {
			return nil, fmt.Errorf("can't marshal rxpdos: %v", err)
		}

		profile = append(profile, txpdos...)
		profile = append(profile, rxpdos...)

		sig := fmt.Sprintf("%x", sha256.Sum256(profile))
		dd := pdomap[sig]

		id := &ESIDeviceID{
			Type:        d.ShortType,
			ProductCode: d.ProductCode,
			RevisionNo:  d.RevisionNo,
			URL:         d.EnglishURL,
			Name:        d.EnglishName,
			Vendor:      d.Vendor,
			VendorID:    d.VendorID,
		}

		if dd == nil {
			// This didn't exist in the map before, so add
			// it and zero out the fields that we don't
			// want anymore.
			dd = d
			pdomap[sig] = d
			d.Names = nil
			d.URLs = nil
			d.ShortType = ""
			d.ProductCode = ""
			d.RevisionNo = ""
			d.EnglishURL = ""
			d.EnglishName = ""
			d.Vendor = ""
			d.VendorID = ""
		}

		dd.IDs = append(dd.IDs, id)
	}

	fmt.Fprintf(os.Stderr, "Started with %d devices, reduced to %d\n", len(devices), len(pdomap))

	devs := []*ESIDevice{}
	for _, v := range pdomap {
		if v != nil {
			// Sort IDs by type + revision no
			slices.SortStableFunc(v.IDs, func(i, j *ESIDeviceID) int {
				if n := cmp.Compare(i.Type, j.Type); n != 0 {
					return n
				}
				return cmp.Compare(i.RevisionNo, j.RevisionNo)
			})

			devs = append(devs, v)
		}
	}

	// sort devices by the first ID's type + revision no.  It's
	// important that this order stays relatively stable to keep
	// diff sizes reasonable.
	slices.SortStableFunc(devs, func(i, j *ESIDevice) int {
		if n := cmp.Compare(i.IDs[0].Type, j.IDs[0].Type); n != 0 {
			return n
		}
		return cmp.Compare(i.IDs[0].RevisionNo, j.IDs[0].RevisionNo)
	})

	return devs, nil
}
