/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package dexchain

import (
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/v2/common"
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/quorum-adapter/quorum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcom "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"path/filepath"
	"strings"
	"testing"
)

var (
	tw *WalletManager
)

func init() {

	tw = testNewWalletManager()
}

func testNewWalletManager() *WalletManager {
	wm := NewWalletManager()

	//读取配置
	absFile := filepath.Join("conf", "TDEX.ini")
	//log.Debug("absFile:", absFile)
	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		panic(err)
	}
	wm.LoadAssetsConfig(c)
	wm.WalletClient.Debug = true
	return wm
}

func TestFixGasLimit(t *testing.T) {
	fixGasLimitStr := "sfsd"
	fixGasLimit := new(big.Int)
	fixGasLimit.SetString(fixGasLimitStr, 10)
	fmt.Printf("fixGasLimit: %d\n", fixGasLimit.Int64())
}

func TestWalletManager_GetAddrBalance(t *testing.T) {
	wm := testNewWalletManager()
	balance, err := wm.GetAddrBalance("0xdbe88debb5eea51e963b1ae7a241374be972e104", "latest")
	if err != nil {
		t.Errorf("GetAddrBalance2 error: %v", err)
		return
	}
	ethB := common.BigIntToDecimals(balance, wm.Decimal())
	log.Infof("ethB: %v", ethB)
}

func TestWalletManager_SetNetworkChainID(t *testing.T) {
	wm := testNewWalletManager()
	id, err := wm.SetNetworkChainID()
	if err != nil {
		t.Errorf("SetNetworkChainID error: %v", err)
		return
	}
	log.Infof("chainID: %d", id)
}

func TestWalletManager_GetTransactionFeeEstimated(t *testing.T) {
	wm := testNewWalletManager()

	from := "0x5cD2554daaB3919Cc1C0b5D66cCE16b19f27B438"
	contractAddress := "0xACa6A122656929b7a9BE61978acc907aF346bee3"
	data, _ := hex.DecodeString("2f2ff15d9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a600000000000000000000000095f28edc0ebbfc7d893dbc4e8d99da10a2291f1e")
	//value, _ := common.StringValueToBigInt("0x98a7d9b8314c0000", 16)


	log.Infof("data: %s", hex.EncodeToString(data))

	txFee, err := wm.GetTransactionFeeEstimated(
		from,
		contractAddress,
		nil,
		data)
	if err != nil {
		t.Errorf("GetTransactionFeeEstimated error: %v", err)
		return
	}
	log.Infof("txfee: %v", txFee)
}

func TestWalletManager_GetGasPrice(t *testing.T) {
	wm := testNewWalletManager()
	price, err := wm.GetGasPrice()
	if err != nil {
		t.Errorf("GetGasPrice error: %v", err)
		return
	}
	log.Infof("price: %v", price.String())
}

func TestWalletManager_GetTransactionCount(t *testing.T) {
	wm := testNewWalletManager()
	count, err := wm.GetTransactionCount("0xdbe88debb5eea51e963b1ae7a241374be972e104")
	if err != nil {
		t.Errorf("GetTransactionCount error: %v", err)
		return
	}
	log.Infof("count: %v", count)
}

func TestWalletManager_IsContract(t *testing.T) {
	wm := testNewWalletManager()
	a, err := wm.IsContract("0x3440f720862aa7dfd4f86ecc78542b3ded900c02")
	log.Infof("IsContract: %v", a)
	if err != nil {
		t.Errorf("IsContract error: %v", err)
		return
	}

	c, _ := wm.IsContract("0x627b11ead4eb39ebe61a70ab3d6fe145e5d06ab6")
	log.Infof("IsContract: %v", c)

}

func TestWalletManager_DecodeReceiptLogResult(t *testing.T) {
	wm := testNewWalletManager()
	abiJSON := `
[{"inputs":[{"internalType":"contract KeyValueStorage","name":"storage_","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"implementation","type":"address"}],"name":"Upgraded","type":"event"},{"payable":true,"stateMutability":"payable","type":"fallback"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getOwner","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"implementation","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"impl","type":"address"}],"name":"upgradeTo","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`
	logJSON := `
			{
                "logIndex": "0x0",
                "transactionIndex": "0x0",
                "transactionHash": "0x6a949727089705103e873c5dc9ebfaac79deb5fe5df0b9f02672988336130af9",
                "blockHash": "0xd80805f3b261f8dc9fd95a60030615c20ff1ca29ecb34101faf91512aedd9f2c",
                "blockNumber": "0x4b",
                "address": "0xf8afe0a06e27ddbd5ec8adbbd5cee5220c3d4d85",
                "data": "0x",
                "topics": [
                    "0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b",
                    "0x00000000000000000000000044f64ef4bc4952b133a9c4b07157770f048eebe9"
                ],
                "type": "mined"
            }
`
	var logObj types.Log
	err := logObj.UnmarshalJSON([]byte(logJSON))
	if err != nil {
		t.Errorf("UnmarshalJSON error: %v", err)
		return
	}

	abiInstance, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		t.Errorf("abi.JSON error: %v", err)
		return
	}

	rMap, name, rJSON, err := wm.DecodeReceiptLogResult(abiInstance, logObj)
	if err != nil {
		t.Errorf("DecodeReceiptLogResult error: %v", err)
		return
	}
	log.Infof("rMap: %+v", rMap)
	log.Infof("name: %+v", name)
	log.Infof("rJSON: %s", rJSON)
}

func TestWalletManager_GetAddrBalanceMulti(t *testing.T) {
	wm := testNewWalletManager()
	for i := 0; i < 10000; i++ {
		balance, err := wm.GetAddrBalance("0xDBe88debb5EEa51e963B1aE7a241374BE972E104", "0x6cb53")
		if err != nil {
			t.Errorf("GetAddrBalance2 error: %v", err)
			return
		}
		ethB := common.BigIntToDecimals(balance, wm.Decimal())
		log.Infof("ethB: %v", ethB)
	}
}


func TestWalletManager_GetTransactionFeeEstimatedMulti(t *testing.T) {
	wm := testNewWalletManager()
	data, _ := hex.DecodeString("f305d719000000000000000000000000c1b3702b8b2778cd8315efd40ca6c1addf094c8000000000000000000000000000000000000000000000000000000000832156000000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000b000000000000000000000000373944b86bc07887f2cdcc5ef5e055ee33ac2d3c00000000000000000000000000000000000000000000000000000000609e42a6")
	//log.Infof("data: %s", hex.EncodeToString(data))
	value, _ := common.StringValueToBigInt("0x98a7d9b8314c0000", 16)
	for i := 0; i < 10000; i++ {
		txFee, err := wm.GetTransactionFeeEstimated(
			"0x373944b86bc07887f2cdcc5ef5e055ee33ac2d3c",
			"0x59562acce01eee481a5fdfbd45fc50adfd4538d7",
			value,
			data)
		if err != nil {
			t.Errorf("GetTransactionFeeEstimated error: %v", err)
			return
		}
		log.Infof("txfee: %v", txFee)
	}
}


func TestWalletManager_EthCall(t *testing.T) {
	wm := testNewWalletManager()

	from := "0x373944b86bc07887f2cdcc5ef5e055ee33ac2d3c"
	contractAddress := "0x59562acce01eee481a5fdfbd45fc50adfd4538d7"
	data, _ := hex.DecodeString("0xf305d719000000000000000000000000c1b3702b8b2778cd8315efd40ca6c1addf094c8000000000000000000000000000000000000000000000000000000019de1377000000000000000000000000000000000000000000000000000000000000000457000000000000000000000000000000000000000000000000000000000000000b000000000000000000000000373944b86bc07887f2cdcc5ef5e055ee33ac2d3c0000000000000000000000000000000000000000000000000000000060a32954")
	value, _ := common.StringValueToBigInt("0x98a7d9b8314c0000", 16)

	callMsg := quorum.CallMsg{
		From: ethcom.HexToAddress(from),
		To:   ethcom.HexToAddress(contractAddress),
		Data: data,
		Value: value,
	}

	result, err := wm.EthCall(callMsg, "latest")
	if err != nil {
		t.Errorf("EthCall error: %v", err)
		return
	}

	log.Infof("result: %s", result)
}
