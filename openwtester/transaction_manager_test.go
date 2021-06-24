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

package openwtester

import (
	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/v2/openw"
	"path/filepath"
	"testing"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
)

func TestWalletManager_GetAssetsAccountBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WJtr6eEagNejQ1iDtrNTxf1nhvzm3Hm2zq"
	//accountID := "AX9pgUb3eSyhF5dHjT9CGsmyjFeV8DdWW9FcSbvVwk3y"
	accountID := "Abact92ucAteDmtinTd7Upv6x9pPK4WSp5irN13eiyTd"

	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func TestWalletManager_GetAssetsAccountTokenBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WBGYxZ6yEX582Mx8mGvygXevdLVc7NQnLM"
	accountID := "9EfTQiMEaKSMd1CjxMXRMMxukrwckxdBZpiEkS2B3avD"
	//accountID := "CxE3ds4JdTHXV1f2xSsE6qahgfReKR9iPmFPcBmTfaKP"

	contract := openwallet.SmartContract{
		Address:  "0x4092678e4e78230f46a1534c0fbc8fa39780892b",
		Symbol:   "TDEX",
		Name:     "OCoin",
		Token:    "OCN",
		Decimals: 18,
	}

	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance.Balance)
}

func TestWalletManager_GetEstimateFeeRate(t *testing.T) {
	tm := testInitWalletManager()
	coin := openwallet.Coin{
		Symbol: "TDEX",
	}
	feeRate, unit, err := tm.GetEstimateFeeRate(coin)
	if err != nil {
		log.Error("GetEstimateFeeRate failed, unexpected error:", err)
		return
	}
	log.Std.Info("feeRate: %s %s/%s", feeRate, coin.Symbol, unit)
}

func TestGetAddressBalance(t *testing.T) {
	symbol := "VSYS"
	assetsMgr, err := openw.GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}
	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)
	bs := assetsMgr.GetBlockScanner()

	addrs := []string{
		"AR5D3fGVWDz32wWCnVbwstsMW8fKtWdzNFT",
		"AR9qbgbsmLh3ADSU9ngR22J2HpD5D9ncTCg",
		"ARAA8AnUYa4kWwWkiZTTyztG5C6S9MFTx11",
		"ARCUYWyLvGDTrhZ6K9jjMh9B5iRVEf3vRzs",
		"ARGehumz77nGcfkQrPjK4WUyNevvU9NCNqQ",
		"ARJdaB9Fo6Sk2nxBrQP2p4woWotPxjaebCv",
	}

	balances, err := bs.GetBalanceByAddress(addrs...)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	for _, b := range balances {
		log.Infof("balance[%s] = %s", b.Address, b.Balance)
		log.Infof("UnconfirmBalance[%s] = %s", b.Address, b.UnconfirmBalance)
		log.Infof("ConfirmBalance[%s] = %s", b.Address, b.ConfirmBalance)
	}
}
