package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

// GetIPAFilePath 获取ipa文件路径
func GetIPAFilePath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	ipaFile, err := FindIpaFile(dir)
	if err != nil {
		return "", err
	}

	return ipaFile, nil
}

// FindIpaFile 查找ipa文件
func FindIpaFile(dir string) (string, error) {
	var ipaFile string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".ipa" {
			ipaFile = path
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if ipaFile == "" {
		return "", fmt.Errorf("未找到.ipa 文件")
	}
	return ipaFile, nil
}

// GenerateRandomString 随机字符串
func GenerateRandomString(length int) string {
	strS := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLen := len(strS)
	pool := make([]byte, charsetLen)
	for i := 0; i < charsetLen; i++ {
		pool[i] = strS[i]
	}

	str := make([]byte, length)
	for i := 0; i < length; i++ {
		index := rand.Intn(charsetLen)
		str[i] = pool[index]

		pool[index] = pool[charsetLen-1]
		charsetLen--
	}

	return string(str)
}
