package main

import (
	"fmt"
	"time"
)

func main() {
	msTimestamp := time.Now().UnixNano() / int64(time.Millisecond)
	timeStr := time.Unix(0, msTimestamp*int64(time.Millisecond)).Format("20060102150405")
	fileName := fmt.Sprintf("%s%s", timeStr, GenerateRandomString(8))

	ipaPath, err := GetIPAFilePath()
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

	err = NewRenameFile(fileName, ipaPath).Rename()
	if err != nil {
		fmt.Println("重命名文件失败:", err)
	}

	fmt.Println("运行完毕,回车关闭窗口!")
	fmt.Scanln()
}
