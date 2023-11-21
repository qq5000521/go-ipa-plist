package main

import "os"

type RenameFile struct {
	fileName   string
	oldIpaName string
}

func NewRenameFile(fileName string, oldIpaName string) *RenameFile {
	return &RenameFile{
		oldIpaName: oldIpaName,
		fileName:   fileName,
	}
}

func (m *RenameFile) Rename() error {
	err := os.Rename(m.oldIpaName, m.fileName+".ipa")
	if err != nil {
		return err
	}

	return nil
}
