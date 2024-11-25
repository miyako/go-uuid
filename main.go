package main

import (
	"flag"
	"fmt"

	"github.com/gofrs/uuid/v5"
)

func main() {

	mode := flag.Int("mode", 7, "mode")
	name := flag.String("name", "", "name")
	namespace := flag.String("ns", "dns", "ns")

	flag.Parse()

	var ns uuid.UUID
	switch *namespace {
	case "dns":
		ns = uuid.NamespaceDNS
	case "url":
		ns = uuid.NamespaceURL
	case "oid":
		ns = uuid.NamespaceOID
	default:
		ns = uuid.NamespaceX500
	}

	switch *mode {
	case 6:
		uid, err := uuid.NewV6()
		if err == nil {
			fmt.Println(uid)
		}
	case 5:
		uid := uuid.NewV5(ns, *name)
		// log.Printf("generated Version 5 UUID %v ns %v name %v", uid, *namespace, *name)
		fmt.Println(uid)
	case 4:
		uid, err := uuid.NewV4()
		if err == nil {
			// log.Printf("generated Version 4 UUID %v", uid)
			fmt.Println(uid)
		}
	case 3:
		uid := uuid.NewV3(ns, *name)
		// log.Printf("generated Version 3 UUID %v ns %v name %v", uid, *namespace, *name)
		fmt.Println(uid)
	case 1:
		uid, err := uuid.NewV1()
		if err == nil {
			// log.Printf("generated Version 1 UUID %v", uid)
			fmt.Println(uid)
		}
	default:
		uid, err := uuid.NewV7()
		if err == nil {
			// log.Printf("generated Version 7 UUID %v", uid)
			fmt.Println(uid)
		}
	}

}
