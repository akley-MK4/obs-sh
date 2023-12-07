package implement

import (
	_ "embed"
	"log"
)

func ShellLaunch(encAppData []byte, accKey, outDir string, enableClean bool, kwArgs map[string]interface{}) error {
	if accKey == "" {
		accKey = defaultAccKey
	}
	if outDir == "" {
		outDir = defaultExecDir
	}

	if err := decryptFileToExecDir(accKey, encAppData, outDir, enableClean); err != nil {
		return err
	}

	log.Println("Successfully parsed the app")
	cmdStart, _, startErr := startProcess(kwArgs)
	if startErr != nil {
		return startErr
	}

	log.Println("Successfully ran the app")
	if err := cmdStart.Wait(); err != nil {
		log.Println("Failed to wait app, ", err.Error())
	}

	return nil
}
