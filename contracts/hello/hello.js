var browser_untitled_sol_helloworldContract = web3.eth.contract([{"constant":false,"inputs":[],"name":"kill","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_newgreeting","type":"string"}],"name":"setGreeting","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"greet","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getGreeting","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]);
var browser_untitled_sol_helloworld = browser_untitled_sol_helloworldContract.new(
    {
        from: web3.eth.accounts[0],
        data: '0x6060604052341561000f57600080fd5b5b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b61046e806100616000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806341c0e1b51461005f578063a413686214610074578063cfae3217146100d1578063fe50cc7214610160575b600080fd5b341561006a57600080fd5b6100726101ef565b005b341561007f57600080fd5b6100cf600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610281565b005b34156100dc57600080fd5b6100e461029c565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101255780820151818401525b602081019050610109565b50505050905090810190601f1680156101525780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561016b57600080fd5b6101736102e0565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101b45780820151818401525b602081019050610198565b50505050905090810190601f1680156101e15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561027e576000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b8060019080519060200190610297929190610389565b505b50565b6102a4610409565b6040805190810160405280600e81526020017f68656c6c6f20736f6c696469747900000000000000000000000000000000000081525090505b90565b6102e8610409565b60018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561037e5780601f106103535761010080835404028352916020019161037e565b820191906000526020600020905b81548152906001019060200180831161036157829003601f168201915b505050505090505b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106103ca57805160ff19168380011785556103f8565b828001600101855582156103f8579182015b828111156103f75782518255916020019190600101906103dc565b5b509050610405919061041d565b5090565b602060405190810160405280600081525090565b61043f91905b8082111561043b576000816000905550600101610423565b5090565b905600a165627a7a72305820efe3b0201c414f3b16373e9af6378d91f9af89a37876888c144a4629356185ae0029',
        gas: '1000000'
    }, function (e, contract){
        console.log(e, contract);
        if (typeof contract.address !== 'undefined') {
            console.log('Contract mined! address: ' + contract.address + ' transactionHash: ' + contract.transactionHash);
        }
    })