package cmds

import (
	"ethDemo/cmds/ac"
	"ethDemo/cmds/km"
	"testing"
)

// TestKeyManager 地址管理功能测试
func TestKeyManager(t *testing.T) {
	t.Run("Launch", func(t *testing.T) {
		t.Helper()
		km.Run()
	})
}

// TestAccountManager 账户批量创建功能测试
func TestAccountManager(t *testing.T) {
	t.Run("Launch", func(t *testing.T) {
		t.Helper()
		ac.Run()
	})
}
