package main

import (
	"github.com/dylenfu/eth-libs/types"
	"log"
)

const logo = `
	THATH_LRC
`

const (
	AbiFilePath = "github.com/dylenfu/eth-libs/contracts/lrc/abi.txt"
)

// [{"constant":true,"inputs":[{"name":"orderHash","type":"bytes32"}],"name":"getOrderCancelled","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"bytes32"}],"name":"filled","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"bytes32"}],"name":"cancelled","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"ringIndex","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"addresses","type":"address[3]"},{"name":"orderValues","type":"uint256[7]"},{"name":"buyNoMoreThanAmountB","type":"bool"},{"name":"marginSplitPercentage","type":"uint8"},{"name":"v","type":"uint8"},{"name":"r","type":"bytes32"},{"name":"s","type":"bytes32"}],"name":"cancelOrder","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"RATE_RATIO_SCALE","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"ringminer","type":"address"},{"name":"ringSize","type":"uint256"},{"name":"vList","type":"uint8[]"},{"name":"rList","type":"bytes32[]"},{"name":"sList","type":"bytes32[]"}],"name":"calculateRinghash","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"lrcTokenAddress","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"tokenRegistryAddress","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"addressList","type":"address[2][]"},{"name":"uintArgsList","type":"uint256[7][]"},{"name":"uint8ArgsList","type":"uint8[2][]"},{"name":"buyNoMoreThanAmountBList","type":"bool[]"},{"name":"vList","type":"uint8[]"},{"name":"rList","type":"bytes32[]"},{"name":"sList","type":"bytes32[]"},{"name":"ringminer","type":"address"},{"name":"feeRecepient","type":"address"},{"name":"throwIfLRCIsInsuffcient","type":"bool"}],"name":"submitRing","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"delegateAddress","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"orderHash","type":"bytes32"}],"name":"getOrderFilled","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"maxRingSize","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"ringhashRegistryAddress","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"cutoffs","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"rateRatioCVSThreshold","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"_lrcTokenAddress","type":"address"},{"name":"_tokenRegistryAddress","type":"address"},{"name":"_ringhashRegistryAddress","type":"address"},{"name":"_delegateAddress","type":"address"},{"name":"_maxRingSize","type":"uint256"},{"name":"_rateRatioCVSThreshold","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"payable":true,"stateMutability":"payable","type":"fallback"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_ringIndex","type":"uint256"},{"indexed":false,"name":"_time","type":"uint256"},{"indexed":false,"name":"_blocknumber","type":"uint256"},{"indexed":true,"name":"_ringhash","type":"bytes32"},{"indexed":true,"name":"_miner","type":"address"},{"indexed":true,"name":"_feeRecepient","type":"address"},{"indexed":false,"name":"_ringhashFound","type":"bool"}],"name":"RingMined","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_ringIndex","type":"uint256"},{"indexed":false,"name":"_time","type":"uint256"},{"indexed":false,"name":"_blocknumber","type":"uint256"},{"indexed":true,"name":"_ringhash","type":"bytes32"},{"indexed":false,"name":"_prevOrderHash","type":"bytes32"},{"indexed":true,"name":"_orderHash","type":"bytes32"},{"indexed":false,"name":"_nextOrderHash","type":"bytes32"},{"indexed":false,"name":"_amountS","type":"uint256"},{"indexed":false,"name":"_amountB","type":"uint256"},{"indexed":false,"name":"_lrcReward","type":"uint256"},{"indexed":false,"name":"_lrcFee","type":"uint256"}],"name":"OrderFilled","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_time","type":"uint256"},{"indexed":false,"name":"_blocknumber","type":"uint256"},{"indexed":true,"name":"_orderHash","type":"bytes32"},{"indexed":false,"name":"_amountCancelled","type":"uint256"}],"name":"OrderCancelled","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_time","type":"uint256"},{"indexed":false,"name":"_blocknumber","type":"uint256"},{"indexed":true,"name":"_address","type":"address"},{"indexed":false,"name":"_cutoff","type":"uint256"}],"name":"CutoffTimestampChanged","type":"event"}]

func main() {
	tabi := types.NewAbi(AbiFilePath)
	for k, v := range tabi.Events {
		log.Println(k, v.Id().String())
	}
}
