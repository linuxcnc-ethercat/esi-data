#!/bin/sh

set -e

dir=`pwd`

if [ $(basename $dir) = "scripts" ]; then
	echo "Please run 'scripts/update-esi.sh' from the root of this repository:"
       	echo 
	echo "  cd .."
	echo "  scripts/update-esi.sh"
	exit 0
fi

esiyml=$dir/esi.yml
beckhoffurl="https://www.beckhoff.com/en-en/download/128205835"
omronurl1="https://assets.omron.eu/downloads/ddf/en/v1/r88d-knxxx-ect(-l)_ethercat_esi_file_en.zip"
omronurl2="https://assets.omron.eu/downloads/ddf/en/v3/r88d-1snxxx-ect_ethercat_esi_file_en.zip"
useragent="User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0"

curl -H "$useragent" -L "$beckhoffurl" -o $dir/beckhoff.zip
curl -L "$omronurl1" -o $dir/omron1.zip
curl -L "$omronurl2" -o $dir/omron2.zip

# Unzip, but don't overwrite files.
unzip -nj $dir/beckhoff.zip -d $dir/tmpesi
unzip -nj $dir/omron1.zip -d $dir/tmpesi
unzip -nj $dir/omron2.zip -d $dir/tmpesi

cd esidecoder
go build esidecoder.go
cd $dir

esidecoder/esidecoder --esi_directory=$dir/tmpesi --output=$esiyml --pdos --objects
