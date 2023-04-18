package etherscanClient

import (
	"fmt"
	"github.com/nanmu42/etherscan-api"
	"testing"
)

func TestGetBalance(t *testing.T) {
	t.Run("GetBalance-Correct", func(t *testing.T) {
		t.Helper()
		client := new(ScanManage)
		client.NewClient(etherscan.Goerli, "X8TFJ8QNJXSEQBRVXKNS2MTF2U4QJZAVH9")
		balance := client.getAccountBalance("0x1D9b2905b2EC7d9F64022c6e698c0d622A35225c")
		fmt.Println("balance is ", balance)
	})
}

func TestGetAccountsBalance(t *testing.T) {
	t.Run("GetAccountsBalance-Correct", func(t *testing.T) {
		t.Helper()
		client := new(ScanManage)
		client.NewClient(etherscan.Goerli, "X8TFJ8QNJXSEQBRVXKNS2MTF2U4QJZAVH9")
		client.NewWallet([]string{
			"0x1D9b2905b2EC7d9F64022c6e698c0d622A35225c",
			"0xF0401c429B40E16165f2D572f785dB7E5a12d866",
			"0xB1D8F3419b545cd86b0eC19B30c5ea96A79acD8F",
		})

		//验证：钱包地址余额显示正确
		for k, v := range client.wallet {
			fmt.Println("address is ", k, "balance is ", v)
		}
	})
}

func TestUpdateAccountsBalance(t *testing.T) {
	t.Run("UpdateAccountsBalance-Correct", func(t *testing.T) {
		t.Helper()
		client := new(ScanManage)
		client.NewClient(etherscan.Goerli, "X8TFJ8QNJXSEQBRVXKNS2MTF2U4QJZAVH9")
		client.NewWallet([]string{
			"0x1D9b2905b2EC7d9F64022c6e698c0d622A35225c",
			"0xF0401c429B40E16165f2D572f785dB7E5a12d866",
			"0xB1D8F3419b545cd86b0eC19B30c5ea96A79acD8F",
		})

		//验证：钱包地址余额显示正确
		for k, v := range client.wallet {
			fmt.Println("address is ", k, "balance is ", v)
		}

		//更新钱包余额
		client.updateAccountsBalance()

		//验证：钱包地址余额显示正确
		for k, v := range client.wallet {
			fmt.Println("address is ", k, "balance is ", v)
		}
	})
}
