package main

import (
	_ "embed"
	"flag"
	"github.com/akley-MK4/obs-sh/implement"
	"log"
	"os"
)

//go:embed static/enc_obs_app
var encAppData []byte

func main() {
	accKey := flag.String("acc_key", "", "acc_key=")
	outDir := flag.String("out_dir", "", "out_dir=")
	flag.Parse()

	if err := implement.ShellLaunch(encAppData, *accKey, *outDir); err != nil {
		log.Println("Failed to execute the app, ", err.Error())
		os.Exit(1)
	}

	log.Println("Successfully executed the app")
	os.Exit(0)
}
