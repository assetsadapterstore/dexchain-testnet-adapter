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
	"github.com/blocktree/openwallet/v2/openwallet"
	"testing"
)

func TestWalletManager_GetTokenBalanceByAddress(t *testing.T) {
	wm := testNewWalletManager()

	contract := openwallet.SmartContract{
		Address:  "0xd6fddf404a5cdadef3b857d8b3118a67165315fc",
		Symbol:   "DEX",
		Name:     "WWW",
		Token:    "WWW",
		Decimals: 18,
	}

	tokenBalances, err := wm.ContractDecoder.GetTokenBalanceByAddress(contract, "0xDBe88debb5EEa51e963B1aE7a241374BE972E104")
	if err != nil {
		t.Errorf("GetTokenBalanceByAddress unexpected error: %v", err)
		return
	}
	for _, b := range tokenBalances {
		t.Logf("token balance: %+v", b.Balance)
	}
}

func TestWalletManager_GetTokenBalanceByAddressMulti(t *testing.T) {
	wm := testNewWalletManager()

	contract := openwallet.SmartContract{
		Address:  "0xd6fddf404a5cdadef3b857d8b3118a67165315fc",
		Symbol:   "DEX",
		Name:     "WWW",
		Token:    "WWW",
		Decimals: 18,
	}
	for i := 0; i < 10000; i++ {
		tokenBalances, err := wm.ContractDecoder.GetTokenBalanceByAddress(contract, "0xDBe88debb5EEa51e963B1aE7a241374BE972E104")
		if err != nil {
			t.Errorf("GetTokenBalanceByAddress unexpected error: %v", err)
			return
		}
		for _, b := range tokenBalances {
			t.Logf("token balance: %+v", b.Balance)
		}
	}
}
