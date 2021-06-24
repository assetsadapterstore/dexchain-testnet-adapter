module github.com/blocktree/dexchain-test-adapter

go 1.12

require (
	github.com/astaxie/beego v1.12.1
	github.com/blocktree/openwallet/v2 v2.0.7
	github.com/blocktree/quorum-adapter v1.4.2-0.20210623101640-c4cd54812ec3
	github.com/ethereum/go-ethereum v1.9.9
)

//replace github.com/blocktree/quorum-adapter => ../quorum-adapter

//replace github.com/blocktree/openwallet/v2 => ../../openwallet
