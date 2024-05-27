package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	n, _ := new(big.Int).SetString("5f85a133692139651055c4d4cf04d7d66a3c002a3779ec6746a6b9fa558514653de0cd333b1796a8a5eb9e3eefc92a402af1f080447ae498a93a1923a7bd8853", 16)
	
	e, _ := new(big.Int).SetString("10001", 16)
	
	// example data
	headers := map[string]string {"CM2-User":"5154741620","CM2-Nonce":"145328524","CM2-Signature":"16a576cca2a89ad303a0900c61502da4986c39fe37e692f7df1c3580c18b041f3cbaf88d1eb9637df1ead4266fbdcb4b416321fbae59a1773b5394ce8e8e9a87"}

	nonce, _ := strconv.ParseUint(headers["CM2-Nonce"], 10, 64)
	userId, _ := strconv.ParseUint(headers["CM2-User"], 10, 64)
	signature, _ := new(big.Int).SetString(headers["CM2-Signature"], 16)
	
	expectedResult := new(big.Int).SetUint64(userId + nonce)
	result := new(big.Int).Exp(signature, e, n)
	
	success := result.Cmp(expectedResult) == 0
	fmt.Println(success)
}