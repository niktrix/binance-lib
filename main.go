package main

import (
	"fmt"
	"log"

	sdk "github.com/binance-chain/go-sdk/client"
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
	"github.com/binance-chain/go-sdk/types/msg"
)

var (
	toaddress   = "tbnb15m0ddk8jnn7me2p8qcp244c20xmry0dmur64xu"
	mnemonic    = "add your mnemonic" // https://www.binance.com/en/support/articles/360023912272
	amount      = int64(1)            // Amount to transfer
	chain       = "testnet"           // change to mainnet while doin mainnet tx
	rpcURL      = "testnet-dex.binance.org"
	networkType = types.TestNetwork // types.ProdNetwork
)

func main() {
	var transfers []msg.Transfer
	t := msg.Transfer{}
	addressBytes, err := types.GetFromBech32(toaddress, "tbnb")
	if err != nil {
		log.Println("Wrong  address", err)
	}
	t.ToAddr = types.AccAddress(addressBytes)
	t.Coins = append(t.Coins, types.Coin{Denom: "BNB", Amount: amount})

	km, _ := keys.NewMnemonicKeyManager(mnemonic)
	client, err := sdk.NewDexClient(rpcURL, networkType, km)
	if err != nil {
		log.Println("Error getting client", err)
	}

	result, err := client.SendToken(append(transfers, t), true)
	if err != nil {
		log.Println("Error sending Tx", err)
	}
	fmt.Println("From Address", km.GetAddr())
	fmt.Println(result)
}
