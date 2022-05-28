# ethDemo


# GoEthereumBook 01：初始化客户端

教程地址

[https://goethereumbook.org/zh/client-setup/](https://goethereumbook.org/zh/client-setup/)

开发环境：goland

go version

```bash
go version go1.17.6 darwin/amd64
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

# GoEthereumBook 02：初始化客户端

### 课程内容1：获取ETH地址对应的Balance

``` go
client.BalanceAt(context.Background(), account, nil)
```

### 课程内容2：如何将Balance转换成eth币

使用wei转换至币的方法：converWeiToBi

### 课程内容3：如何根据地址和区块号，查询该区块号上的Balance

这个部分还未实验成功。原因是我使用了infura的服务，他们查询存档服务是收费的。Todo 实验其他方式查询区块告诉的Balance

### 本次添加的代码模块如下：

```
KovanTestEthValue := getBalanceFromAddress(*client, kovanTestEthAddress)
func getBalanceFromAddress(client ethclient.Client, address string) *big.Float 
func convertWeiToBi(balance *big.Int) (ethValue *big.Float) 

```

