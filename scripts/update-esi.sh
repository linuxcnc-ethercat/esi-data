#!/bin/sh

set -e

dir=`pwd`
xmldir=$dir/tmpesi

if [ $(basename $dir) = "scripts" ]; then
	echo "Please run 'scripts/update-esi.sh' from the root of this repository:"
       	echo 
	echo "  cd .."
	echo "  scripts/update-esi.sh"
	exit 0
fi

fetch() {
    file=$1
    url=$2
    useragent="User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0"

    if [ ! -e $file ]; then
	curl -H "$useragent" -L $url -o $file
    fi
}

esiyml=$dir/esi.yml
fetch $dir/beckhoff.zip "https://www.beckhoff.com/en-en/download/128205835"
fetch $dir/omron1.zip "https://assets.omron.eu/downloads/ddf/en/v1/r88d-knxxx-ect(-l)_ethercat_esi_file_en.zip"
fetch $dir/omron2.zip "https://assets.omron.eu/downloads/ddf/en/v3/r88d-1snxxx-ect_ethercat_esi_file_en.zip"
fetch $xmldir/EpoCAT_FR1000.xml "https://www.bausano.net/images/epocat/EpoCAT_FR1000.xml"
fetch $dir/smc-ex260.zip "https://content2.smcetech.com/files/si/SMC_EX260_Serial_Interface_Configuration_Files.zip"
fetch $dir/delta1.zip "https://downloadcenter.deltaww.com/downloadCenterCounter.aspx?DID=19032&DocPath=1&hl=en-US"
fetch $dir/leadshine1.zip " https://www.leadshine.com/upfiles/downloads/bc4d3fde51ab7b2d92d478956ca4aa9e_1650879535067.zip"

# Unzip, but don't overwrite files.
unzip -nj $dir/beckhoff.zip -d $xmldir
unzip -nj $dir/omron1.zip -d $xmldir
unzip -nj $dir/omron2.zip -d $xmldir
unzip -nj $dir/smc-ex260.zip -d $xmldir '*EtherCAT*.xml'
unzip -nj $dir/delta1.zip -d $xmldir
unzip -nj $dir/leadshine1.zip -d $xmldir


cd esidecoder
go build esidecoder.go
cd $dir

esidecoder/esidecoder --esi_directory=$dir/tmpesi --output=$esiyml --pdos --objects
