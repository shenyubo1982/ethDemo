package chainClient

//  meta链测试脚本
import (
	"ethDemo/util"
	"testing"
)

// 元生链 实验室 测试
type EthKovanTest struct {
	// test to eth chain in kovan network
	ConfigFile          string
	blockInfoRequestUrl string
}

func NewEthKovanTest() *EthKovanTest {
	instance := new(EthKovanTest)
	//添加测试需要的常量
	instance.ConfigFile = "kovanTest.yaml"
	instance.blockInfoRequestUrl = ""
	return instance
}

func (k EthKovanTest) IsConnected(t *testing.T) bool {
	myChainConfig := util.NewChainTestYaml(k.ConfigFile)
	myChainClient := Launch(myChainConfig.YamlContent)
	if myChainClient == nil {
		t.Errorf("Can't get Client")
		return false
	}
	return true
}

func (k EthKovanTest) CallContract(t *testing.T) {
	//TODO implement me
	panic("implement me")
}

func (k EthKovanTest) CheckChainNum(t *testing.T) {
	//TODO implement me
	panic("implement me")
}

func (k EthKovanTest) Transact(t *testing.T) {
	//TODO implement me
	panic("implement me")
}

func (k EthKovanTest) CreateAccount(newCount int, t *testing.T) {
	//TODO implement me
	panic("implement me")
}

func (k EthKovanTest) PressureAttack(t *testing.T) {
	//TODO implement me
	panic("implement me")
}

//
//  TestLaunch
//  @Description: 区块链网络连接测试
//  @param t
//
func TestLaunchKovan(t *testing.T) {
	t.Run("LaunchKovan", func(t *testing.T) {
		t.Helper()
		var kovanChainOptionTest ChainTestingCase
		kovanChainOptionTest = NewEthKovanTest()
		kovanChainOptionTest.IsConnected(t)
	})
}
