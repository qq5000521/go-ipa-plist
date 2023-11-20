package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type PlistReplacer struct {
	infoMap map[string]interface{}
}

func NewPlistReplacer(infoMap map[string]interface{}) *PlistReplacer {
	return &PlistReplacer{infoMap: infoMap}
}
func (r *PlistReplacer) Replace(plistFilePath, appName string) error {
	plistFile, err := os.Open(plistFilePath)
	if err != nil {
		return err
	}
	defer plistFile.Close()
	content, err := io.ReadAll(plistFile)
	if err != nil {
		return err
	}

	newContent := strings.Replace(string(content), "huandiao", appName, -1)
	newContent = strings.Replace(newContent, "com.udcs.jplay", r.infoMap["CFBundleIdentifier"].(string), -1)
	fmt.Println("获取的BundleId:", r.infoMap["CFBundleIdentifier"].(string))

	var replaceString string
	fmt.Print("请输入APP名字(回车自动读取ipa包内的名称): ")
	fmt.Scanln(&replaceString)

	for len(replaceString) == 0 {
		// 如果用户没有输入要替换的字符串，则尝试从 Info.plist 文件中获取 CFBundleDisplayName 字段的值
		if r.infoMap["CFBundleDisplayName"] != nil {
			replaceString = r.infoMap["CFBundleDisplayName"].(string)
			fmt.Println("获取的应用名字:", replaceString)
		} else {
			fmt.Print("未读取到app名字,请手动输入: ")
			fmt.Scanln(&replaceString)
		}
	}

	fmt.Println("将应用名字替换为:", replaceString)
	newContent = strings.Replace(newContent, "王者传奇", replaceString, -1)

	err = os.WriteFile(appName+".plist", []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
