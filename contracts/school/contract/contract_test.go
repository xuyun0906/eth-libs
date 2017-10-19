package contract_test

import (
	"testing"
	"github.com/dylenfu/eth-libs/types"
	"github.com/dylenfu/eth-libs/common"
	. "github.com/dylenfu/eth-libs/contracts/school/contract"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

//
func TestUnpackBabyEvent(t *testing.T) {
	tabi := types.NewAbi(AbiFilePath)
	data := "0x0000000000000000000000000c0b638ffccb4bdc4c0d0d5fef062fc512c9251100000000000000000000000096124db0972e3522a9b3910578b3f2e1a50159c700000000000000000000000086324df0972e3522a9b3910578b3f2e1a50132d5"
	evt := &BabyEvent{}
	err := common.UnpackEvent(tabi.Events["BabyEvent"].Inputs, evt, hexutil.MustDecode(data), []string{})
	if err != nil {
		t.Error(err)
	}

	for _, v := range evt.Addresses {
		t.Log(v.Hex())
	}
}

func TestUnpackChildEvent(t *testing.T) {
	tabi := types.NewAbi(AbiFilePath)
	data := "0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000c0b638ffccb4bdc4c0d0d5fef062fc512c9251100000000000000000000000096124db0972e3522a9b3910578b3f2e1a50159c700000000000000000000000086324df0972e3522a9b3910578b3f2e1a50132d5"
	evt := &ChildEvent{}
	err := common.UnpackEvent(tabi.Events["ChildEvent"].Inputs, evt, hexutil.MustDecode(data), []string{})
	if err != nil {
		t.Error(err)
	}

	for _, v := range evt.AddressList {
		for _,v1 := range v {
			t.Log(v1.Hex())
		}
	}
}

func TestUnpackStudentEvent(t *testing.T) {
	tabi := types.NewAbi(AbiFilePath)
	data := "0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000030000000000000000000000000c0b638ffccb4bdc4c0d0d5fef062fc512c9251100000000000000000000000096124db0972e3522a9b3910578b3f2e1a50159c700000000000000000000000096124db0972e3522a9b3910578b3f2e1a50159c70000000000000000000000000c0b638ffccb4bdc4c0d0d5fef062fc512c9251100000000000000000000000086324df0972e3522a9b3910578b3f2e1a50132d500000000000000000000000096124db0972e3522a9b3910578b3f2e1a50159c7"
	evt := &StudentEvent{}
	err := common.UnpackEvent(tabi.Events["StudentEvent"].Inputs, evt, hexutil.MustDecode(data), []string{})
	if err != nil {
		panic(err)
	}

	for _, v := range evt.AddressList {
		for _,v1 := range v {
			t.Log(v1.Hex())
		}
	}
}
