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
	echo "Fetching into $(basename $file)..."
	curl -H "$useragent" -L $url -o $file
    fi
}

esiyml=$dir/esi.yml
fetch $dir/beckhoff.zip "https://www.beckhoff.com/en-en/download/128205835"
fetch $dir/omron1.zip "https://assets.omron.eu/downloads/ddf/en/v1/r88d-knxxx-ect(-l)_ethercat_esi_file_en.zip"
fetch $dir/omron2.zip "https://assets.omron.eu/downloads/ddf/en/v3/r88d-1snxxx-ect_ethercat_esi_file_en.zip"
fetch $dir/smc-ex260.zip "https://content2.smcetech.com/files/si/SMC_EX260_Serial_Interface_Configuration_Files.zip"
fetch $dir/delta1.zip "https://downloadcenter.deltaww.com/downloadCenterCounter.aspx?DID=19032&DocPath=1&hl=en-US"
fetch $dir/delta2.zip "https://downloadcenter.deltaww.com/downloadCenterCounter.aspx?DID=21858&DocPath=1&hl=en-US"
fetch $dir/delta3.zip "https://downloadcenter.deltaww.com/downloadCenterCounter.aspx?DID=4796&DocPath=1&hl=en-US"
fetch $dir/leadshine1.zip "https://www.leadshine.com/upfiles/downloads/bc4d3fde51ab7b2d92d478956ca4aa9e_1650879535067.zip"
fetch $xmldir/EpoCAT_FR.xml "https://www.bausano.net/images/epocat/EpoCAT_FR1000.xml"
fetch $xmldir/easyio.xml "https://www.bausano.net/images/arduino-easyio/EasyIO_ESI_V1_0.xml"
fetch $xmldir/epocat-io1616.xml "https://www.bausano.net/images/epocat-io1616/EpoCAT_IO1616_Esi_V1_1.xml"
fetch $xmldir/yaskawa_sgd7s.xml "https://www.yaskawa.com/delegate/getAttachment?documentId=Yaskawa_Sigma-7_CoE_ESI_Files&cmd=documents&documentName=Yaskawa_SGD7S-xxxxA0x.xml"
fetch $xmldir/yaskawa_sgd7s-400v.xml "https://www.yaskawa.com/delegate/getAttachment?documentId=SW.Sigma-7.01&cmd=documents&documentName=Yaskawa_SGD7S-xxxDA0xxxxF64.xml"
fetch $xmldir/yaskawa_sgd7w-400v.xml "https://www.yaskawa.com/delegate/getAttachment?documentId=SW.Sigma-7.02&cmd=documents&documentName=Yaskawa_SGD7W-xxxDA0x.xml"
fetch $xmldir/yaskawa_sies3.xml "https://www.yaskawa.com/delegate/getAttachment?documentId=ESI_SIES3_OPT_V_1_03_01&cmd=documents&documentName=ESI_SIES3_OPT_V_1_03_01.xml"
fetch $dir/yaskawa-sigma5.zip "https://mobile.yaskawa.com/delegate/getAttachment?documentId=Yaskawa_CoE_ESI_Files&cmd=documents&documentName=Yaskawa_CoE_ESI_Files.zip"
fetch $dir/hitachi-eh-ioca.zip "https://automation.hitachi-industrial.eu/_Resources/Persistent/4/3/9/4/439459d98d2a33cf4dc4637537182d45999d30e4/EH-IOCA.zip"
fetch $dir/hitachi-p1-ect.zip 'https://automation.hitachi-industrial.eu/_Resources/Static/Packages/Moon.HitachiEurope/Downloads/automation/%5B2%5D%20Software/%5B5%5D%20Configuration%20Files/%5B3%5D%20ESI%20Files/P1-ECT/P1-ECT_ESI.zip'
fetch $dir/hitachi-rio2-eca.zip 'https://automation.hitachi-industrial.eu/_Resources/Static/Packages/Moon.HitachiEurope/Downloads/automation/%5B2%5D%20Software/%5B5%5D%20Configuration%20Files/%5B3%5D%20ESI%20Files/RIO2-ECA/20180115.zip'
fetch $dir/hitachi-wj-ect.zip 'https://automation.hitachi-industrial.eu/_Resources/Static/Packages/Moon.HitachiEurope/Downloads/automation/%5B2%5D%20Software/%5B5%5D%20Configuration%20Files/%5B3%5D%20ESI%20Files/WJ-ECT/WJ-ECT_ESI%20V1.03.xml.zip'
fetch $dir/boschrexroth1.zip 'https://www.boschrexroth.com/media/e043994b-0459-420e-9a3e-8a5e6ffb375c'
fetch $dir/rtelligent.zip 'http://www.rtelligent.net/upload/wenjian/XML.zip'

# Unzip, but don't overwrite files.
unzip -nj $dir/beckhoff.zip -d $xmldir
unzip -nj $dir/omron1.zip -d $xmldir
unzip -nj $dir/omron2.zip -d $xmldir
unzip -nj $dir/smc-ex260.zip -d $xmldir '*EtherCAT*.xml'
unzip -nj $dir/delta1.zip -d $xmldir
unzip -nj $dir/delta2.zip -d $xmldir
unzip -nj $dir/delta3.zip -d $xmldir
unzip -nj $dir/leadshine1.zip -d $xmldir
unzip -nj $dir/yaskawa-sigma5.zip -d $xmldir
unzip -nj $dir/hitachi-eh-ioca.zip -d $xmldir
unzip -nj $dir/hitachi-p1-ect.zip -d $xmldir
unzip -nj $dir/hitachi-rio2-eca.zip -d $xmldir
unzip -nj $dir/hitachi-wj-ect.zip -d $xmldir
unzip -nj $dir/boschrexroth1.zip -d $xmldir
unzip -nj $dir/rtelligent.zip -d $xmldir

cd esidecoder
go build esidecoder.go
cd $dir

esidecoder/esidecoder --esi_directory=$dir/tmpesi --output=$esiyml --pdos --objects
