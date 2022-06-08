package business

import "testing"

// ChainTestingCase
//  @Description: 区块链业务功能测试通用接口
//   接口定义必须要测试的功能或模块。作为测试标准和今后通用化规划
type ChainTestingCase interface {
	IsConnected(t *testing.T) bool            //连接网络是否成功
	CallContract(t *testing.T)                //调用智能合约调用
	CheckChainNum(t *testing.T)               //确认链上最新的区块高度
	Transact(t *testing.T)                    //交易
	CreateAccount(newCount int, t *testing.T) //创建账号
	PressureAttack(t *testing.T)              //压力测试
}
