package main

import (
	"flag"
	"github.com/akley-MK4/obs-sh/internal"
	"log"
	"os"
)

func main() {
	accKey := flag.String("acc_key", "", "acc_key=")
	filePath := flag.String("file_path", "", "file_path=")
	outDir := flag.String("out_dir", "", "out_dir=")
	flag.Parse()

	if err := internal.EncryptFileToOutDir(*accKey, *filePath, *outDir); err != nil {
		log.Println("Failed to encrypt file to the output directory, ", err.Error())
		os.Exit(1)
	}

	log.Println("Successfully encrypted the file to the output directory")
	os.Exit(0)
}
