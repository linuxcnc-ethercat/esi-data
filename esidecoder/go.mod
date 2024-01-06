module github.com/linuxcnc-ethercat/esi-data/esidecoder

go 1.21.5

replace github.com/linuxcnc-ethercat/esi-data/esi => /home/scott/git/esi-data/esi

require (
	github.com/linuxcnc-ethercat/esi-data/esi v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.19.0
	gopkg.in/yaml.v3 v3.0.1
)

require golang.org/x/text v0.14.0 // indirect
