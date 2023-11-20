package main

import (
	"archive/zip"
	"errors"
	"howett.net/plist"
	"io"
	"os"
	"strings"
)

type IPAParser struct {
	ipaFilePath string
}

func NewIPAParser(ipaFilePath string) *IPAParser {
	return &IPAParser{ipaFilePath: ipaFilePath}
}

func (p *IPAParser) GetInfoPlist() (map[string]interface{}, error) {
	ipaFile, err := os.Open(p.ipaFilePath)
	if err != nil {
		return nil, err
	}
	defer ipaFile.Close()
	ipaFileStat, err := ipaFile.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := ipaFileStat.Size()

	zipReader, err := zip.NewReader(ipaFile, fileSize)
	if err != nil {
		return nil, err
	}

	for _, file := range zipReader.File {
		filePath := file.Name
		if strings.HasPrefix(filePath, "Payload/") && strings.HasSuffix(filePath, ".app/Info.plist") {
			infoFile, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer infoFile.Close()

			infoPlistBytes, err := io.ReadAll(infoFile)
			if err != nil {
				return nil, err
			}

			var infoMap map[string]interface{}
			_, err = plist.Unmarshal(infoPlistBytes, &infoMap)
			if err != nil {
				return nil, err
			}

			return infoMap, nil
		}
	}

	return nil, errors.New("在IPA文件中没有找到:Info.plist")
}
