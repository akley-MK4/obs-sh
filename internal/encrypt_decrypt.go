package internal

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

func EncryptFileToOutDir(accKey string, filePath string, outDir string) error {
	data, readErr := ioutil.ReadFile(filePath)
	if readErr != nil {
		return readErr
	}

	if accKey == "" {
		accKey = defaultAccKey
	}
	encryptedData, encErr := aesEncryptCBC(data, []byte(accKey))
	if encErr != nil {
		return encErr
	}

	_, createDirErr := createDirectory(outDir)
	if createDirErr != nil {
		return createDirErr
	}

	outFilePath := path.Join(outDir, fmt.Sprintf("enc_%v", defaultExecName))

	f, openErr := os.OpenFile(outFilePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if openErr != nil {
		return openErr
	}
	defer func() {
		_ = f.Close()
	}()

	_, writeErr := f.Write(encryptedData)
	if writeErr != nil {
		return writeErr
	}

	return nil
}

func decryptFileToExecDir(accKey string, encData []byte, outDir string) error {
	decData, decErr := aesDecryptCBC(encData, []byte(accKey))
	if decErr != nil {
		return decErr
	}

	_, statErr := os.Stat(outDir)
	if os.IsNotExist(statErr) {
		_, createDirErr := createDirectory(outDir)
		if createDirErr != nil {
			return createDirErr
		}
	}

	outExecFilePath = path.Join(outDir, defaultExecName)
	osType := fmt.Sprintf("%v", runtime.GOOS)
	if osType == "windows" {
		outExecFilePath += ".exe"
	}

	_, statFileErr := os.Stat(outExecFilePath)
	if !os.IsNotExist(statFileErr) {
		if err := os.Remove(outExecFilePath); err != nil {
			return err
		}
	}

	f, openErr := os.OpenFile(outExecFilePath, os.O_CREATE|os.O_RDWR|os.O_EXCL, os.ModePerm)
	if openErr != nil {
		return openErr
	}
	defer func() {
		_ = f.Close()
	}()

	_, writeErr := f.Write(decData)
	if writeErr != nil {
		return writeErr
	}

	return nil
}

func aesEncryptCBC(origData []byte, key []byte) (retEncrypted []byte, retErr error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		retErr = err
		return
	}
	blockSize := block.BlockSize()
	origData = pkcs5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	retEncrypted = make([]byte, len(origData))
	blockMode.CryptBlocks(retEncrypted, origData)
	return
}

func aesDecryptCBC(encrypted []byte, key []byte) (retEncrypted []byte, retErr error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		retErr = err
		return
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	retEncrypted = make([]byte, len(encrypted))
	blockMode.CryptBlocks(retEncrypted, encrypted)
	retEncrypted = pkcs5UnPadding(retEncrypted)
	return
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func createDirectory(p string) (bool, error) {
	_, statErr := os.Stat(p)
	if os.IsNotExist(statErr) {
		if mkErr := os.MkdirAll(p, os.ModePerm); mkErr != nil {
			return false, mkErr
		}
		return true, nil
	}
	if statErr != nil {
		return false, statErr
	}

	return false, nil
}
