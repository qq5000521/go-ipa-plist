package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
	msTimestamp := time.Now().UnixNano() / int64(time.Millisecond)
	timeStr := time.Unix(0, msTimestamp*int64(time.Millisecond)).Format("20060102150405")
	fileName := fmt.Sprintf("%s%s", timeStr, generateRandomString(8))
	ipaPath, err := getIPAFilePath()
	if err != nil {
		fmt.Println("获取 IPA 文件路径失败:", err)
	}

	infoMap, err := NewIPAParser(ipaPath).GetInfoPlist()
	if err != nil {
		fmt.Println("解析 Info.plist 失败:", err)
	}

	err = NewPlistReplacer(infoMap).Replace("setting.plist", fileName)
	if err != nil {
		fmt.Println("替换 Info.plist 失败:", err)
	}

	err = NewFileMover(fileName, ipaPath).Move()
	if err != nil {
		fmt.Println("移动文件失败:", err)
	}

	fmt.Println("运行完毕,回车关闭窗口!")
	fmt.Scanln()
}
func getIPAFilePath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	ipaFile, err := findIpaFile(dir)
	if err != nil {
		return "", err
	}

	return ipaFile, nil
}

func findIpaFile(dir string) (string, error) {
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

func generateRandomString(length int) string {
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
