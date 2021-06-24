package openwtester

import (
	"github.com/blocktree/dexchain-test-adapter/dexchain"
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openw"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	openw.RegAssets(dexchain.Symbol, dexchain.NewWalletManager())
}
