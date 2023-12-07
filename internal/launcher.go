package internal

import (
	_ "embed"
	"log"
)

//go:embed static/es_app
var encAppData []byte

func ShellLaunch(accKey, outDir string) error {
	if accKey == "" {
		accKey = defaultAccKey
	}
	if outDir == "" {
		outDir = defaultExecDir
	}

	if err := decryptFileToExecDir(accKey, encAppData, outDir); err != nil {
		return err
	}

	log.Println("Successfully parsed the app")
	cmdStart, _, startErr := startProcess(nil)
	if startErr != nil {
		return startErr
	}

	log.Println("Successfully ran the app")
	if err := cmdStart.Wait(); err != nil {
		log.Println("Failed to wait app, ", err.Error())
	}

	return nil
}
