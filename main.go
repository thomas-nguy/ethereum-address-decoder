package main

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

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

	va ,_ := new(big.Int).SetString("85",10)
	ra ,_ := new(big.Int).SetString("61965057129201923018618051904524766544240930237387902478223956349334883498198",10)
	sa ,_ := new(big.Int).SetString("36092247489302229830963641029442892947983301809207198407845437301669493599802",10)


	// to decode data, translate base64 to hex
	data, _ := hex.DecodeString("c41cc2700000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002a63726f3137367138616d36633868736b3472326b3476736d7138306861706b6b7535337979367568666d00000000000000000000000000000000000000000000")
	address := common.HexToAddress("0x6b1b50c2223eb31E0d4683b046ea9C6CB0D0ea4F")

	vala ,_ := new(big.Int).SetString("102030243391546367224",10)
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

	signer:= types.NewEIP155Signer(big.NewInt(25));
	sender, _ := signer.Sender(tx)

	fmt.Printf("sender address 0x%x\n", sender)

}