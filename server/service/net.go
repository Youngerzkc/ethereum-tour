package service

import (
	"log"
)

func (es *EtherService) GetVersion() (string, error) {
	var version string
	if err := es.call(&version, "net_version"); err!= nil {
		log.Fatal(err)
		return version, err
	}
	return version, nil
}