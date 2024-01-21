package esi

import (
	"encoding/xml"
)

type ESIName struct {
	Name       string `xml:",chardata"`
	LanguageID string `xml:"LcId,attr"`
}

type ESIDeviceURL struct {
	URL        string `xml:",chardata"`
	LanguageID string `xml:"LcId,attr"`
}

type ESIDeviceInfo struct {
	Current         string              `xml:"Electrical>EBusCurrent" yaml:",omitempty"`
	RequestTimeout  int                 `xml:"Mailbox>Timeout>RequestTimeout" yaml:",omitempty"`
	ResponseTimeout int                 `xml:"Mailbox>Timeout>ResponseTimeout" yaml:",omitempty"`
	Ports           []ESIDeviceInfoPort `xml:"Port" yaml:",omitempty"`
}

type ESIDeviceInfoPort struct {
	Type  string
	Label string
}

type ESIProfile struct {
	ProfileNo   int            `xml:"ProfileNo" yaml:",omitempty"`
	Channels    int            `xml:"ChannelCount" yaml:",omitempty"`
	AddInfo     int            `xml:"AddInfo" yaml:",omitempty"` // what is "AddInfo?"
	DataTypes   []*ESIDataType `xml:"Dictionary>DataTypes>DataType"`
	DataTypeMap map[string]*ESIDataType
	Objects     []*ESIObject `xml:"Dictionary>Objects>Object" yaml:"Objects,omitempty"`
}

type ESIDataType struct {
	Name     string                `xml:"Name"`
	BitSize  int                   `xml:"BitSize"`
	SubItems []*ESIDataTypeSubItem `xml:"SubItem"`
}

type ESIDataTypeSubItem struct {
	SubIndex  int    `xml:"SubIdx"`
	Name      string `xml:"Name"`
	Type      string `xml:"Type"`
	BitSize   int    `xml:"BitSize"`
	BitOffset int    `xml:"BitOffs"`
	Access    string `xml:"Flags>Access"`
	Category  string `xml:"Flags>Category"`
}

type ESIObject struct {
	Index    string              `xml:"Index" yaml:",omitempty"`
	Name     string              `xml:"Name" yaml:",omitempty"`
	TypeName string              `xml:"Type" yaml:",omitempty"`
	BitSize  int                 `xml:"BitSize" yaml:",omitempty"`
	SubItems []*ESIObjectSubItem `xml:"Info>SubItem" yaml:",omitempty"`
	// Info > (DefaultData | SubItem | ...
	// Flags > (Access | Category | ...
}

type ESIObjectSubItem struct {
	Name        string `xml:"Name"\ yaml:",omitempty"`
	DefaultData string `xml:"Info>DefaultData" yaml:",omitempty"`
	Type        string `yaml:"Type,omitempty"`
	BitSize     int    `yaml:"BitSize,omitempty"`
	BitOffset   int    `yaml:"BitOffs,omitempty"`
	Access      string `yaml:"Access,omitempty"`
	Category    string `yaml:"Category,omitempty"`
}

type ESIDeviceID struct {
	Type        string `yaml:"Type,omitempty"`
	ProductCode string `yaml:"ProductCode,omitempty"`
	RevisionNo  string `yaml:"RevisionNo,omitempty"`
	URL         string `yaml:"URL,omitempty"`
	Name        string `yaml:"Name,omitempty"`
	Vendor      string `yaml:"Vendor,omitempty"`
	VendorID    string `yaml:"VendorID,omitempty"`
}

type ESIDevice struct {
	//Physics string `xml:"Physics,attr" yaml:",omitempty"`

	Type struct {
		Type            string `xml:",chardata"`
		ProductCode     string `xml:"ProductCode,attr"`
		RevisionNo      string `xml:"RevisionNo,attr"`
		CheckRevisionNo string `xml:"CheckRevisionNo,attr"`
	} `xml:"Type" yaml:"-"`

	Names     []ESIName      `xml:"Name" yaml:",omitempty"`
	URLs      []ESIDeviceURL `xml:"URL" yaml:",omitempty"`
	Info      ESIDeviceInfo  `xml:"Info" yaml:"Info,omitempty"`
	GroupType string         `xml:"GroupType" yaml:",omitempty"` // maps to ESIGroup.Type
	//Fmmu      []string       `xml:"Fmmu" yaml:",omitempty,flow"` // "Inputs"/"Outputs"

	ShortType   string         `yaml:"Type,omitempty"`
	ProductCode string         `yaml:"ProductCode,omitempty"`
	RevisionNo  string         `yaml:"RevisionNo,omitempty"`
	EnglishURL  string         `yaml:"URL,omitempty"`
	EnglishName string         `yaml:"Name,omitempty"`
	GroupName   string         `yaml:"DeviceGroup,omitempty"`
	Vendor      string         `yaml:"Vendor,omitempty"`
	VendorID    string         `yaml:"VendorID,omitempty"`
	Profile     ESIProfile     `yaml:"Profile,omitempty"`
	TxPDOs      []*ESIPDO      `xml:"TxPdo" yaml:"TxPDOs,omitempty"`
	RxPDOs      []*ESIPDO      `xml:"RxPdo" yaml:"RxPDOs,omitempty"`
	IDs         []*ESIDeviceID `yaml:"IDs"`
}

type ESIPDO struct {
	Index     string         `xml:"Index" yaml:"Index,omitempty"`
	PDOName   string         `yaml:"Name,omitempty"`
	LangNames []*ESILangName `xml:"Name"`
	Ref       string         `xml:"Ref,attr" yaml:"Ref,omitempty"`
	Chn       string         `xml:"Chn,attr" yaml:"Chn,omitempty"`
	Fixed     string         `xml:"Fixed,attr" yaml:"Fixed,omitempty"`
	Sm        string         `xml:"Sm,attr" yaml:"Sm,omitempty"`
	Entries   []*ESIPDOEntry `xml:"Entry" yaml:"Entry,omitempty"`
}

type ESIPDOEntry struct {
	PDOName   string           `yaml:"Name,omitempty"`
	LangNames []*ESILangName   `xml:"Name" yaml:"Names,omitempty"`
	Index     string           `xml:"Index" yaml:"Index,omitempty"`
	SubIndex  string           `xml:"SubIndex" yaml:"SubIndex,omitempty"`
	BitLen    int              `xml:"BitLen" yaml:"BitLen,omitempty"`
	Dtype     ESIPDOEntryData `xml:"DataType" yaml:",omitempty"`
	DataType  string           `xml:"-" yaml:"DataType,omitempty"`
	DataScale string           `xml:"-" yaml:"DataScale,omitempty"`
	Comment   string           `xml:"Comment" yaml:"Comment,omitempty"`
}

type ESIPDOEntryData struct {
	DataType  string `xml:",chardata"`
	DataScale string `xml:"DScale,attr"`
}

type ESIGroup struct {
	Type        string    `xml:"Type"`
	Names       []ESIName `xml:"Name"`
	EnglishName string
}

type ESILangName struct {
	Name       string `xml:",chardata"`
	LanguageID string `xml:"LcId,attr"`
}

type ESIData struct {
	EtherCATInfo xml.Name       `xml:"EtherCATInfo"`
	VendorID     string         `xml:"Vendor>Id"`
	VendorNames  []*ESILangName `xml:"Vendor>Name"`
	VendorName   string
	Groups       []*ESIGroup  `xml:"Descriptions>Groups>Group"`
	Devices      []*ESIDevice `xml:"Descriptions>Devices>Device"`
}
