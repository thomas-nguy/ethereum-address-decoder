package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const abiJSON = `[{"inputs":[{"internalType":"string","name":"recipient","type":"string"}],"name":"send_cro_to_crypto_org","outputs":[],"stateMutability":"payable","type":"function"}]`

func main() {

	/**
	Example Legacy tx:
	{
		"@type": "/ethermint.evm.v1.MsgEthereumTx",
		"data": {
			"@type": "/ethermint.evm.v1.LegacyTx",
			"nonce": "332",
			"gas_price": "5000000000000",
			"gas": "33578",
			"to": "0x6b1b50c2223eb31E0d4683b046ea9C6CB0D0ea4F",
			"value": "102030243391546367224",
			"data": "xBzCcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACpjcm8xNzZxOGFtNmM4aHNrNHIyazR2c21xODBoYXBra3U1M3l5NnVoZm0AAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			"v": "VQ==",
			"r": "iP76vlA2inv1keUSgS4WA9o3rJQn+H57mnDNJO/FeNY=",
			"s": "T8t898tl+P4aFmsop/KR2JgG8rt8EG7Znhh3IgFN7jo="
		},
		"size": 247,
		"hash": "0xcdd1b6fe327e5c17e11be724768452c41bc9e6906cb105ded82a42cdacbbaaef",
		"from": ""
	}
	*/

	// To decode v,r,s first translate base64 to hex https://base64.guru/converter/decode/hex
	// then convert hex to decimal https://www.rapidtables.com/convert/number/hex-to-decimal.html
	v, _ := base64.StdEncoding.DecodeString("VQ==")
	v_h := hex.EncodeToString(v)
	r, _ := base64.StdEncoding.DecodeString("iP76vlA2inv1keUSgS4WA9o3rJQn+H57mnDNJO/FeNY=")
	r_h := hex.EncodeToString(r)
	s, _ := base64.StdEncoding.DecodeString("T8t898tl+P4aFmsop/KR2JgG8rt8EG7Znhh3IgFN7jo=")
	s_h := hex.EncodeToString(s)

	va, _ := new(big.Int).SetString(v_h, 16)
	ra, _ := new(big.Int).SetString(r_h, 16)
	sa, _ := new(big.Int).SetString(s_h, 16)

	// to decode data, translate base64 to hex
	data_base64, _ := base64.StdEncoding.DecodeString("xBzCcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACpjcm8xNzZxOGFtNmM4aHNrNHIyazR2c21xODBoYXBra3U1M3l5NnVoZm0AAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	data_h := hex.EncodeToString(data_base64)
	data, _ := hex.DecodeString(data_h)

	address := common.HexToAddress("0x6b1b50c2223eb31E0d4683b046ea9C6CB0D0ea4F")

	vala, _ := new(big.Int).SetString("102030243391546367224", 10)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    uint64(332),
		To:       &address,
		Value:    vala,
		Gas:      33578,
		GasPrice: big.NewInt(int64(5000000000000)),
		Data:     data,
		V:        va,
		R:        ra,
		S:        sa,
	})

	signer := types.NewEIP155Signer(big.NewInt(25))
	sender, _ := signer.Sender(tx)

	// load contract ABI
	abi, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		log.Fatal(err)
	}
	recipient, _ := abi.Methods["send_cro_to_crypto_org"].Inputs.Unpack(data[4:])

	// get ibc timeout timestamp
	blocktimeString := "2022-03-28T15:45:02.835813016Z"
	blocktime, err := time.Parse(time.RFC3339Nano, blocktimeString)
	if err != nil {
		fmt.Println("Could not parse time:", err)
	}

	fmt.Printf("sender address: 0x%x\n", sender)
	fmt.Printf("recipient address: %v\n", recipient[0])
	fmt.Printf("timeout timestamp: %v\n", blocktime.UnixNano()+86400000000000)

}
