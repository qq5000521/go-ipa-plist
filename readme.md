## Go脚本(推荐用这个版本20231121日更新)

### 功能：go版本会自动读取CFBundleIdentifier和CFBundleDisplayName的值，不用手动输入，自动替换ipa文件名，自动修改模板并且创建plist文件。

### 注意：setting.plist模板文件不要删除，可以改掉里面的url地址。替换字符huandiao这个不要改，会自动替换新生成的文件名。

### 用法：将签好名的ipa包放在与setting.plist同级目录下，(windows系统)双击执行[go-ipa-list.exe](go-ipa-list.exe),其他系统自己go build一下。完成后上传两个文件到对应的url目录即可。
