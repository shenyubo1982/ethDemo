# ethDemo


<a href="https://img.shields.io/badge/go-1.17.6-orange" alt="Contributors">
<img src="https://img.shields.io/badge/go-1.17.6-orange" /></a>

<!--START_SECTION:waka-->
<!--END_SECTION:waka-->

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

ethclient.Dial 参数是你的以太坊节点 url。若您没有现有以太坊客户端，您可以连接到 infura 网关。Infura 管理着一批安全，可靠，可扩展的以太坊[geth 和 parity]节点，并且在接入以太坊网络时降低了新人的入门门槛。

我们使用了 JetBrains 的 goLand，所以使用 go.mod 管理外部的包参考方法如下

当前 IDE 开发环境的版本

<img src="https://s2.loli.net/2022/05/28/RBx6jDHzeyKEfvb.png" alt="Untitled.png" style="zoom:45%;" />

**遇到的问题**

- 代码中 import [github.com/ethereum/go-ethereum/ethclient](http://github.com/ethereum/go-ethereum/ethclient) 无法找到对应的包

  解决方法

  - 使用 go 的 mod 管理方式来下载包和管理

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

### 课程内容 1：获取 ETH 地址对应的 Balance

```go
client.BalanceAt(context.Background(), account, nil)
```

### 课程内容 2：如何将 Balance 转换成 eth 币

使用 wei 转换至币的方法：converWeiToBi

### 课程内容 3：如何根据地址和区块号，查询该区块号上的 Balance

这个部分还未实验成功。原因是我使用了 infura 的服务，他们查询存档服务是收费的。Todo 实验其他方式查询区块告诉的 Balance

### 本次添加的代码模块如下：

```
KovanTestEthValue := getBalanceFromAddress(*client, kovanTestEthAddress)
func getBalanceFromAddress(client ethclient.Client, address string) *big.Float
func convertWeiToBi(balance *big.Int) (ethValue *big.Float)

```

# 2022.06.03 变更内容

1. 项目文件结构变动
   - abi：存放已部署完成的合约 abi 文件，还有根据脚本生成的配套.go 文件。其实我们 go 项目中，只会用到 go 文件
   - chainClient：我们的客户端与区块链网络连接器。目前我们都是连接基于 eth 的网络，主要用于我们 metainchain 的网络测试，也可以用于 eth 测试网络
   - config：配置文件。我们测试场景需要测试不同的网络不同的特性、不同的账号。目前可以根据网络区分不同的文件读取，获取测试需要的基本信息。
   - scripts：shell 自动化脚本库。用来执行一些常用的脚本命令。abi2go.sh 是将 abi 文件下的 abi 文件编译成 go 文件，今后考虑将所有的脚本都搬至此文件架下
   - testing：目前空，今后会将自动化脚本运行的报告放置此处。
   - util：基建模块。目前包括了 yaml 读取
   - wiki：全局性的说明文档搬至此处
2. 关于测试方法
   - 目前考虑使用简单的 go test 方式进行测试。每一个目录下放置一个 test 做模块的单体测试。
   - 关于区块链的功能测试，放置在 chainClient 目录下。目前使用 client_test
   - 执行测试的方式可以直接在命令行调用 test.sh 会执行目录下所有的 test 文件，并生成测试报告包含覆盖率。但是这个覆盖率只是我们自己编码的覆盖率。
   - 可以尝试在 test 文件中编写测试用例，然后自动化运行测试脚本，后续成熟了也可以集成 CI
3. 使用方法还没有具体说明清楚待补充。

待办事项：

在 chainClient 中，添加常用的区块链功能测试模块化，尽可能抽象处理达到通用的效果，包括以下内容。

- 创建账号
- 区块高度
- 查询交易
- 交易
  - 转账（from to)
  - 合约调用（合约方法和参数可以泛型代入）
  - NFT 相关功能部分还不明白，待确认。
- 共识
- 性能

# 2022.06.10 mac mini workspace test.
# 2022.07.04 add new feat:from terminal to show eth address with local keystore folder.

​
