# 如何利用go-ethereum 调用已部署的合约中的某个方法

我们使用的Demo场景：使用元生链网络，调用`EvidenceContract`合约中`AddInfo`方法。 

## 准备环境

1. you have to installed Go environment.
   
    go version go1.17.6 darwin/amd64
    
2. chose your IDE.
   
     GoLand or anyway
    
3. Install The Ethereum environment for Go
    - Install ethereum In MacOs
      
        open [https://geth.ethereum.org/docs/install-and-build/installing-geth#macos-via-homebrew](https://geth.ethereum.org/docs/install-and-build/installing-geth#macos-via-homebrew) , and step by step.
        
        # MacOS via Homebrew
        
        The easiest way to install go-ethereum is to use the Geth Homebrew tap. The first step is to check that Homebrew is installed. The following command should return a version number.
        
        `brew -v`
        
        If a version number is returned, then Homebrew is installed. If not, Homebrew can be installed by following the instructions [here](https://brew.sh/). With Homebrew installed, the following commands add the Geth tap and install Geth:
        
        `brew tap ethereum/ethereum
        brew install ethereum`
        
        The previous command installs the latest stable release. Developers that wish to install the most up-to-date version can install the Geth repository’s master branch by adding the `--devel` parameter to the install command:
        
        `brew install ethereum --devel`
        
        These commands install the core Geth software and the following developer tools: `clef`, `devp2p`, `abigen`, `bootnode`, `evm`, `rlpdump` and `puppeth`. The binaries for each of these tools are saved in `/usr/local/bin/`. The full list of command line options can be viewed [here](https://geth.ethereum.org/docs/interface/command-line-options) or in the terminal by running `geth --help`.
        
        Updating an existing Geth installation to the latest version can be achieved by stopping the node and running the following commands:
        
        `brew update 
        brew upgrade 
        brew reinstall ethereum`
        
        When the node is started again, Geth will automatically use all the data from the previous version and sync the blocks that were missed while the node was offline.
        
        When finished. you can config your Geth bin file path.
        
        ```bash
        # config your ethereum path like this .(eth path 1.10.18)
        export PATH="/usr/local/Cellar/ethereum/1.10.18/bin:$PATH"
        ```
        
    - If you have already published your contract in Chain.
        - import your contract.abi
        - after you will create your contract.go

## 核心环节

1. 成功连接区块链网络客户端
   
    ```go
    client, err := ethclient.Dial(chainUrl)
    ```
    
2. 完成部署合约
   
    建议使用remix部署。关于合约生成可以参考团队内部的培训资料，此处略过，后续再添加学习资料
    
    ```go
    // 假设我们完成了部署，并获得合约的地址，并配置在我们的代码中
    contractAddressHex := "0x03Bc2D794B2FcDA47a9dBb1d43B1fA7B05260282"
    ```
    
3. 生成合约的go文件
    - 检查abigen工具包是安装成功
      如果你在准备中已安装了geth和配置了geth的bin路径，那可以直接使用abigen命令。执行后会生成1个go文件。
      
        ```bash
        abigen --abi=EvidenceContract.abi --pkg=evidencecountract --out=EvidenceContract.go
        # EvidenceContract.abi是你部署后获得的abi文件，放在你的go项目目录下。
        # evidencecountract 是合约名称的包名（这里故意使用错别字 country ，为了更好地演示代码中哪里可以使用它！）
        # 细心的你会发现：在我们的go项目中，可以这样使用它
        # instance, err := evidencecountract.NewEvidencecountract(address, &client)
        # EvidenceContract.go 是生成的结果文件，放在你的go项目中。
        ```
    
4. 配置信息
    1. 合约地址（部署在区块链上的合约地址）
    2. 合约部署时使用的账户私钥Hex（发起者私钥）
    3. 合约部署时使用的账户地址（发起者地址）
5. 配置合约的abi
   
    将abi文件放置在go项目中，建议放在固定目录中。
    
    ├── abi
     │   ├── EvidenceContract.abi
     │   ├── EvidenceContract.go
    
6. 生成合约实例
    - code
      
        ```go
        instance, err := evidencecountract.NewEvidencecountract(address, &client)
        iferr != nil {
        	panic(err)
        }
        ```
    
7. 用签名者的私钥在指定的区块链网络中创建一个交易，获得auth（`auth参考：TransactOpts`）
    - Code
      
        ```go
        auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(77))
        iferr != nil {
           log.Fatal(err)
        }
        ```
    
8. 创建一个transactOpts，并设定结构体的内容
    - Code
      
        ```go
        auth.GasLimit = 8100000
        auth.GasPrice = big.NewInt(1000000000)
        
        transactOpts := &bind.TransactOpts{
           From:   fromAddress,
           Nonce:  auth.Nonce,
           Signer: auth.Signer,
           Value:    big.NewInt(100),
           GasPrice: auth.GasPrice,
           GasLimit: auth.GasLimit,
           Context:  auth.Context,
           NoSend:   false,
        }
        ```
    
9. 使用合约实例，调用合约中的具体方法
    - Code
      
        ```go
        tx, err := instance.AddInfo(transactOpts, title, name, content)
        iferr != nil {
           log.Fatal(err)
        }
        ```
        

## 运行结果

1. 测试代码
   
    ```go
    // 调用合约功能测试
    func TestCallContract(t *testing.T) {
    	t.Run("CAllContract", func(t *testing.T) {
    		t.Helper()
    		chainUrl := "http://172.17.4.13:7755"
    		ansClient, err := launch(chainUrl)
    		if ansClient == nil || err != nil {
    			t.Fatal("Can't get Client")
    		}
    		privateKeyHex := "794c479028076af7673a6941185af09a51c86a44082b438dbdfca70b6c6829ed"
    		contractAddressHex := "0x03Bc2D794B2FcDA47a9dBb1d43B1fA7B05260282"
    		title := "Title-goLang"
    		name := "bobo-go"
    		content := "content-by-golang"
    		callContract(*ansClient, contractAddressHex, privateKeyHex, title, name, content)
    	})
    }
    ```
    
2. 运行结果
   ![EvidenceContract-RuningPass01.png](https://s2.loli.net/2022/06/01/QxCeJMfnP9H2IyF.png) 

---

## 遇到的问题和解决方法

- 问题描述
  
    **error：no contract code at given address**
    
    - 分析
        - Debug 源代码。发现在调用transact 方法的时候，opts.GasPrice 为空，所以报错。为什么代码在获取auth的时候没有获取到gasPrice 和gaslimit ，待后续继续调查。如果仍然报错，可以继续用这个debug的方法看看哪个参数为空或者不正常，继续手动调整后，再尝试调用合约。
        - **TODO：为什么这些参数不正常还需后续继续调查。**
    
    ```go
    // base.go
    func (c *BoundContract) transact(opts *TransactOpts, contract *common.Address, input []byte) (*types.Transaction, error) {
    	if opts.GasPrice != nil && (opts.GasFeeCap != nil || opts.GasTipCap != nil) {
    		return nil, errors.New("both gasPrice and (maxFeePerGas or maxPriorityFeePerGas) specified")
    	}
    	// Create the transaction
    	var (
    		rawTx *types.Transaction
    		err   error
    	)
    	if opts.GasPrice != nil {
    		rawTx, err = c.createLegacyTx(opts, contract, input)
    	} else {
    		// Only query for basefee if gasPrice not specified
    		if head, errHead := c.transactor.HeaderByNumber(ensureContext(opts.Context), nil); errHead != nil {
    			return nil, errHead
    		} else if head.BaseFee != nil {
    			rawTx, err = c.createDynamicTx(opts, contract, input, head)
    		} else {
    			// Chain is not London ready -> use legacy transaction
    			rawTx, err = c.createLegacyTx(opts, contract, input)
    		}
    	}
    	if err != nil {
    		return nil, err
    	}
    	// Sign the transaction and schedule it for execution
    	if opts.Signer == nil {
    		return nil, errors.New("no signer to authorize the transaction with")
    	}
    	signedTx, err := opts.Signer(opts.From, rawTx)
    	if err != nil {
    		return nil, err
    	}
    	if opts.NoSend {
    		return signedTx, nil
    	}
    	if err := c.transactor.SendTransaction(ensureContext(opts.Context), signedTx); err != nil {
    		return nil, err
    	}
    	return signedTx, nil
    }
    ```
    
    - 解决
      
        修改调用合约时候的GasPrice 和GasLimit
        
        ```go
        // 调整 GasLimit 和 GasPrice
        auth.GasLimit = 8100000
        auth.GasPrice = big.NewInt(1000000000)
        
        //设定transactOpts
        transactOpts := &bind.TransactOpts{
        	From:   fromAddress,
        	Nonce:  auth.Nonce,
        	Signer: auth.Signer,
        	//Value:    big.NewInt(0),
        	Value:    big.NewInt(100),
        	GasPrice: auth.GasPrice,
        	GasLimit: auth.GasLimit,
        	Context:  auth.Context,
        	NoSend:   false,
        }
        ```
        
    - 参考资料
      
        > https://github.com/ethereum/go-ethereum/issues/15930
        I resolved the issue by increasing the gas limit.
        > 

## 名称说明

- Geth：go-ethereum 。可以参考源代码