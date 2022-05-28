# ethDemo


# GoEthereumBook 01：初始化客户端

教程地址

[https://goethereumbook.org/zh/client-setup/](https://goethereumbook.org/zh/client-setup/)

开发环境：goland

go version

```bash
go version go1.17.6 darwin/amd64
```

代码

```go
package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	fmt.Println(client)
	if err != nil {
		log.Fatalf("Oops! There was a problem", err)
	} else {
		fmt.Println("Success! you are connected to the Ethereum Network")
	}
}
```

成功运行效果

```go
go run client.go
&{0xc0001a6000}
Success! you are connected to the Ethereum Network
```

说明

ethclient.Dial 参数是你的以太坊节点url。若您没有现有以太坊客户端，您可以连接到infura网关。Infura管理着一批安全，可靠，可扩展的以太坊[geth和parity]节点，并且在接入以太坊网络时降低了新人的入门门槛。

我们使用了JetBrains的goLand，所以使用go.mod管理外部的包参考方法如下

当前IDE开发环境的版本

<img src="https://s2.loli.net/2022/05/28/RBx6jDHzeyKEfvb.png" alt="Untitled.png" style="zoom:45%;" />

**遇到的问题**

- 代码中 import [github.com/ethereum/go-ethereum/ethclient](http://github.com/ethereum/go-ethereum/ethclient) 无法找到对应的包
  
    解决方法
    
    - 使用go的mod管理方式来下载包和管理
      
        https://github.com/shenyubo1982/ethDemo/commit/c3e011105935466b04ea49d0f81b62e3c67a42c7
        
        ```go
        go get github.com/ethereum/go-ethereum/ethclient
        vi go.mod
        vi go.sum
        ```
    
- 编译运行代码发现
  
    ```bash
    missing go.sum entry for module providing package <package_name>
    ```
    
    - 解决方法
      
        ```bash
        go mod tidy
        # 去除
        ```

