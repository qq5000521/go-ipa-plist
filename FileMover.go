package main

import "os"

type FileMover struct {
	fileName   string
	oldIpaName string
}

func NewFileMover(fileName string, oldIpaName string) *FileMover {
	return &FileMover{
		oldIpaName: oldIpaName,
		fileName:   fileName,
	}
}
func (m *FileMover) Move() error {
	err := os.Rename(m.oldIpaName, m.fileName+".ipa")
	if err != nil {
		return err
	}
	err = os.Rename(m.fileName+".plist", m.fileName+".plist")
	if err != nil {
		return err
	}

	return nil
}
