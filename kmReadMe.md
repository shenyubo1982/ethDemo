# show eth Address from local keyStore folder.
## mac环境编译:
'''
chmod 700 ./build/keyManager.sh
./build/keyManager.sh
'''

## 使用:
在命令行执行编译后的可执行文件
'''
./bin/ethKeyManage
'''

### screenshots


必须文件
.
├── README.md
├── bin
│   └── ethKeyManage
├── build
│   └── keyManager.sh
├── cmds
│   └── km
│       ├── keyManager.go
│       └── keyManager_test.go
├── go.mod
├── go.sum
├── main.go
├── menu
│   └── menu.go
├── util
│   ├── abiTks.go
│   ├── coverage.out
│   ├── fileTks.go
│   ├── hint_test.out
│   ├── reflectTks.go
│   ├── reflectTks_test.go
│   ├── test.sh
│   ├── webRequest.go
│   ├── webRequest_test.go
│   ├── yamlTks.go
│   └── yamlTks_test.go
└── wiki

