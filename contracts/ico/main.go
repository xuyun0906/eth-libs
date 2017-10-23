package main

import (
	"flag"
	"reflect"
	. "github.com/dylenfu/eth-libs/contracts/ico/contract"
	"github.com/dylenfu/eth-libs/types"
	. "github.com/dylenfu/eth-libs/params"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

var (
	tokenA = TokenA.Token.(*IcoImpl)
	tokenB = TokenB.Token.(*IcoImpl)
	handle = &Handle{}
	fn = flag.String("fn", "deposit", "chose function")
)

type Handle struct{}

func main() {
	flag.Parse()
	reflect.ValueOf(handle).MethodByName(*fn).Call([]reflect.Value{})
}

func (h *Handle) BalanceOf() {
	var result types.HexNumber

	account := types.Str2Address(Account1)
	err := tokenA.BalanceOf.Call(&result, "latest", account)
	if err != nil {
		panic(err)
	}

	log.Println(result)
}

/*
"0x48ff2269e58a373120FFdBBdEE3FBceA854AC30A"
"0xb5fab0b11776aad5ce60588c16bd59dcfd61a1c2"
*/
func (h *Handle) Deposit() {
	var result types.HexNumber
	value := big.NewInt(1000000000000)

	account := types.Str2Address("0xb5fab0b11776aad5ce60588c16bd59dcfd61a1c2")
	err := tokenB.Deposit.SendTransaction(&result, account, value)
	if err != nil {
		panic(err)
	}

	log.Println(result)
}

func (h *Handle) Approve() {
	var result bool

	// 合约里面msg.sender是看谁对transaction签名过即为sender
	to := common.HexToAddress(Account1)
	amount := big.NewInt(100000000)
	if err := tokenA.Transfer.SignAndSendTransaction(Account2, &result, to, amount); err != nil {
		panic(err)
	}

	log.Println(result)
}

func (h *Handle) Allowance() {
	var result types.HexNumber

	owner := types.Str2Address(Account1)
	spender := types.Str2Address(Account2)
	err := tokenA.Allowance.Call(&result, "latest", owner, spender)
	if err != nil {
		panic(err)
	}

	log.Println(result)
}

func (h *Handle) Transfer() {
	var result string

	// 合约里面msg.sender是看谁对transaction签名过即为sender
	to := common.HexToAddress(Account1)
	amount := big.NewInt(100000000)
	if err := tokenA.Transfer.SignAndSendTransaction(Account2, &result, to, amount); err != nil {
		panic(err)
	}

	log.Println(result)
}