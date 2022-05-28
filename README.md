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
	_ "github.com/ethereum/go-ethereum/ethclient/gethclient"
	"log"
)

func main() {

	//client, err := ethclient.Dial("ADD_YOUR_ETHEREUM_NODE_URL")
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

![Untitled](GoEthereumBook%2001%EF%BC%9A%E5%88%9D%E5%A7%8B%E5%8C%96%E5%AE%A2%E6%88%B7%E7%AB%AF%20d5bb2f9deaf64eacb791935dad1d2859/Untitled.png)

**遇到的问题**

- 代码中 import [github.com/ethereum/go-ethereum/ethclient](http://github.com/ethereum/go-ethereum/ethclient) 无法找到对应的包
  
    解决方法
    
    - 使用go的mod管理方式来下载包和管理
      
        ```go
        go get github.com/ethereum/go-ethereum/ethclient
        
        vi go.mod
        =========================
        module ethDemo
        
        go 1.17
        
        require github.com/ethereum/go-ethereum v1.10.18
        
        require (
        	github.com/StackExchange/wmi v0.0.0-20180116203802-5d049714c4a6 // indirect
        	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
        	github.com/deckarep/golang-set v1.8.0 // indirect
        	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
        	github.com/go-ole/go-ole v1.2.1 // indirect
        	github.com/go-stack/stack v1.8.0 // indirect
        	github.com/golang/snappy v0.0.4 // indirect
        	github.com/gorilla/websocket v1.4.2 // indirect
        	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d // indirect
        	github.com/huin/goupnp v1.0.3 // indirect
        	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
        	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
        	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7 // indirect
        	github.com/tklauser/go-sysconf v0.3.5 // indirect
        	github.com/tklauser/numcpus v0.2.2 // indirect
        	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
        	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
        	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 // indirect
        	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
        )
        =====================
        
        vi go.sum
        
        =========================
        cloud.google.com/go v0.26.0/go.mod h1:aQUYkXzVsufM+DwF1aE+0xfcU+56JwCaLick0ClmMTw=
        cloud.google.com/go v0.34.0/go.mod h1:aQUYkXzVsufM+DwF1aE+0xfcU+56JwCaLick0ClmMTw=
        cloud.google.com/go v0.38.0/go.mod h1:990N+gfupTy94rShfmMCWGDn0LpTmnzTp2qbd1dvSRU=
        cloud.google.com/go v0.43.0/go.mod h1:BOSR3VbTLkk6FDC/TcffxP4NF/FFBGA5ku+jvKOP7pg=
        cloud.google.com/go v0.44.1/go.mod h1:iSa0KzasP4Uvy3f1mN/7PiObzGgflwredwwASm/v6AU=
        cloud.google.com/go v0.44.2/go.mod h1:60680Gw3Yr4ikxnPRS/oxxkBccT6SA1yMk63TGekxKY=
        cloud.google.com/go v0.45.1/go.mod h1:RpBamKRgapWJb87xiFSdk4g1CME7QZg3uwTez+TSTjc=
        cloud.google.com/go v0.46.3/go.mod h1:a6bKKbmY7er1mI7TEI4lsAkts/mkhTSZK8w33B4RAg0=
        cloud.google.com/go v0.50.0/go.mod h1:r9sluTvynVuxRIOHXQEHMFffphuXHOMZMycpNR5e6To=
        cloud.google.com/go v0.51.0/go.mod h1:hWtGJ6gnXH+KgDv+V0zFGDvpi07n3z8ZNj3T1RW0Gcw=
        cloud.google.com/go/bigquery v1.0.1/go.mod h1:i/xbL2UlR5RvWAURpBYZTtm/cXjCha9lbfbpx4poX+o=
        cloud.google.com/go/bigquery v1.3.0/go.mod h1:PjpwJnslEMmckchkHFfq+HTD2DmtT67aNFKH1/VBDHE=
        cloud.google.com/go/bigtable v1.2.0/go.mod h1:JcVAOl45lrTmQfLj7T6TxyMzIN/3FGGcFm+2xVAli2o=
        cloud.google.com/go/datastore v1.0.0/go.mod h1:LXYbyblFSglQ5pkeyhO+Qmw7ukd3C+pD7TKLgZqpHYE=
        cloud.google.com/go/pubsub v1.0.1/go.mod h1:R0Gpsv3s54REJCy4fxDixWD93lHJMoZTyQ2kNxGRt3I=
        cloud.google.com/go/pubsub v1.1.0/go.mod h1:EwwdRX2sKPjnvnqCa270oGRyludottCI76h+R3AArQw=
        cloud.google.com/go/storage v1.0.0/go.mod h1:IhtSnM/ZTZV8YYJWCY8RULGVqBDmpoyjwiyrjsg+URw=
        cloud.google.com/go/storage v1.5.0/go.mod h1:tpKbwo567HUNpVclU5sGELwQWBDZ8gh0ZeosJ0Rtdos=
        collectd.org v0.3.0/go.mod h1:A/8DzQBkF6abtvrT2j/AU/4tiBgJWYyh0y/oB/4MlWE=
        dmitri.shuralyov.com/gpu/mtl v0.0.0-20190408044501-666a987793e9/go.mod h1:H6x//7gZCb22OMCxBHrMx7a5I7Hp++hsVxbQ4BYO7hU=
        github.com/Azure/azure-sdk-for-go/sdk/azcore v0.21.1/go.mod h1:fBF9PQNqB8scdgpZ3ufzaLntG0AG7C1WjPMsiFOmfHM=
        github.com/Azure/azure-sdk-for-go/sdk/internal v0.8.3/go.mod h1:KLF4gFr6DcKFZwSuH8w8yEK6DpFl3LP5rhdvAb7Yz5I=
        github.com/Azure/azure-sdk-for-go/sdk/storage/azblob v0.3.0/go.mod h1:tPaiy8S5bQ+S5sOiDlINkp7+Ef339+Nz5L5XO+cnOHo=
        github.com/BurntSushi/toml v0.3.1/go.mod h1:xHWCNGjB5oqiDr8zfno3MHue2Ht5sIBksp03qcyfWMU=
        github.com/BurntSushi/xgb v0.0.0-20160522181843-27f122750802/go.mod h1:IVnqGOEym/WlBOVXweHU+Q+/VP0lqqI8lqeDx9IjBqo=
        github.com/DATA-DOG/go-sqlmock v1.3.3/go.mod h1:f/Ixk793poVmq4qj/V1dPUg2JEAKC73Q5eFN3EC/SaM=
        github.com/OneOfOne/xxhash v1.2.2/go.mod h1:HSdplMjZKSmBqAxg5vPj2TmRDmfkzw+cTzAElWljhcU=
        github.com/StackExchange/wmi v0.0.0-20180116203802-5d049714c4a6 h1:fLjPD/aNc3UIOA6tDi6QXUemppXK3P9BI7mr2hd6gx8=
        github.com/StackExchange/wmi v0.0.0-20180116203802-5d049714c4a6/go.mod h1:3eOhrUMpNV+6aFIbp5/iudMxNCF27Vw2OZgy4xEx0Fg=
        github.com/VictoriaMetrics/fastcache v1.6.0 h1:C/3Oi3EiBCqufydp1neRZkqcwmEiuRT9c3fqvvgKm5o=
        github.com/VictoriaMetrics/fastcache v1.6.0/go.mod h1:0qHz5QP0GMX4pfmMA/zt5RgfNuXJrTP0zS7DqpHGGTw=
        github.com/ajstarks/svgo v0.0.0-20180226025133-644b8db467af/go.mod h1:K08gAheRH3/J6wwsYMMT4xOr94bZjxIelGM0+d/wbFw=
        github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc/go.mod h1:LOuyumcjzFXgccqObfd/Ljyb9UuFJ6TxHnclSeseNhc=
        github.com/alecthomas/units v0.0.0-20151022065526-2efee857e7cf/go.mod h1:ybxpYRFXyAe+OPACYpWeL0wqObRcbAqCMya13uyzqw0=
        github.com/allegro/bigcache v1.2.1-0.20190218064605-e24eb225f156/go.mod h1:Cb/ax3seSYIx7SuZdm2G2xzfwmv3TPSk2ucNfQESPXM=
        github.com/andreyvit/diff v0.0.0-20170406064948-c7f18ee00883/go.mod h1:rCTlJbsFo29Kk6CurOXKm700vrz8f0KW0JNfpkRJY/8=
        github.com/apache/arrow/go/arrow v0.0.0-20191024131854-af6fa24be0db/go.mod h1:VTxUBvSJ3s3eHAg65PNgrsn5BtqCRPdmyXh6rAfdxN0=
        github.com/aws/aws-sdk-go-v2 v1.2.0/go.mod h1:zEQs02YRBw1DjK0PoJv3ygDYOFTre1ejlJWl8FwAuQo=
        github.com/aws/aws-sdk-go-v2/config v1.1.1/go.mod h1:0XsVy9lBI/BCXm+2Tuvt39YmdHwS5unDQmxZOYe8F5Y=
        github.com/aws/aws-sdk-go-v2/credentials v1.1.1/go.mod h1:mM2iIjwl7LULWtS6JCACyInboHirisUUdkBPoTHMOUo=
        github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.0.2/go.mod h1:3hGg3PpiEjHnrkrlasTfxFqUsZ2GCk/fMUn4CbKgSkM=
        github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.0.2/go.mod h1:45MfaXZ0cNbeuT0KQ1XJylq8A6+OpVV2E5kvY/Kq+u8=
        github.com/aws/aws-sdk-go-v2/service/route53 v1.1.1/go.mod h1:rLiOUrPLW/Er5kRcQ7NkwbjlijluLsrIbu/iyl35RO4=
        github.com/aws/aws-sdk-go-v2/service/sso v1.1.1/go.mod h1:SuZJxklHxLAXgLTc1iFXbEWkXs7QRTQpCLGaKIprQW0=
        github.com/aws/aws-sdk-go-v2/service/sts v1.1.1/go.mod h1:Wi0EBZwiz/K44YliU0EKxqTCJGUfYTWXrrBwkq736bM=
        github.com/aws/smithy-go v1.1.0/go.mod h1:EzMw8dbp/YJL4A5/sbhGddag+NPT7q084agLbB9LgIw=
        github.com/beorn7/perks v0.0.0-20180321164747-3a771d992973/go.mod h1:Dwedo/Wpr24TaqPxmxbtue+5NUziq4I4S80YR8gNf3Q=
        github.com/beorn7/perks v1.0.0/go.mod h1:KWe93zE9D1o94FZ5RNwFwVgaQK1VOXiVxmqh+CedLV8=
        github.com/bmizerany/pat v0.0.0-20170815010413-6226ea591a40/go.mod h1:8rLXio+WjiTceGBHIoTvn60HIbs7Hm7bcHjyrSqYB9c=
        github.com/boltdb/bolt v1.3.1/go.mod h1:clJnj/oiGkjum5o1McbSZDSLxVThjynRyGBgiAx27Ps=
        github.com/btcsuite/btcd/btcec/v2 v2.2.0 h1:fzn1qaOt32TuLjFlkzYSsBC35Q3KUjT1SwPxiMSCF5k=
        github.com/btcsuite/btcd/btcec/v2 v2.2.0/go.mod h1:U7MHm051Al6XmscBQ0BoNydpOTsFAn707034b5nY8zU=
        github.com/btcsuite/btcd/chaincfg/chainhash v1.0.1 h1:q0rUy8C/TYNBQS1+CGKw68tLOFYSNEs0TFnxxnS9+4U=
        github.com/btcsuite/btcd/chaincfg/chainhash v1.0.1/go.mod h1:7SFka0XMvUgj3hfZtydOrQY2mwhPclbT2snogU7SQQc=
        github.com/c-bata/go-prompt v0.2.2/go.mod h1:VzqtzE2ksDBcdln8G7mk2RX9QyGjH+OVqOCSiVIqS34=
        github.com/census-instrumentation/opencensus-proto v0.2.1/go.mod h1:f6KPmirojxKA12rnyqOA5BBL4O983OfeGPqjHWSTneU=
        github.com/cespare/cp v0.1.0/go.mod h1:SOGHArjBr4JWaSDEVpWpo/hNg6RoKrls6Oh40hiwW+s=
        github.com/cespare/xxhash v1.1.0 h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=
        github.com/cespare/xxhash v1.1.0/go.mod h1:XrSqR1VqqWfGrhpAt58auRo0WTKS1nRRg3ghfAqPWnc=
        github.com/cespare/xxhash/v2 v2.1.1 h1:6MnRN8NT7+YBpUIWxHtefFZOKTAPgGjpQSxqLNn0+qY=
        github.com/cespare/xxhash/v2 v2.1.1/go.mod h1:VGX0DQ3Q6kWi7AoAeZDth3/j3BFtOZR5XLFGgcrjCOs=
        github.com/chzyer/logex v1.1.10/go.mod h1:+Ywpsq7O8HXn0nuIou7OrIPyXbp3wmkHB+jjWRnGsAI=
        github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e/go.mod h1:nSuG5e5PlCu98SY8svDHJxuZscDgtXS6KTTbou5AhLI=
        github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1/go.mod h1:Q3SI9o4m/ZMnBNeIyt5eFwwo7qiLfzFZmjNmxjkiQlU=
        github.com/client9/misspell v0.3.4/go.mod h1:qj6jICC3Q7zFZvVWo7KLAzC3yx5G7kyvSDkc90ppPyw=
        github.com/cloudflare/cloudflare-go v0.14.0/go.mod h1:EnwdgGMaFOruiPZRFSgn+TsQ3hQ7C/YWzIGLeu5c304=
        github.com/consensys/bavard v0.1.8-0.20210406032232-f3452dc9b572/go.mod h1:Bpd0/3mZuaj6Sj+PqrmIquiOKy397AKGThQPaGzNXAQ=
        github.com/consensys/gnark-crypto v0.4.1-0.20210426202927-39ac3d4b3f1f/go.mod h1:815PAHg3wvysy0SyIqanF8gZ0Y1wjk/hrDHD/iT88+Q=
        github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d/go.mod h1:maD7wRr/U5Z6m/iR4s+kqSMx2CaBsrgA7czyZG/E6dU=
        github.com/creack/pty v1.1.9/go.mod h1:oKZEueFk5CKHvIhNR5MUki03XCEU+Q6VDXinZuGJ33E=
        github.com/cyberdelia/templates v0.0.0-20141128023046-ca7fffd4298c/go.mod h1:GyV+0YP4qX0UQ7r2MoYZ+AvYDp12OF5yg4q8rGnyNh4=
        github.com/dave/jennifer v1.2.0/go.mod h1:fIb+770HOpJ2fmN9EPPKOqm1vMGhB+TwXKMZhrIygKg=
        github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
        github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
        github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
        github.com/deckarep/golang-set v1.8.0 h1:sk9/l/KqpunDwP7pSjUg0keiOOLEnOBHzykLrsPppp4=
        github.com/deckarep/golang-set v1.8.0/go.mod h1:5nI87KwE7wgsBU1F4GKAw2Qod7p5kyS383rP6+o6qqo=
        github.com/decred/dcrd/crypto/blake256 v1.0.0 h1:/8DMNYp9SGi5f0w7uCm6d6M4OU2rGFK09Y2A4Xv7EE0=
        github.com/decred/dcrd/crypto/blake256 v1.0.0/go.mod h1:sQl2p6Y26YV+ZOcSTP6thNdn47hh8kt6rqSlvmrXFAc=
        github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 h1:YLtO71vCjJRCBcrPMtQ9nqBsqpA1m5sE92cU+pd5Mcc=
        github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1/go.mod h1:hyedUtir6IdtD/7lIxGeCxkaw7y45JueMRL4DIyJDKs=
        github.com/deepmap/oapi-codegen v1.6.0/go.mod h1:ryDa9AgbELGeB+YEXE1dR53yAjHwFvE9iAUlWl9Al3M=
        github.com/deepmap/oapi-codegen v1.8.2/go.mod h1:YLgSKSDv/bZQB7N4ws6luhozi3cEdRktEqrX88CvjIw=
        github.com/dgrijalva/jwt-go v3.2.0+incompatible/go.mod h1:E3ru+11k8xSBh+hMPgOLZmtrrCbhqsmaPHjLKYnJCaQ=
        github.com/dgryski/go-bitstream v0.0.0-20180413035011-3522498ce2c8/go.mod h1:VMaSuZ+SZcx/wljOQKvp5srsbCiKDEb6K2wC4+PiBmQ=
        github.com/dgryski/go-sip13 v0.0.0-20181026042036-e10d5fee7954/go.mod h1:vAd38F8PWV+bWy6jNmig1y/TA+kYO4g3RSRF0IAv0no=
        github.com/dlclark/regexp2 v1.4.1-0.20201116162257-a2a8dda75c91/go.mod h1:2pZnwuY/m+8K6iRw6wQdMtk+rH5tNGR1i55kozfMjCc=
        github.com/dnaeon/go-vcr v1.1.0/go.mod h1:M7tiix8f0r6mKKJ3Yq/kqU1OYf3MnfmBWVbPx/yU9ko=
        github.com/dnaeon/go-vcr v1.2.0/go.mod h1:R4UdLID7HZT3taECzJs4YgbbH6PIGXB6W/sc5OLb6RQ=
        github.com/docker/docker v1.6.2/go.mod h1:eEKB0N0r5NX/I1kEveEz05bcu8tLC/8azJZsviup8Sk=
        github.com/dop251/goja v0.0.0-20220405120441-9037c2b61cbf/go.mod h1:R9ET47fwRVRPZnOGvHxxhuZcbrMCuiqOz3Rlrh4KSnk=
        github.com/dop251/goja_nodejs v0.0.0-20210225215109-d91c329300e7/go.mod h1:hn7BA7c8pLvoGndExHudxTDKZ84Pyvv+90pbBjbTz0Y=
        github.com/eclipse/paho.mqtt.golang v1.2.0/go.mod h1:H9keYFcgq3Qr5OUJm/JZI/i6U7joQ8SYLhZwfeOo6Ts=
        github.com/edsrzf/mmap-go v1.0.0 h1:CEBF7HpRnUCSJgGUb5h1Gm7e3VkmVDrR8lvWVLtrOFw=
        github.com/edsrzf/mmap-go v1.0.0/go.mod h1:YO35OhQPt3KJa3ryjFM5Bs14WD66h8eGKpfaBNrHW5M=
        github.com/envoyproxy/go-control-plane v0.9.1-0.20191026205805-5f8ba28d4473/go.mod h1:YTl/9mNaCwkRvm6d1a2C3ymFceY/DCBVvsKhRF0iEA4=
        github.com/envoyproxy/protoc-gen-validate v0.1.0/go.mod h1:iSmxcyjqTsJpI2R4NaDN7+kN2VEUnK/pcBlmesArF7c=
        github.com/ethereum/go-ethereum v1.10.18 h1:hLEd5M+UD0GJWPaROiYMRgZXl6bi5YwoTJSthsx5CZw=
        github.com/ethereum/go-ethereum v1.10.18/go.mod h1:RD3NhcSBjZpj3k+SnQq24wBrmnmie78P5R/P62iNBD8=
        github.com/fatih/color v1.7.0/go.mod h1:Zm6kSWBoL9eyXnKyktHP6abPY2pDugNf5KwzbycvMj4=
        github.com/fjl/gencodec v0.0.0-20220412091415-8bb9e558978c/go.mod h1:AzA8Lj6YtixmJWL+wkKoBGsLWy9gFrAzi4g+5bCKwpY=
        github.com/fjl/memsize v0.0.0-20190710130421-bcb5799ab5e5 h1:FtmdgXiUlNeRsoNMFlKLDt+S+6hbjVMEW6RGQ7aUf7c=
        github.com/fjl/memsize v0.0.0-20190710130421-bcb5799ab5e5/go.mod h1:VvhXpOYNQvB+uIk2RvXzuaQtkQJzzIx6lSBe1xv7hi0=
        github.com/fogleman/gg v1.2.1-0.20190220221249-0403632d5b90/go.mod h1:R/bRT+9gY/C5z7JzPU0zXsXHKM4/ayA+zqcVNZzPa1k=
        github.com/fsnotify/fsnotify v1.4.7/go.mod h1:jwhsz4b93w/PPRr/qN1Yymfu8t87LnFCMoQvtojpjFo=
        github.com/fsnotify/fsnotify v1.4.9 h1:hsms1Qyu0jgnwNXIxa+/V/PDsU6CfLf6CNO8H7IWoS4=
        github.com/fsnotify/fsnotify v1.4.9/go.mod h1:znqG4EE+3YCdAaPaxE2ZRY/06pZUdp0tY4IgpuI1SZQ=
        github.com/garslo/gogen v0.0.0-20170306192744-1d203ffc1f61/go.mod h1:Q0X6pkwTILDlzrGEckF6HKjXe48EgsY/l7K7vhY4MW8=
        github.com/gballet/go-libpcsclite v0.0.0-20190607065134-2772fd86a8ff h1:tY80oXqGNY4FhTFhk+o9oFHGINQ/+vhlm8HFzi6znCI=
        github.com/gballet/go-libpcsclite v0.0.0-20190607065134-2772fd86a8ff/go.mod h1:x7DCsMOv1taUwEWCzT4cmDeAkigA5/QCwUodaVOe8Ww=
        github.com/getkin/kin-openapi v0.53.0/go.mod h1:7Yn5whZr5kJi6t+kShccXS8ae1APpYTW6yheSwk8Yi4=
        github.com/getkin/kin-openapi v0.61.0/go.mod h1:7Yn5whZr5kJi6t+kShccXS8ae1APpYTW6yheSwk8Yi4=
        github.com/ghodss/yaml v1.0.0/go.mod h1:4dBDuWmgqj2HViK6kFavaiC9ZROes6MMH2rRYeMEF04=
        github.com/glycerine/go-unsnap-stream v0.0.0-20180323001048-9f0cb55181dd/go.mod h1:/20jfyN9Y5QPEAprSgKAUr+glWDY39ZiUEAYOEv5dsE=
        github.com/glycerine/goconvey v0.0.0-20190410193231-58a59202ab31/go.mod h1:Ogl1Tioa0aV7gstGFO7KhffUsb9M4ydbEbbxpcEDc24=
        github.com/go-chi/chi/v5 v5.0.0/go.mod h1:BBug9lr0cqtdAhsu6R4AAdvufI0/XBzAQSsUqJpoZOs=
        github.com/go-gl/glfw v0.0.0-20190409004039-e6da0acd62b1/go.mod h1:vR7hzQXu2zJy9AVAgeJqvqgH9Q5CA+iKCZ2gyEVpxRU=
        github.com/go-gl/glfw/v3.3/glfw v0.0.0-20191125211704-12ad95a8df72/go.mod h1:tQ2UAYgL5IevRw8kRxooKSPJfGvJ9fJQFa0TUsXzTg8=
        github.com/go-kit/kit v0.8.0/go.mod h1:xBxKIO96dXMWWy0MnWVtmwkA9/13aqxPnvrjFYMA2as=
        github.com/go-logfmt/logfmt v0.3.0/go.mod h1:Qt1PoO58o5twSAckw1HlFXLmHsOX5/0LbT9GBnD5lWE=
        github.com/go-logfmt/logfmt v0.4.0/go.mod h1:3RMwSq7FuexP4Kalkev3ejPJsZTpXXBr9+V4qmtdjCk=
        github.com/go-ole/go-ole v1.2.1 h1:2lOsA72HgjxAuMlKpFiCbHTvu44PIVkZ5hqm3RSdI/E=
        github.com/go-ole/go-ole v1.2.1/go.mod h1:7FAglXiTm7HKlQRDeOQ6ZNUHidzCWXuZWq/1dTyBNF8=
        github.com/go-openapi/jsonpointer v0.19.5/go.mod h1:Pl9vOtqEWErmShwVjC8pYs9cog34VGT37dQOVbmoatg=
        github.com/go-openapi/swag v0.19.5/go.mod h1:POnQmlKehdgb5mhVOsnJFsivZCEZ/vjK9gh66Z9tfKk=
        github.com/go-sourcemap/sourcemap v2.1.3+incompatible/go.mod h1:F8jJfvm2KbVjc5NqelyYJmf/v5J0dwNLS2mL4sNA1Jg=
        github.com/go-sql-driver/mysql v1.4.1/go.mod h1:zAC/RDZ24gD3HViQzih4MyKcchzm+sOG5ZlKdlhCg5w=
        github.com/go-stack/stack v1.8.0 h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=
        github.com/go-stack/stack v1.8.0/go.mod h1:v0f6uXyyMGvRgIKkXu+yp6POWl0qKG85gN/melR3HDY=
        github.com/gofrs/uuid v3.3.0+incompatible/go.mod h1:b2aQJv3Z4Fp6yNu3cdSllBxTCLRxnplIgP/c0N/04lM=
        github.com/gogo/protobuf v1.1.1/go.mod h1:r8qH/GZQm5c6nD/R0oafs1akxWv10x8SbQlK7atdtwQ=
        github.com/gogo/protobuf v1.3.1/go.mod h1:SlYgWuQ5SjCEi6WLHjHCa1yvBfUnHcTbrrZtXPKa29o=
        github.com/golang-jwt/jwt/v4 v4.3.0 h1:kHL1vqdqWNfATmA0FNMdmZNMyZI1U6O31X4rlIPoBog=
        github.com/golang-jwt/jwt/v4 v4.3.0/go.mod h1:/xlHOz8bRuivTWchD4jCa+NbatV+wEUSzwAxVc6locg=
        github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0/go.mod h1:E/TSTwGwJL78qG/PmXZO1EjYhfJinVAhrmmHX6Z8B9k=
        github.com/golang/geo v0.0.0-20190916061304-5b978397cfec/go.mod h1:QZ0nwyI2jOfgRAoBvP+ab5aRr7c9x7lhGEJrKvBwjWI=
        github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b/go.mod h1:SBH7ygxi8pfUlaOkMMuAQtPIUF8ecWP5IEl/CR7VP2Q=
        github.com/golang/groupcache v0.0.0-20190702054246-869f871628b6/go.mod h1:cIg4eruTrX1D+g88fzRXU5OdNfaM+9IcxsU14FzY7Hc=
        github.com/golang/groupcache v0.0.0-20191227052852-215e87163ea7/go.mod h1:cIg4eruTrX1D+g88fzRXU5OdNfaM+9IcxsU14FzY7Hc=
        github.com/golang/mock v1.1.1/go.mod h1:oTYuIxOrZwtPieC+H1uAHpcLFnEyAGVDL/k47Jfbm0A=
        github.com/golang/mock v1.2.0/go.mod h1:oTYuIxOrZwtPieC+H1uAHpcLFnEyAGVDL/k47Jfbm0A=
        github.com/golang/mock v1.3.1/go.mod h1:sBzyDLLjw3U8JLTeZvSv8jJB+tU5PVekmnlKIyFUx0Y=
        github.com/golang/protobuf v1.2.0/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
        github.com/golang/protobuf v1.3.1/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
        github.com/golang/protobuf v1.3.2/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
        github.com/golang/protobuf v1.4.0-rc.1/go.mod h1:ceaxUfeHdC40wWswd/P6IGgMaK3YpKi5j83Wpe3EHw8=
        github.com/golang/protobuf v1.4.0-rc.1.0.20200221234624-67d41d38c208/go.mod h1:xKAWHe0F5eneWXFV3EuXVDTCmh+JuBKY0li0aMyXATA=
        github.com/golang/protobuf v1.4.0-rc.2/go.mod h1:LlEzMj4AhA7rCAGe4KMBDvJI+AwstrUpVNzEA03Pprs=
        github.com/golang/protobuf v1.4.0-rc.4.0.20200313231945-b860323f09d0/go.mod h1:WU3c8KckQ9AFe+yFwt9sWVRKCVIyN9cPHBJSNnbL67w=
        github.com/golang/protobuf v1.4.0/go.mod h1:jodUvKwWbYaEsadDk5Fwe5c77LiNKVO9IDvqG2KuDX0=
        github.com/golang/protobuf v1.4.2/go.mod h1:oDoupMAO8OvCJWAcko0GGGIgR6R6ocIYbsSw735rRwI=
        github.com/golang/protobuf v1.4.3/go.mod h1:oDoupMAO8OvCJWAcko0GGGIgR6R6ocIYbsSw735rRwI=
        github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db/go.mod h1:/XxbfmMg8lxefKM7IXC3fBNl/7bRcc72aCRzEWrmP2Q=
        github.com/golang/snappy v0.0.3/go.mod h1:/XxbfmMg8lxefKM7IXC3fBNl/7bRcc72aCRzEWrmP2Q=
        github.com/golang/snappy v0.0.4 h1:yAGX7huGHXlcLOEtBnF4w7FQwA26wojNCwOYAEhLjQM=
        github.com/golang/snappy v0.0.4/go.mod h1:/XxbfmMg8lxefKM7IXC3fBNl/7bRcc72aCRzEWrmP2Q=
        github.com/golangci/lint-1 v0.0.0-20181222135242-d2cdd8c08219/go.mod h1:/X8TswGSh1pIozq4ZwCfxS0WA5JGXguxk94ar/4c87Y=
        github.com/google/btree v0.0.0-20180813153112-4030bb1f1f0c/go.mod h1:lNA+9X1NB3Zf8V7Ke586lFgjr2dZNuvo3lPJSGZ5JPQ=
        github.com/google/btree v1.0.0/go.mod h1:lNA+9X1NB3Zf8V7Ke586lFgjr2dZNuvo3lPJSGZ5JPQ=
        github.com/google/flatbuffers v1.11.0/go.mod h1:1AeVuKshWv4vARoZatz6mlQ0JxURH0Kv5+zNeJKJCa8=
        github.com/google/go-cmp v0.2.0/go.mod h1:oXzfMopK8JAjlY9xF4vHSVASa0yLyX7SntLO5aqRK0M=
        github.com/google/go-cmp v0.3.0/go.mod h1:8QqcDgzrUqlUb/G2PQTWiueGozuR1884gddMywk6iLU=
        github.com/google/go-cmp v0.3.1/go.mod h1:8QqcDgzrUqlUb/G2PQTWiueGozuR1884gddMywk6iLU=
        github.com/google/go-cmp v0.4.0/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
        github.com/google/go-cmp v0.4.1/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
        github.com/google/go-cmp v0.5.4/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
        github.com/google/gofuzz v1.1.1-0.20200604201612-c04b05f3adfa/go.mod h1:dBl0BpW6vV/+mYPU4Po3pmUjxk6FQPldtuIdl/M65Eg=
        github.com/google/martian v2.1.0+incompatible/go.mod h1:9I4somxYTbIHy5NJKHRl3wXiIaQGbYVAs8BPL6v8lEs=
        github.com/google/pprof v0.0.0-20181206194817-3ea8567a2e57/go.mod h1:zfwlbNMJ+OItoe0UupaVj+oy1omPYYDuagoSzA8v9mc=
        github.com/google/pprof v0.0.0-20190515194954-54271f7e092f/go.mod h1:zfwlbNMJ+OItoe0UupaVj+oy1omPYYDuagoSzA8v9mc=
        github.com/google/pprof v0.0.0-20191218002539-d4f498aebedc/go.mod h1:ZgVRPoUq/hfqzAqh7sHMqb3I9Rq5C59dIz2SbBwJ4eM=
        github.com/google/renameio v0.1.0/go.mod h1:KWCgfxg9yswjAJkECMjeO8J8rahYeXnNhOm40UhjYkI=
        github.com/google/uuid v1.2.0 h1:qJYtXnJRWmpe7m/3XlyhrsLrEURqHRM2kxzoxXqyUDs=
        github.com/google/uuid v1.2.0/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
        github.com/googleapis/gax-go/v2 v2.0.4/go.mod h1:0Wqv26UfaUD9n4G6kQubkQ+KchISgw+vpHVxEJEs9eg=
        github.com/googleapis/gax-go/v2 v2.0.5/go.mod h1:DWXyrwAJ9X0FpwwEdw+IPEYBICEFu5mhpdKc/us6bOk=
        github.com/gopherjs/gopherjs v0.0.0-20181017120253-0766667cb4d1/go.mod h1:wJfORRmW1u3UXTncJ5qlYoELFm8eSnnEO6hX4iZ3EWY=
        github.com/gorilla/mux v1.8.0/go.mod h1:DVbg23sWSpFRCP0SfiEN6jmj59UnW/n46BH5rLB71So=
        github.com/gorilla/websocket v1.4.2 h1:+/TMaTYc4QFitKJxsQ7Yye35DkWvkdLcvGKqM+x0Ufc=
        github.com/gorilla/websocket v1.4.2/go.mod h1:YR8l580nyteQvAITg2hZ9XVh4b55+EU/adAjf1fMHhE=
        github.com/graph-gophers/graphql-go v1.3.0/go.mod h1:9CQHMSxwO4MprSdzoIEobiHpoLtHm77vfxsvsIN5Vuc=
        github.com/hashicorp/go-bexpr v0.1.10 h1:9kuI5PFotCboP3dkDYFr/wi0gg0QVbSNz5oFRpxn4uE=
        github.com/hashicorp/go-bexpr v0.1.10/go.mod h1:oxlubA2vC/gFVfX1A6JGp7ls7uCDlfJn732ehYYg+g0=
        github.com/hashicorp/golang-lru v0.5.0/go.mod h1:/m3WP610KZHVQ1SGc6re/UDhFvYD7pJ4Ao+sR/qLZy8=
        github.com/hashicorp/golang-lru v0.5.1/go.mod h1:/m3WP610KZHVQ1SGc6re/UDhFvYD7pJ4Ao+sR/qLZy8=
        github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d h1:dg1dEPuWpEqDnvIw251EVy4zlP8gWbsGj4BsUKCRpYs=
        github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d/go.mod h1:iADmTwqILo4mZ8BN3D2Q6+9jd8WM5uGBxy+E8yxSoD4=
        github.com/holiman/bloomfilter/v2 v2.0.3 h1:73e0e/V0tCydx14a0SCYS/EWCxgwLZ18CZcZKVu0fao=
        github.com/holiman/bloomfilter/v2 v2.0.3/go.mod h1:zpoh+gs7qcpqrHr3dB55AMiJwo0iURXE7ZOP9L9hSkA=
        github.com/holiman/uint256 v1.2.0 h1:gpSYcPLWGv4sG43I2mVLiDZCNDh/EpGjSk8tmtxitHM=
        github.com/holiman/uint256 v1.2.0/go.mod h1:y4ga/t+u+Xwd7CpDgZESaRcWy0I7XMlTMA25ApIH5Jw=
        github.com/hpcloud/tail v1.0.0/go.mod h1:ab1qPbhIpdTxEkNHXyeSf5vhxWSCs/tWer42PpOxQnU=
        github.com/huin/goupnp v1.0.3 h1:N8No57ls+MnjlB+JPiCVSOyy/ot7MJTqlo7rn+NYSqQ=
        github.com/huin/goupnp v1.0.3/go.mod h1:ZxNlw5WqJj6wSsRK5+YfflQGXYfccj5VgQsMNixHM7Y=
        github.com/huin/goutil v0.0.0-20170803182201-1ca381bf3150/go.mod h1:PpLOETDnJ0o3iZrZfqZzyLl6l7F3c6L1oWn7OICBi6o=
        github.com/ianlancetaylor/demangle v0.0.0-20181102032728-5e5cf60278f6/go.mod h1:aSSvb/t6k1mPoxDqO4vJh6VOCGPwU4O0C2/Eqndh1Sc=
        github.com/inconshreveable/mousetrap v1.0.0/go.mod h1:PxqpIevigyE2G7u3NXJIT2ANytuPF1OarO4DADm73n8=
        github.com/influxdata/flux v0.65.1/go.mod h1:J754/zds0vvpfwuq7Gc2wRdVwEodfpCFM7mYlOw2LqY=
        github.com/influxdata/influxdb v1.8.3/go.mod h1:JugdFhsvvI8gadxOI6noqNeeBHvWNTbfYGtiAn+2jhI=
        github.com/influxdata/influxdb-client-go/v2 v2.4.0/go.mod h1:vLNHdxTJkIf2mSLvGrpj8TCcISApPoXkaxP8g9uRlW8=
        github.com/influxdata/influxql v1.1.1-0.20200828144457-65d3ef77d385/go.mod h1:gHp9y86a/pxhjJ+zMjNXiQAA197Xk9wLxaz+fGG+kWk=
        github.com/influxdata/line-protocol v0.0.0-20180522152040-32c6aa80de5e/go.mod h1:4kt73NQhadE3daL3WhR5EJ/J2ocX0PZzwxQ0gXJ7oFE=
        github.com/influxdata/line-protocol v0.0.0-20200327222509-2487e7298839/go.mod h1:xaLFMmpvUxqXtVkUJfg9QmT88cDaCJ3ZKgdZ78oO8Qo=
        github.com/influxdata/line-protocol v0.0.0-20210311194329-9aa0e372d097/go.mod h1:xaLFMmpvUxqXtVkUJfg9QmT88cDaCJ3ZKgdZ78oO8Qo=
        github.com/influxdata/promql/v2 v2.12.0/go.mod h1:fxOPu+DY0bqCTCECchSRtWfc+0X19ybifQhZoQNF5D8=
        github.com/influxdata/roaring v0.4.13-0.20180809181101-fc520f41fab6/go.mod h1:bSgUQ7q5ZLSO+bKBGqJiCBGAl+9DxyW63zLTujjUlOE=
        github.com/influxdata/tdigest v0.0.0-20181121200506-bf2b5ad3c0a9/go.mod h1:Js0mqiSBE6Ffsg94weZZ2c+v/ciT8QRHFOap7EKDrR0=
        github.com/influxdata/usage-client v0.0.0-20160829180054-6d3895376368/go.mod h1:Wbbw6tYNvwa5dlB6304Sd+82Z3f7PmVZHVKU637d4po=
        github.com/jackpal/go-nat-pmp v1.0.2 h1:KzKSgb7qkJvOUTqYl9/Hg/me3pWgBmERKrTGD7BdWus=
        github.com/jackpal/go-nat-pmp v1.0.2/go.mod h1:QPH045xvCAeXUZOxsnwmrtiCoxIr9eob+4orBN1SBKc=
        github.com/jedisct1/go-minisign v0.0.0-20190909160543-45766022959e/go.mod h1:G1CVv03EnqU1wYL2dFwXxW2An0az9JTl/ZsqXQeBlkU=
        github.com/jmespath/go-jmespath v0.4.0/go.mod h1:T8mJZnbsbmF+m6zOOFylbeCJqk5+pHWvzYPziyZiYoo=
        github.com/jmespath/go-jmespath/internal/testify v1.5.1/go.mod h1:L3OGu8Wl2/fWfCI6z80xFu9LTZmf1ZRjMHUOPmWr69U=
        github.com/json-iterator/go v1.1.6/go.mod h1:+SdeFBvtyEkXs7REEP0seUULqWtbJapLOCVDaaPEHmU=
        github.com/jstemmer/go-junit-report v0.0.0-20190106144839-af01ea7f8024/go.mod h1:6v2b51hI/fHJwM22ozAgKL4VKDeJcHhJFhtBdhmNjmU=
        github.com/jstemmer/go-junit-report v0.9.1/go.mod h1:Brl9GWCQeLvo8nXZwPNNblvFj/XSXhF0NWZEnDohbsk=
        github.com/jsternberg/zap-logfmt v1.0.0/go.mod h1:uvPs/4X51zdkcm5jXl5SYoN+4RK21K8mysFmDaM/h+o=
        github.com/jtolds/gls v4.20.0+incompatible/go.mod h1:QJZ7F/aHp+rZTRtaJ1ow/lLfFfVYBRgL+9YlvaHOwJU=
        github.com/julienschmidt/httprouter v1.2.0/go.mod h1:SYymIcj16QtmaHHD7aYtjjsJG7VTCxuUUipMqKk8s4w=
        github.com/jung-kurt/gofpdf v1.0.3-0.20190309125859-24315acbbda5/go.mod h1:7Id9E/uU8ce6rXgefFLlgrJj/GYY22cpxn+r32jIOes=
        github.com/jwilder/encoding v0.0.0-20170811194829-b4e1701a28ef/go.mod h1:Ct9fl0F6iIOGgxJ5npU/IUOhOhqlVrGjyIZc8/MagT0=
        github.com/karalabe/usb v0.0.2/go.mod h1:Od972xHfMJowv7NGVDiWVxk2zxnWgjLlJzE+F4F7AGU=
        github.com/kisielk/errcheck v1.2.0/go.mod h1:/BMXB+zMLi60iA8Vv6Ksmxu/1UDYcXs4uQLJ+jE2L00=
        github.com/kisielk/gotool v1.0.0/go.mod h1:XhKaO+MFFWcvkIS/tQcRk01m1F5IRFswLeQ+oQHNcck=
        github.com/klauspost/compress v1.4.0/go.mod h1:RyIbtBH6LamlWaDj8nUwkbUhJ87Yi3uG0guNDohfE1A=
        github.com/klauspost/cpuid v0.0.0-20170728055534-ae7887de9fa5/go.mod h1:Pj4uuM528wm8OyEC2QMXAi2YiTZ96dNQPGgoMS4s3ek=
        github.com/klauspost/crc32 v0.0.0-20161016154125-cb6bfca970f6/go.mod h1:+ZoRqAPRLkC4NPOvfYeR5KNOrY6TD+/sAC3HXPZgDYg=
        github.com/klauspost/pgzip v1.0.2-0.20170402124221-0bf5dcad4ada/go.mod h1:Ch1tH69qFZu15pkjo5kYi6mth2Zzwzt50oCQKQE9RUs=
        github.com/konsorten/go-windows-terminal-sequences v1.0.1/go.mod h1:T0+1ngSBFLxvqU3pZ+m/2kptfBszLMUkC4ZK/EgS/cQ=
        github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515/go.mod h1:+0opPa2QZZtGFBFZlji/RkVcI2GknAs/DXo4wKdlNEc=
        github.com/kr/pretty v0.1.0/go.mod h1:dAy3ld7l9f0ibDNOQOHHMYYIIbhfbHSm3C4ZsoJORNo=
        github.com/kr/pretty v0.2.1/go.mod h1:ipq/a2n7PKx3OHsz4KJII5eveXtPO4qwEXGdVfWzfnI=
        github.com/kr/pty v1.1.1/go.mod h1:pFQYn66WHrOpPYNljwOMqo10TkYh1fy3cYio2l3bCsQ=
        github.com/kr/text v0.1.0/go.mod h1:4Jbv+DJW3UT/LiOwJeYQe1efqtUx/iVham/4vfdArNI=
        github.com/kr/text v0.2.0/go.mod h1:eLer722TekiGuMkidMxC/pM04lWEeraHUUmBw8l2grE=
        github.com/kylelemons/godebug v0.0.0-20170224010052-a616ab194758/go.mod h1:B69LEHPfb2qLo0BaaOLcbitczOKLWTsrBG9LczfCD4k=
        github.com/kylelemons/godebug v1.1.0/go.mod h1:9/0rRGxNHcop5bhtWyNeEfOS8JIWk580+fNqagV/RAw=
        github.com/labstack/echo/v4 v4.2.1/go.mod h1:AA49e0DZ8kk5jTOOCKNuPR6oTnBS0dYiM4FW1e6jwpg=
        github.com/labstack/gommon v0.3.0/go.mod h1:MULnywXg0yavhxWKc+lOruYdAhDwPK9wf0OL7NoOu+k=
        github.com/leanovate/gopter v0.2.9/go.mod h1:U2L/78B+KVFIx2VmW6onHJQzXtFb+p5y3y2Sh+Jxxv8=
        github.com/lib/pq v1.0.0/go.mod h1:5WUZQaWbwv1U+lTReE5YruASi9Al49XbQIvNi/34Woo=
        github.com/mailru/easyjson v0.0.0-20190614124828-94de47d64c63/go.mod h1:C1wdFJiN94OJF2b5HbByQZoLdCWB1Yqtg26g4irojpc=
        github.com/mailru/easyjson v0.0.0-20190626092158-b2ccc519800e/go.mod h1:C1wdFJiN94OJF2b5HbByQZoLdCWB1Yqtg26g4irojpc=
        github.com/matryer/moq v0.0.0-20190312154309-6cfb0558e1bd/go.mod h1:9ELz6aaclSIGnZBoaSLZ3NAl1VTufbOrXBPvtcy6WiQ=
        github.com/mattn/go-colorable v0.0.9/go.mod h1:9vuHe8Xs5qXnSaW/c/ABM9alt+Vo+STaOChaDxuIBZU=
        github.com/mattn/go-colorable v0.1.2/go.mod h1:U0ppj6V5qS13XJ6of8GYAs25YV2eR4EVcfRqFIhoBtE=
        github.com/mattn/go-colorable v0.1.7/go.mod h1:u6P/XSegPjTcexA+o6vUJrdnUu04hMope9wVRipJSqc=
        github.com/mattn/go-colorable v0.1.8 h1:c1ghPdyEDarC70ftn0y+A/Ee++9zz8ljHG1b13eJ0s8=
        github.com/mattn/go-colorable v0.1.8/go.mod h1:u6P/XSegPjTcexA+o6vUJrdnUu04hMope9wVRipJSqc=
        github.com/mattn/go-isatty v0.0.4/go.mod h1:M+lRXTBqGeGNdLjl/ufCoiOlB5xdOkqRJdNxMWT7Zi4=
        github.com/mattn/go-isatty v0.0.8/go.mod h1:Iq45c/XA43vh69/j3iqttzPXn0bhXyGjM0Hdxcsrc5s=
        github.com/mattn/go-isatty v0.0.9/go.mod h1:YNRxwqDuOph6SZLI9vUUz6OYw3QyUt7WiY2yME+cCiQ=
        github.com/mattn/go-isatty v0.0.12 h1:wuysRhFDzyxgEmMf5xjvJ2M9dZoWAXNNr5LSBS7uHXY=
        github.com/mattn/go-isatty v0.0.12/go.mod h1:cbi8OIDigv2wuxKPP5vlRcQ1OAZbq2CE4Kysco4FUpU=
        github.com/mattn/go-runewidth v0.0.3/go.mod h1:LwmH8dsx7+W8Uxz3IHJYH5QSwggIsqBzpuz5H//U1FU=
        github.com/mattn/go-runewidth v0.0.9 h1:Lm995f3rfxdpd6TSmuVCHVb/QhupuXlYr8sCI/QdE+0=
        github.com/mattn/go-runewidth v0.0.9/go.mod h1:H031xJmbD/WCDINGzjvQ9THkh0rPKHF+m2gUSrubnMI=
        github.com/mattn/go-sqlite3 v1.11.0/go.mod h1:FPy6KqzDD04eiIsT53CuJW3U88zkxoIYsOqkbpncsNc=
        github.com/mattn/go-tty v0.0.0-20180907095812-13ff1204f104/go.mod h1:XPvLUNfbS4fJH25nqRHfWLMa1ONC8Amw+mIA639KxkE=
        github.com/matttproud/golang_protobuf_extensions v1.0.1/go.mod h1:D8He9yQNgCq6Z5Ld7szi9bcBfOoFv/3dc6xSMkL2PC0=
        github.com/mitchellh/mapstructure v1.4.1 h1:CpVNEelQCZBooIPDn+AR3NpivK/TIKU8bDxdASFVQag=
        github.com/mitchellh/mapstructure v1.4.1/go.mod h1:bFUtVrKA4DC2yAKiSyO/QUcy7e+RRV2QTWOzhPopBRo=
        github.com/mitchellh/pointerstructure v1.2.0 h1:O+i9nHnXS3l/9Wu7r4NrEdwA2VFTicjUEN1uBnDo34A=
        github.com/mitchellh/pointerstructure v1.2.0/go.mod h1:BRAsLI5zgXmw97Lf6s25bs8ohIXc3tViBH44KcwB2g4=
        github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd/go.mod h1:6dJC0mAP4ikYIbvyc7fijjWJddQyLn8Ig3JB5CqoB9Q=
        github.com/modern-go/reflect2 v1.0.1/go.mod h1:bx2lNnkwVCuqBIxFjflWJWanXIb3RllmbCylyMrvgv0=
        github.com/modocache/gover v0.0.0-20171022184752-b58185e213c5/go.mod h1:caMODM3PzxT8aQXRPkAt8xlV/e7d7w8GM5g0fa5F0D8=
        github.com/mschoch/smat v0.0.0-20160514031455-90eadee771ae/go.mod h1:qAyveg+e4CE+eKJXWVjKXM4ck2QobLqTDytGJbLLhJg=
        github.com/mwitkow/go-conntrack v0.0.0-20161129095857-cc309e4a2223/go.mod h1:qRWi+5nqEBWmkhHvq77mSJWrCKwh8bxhgT7d/eI7P4U=
        github.com/naoina/go-stringutil v0.1.0/go.mod h1:XJ2SJL9jCtBh+P9q5btrd/Ylo8XwT/h1USek5+NqSA0=
        github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416/go.mod h1:NBIhNtsFMo3G2szEBne+bO4gS192HuIYRqfvOWb4i1E=
        github.com/nxadm/tail v1.4.4 h1:DQuhQpB1tVlglWS2hLQ5OV6B5r8aGxSrPc5Qo6uTN78=
        github.com/nxadm/tail v1.4.4/go.mod h1:kenIhsEOeOJmVchQTgglprH7qJGnHDVpk1VPCcaMI8A=
        github.com/oklog/ulid v1.3.1/go.mod h1:CirwcVhetQ6Lv90oh/F+FBtV6XMibvdAFo93nm5qn4U=
        github.com/olekukonko/tablewriter v0.0.5 h1:P2Ga83D34wi1o9J6Wh1mRuqd4mF/x/lgBS7N7AbDhec=
        github.com/olekukonko/tablewriter v0.0.5/go.mod h1:hPp6KlRPjbx+hW8ykQs1w3UBbZlj6HuIJcUGPhkA7kY=
        github.com/onsi/ginkgo v1.6.0/go.mod h1:lLunBs/Ym6LB5Z9jYTR76FiuTmxDTDusOGeTQH+WWjE=
        github.com/onsi/ginkgo v1.10.3/go.mod h1:lLunBs/Ym6LB5Z9jYTR76FiuTmxDTDusOGeTQH+WWjE=
        github.com/onsi/ginkgo v1.12.1/go.mod h1:zj2OWP4+oCPe1qIXoGWkgMRwljMUYCdkwsT2108oapk=
        github.com/onsi/ginkgo v1.14.0 h1:2mOpI4JVVPBN+WQRa0WKH2eXR+Ey+uK4n7Zj0aYpIQA=
        github.com/onsi/ginkgo v1.14.0/go.mod h1:iSB4RoI2tjJc9BBv4NKIKWKya62Rps+oPG/Lv9klQyY=
        github.com/onsi/gomega v1.7.1/go.mod h1:XdKZgCCFLUoM/7CFJVPcG8C1xQ1AJ0vpAezJrB7JYyY=
        github.com/onsi/gomega v1.10.1 h1:o0+MgICZLuZ7xjH7Vx6zS/zcu93/BEp1VwkIW1mEXCE=
        github.com/onsi/gomega v1.10.1/go.mod h1:iN09h71vgCQne3DLsj+A5owkum+a2tYe+TOCB1ybHNo=
        github.com/opentracing/opentracing-go v1.0.2/go.mod h1:UkNAQd3GIcIGf0SeVgPpRdFStlNbqXla1AfSYxPUl2o=
        github.com/opentracing/opentracing-go v1.0.3-0.20180606204148-bd9c31933947/go.mod h1:UkNAQd3GIcIGf0SeVgPpRdFStlNbqXla1AfSYxPUl2o=
        github.com/opentracing/opentracing-go v1.1.0/go.mod h1:UkNAQd3GIcIGf0SeVgPpRdFStlNbqXla1AfSYxPUl2o=
        github.com/paulbellamy/ratecounter v0.2.0/go.mod h1:Hfx1hDpSGoqxkVVpBi/IlYD7kChlfo5C6hzIHwPqfFE=
        github.com/peterh/liner v1.0.1-0.20180619022028-8c1271fcf47f/go.mod h1:xIteQHvHuaLYG9IFj6mSxM0fCKrs34IrEQUhOYuGPHc=
        github.com/peterh/liner v1.1.1-0.20190123174540-a2c9a5303de7/go.mod h1:CRroGNssyjTd/qIG2FyxByd2S8JEAZXBl4qUrZf8GS0=
        github.com/philhofer/fwd v1.0.0/go.mod h1:gk3iGcWd9+svBvR0sR+KPcfE+RNWozjowpeBVG3ZVNU=
        github.com/pierrec/lz4 v2.0.5+incompatible/go.mod h1:pdkljMzZIN41W+lC3N2tnIh5sFi+IEE17M5jbnwPHcY=
        github.com/pkg/errors v0.8.0/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
        github.com/pkg/errors v0.8.1/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
        github.com/pkg/errors v0.9.1 h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=
        github.com/pkg/errors v0.9.1/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
        github.com/pkg/term v0.0.0-20180730021639-bffc007b7fd5/go.mod h1:eCbImbZ95eXtAUIbLAuAVnBnwf83mjf6QIVH8SHYwqQ=
        github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
        github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
        github.com/prometheus/client_golang v0.9.1/go.mod h1:7SWBe2y4D6OKWSNQJUaRYU/AaXPKyh/dDVn+NZz0KFw=
        github.com/prometheus/client_golang v1.0.0/go.mod h1:db9x61etRT2tGnBNRi70OPL5FsnadC4Ky3P0J6CfImo=
        github.com/prometheus/client_model v0.0.0-20180712105110-5c3871d89910/go.mod h1:MbSGuTsp3dbXC40dX6PRTWyKYBIrTGTE9sqQNg2J8bo=
        github.com/prometheus/client_model v0.0.0-20190129233127-fd36f4220a90/go.mod h1:xMI15A0UPsDsEKsMN9yxemIoYk6Tm2C1GtYGdfGttqA=
        github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4/go.mod h1:xMI15A0UPsDsEKsMN9yxemIoYk6Tm2C1GtYGdfGttqA=
        github.com/prometheus/common v0.0.0-20181113130724-41aa239b4cce/go.mod h1:daVV7qP5qjZbuso7PdcryaAu0sAZbrN9i7WWcTMWvro=
        github.com/prometheus/common v0.4.1/go.mod h1:TNfzLD0ON7rHzMJeJkieUDPYmFC7Snx/y86RQel1bk4=
        github.com/prometheus/common v0.6.0/go.mod h1:eBmuwkDJBwy6iBfxCBob6t6dR6ENT/y+J+Zk0j9GMYc=
        github.com/prometheus/procfs v0.0.0-20181005140218-185b4288413d/go.mod h1:c3At6R/oaqEKCNdg8wHV1ftS6bRYblBhIjjI8uT2IGk=
        github.com/prometheus/procfs v0.0.2/go.mod h1:TjEm7ze935MbeOT/UhFTIMYKhuLP4wbCsTZCD3I8kEA=
        github.com/prometheus/tsdb v0.7.1 h1:YZcsG11NqnK4czYLrWd9mpEuAJIHVQLwdrleYfszMAA=
        github.com/prometheus/tsdb v0.7.1/go.mod h1:qhTCs0VvXwvX/y3TZrWD7rabWM+ijKTux40TwIPHuXU=
        github.com/retailnext/hllpp v1.0.1-0.20180308014038-101a6d2f8b52/go.mod h1:RDpi1RftBQPUCDRw6SmxeaREsAaRKnOclghuzp/WRzc=
        github.com/rjeczalik/notify v0.9.1 h1:CLCKso/QK1snAlnhNR/CNvNiFU2saUtjV0bx3EwNeCE=
        github.com/rjeczalik/notify v0.9.1/go.mod h1:rKwnCoCGeuQnwBtTSPL9Dad03Vh2n40ePRrjvIXnJho=
        github.com/rogpeppe/go-internal v1.3.0/go.mod h1:M8bDsm7K2OlrFYOpmOWEs/qY81heoFRclV5y23lUDJ4=
        github.com/rs/cors v1.7.0 h1:+88SsELBHx5r+hZ8TCkggzSstaWNbDvThkVK8H6f9ik=
        github.com/rs/cors v1.7.0/go.mod h1:gFx+x8UowdsKA9AchylcLynDq+nNFfI8FkUZdN/jGCU=
        github.com/russross/blackfriday/v2 v2.0.1/go.mod h1:+Rmxgy9KzJVeS9/2gXHxylqXiyQDYRxCVz55jmeOWTM=
        github.com/segmentio/kafka-go v0.1.0/go.mod h1:X6itGqS9L4jDletMsxZ7Dz+JFWxM6JHfPOCvTvk+EJo=
        github.com/segmentio/kafka-go v0.2.0/go.mod h1:X6itGqS9L4jDletMsxZ7Dz+JFWxM6JHfPOCvTvk+EJo=
        github.com/sergi/go-diff v1.0.0/go.mod h1:0CfEIISq7TuYL3j771MWULgwwjU+GofnZX9QAmXWZgo=
        github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible h1:Bn1aCHHRnjv4Bl16T8rcaFjYSrGrIZvpiGO6P3Q4GpU=
        github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible/go.mod h1:5b4v6he4MtMOwMlS0TUMTu2PcXUg8+E1lC7eC3UO/RA=
        github.com/shurcooL/sanitized_anchor_name v1.0.0/go.mod h1:1NzhyTcUVG4SuEtjjoZeVRXNmyL/1OwPU0+IJeTBvfc=
        github.com/sirupsen/logrus v1.2.0/go.mod h1:LxeOpSwHxABJmUn/MG1IvRgCAasNZTLOkJPxbbu5VWo=
        github.com/smartystreets/assertions v0.0.0-20180927180507-b2de0cb4f26d/go.mod h1:OnSkiWE9lh6wB0YB77sQom3nweQdgAjqCqsofrRNTgc=
        github.com/smartystreets/goconvey v1.6.4/go.mod h1:syvi0/a8iFYH4r/RixwvyeAJjdLS9QV7WQ/tjFTllLA=
        github.com/spaolacci/murmur3 v0.0.0-20180118202830-f09979ecbc72/go.mod h1:JwIasOWyU6f++ZhiEuf87xNszmSA2myDM2Kzu9HwQUA=
        github.com/spf13/cast v1.3.0/go.mod h1:Qx5cxh0v+4UWYiBimWS+eyWzqEqokIECu5etghLkUJE=
        github.com/spf13/cobra v0.0.3/go.mod h1:1l0Ry5zgKvJasoi3XT1TypsSe7PqH0Sj9dhYf7v3XqQ=
        github.com/spf13/pflag v1.0.3/go.mod h1:DYY7MBk1bdzusC3SYhjObp+wFpr4gzcvqqNjLnInEg4=
        github.com/status-im/keycard-go v0.0.0-20190316090335-8537d3370df4 h1:Gb2Tyox57NRNuZ2d3rmvB3pcmbu7O1RS3m8WRx7ilrg=
        github.com/status-im/keycard-go v0.0.0-20190316090335-8537d3370df4/go.mod h1:RZLeN1LMWmRsyYjvAu+I6Dm9QmlDaIIt+Y+4Kd7Tp+Q=
        github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
        github.com/stretchr/objx v0.1.1/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
        github.com/stretchr/testify v1.2.0/go.mod h1:a8OnRcib4nhh0OaRAV+Yts87kKdq0PP7pXfy6kDkUVs=
        github.com/stretchr/testify v1.2.2/go.mod h1:a8OnRcib4nhh0OaRAV+Yts87kKdq0PP7pXfy6kDkUVs=
        github.com/stretchr/testify v1.3.0/go.mod h1:M5WIy9Dh21IEIfnGCwXGc5bZfKNJtfHm1UVUgZn+9EI=
        github.com/stretchr/testify v1.4.0/go.mod h1:j7eGeouHqKxXV5pUuKE4zz7dFj8WfuZ+81PSLYec5m4=
        github.com/stretchr/testify v1.5.1/go.mod h1:5W2xD1RspED5o8YsWQXVCued0rvSQ+mT+I5cxcmMvtA=
        github.com/stretchr/testify v1.7.0 h1:nwc3DEeHmmLAfoZucVR881uASk0Mfjw8xYJ99tb5CcY=
        github.com/stretchr/testify v1.7.0/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
        github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7 h1:epCh84lMvA70Z7CTTCmYQn2CKbY8j86K7/FAIr141uY=
        github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7/go.mod h1:q4W45IWZaF22tdD+VEXcAWRA037jwmWEB5VWYORlTpc=
        github.com/tinylib/msgp v1.0.2/go.mod h1:+d+yLhGm8mzTaHzB+wgMYrodPfmZrzkirds8fDWklFE=
        github.com/tklauser/go-sysconf v0.3.5 h1:uu3Xl4nkLzQfXNsWn15rPc/HQCJKObbt1dKJeWp3vU4=
        github.com/tklauser/go-sysconf v0.3.5/go.mod h1:MkWzOF4RMCshBAMXuhXJs64Rte09mITnppBXY/rYEFI=
        github.com/tklauser/numcpus v0.2.2 h1:oyhllyrScuYI6g+h/zUvNXNp1wy7x8qQy3t/piefldA=
        github.com/tklauser/numcpus v0.2.2/go.mod h1:x3qojaO3uyYt0i56EW/VUYs7uBvdl2fkfZFu0T9wgjM=
        github.com/tyler-smith/go-bip39 v1.0.1-0.20181017060643-dbb3b84ba2ef h1:wHSqTBrZW24CsNJDfeh9Ex6Pm0Rcpc7qrgKBiL44vF4=
        github.com/tyler-smith/go-bip39 v1.0.1-0.20181017060643-dbb3b84ba2ef/go.mod h1:sJ5fKU0s6JVwZjjcUEX2zFOnvq0ASQ2K9Zr6cf67kNs=
        github.com/urfave/cli/v2 v2.3.0/go.mod h1:LJmUH05zAU44vOAcrfzZQKsZbVcdbOG8rtL3/XcUArI=
        github.com/valyala/bytebufferpool v1.0.0/go.mod h1:6bBcMArwyJ5K/AmCkWv1jt77kVWyCJ6HpOuEn7z0Csc=
        github.com/valyala/fasttemplate v1.0.1/go.mod h1:UQGH1tvbgY+Nz5t2n7tXsz52dQxojPUpymEIMZ47gx8=
        github.com/valyala/fasttemplate v1.2.1/go.mod h1:KHLXt3tVN2HBp8eijSv/kGJopbvo7S+qRAEEKiv+SiQ=
        github.com/willf/bitset v1.1.3/go.mod h1:RjeCKbqT1RxIR/KWY6phxZiaY1IyutSBfGjNPySAYV4=
        github.com/xlab/treeprint v0.0.0-20180616005107-d6fb6747feb6/go.mod h1:ce1O1j6UtZfjr22oyGxGLbauSBp2YVXpARAosm7dHBg=
        github.com/yuin/goldmark v1.2.1/go.mod h1:3hX8gzYuyVAZsxl0MRgGTJEmQBFcNTphYh9decYSb74=
        github.com/yuin/goldmark v1.4.1/go.mod h1:mwnBkeHKe2W/ZEtQ+71ViKU8L12m81fl3OWwC1Zlc8k=
        go.opencensus.io v0.21.0/go.mod h1:mSImk1erAIZhrmZN+AvHh14ztQfjbGwt4TtuofqLduU=
        go.opencensus.io v0.22.0/go.mod h1:+kGneAE2xo2IficOXnaByMWTGM9T73dGwxeWcUqIpI8=
        go.opencensus.io v0.22.2/go.mod h1:yxeiOL68Rb0Xd1ddK5vPZ/oVn4vY4Ynel7k9FzqtOIw=
        go.uber.org/atomic v1.3.2/go.mod h1:gD2HeocX3+yG+ygLZcrzQJaqmWj9AIm7n08wl/qW/PE=
        go.uber.org/multierr v1.1.0/go.mod h1:wR5kodmAFQ0UK8QlbwjlSNy0Z68gJhDJUG5sjR94q/0=
        go.uber.org/zap v1.9.1/go.mod h1:vwi/ZaCAaUcBkycHslxD9B2zi4UTXhF60s6SWpuDF0Q=
        golang.org/x/crypto v0.0.0-20180904163835-0709b304e793/go.mod h1:6SG95UA2DQfeDnfUPMdvaQW0Q7yPrPDi9nlGo2tz2b4=
        golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2/go.mod h1:djNgcEr1/C05ACkg1iLfiJU5Ep61QUkGW8qpdssI0+w=
        golang.org/x/crypto v0.0.0-20190510104115-cbcb75029529/go.mod h1:yigFU9vqHzYiE8UmvKecakEJjdnWj3jj499lnFckfCI=
        golang.org/x/crypto v0.0.0-20190605123033-f99c8df09eb5/go.mod h1:yigFU9vqHzYiE8UmvKecakEJjdnWj3jj499lnFckfCI=
        golang.org/x/crypto v0.0.0-20190909091759-094676da4a83/go.mod h1:yigFU9vqHzYiE8UmvKecakEJjdnWj3jj499lnFckfCI=
        golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550/go.mod h1:yigFU9vqHzYiE8UmvKecakEJjdnWj3jj499lnFckfCI=
        golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9/go.mod h1:LzIPMQfyMNhhGPhUkYOs5KpL4U8rLKemX1yGLhDgUto=
        golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a/go.mod h1:LzIPMQfyMNhhGPhUkYOs5KpL4U8rLKemX1yGLhDgUto=
        golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad/go.mod h1:jdWPYTVW3xRLrWPugEBEK3UY2ZEsg3UU495nc5E+M+I=
        golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2/go.mod h1:T9bdIzuCu7OtxOm1hfPfRQxPLYneinmdGuTeoZ9dtd4=
        golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 h1:7I4JAnoQBe7ZtJcBaYHi5UtiO8tQHbUSXxL+pnGRANg=
        golang.org/x/crypto v0.0.0-20210921155107-089bfa567519/go.mod h1:GvvjBRRGRdwPK5ydBHafDWAxML/pGHZbMvKqRZ5+Abc=
        golang.org/x/exp v0.0.0-20180321215751-8460e604b9de/go.mod h1:CJ0aWSM057203Lf6IL+f9T1iT9GByDxfZKAQTCR3kQA=
        golang.org/x/exp v0.0.0-20180807140117-3d87b88a115f/go.mod h1:CJ0aWSM057203Lf6IL+f9T1iT9GByDxfZKAQTCR3kQA=
        golang.org/x/exp v0.0.0-20190121172915-509febef88a4/go.mod h1:CJ0aWSM057203Lf6IL+f9T1iT9GByDxfZKAQTCR3kQA=
        golang.org/x/exp v0.0.0-20190125153040-c74c464bbbf2/go.mod h1:CJ0aWSM057203Lf6IL+f9T1iT9GByDxfZKAQTCR3kQA=
        golang.org/x/exp v0.0.0-20190306152737-a1d7652674e8/go.mod h1:CJ0aWSM057203Lf6IL+f9T1iT9GByDxfZKAQTCR3kQA=
        golang.org/x/exp v0.0.0-20190510132918-efd6b22b2522/go.mod h1:ZjyILWgesfNpC6sMxTJOJm9Kp84zZh5NQWvqDGG3Qr8=
        golang.org/x/exp v0.0.0-20190829153037-c13cbed26979/go.mod h1:86+5VVa7VpoJ4kLfm080zCjGlMRFzhUhsZKEZO7MGek=
        golang.org/x/exp v0.0.0-20191030013958-a1ab85dbe136/go.mod h1:JXzH8nQsPlswgeRAPE3MuO9GYsAcnJvJ4vnMwN/5qkY=
        golang.org/x/exp v0.0.0-20191129062945-2f5052295587/go.mod h1:2RIsYlXP63K8oxa1u096TMicItID8zy7Y6sNkU49FU4=
        golang.org/x/exp v0.0.0-20191227195350-da58074b4299/go.mod h1:2RIsYlXP63K8oxa1u096TMicItID8zy7Y6sNkU49FU4=
        golang.org/x/exp v0.0.0-20220426173459-3bcf042a4bf5/go.mod h1:lgLbSvA5ygNOMpwM/9anMpWVlVJ7Z+cHWq/eFuinpGE=
        golang.org/x/image v0.0.0-20180708004352-c73c2afc3b81/go.mod h1:ux5Hcp/YLpHSI86hEcLt0YII63i6oz57MZXIpbrjZUs=
        golang.org/x/image v0.0.0-20190227222117-0694c2d4d067/go.mod h1:kZ7UVZpmo3dzQBMxlp+ypCbDeSB+sBbTgSJuh5dn5js=
        golang.org/x/image v0.0.0-20190802002840-cff245a6509b/go.mod h1:FeLwcggjj3mMvU+oOTbSwawSJRM1uh48EjtB4UJZlP0=
        golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3/go.mod h1:UVdnD1Gm6xHRNCYTkRU2/jEulfH38KcIWyp/GAMgvoE=
        golang.org/x/lint v0.0.0-20190227174305-5b3e6a55c961/go.mod h1:wehouNa3lNwaWXcvxsM5YxQ5yQlVC4a0KAMCusXpPoU=
        golang.org/x/lint v0.0.0-20190301231843-5614ed5bae6f/go.mod h1:UVdnD1Gm6xHRNCYTkRU2/jEulfH38KcIWyp/GAMgvoE=
        golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3/go.mod h1:6SW0HCj/g11FgYtHlgUYUwCkIfeOF89ocIRzGO/8vkc=
        golang.org/x/lint v0.0.0-20190409202823-959b441ac422/go.mod h1:6SW0HCj/g11FgYtHlgUYUwCkIfeOF89ocIRzGO/8vkc=
        golang.org/x/lint v0.0.0-20190909230951-414d861bb4ac/go.mod h1:6SW0HCj/g11FgYtHlgUYUwCkIfeOF89ocIRzGO/8vkc=
        golang.org/x/lint v0.0.0-20190930215403-16217165b5de/go.mod h1:6SW0HCj/g11FgYtHlgUYUwCkIfeOF89ocIRzGO/8vkc=
        golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f/go.mod h1:5qLYkcX4OjUUV8bRuDixDT3tpyyb+LUpUlRWLxfhWrs=
        golang.org/x/mobile v0.0.0-20190312151609-d3739f865fa6/go.mod h1:z+o9i4GpDbdi3rU15maQ/Ox0txvL9dWGYEHz965HBQE=
        golang.org/x/mobile v0.0.0-20190719004257-d2bd2a29d028/go.mod h1:E/iHnbuqvinMTCcRqshq8CkpyQDoeVncDDYHnLhea+o=
        golang.org/x/mod v0.0.0-20190513183733-4bf6d317e70e/go.mod h1:mXi4GBBbnImb6dmsKGUJ2LatrhH/nqhxcFungHvyanc=
        golang.org/x/mod v0.1.0/go.mod h1:0QHyrYULN0/3qlju5TqG8bIK38QM8yzMo5ekMj3DlcY=
        golang.org/x/mod v0.1.1-0.20191105210325-c90efee705ee/go.mod h1:QqPTAvyqsEbceGzBzNggFXnrqF1CaUcvgkdR5Ot7KZg=
        golang.org/x/mod v0.3.0/go.mod h1:s0Qsj1ACt9ePp/hMypM3fl4fZqREWJwdYDEqhRiZZUA=
        golang.org/x/mod v0.4.2/go.mod h1:s0Qsj1ACt9ePp/hMypM3fl4fZqREWJwdYDEqhRiZZUA=
        golang.org/x/mod v0.5.1/go.mod h1:5OXOZSfqPIIbmVBIIKWRFfZjPR0E5r58TLhUjH0a2Ro=
        golang.org/x/mod v0.6.0-dev.0.20211013180041-c96bc1413d57/go.mod h1:3p9vT2HGsQu2K1YbXdKPJLVgG5VJdoTa1poYQBtP1AY=
        golang.org/x/net v0.0.0-20180724234803-3673e40ba225/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
        golang.org/x/net v0.0.0-20180826012351-8a410e7b638d/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
        golang.org/x/net v0.0.0-20180906233101-161cd47e91fd/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
        golang.org/x/net v0.0.0-20181114220301-adae6a3d119a/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
        golang.org/x/net v0.0.0-20190108225652-1e06a53dbb7e/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
        golang.org/x/net v0.0.0-20190213061140-3a22650c66bd/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
        golang.org/x/net v0.0.0-20190311183353-d8887717615a/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=
        golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=
        golang.org/x/net v0.0.0-20190501004415-9ce7a6920f09/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=
        golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=
        golang.org/x/net v0.0.0-20190603091049-60506f45cf65/go.mod h1:HSz+uSET+XFnRR8LxR5pz3Of3rY3CfYBVs4xY44aLks=
        golang.org/x/net v0.0.0-20190613194153-d28f0bde5980/go.mod h1:z5CRVTTTmAJ677TzLLGU+0bjPO0LkuOLi4/5GtJWs/s=
        golang.org/x/net v0.0.0-20190620200207-3b0461eec859/go.mod h1:z5CRVTTTmAJ677TzLLGU+0bjPO0LkuOLi4/5GtJWs/s=
        golang.org/x/net v0.0.0-20190724013045-ca1201d0de80/go.mod h1:z5CRVTTTmAJ677TzLLGU+0bjPO0LkuOLi4/5GtJWs/s=
        golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553/go.mod h1:z5CRVTTTmAJ677TzLLGU+0bjPO0LkuOLi4/5GtJWs/s=
        golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7/go.mod h1:qpuaurCH72eLCgpAm/N6yyVIVM9cpaDIP3A8BGJEC5A=
        golang.org/x/net v0.0.0-20200813134508-3edf25e44fcc/go.mod h1:/O7V0waA8r7cgGh81Ro3o1hOxt32SMVPicZroKQ2sZA=
        golang.org/x/net v0.0.0-20200822124328-c89045814202/go.mod h1:/O7V0waA8r7cgGh81Ro3o1hOxt32SMVPicZroKQ2sZA=
        golang.org/x/net v0.0.0-20201010224723-4f7140c49acb/go.mod h1:sp8m0HH+o8qH0wwXwYZr8TS3Oi6o0r6Gce1SSxlDquU=
        golang.org/x/net v0.0.0-20201021035429-f5854403a974/go.mod h1:sp8m0HH+o8qH0wwXwYZr8TS3Oi6o0r6Gce1SSxlDquU=
        golang.org/x/net v0.0.0-20210119194325-5f4716e94777/go.mod h1:m0MpNAwzfU5UDzcl9v0D8zg8gWTRqZa9RBIspLL5mdg=
        golang.org/x/net v0.0.0-20210220033124-5f55cee0dc0d/go.mod h1:m0MpNAwzfU5UDzcl9v0D8zg8gWTRqZa9RBIspLL5mdg=
        golang.org/x/net v0.0.0-20210226172049-e18ecbb05110/go.mod h1:m0MpNAwzfU5UDzcl9v0D8zg8gWTRqZa9RBIspLL5mdg=
        golang.org/x/net v0.0.0-20210610132358-84b48f89b13b/go.mod h1:9nx3DQGgdP8bBQD5qxJ1jj9UTztislL4KSBs9R2vV5Y=
        golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d/go.mod h1:9nx3DQGgdP8bBQD5qxJ1jj9UTztislL4KSBs9R2vV5Y=
        golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f h1:OfiFi4JbukWwe3lzw+xunroH1mnC1e2Gy5cxNJApiSY=
        golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f/go.mod h1:9nx3DQGgdP8bBQD5qxJ1jj9UTztislL4KSBs9R2vV5Y=
        golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be/go.mod h1:N/0e6XlmueqKjAGxoOufVs8QHGRruUQn6yWY3a++T0U=
        golang.org/x/oauth2 v0.0.0-20190226205417-e64efc72b421/go.mod h1:gOpvHmFTYa4IltrdGE7lF6nIHvwfUNPOp7c8zoXwtLw=
        golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45/go.mod h1:gOpvHmFTYa4IltrdGE7lF6nIHvwfUNPOp7c8zoXwtLw=
        golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6/go.mod h1:gOpvHmFTYa4IltrdGE7lF6nIHvwfUNPOp7c8zoXwtLw=
        golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d/go.mod h1:gOpvHmFTYa4IltrdGE7lF6nIHvwfUNPOp7c8zoXwtLw=
        golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
        golang.org/x/sync v0.0.0-20181108010431-42b317875d0f/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
        golang.org/x/sync v0.0.0-20181221193216-37e7f081c4d4/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
        golang.org/x/sync v0.0.0-20190227155943-e225da77a7e6/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
        golang.org/x/sync v0.0.0-20190423024810-112230192c58/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
        golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
        golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
        golang.org/x/sync v0.0.0-20210220032951-036812b2e83c h1:5KslGYwFpkhGh+Q16bwMP3cOontH8FOep7tGV86Y7SQ=
        golang.org/x/sync v0.0.0-20210220032951-036812b2e83c/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
        golang.org/x/sys v0.0.0-20180830151530-49385e6e1522/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
        golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
        golang.org/x/sys v0.0.0-20180909124046-d0be0721c37e/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
        golang.org/x/sys v0.0.0-20181107165924-66b7b1311ac8/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
        golang.org/x/sys v0.0.0-20181116152217-5ac8a444bdc5/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
        golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
        golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
        golang.org/x/sys v0.0.0-20190312061237-fead79001313/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20190412213103-97732733099d/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20190502145724-3ef323f4f1fd/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20190507160741-ecd444e8653b/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20190606165138-5da285871e9c/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20190624142023-c5567b49c5d0/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20190726091711-fc99dfbffb4e/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20190904154756-749cb33beabd/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20191005200804-aed5e4c7ecf9/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20191026070338-33540a1f6037/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20191120155948-bd437916bb0e/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20191204072324-ce4227a45e2e/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20191228213918-04cbcbbfeed8/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20200107162124-548cf772de50/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20200116001909-b77594299b42/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20200223170610-d5e6a3e2c0ae/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20200519105757-fe76b779f299/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20200814200057-3d37ad5750ed/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20200826173525-f9321e4c35a6/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20201119102817-f84b799fce68/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20210316164454-77fc1eacc6aa/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20210324051608-47abb6519492/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20210420205809-ac73e9fd8988/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20210423082822-04245dca01da/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
        golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
        golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 h1:id054HUawV2/6IGm2IV8KZQjqtwAOo2CYlOToYqa0d0=
        golang.org/x/sys v0.0.0-20211019181941-9d821ace8654/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
        golang.org/x/term v0.0.0-20201117132131-f5c789dd3221/go.mod h1:Nr5EML6q2oocZ2LXRh80K7BxOlk5/8JxuGnuhpl+muw=
        golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1/go.mod h1:bj7SfCRtBDWHUb9snDiAeCFNEtKQo2Wmx5Cou7ajbmo=
        golang.org/x/text v0.3.0/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
        golang.org/x/text v0.3.1-0.20180807135948-17ff2d5776d2/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
        golang.org/x/text v0.3.2/go.mod h1:bEr9sfX3Q8Zfm5fL9x+3itogRgK3+ptLWKqgva+5dAk=
        golang.org/x/text v0.3.3/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
        golang.org/x/text v0.3.4/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
        golang.org/x/text v0.3.5/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
        golang.org/x/text v0.3.6/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
        golang.org/x/text v0.3.7 h1:olpwvP2KacW1ZWvsR7uQhoyTYvKAupfQrRGBFM352Gk=
        golang.org/x/text v0.3.7/go.mod h1:u+2+/6zg+i71rQMx5EYifcz6MCKuco9NR6JIITiCfzQ=
        golang.org/x/time v0.0.0-20181108054448-85acf8d2951c/go.mod h1:tRJNPiyCQ0inRvYxbN9jk5I+vvW/OXSQhTDSoE431IQ=
        golang.org/x/time v0.0.0-20190308202827-9d24e82272b4/go.mod h1:tRJNPiyCQ0inRvYxbN9jk5I+vvW/OXSQhTDSoE431IQ=
        golang.org/x/time v0.0.0-20201208040808-7e3f01d25324/go.mod h1:tRJNPiyCQ0inRvYxbN9jk5I+vvW/OXSQhTDSoE431IQ=
        golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba h1:O8mE0/t419eoIwhTFpKVkHiTs/Igowgfkj25AcZrtiE=
        golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba/go.mod h1:tRJNPiyCQ0inRvYxbN9jk5I+vvW/OXSQhTDSoE431IQ=
        golang.org/x/tools v0.0.0-20180525024113-a5b4c53f6e8b/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
        golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
        golang.org/x/tools v0.0.0-20181030221726-6c7e314b6563/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
        golang.org/x/tools v0.0.0-20190114222345-bf090417da8b/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
        golang.org/x/tools v0.0.0-20190206041539-40960b6deb8e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
        golang.org/x/tools v0.0.0-20190226205152-f727befe758c/go.mod h1:9Yl7xja0Znq3iFh3HoIrodX9oNMXvdceNzlUR8zjMvY=
        golang.org/x/tools v0.0.0-20190311212946-11955173bddd/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
        golang.org/x/tools v0.0.0-20190312151545-0bb0c0a6e846/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
        golang.org/x/tools v0.0.0-20190312170243-e65039ee4138/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
        golang.org/x/tools v0.0.0-20190328211700-ab21143f2384/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
        golang.org/x/tools v0.0.0-20190425150028-36563e24a262/go.mod h1:RgjU9mgBXZiqYHBnxXauZ1Gv1EHHAz9KjViQ78xBX0Q=
        golang.org/x/tools v0.0.0-20190506145303-2d16b83fe98c/go.mod h1:RgjU9mgBXZiqYHBnxXauZ1Gv1EHHAz9KjViQ78xBX0Q=
        golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135/go.mod h1:RgjU9mgBXZiqYHBnxXauZ1Gv1EHHAz9KjViQ78xBX0Q=
        golang.org/x/tools v0.0.0-20190606124116-d0a3d012864b/go.mod h1:/rFqwRUd4F7ZHNgwSSTFct+R/Kf4OFW1sUzUTQQTgfc=
        golang.org/x/tools v0.0.0-20190621195816-6e04913cbbac/go.mod h1:/rFqwRUd4F7ZHNgwSSTFct+R/Kf4OFW1sUzUTQQTgfc=
        golang.org/x/tools v0.0.0-20190628153133-6cdbf07be9d0/go.mod h1:/rFqwRUd4F7ZHNgwSSTFct+R/Kf4OFW1sUzUTQQTgfc=
        golang.org/x/tools v0.0.0-20190816200558-6889da9d5479/go.mod h1:b+2E5dAYhXwXZwtnZ6UAqBI28+e2cm9otk0dWdXHAEo=
        golang.org/x/tools v0.0.0-20190911174233-4f2ddba30aff/go.mod h1:b+2E5dAYhXwXZwtnZ6UAqBI28+e2cm9otk0dWdXHAEo=
        golang.org/x/tools v0.0.0-20191012152004-8de300cfc20a/go.mod h1:b+2E5dAYhXwXZwtnZ6UAqBI28+e2cm9otk0dWdXHAEo=
        golang.org/x/tools v0.0.0-20191113191852-77e3bb0ad9e7/go.mod h1:b+2E5dAYhXwXZwtnZ6UAqBI28+e2cm9otk0dWdXHAEo=
        golang.org/x/tools v0.0.0-20191115202509-3a792d9c32b2/go.mod h1:b+2E5dAYhXwXZwtnZ6UAqBI28+e2cm9otk0dWdXHAEo=
        golang.org/x/tools v0.0.0-20191119224855-298f0cb1881e/go.mod h1:b+2E5dAYhXwXZwtnZ6UAqBI28+e2cm9otk0dWdXHAEo=
        golang.org/x/tools v0.0.0-20191125144606-a911d9008d1f/go.mod h1:b+2E5dAYhXwXZwtnZ6UAqBI28+e2cm9otk0dWdXHAEo=
        golang.org/x/tools v0.0.0-20191126055441-b0650ceb63d9/go.mod h1:b+2E5dAYhXwXZwtnZ6UAqBI28+e2cm9otk0dWdXHAEo=
        golang.org/x/tools v0.0.0-20191216173652-a0e659d51361/go.mod h1:TB2adYChydJhpapKDTa4BR/hXlZSLoq2Wpct/0txZ28=
        golang.org/x/tools v0.0.0-20191227053925-7b8e75db28f4/go.mod h1:TB2adYChydJhpapKDTa4BR/hXlZSLoq2Wpct/0txZ28=
        golang.org/x/tools v0.0.0-20200108203644-89082a384178/go.mod h1:TB2adYChydJhpapKDTa4BR/hXlZSLoq2Wpct/0txZ28=
        golang.org/x/tools v0.1.0/go.mod h1:xkSsbof2nBLbhDlRMhhhyNLN/zl3eTqcnHD5viDpcZ0=
        golang.org/x/tools v0.1.8-0.20211029000441-d6a9af8af023/go.mod h1:nABZi5QlRsZVlzPpHl034qft6wpY4eDcsTt5AaioBiU=
        golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
        golang.org/x/xerrors v0.0.0-20191011141410-1b5146add898/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
        golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
        golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=
        golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
        gonum.org/v1/gonum v0.0.0-20180816165407-929014505bf4/go.mod h1:Y+Yx5eoAFn32cQvJDxZx5Dpnq+c3wtXuadVZAcxbbBo=
        gonum.org/v1/gonum v0.0.0-20181121035319-3f7ecaa7e8ca/go.mod h1:Y+Yx5eoAFn32cQvJDxZx5Dpnq+c3wtXuadVZAcxbbBo=
        gonum.org/v1/gonum v0.6.0/go.mod h1:9mxDZsDKxgMAuccQkewq682L+0eCu4dCN2yonUJTCLU=
        gonum.org/v1/netlib v0.0.0-20181029234149-ec6d1f5cefe6/go.mod h1:wa6Ws7BG/ESfp6dHfk7C6KdzKA7wR7u/rKwOGE66zvw=
        gonum.org/v1/netlib v0.0.0-20190313105609-8cb42192e0e0/go.mod h1:wa6Ws7BG/ESfp6dHfk7C6KdzKA7wR7u/rKwOGE66zvw=
        gonum.org/v1/plot v0.0.0-20190515093506-e2840ee46a6b/go.mod h1:Wt8AAjI+ypCyYX3nZBvf6cAIx93T+c/OS2HFAYskSZc=
        google.golang.org/api v0.4.0/go.mod h1:8k5glujaEP+g9n7WNsDg8QP6cUVNI86fCNMcbazEtwE=
        google.golang.org/api v0.7.0/go.mod h1:WtwebWUNSVBH/HAw79HIFXZNqEvBhG+Ra+ax0hx3E3M=
        google.golang.org/api v0.8.0/go.mod h1:o4eAsZoiT+ibD93RtjEohWalFOjRDx6CVaqeizhEnKg=
        google.golang.org/api v0.9.0/go.mod h1:o4eAsZoiT+ibD93RtjEohWalFOjRDx6CVaqeizhEnKg=
        google.golang.org/api v0.13.0/go.mod h1:iLdEw5Ide6rF15KTC1Kkl0iskquN2gFfn9o9XIsbkAI=
        google.golang.org/api v0.14.0/go.mod h1:iLdEw5Ide6rF15KTC1Kkl0iskquN2gFfn9o9XIsbkAI=
        google.golang.org/api v0.15.0/go.mod h1:iLdEw5Ide6rF15KTC1Kkl0iskquN2gFfn9o9XIsbkAI=
        google.golang.org/appengine v1.1.0/go.mod h1:EbEs0AVv82hx2wNQdGPgUI5lhzA/G0D9YwlJXL52JkM=
        google.golang.org/appengine v1.4.0/go.mod h1:xpcJRLb0r/rnEns0DIKYYv+WjYCduHsrkT7/EB5XEv4=
        google.golang.org/appengine v1.5.0/go.mod h1:xpcJRLb0r/rnEns0DIKYYv+WjYCduHsrkT7/EB5XEv4=
        google.golang.org/appengine v1.6.1/go.mod h1:i06prIuMbXzDqacNJfV5OdTW448YApPu5ww/cMBSeb0=
        google.golang.org/appengine v1.6.5/go.mod h1:8WjMMxjGQR8xUklV/ARdw2HLXBOI7O7uCIDZVag1xfc=
        google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8/go.mod h1:JiN7NxoALGmiZfu7CAH4rXhgtRTLTxftemlI0sWmxmc=
        google.golang.org/genproto v0.0.0-20190307195333-5fe7a883aa19/go.mod h1:VzzqZJRnGkLBvHegQrXjBqPurQTc5/KpmUdxsrq26oE=
        google.golang.org/genproto v0.0.0-20190418145605-e7d98fc518a7/go.mod h1:VzzqZJRnGkLBvHegQrXjBqPurQTc5/KpmUdxsrq26oE=
        google.golang.org/genproto v0.0.0-20190425155659-357c62f0e4bb/go.mod h1:VzzqZJRnGkLBvHegQrXjBqPurQTc5/KpmUdxsrq26oE=
        google.golang.org/genproto v0.0.0-20190502173448-54afdca5d873/go.mod h1:VzzqZJRnGkLBvHegQrXjBqPurQTc5/KpmUdxsrq26oE=
        google.golang.org/genproto v0.0.0-20190716160619-c506a9f90610/go.mod h1:DMBHOl98Agz4BDEuKkezgsaosCRResVns1a3J2ZsMNc=
        google.golang.org/genproto v0.0.0-20190801165951-fa694d86fc64/go.mod h1:DMBHOl98Agz4BDEuKkezgsaosCRResVns1a3J2ZsMNc=
        google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55/go.mod h1:DMBHOl98Agz4BDEuKkezgsaosCRResVns1a3J2ZsMNc=
        google.golang.org/genproto v0.0.0-20190911173649-1774047e7e51/go.mod h1:IbNlFCBrqXvoKpeg0TB2l7cyZUmoaFKYIwrEpbDKLA8=
        google.golang.org/genproto v0.0.0-20191108220845-16a3f7862a1a/go.mod h1:n3cpQtvxv34hfy77yVDNjmbRyujviMdxYliBSkLhpCc=
        google.golang.org/genproto v0.0.0-20191115194625-c23dd37a84c9/go.mod h1:n3cpQtvxv34hfy77yVDNjmbRyujviMdxYliBSkLhpCc=
        google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1/go.mod h1:n3cpQtvxv34hfy77yVDNjmbRyujviMdxYliBSkLhpCc=
        google.golang.org/genproto v0.0.0-20191230161307-f3c370f40bfb/go.mod h1:n3cpQtvxv34hfy77yVDNjmbRyujviMdxYliBSkLhpCc=
        google.golang.org/genproto v0.0.0-20200108215221-bd8f9a0ef82f/go.mod h1:n3cpQtvxv34hfy77yVDNjmbRyujviMdxYliBSkLhpCc=
        google.golang.org/grpc v1.19.0/go.mod h1:mqu4LbDTu4XGKhr4mRzUsmM4RtVoemTSY81AxZiDr8c=
        google.golang.org/grpc v1.20.1/go.mod h1:10oTOabMzJvdu6/UiuZezV6QK5dSlG84ov/aaiqXj38=
        google.golang.org/grpc v1.21.1/go.mod h1:oYelfM1adQP15Ek0mdvEgi9Df8B9CZIaU1084ijfRaM=
        google.golang.org/grpc v1.23.0/go.mod h1:Y5yQAOtifL1yxbo5wqy6BxZv8vAUGQwXBOALyacEbxg=
        google.golang.org/grpc v1.26.0/go.mod h1:qbnxyOmOxrQa7FizSgH+ReBfzJrCY1pSN7KXBS8abTk=
        google.golang.org/protobuf v0.0.0-20200109180630-ec00e32a8dfd/go.mod h1:DFci5gLYBciE7Vtevhsrf46CRTquxDuWsQurQQe4oz8=
        google.golang.org/protobuf v0.0.0-20200221191635-4d8936d0db64/go.mod h1:kwYJMbMJ01Woi6D6+Kah6886xMZcty6N08ah7+eCXa0=
        google.golang.org/protobuf v0.0.0-20200228230310-ab0ca4ff8a60/go.mod h1:cfTl7dwQJ+fmap5saPgwCLgHXTUD7jkjRqWcaiX5VyM=
        google.golang.org/protobuf v1.20.1-0.20200309200217-e05f789c0967/go.mod h1:A+miEFZTKqfCUM6K7xSMQL9OKL/b6hQv+e19PK+JZNE=
        google.golang.org/protobuf v1.21.0/go.mod h1:47Nbq4nVaFHyn7ilMalzfO3qCViNmqZ2kzikPIcrTAo=
        google.golang.org/protobuf v1.23.0/go.mod h1:EGpADcykh3NcUnDUJcl1+ZksZNG86OlYog2l/sGQquU=
        gopkg.in/alecthomas/kingpin.v2 v2.2.6/go.mod h1:FMv+mEhP44yOT+4EoQTLFTRgOQ1FBLkstjWtayDeSgw=
        gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
        gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
        gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c/go.mod h1:JHkPIbrfpd72SG/EVd6muEfDQjcINNoR0C8j2r3qZ4Q=
        gopkg.in/errgo.v2 v2.1.0/go.mod h1:hNsd1EY+bozCKY1Ytp96fpM3vjJbqLJn88ws8XvfDNI=
        gopkg.in/fsnotify.v1 v1.4.7/go.mod h1:Tz8NjZHkW78fSQdbUxIjBTcgA1z1m8ZHf0WmKUhAMys=
        gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce h1:+JknDZhAj8YMt7GC73Ei8pv4MzjDUNPHgQWJdtMAaDU=
        gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce/go.mod h1:5AcXVHNjg+BDxry382+8OKon8SEWiKktQR07RKPsv1c=
        gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=
        gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7/go.mod h1:dt/ZhP58zS4L8KSrWDmTeBkI65Dw0HsyUHuEVlX15mw=
        gopkg.in/urfave/cli.v1 v1.20.0 h1:NdAVW6RYxDif9DhDHaAortIu956m2c0v+09AZBPTbE0=
        gopkg.in/urfave/cli.v1 v1.20.0/go.mod h1:vuBzUtMdQeixQj8LVd+/98pzhxNGQoyuPBlsXHOQNO0=
        gopkg.in/yaml.v2 v2.2.1/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
        gopkg.in/yaml.v2 v2.2.2/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
        gopkg.in/yaml.v2 v2.2.3/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
        gopkg.in/yaml.v2 v2.2.4/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
        gopkg.in/yaml.v2 v2.2.8/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
        gopkg.in/yaml.v2 v2.3.0/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
        gopkg.in/yaml.v2 v2.4.0 h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=
        gopkg.in/yaml.v2 v2.4.0/go.mod h1:RDklbk79AGWmwhnvt/jBztapEOGDOx6ZbXqjP6csGnQ=
        gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
        gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b h1:h8qDotaEPuJATrMmW04NCwg7v22aHH28wwpauUhK9Oo=
        gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
        honnef.co/go/tools v0.0.0-20190102054323-c2f93a96b099/go.mod h1:rf3lG4BRIbNafJWhAfAdb/ePZxsR/4RtNHQocxwk9r4=
        honnef.co/go/tools v0.0.0-20190106161140-3f1c8253044a/go.mod h1:rf3lG4BRIbNafJWhAfAdb/ePZxsR/4RtNHQocxwk9r4=
        honnef.co/go/tools v0.0.0-20190418001031-e561f6794a2a/go.mod h1:rf3lG4BRIbNafJWhAfAdb/ePZxsR/4RtNHQocxwk9r4=
        honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc/go.mod h1:rf3lG4BRIbNafJWhAfAdb/ePZxsR/4RtNHQocxwk9r4=
        honnef.co/go/tools v0.0.1-2019.2.3/go.mod h1:a3bituU0lyd329TUQxRnasdCoJDkEUEAqEt0JzvZhAg=
        honnef.co/go/tools v0.1.3/go.mod h1:NgwopIslSNH47DimFoV78dnkksY2EFtX0ajyb3K/las=
        rsc.io/binaryregexp v0.2.0/go.mod h1:qTv7/COck+e2FymRvadv62gMdZztPaShugOCi3I+8D8=
        rsc.io/pdf v0.1.1/go.mod h1:n8OzWcQ6Sp37PL01nO98y4iUCRdTGarVfzxY20ICaU4=
        
        =======================
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

