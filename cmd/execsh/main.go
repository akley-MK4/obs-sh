package main

import (
	_ "embed"
	"flag"
	"github.com/akley-MK4/obs-sh/internal"
	"log"
	"os"
)

func main() {
	accKey := flag.String("acc_key", "", "acc_key=")
	outDir := flag.String("out_dir", "", "out_dir=")
	flag.Parse()

	if err := internal.ShellLaunch(*accKey, *outDir); err != nil {
		log.Println("Failed to exec app, ", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
