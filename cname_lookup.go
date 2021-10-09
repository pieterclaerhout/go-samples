package main

import (
	"errors"
	"log"
	"net"
	"strings"
)

func lookupCNAME(host string) (string, error) {

	if host == "" {
		return "", errors.New("Empty host name")
	}

	cname, err := net.LookupCNAME(host)
	if err != nil {
		return "", errors.New("Domain name doesn't exist")
	}

	cname = strings.TrimSuffix(cname, ".")
	host = strings.TrimSuffix(host, ".")

	if cname == "" || cname == host {
		return "", errors.New("Domain name is not a CNAME")
	}

	return cname, nil

}

func main() {

	result, err := lookupCNAME("api.yellowduck.be")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(result)

}
