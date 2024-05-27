package main

import (
	"net/http"
	"fmt"
	"strconv"
	"math/big"
)

var n, e *big.Int

func VerifyAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// basic checks
		userIdIn := r.Header.Get("CM2-User")
		if len(userIdIn) > 12 { return }
		
		nonceIn := r.Header.Get("CM2-Nonce")
		if len(nonceIn) > 10 { return }
		
		signatureIn := r.Header.Get("CM2-Signature")
		if len(signatureIn) > 256 { return }
		
		// convert
		nonce, err := strconv.ParseUint(nonceIn, 10, 64)
		if err != nil { return }
		
		userId, err := strconv.ParseUint(userIdIn, 10, 64)
		if err != nil { return }
		
		signature, ok := new(big.Int).SetString(signatureIn, 16)
		if !ok { return }

		// check signature
		expectedResult := new(big.Int).SetUint64(userId + nonce)
		result := new(big.Int).Exp(signature, e, n)

		success := result.Cmp(expectedResult) == 0
		if !success { return }
	
		next.ServeHTTP(w, r)
	})
}


func main() {
	// init vars
	var ok bool
	n, ok = new(big.Int).SetString("5f85a133692139651055c4d4cf04d7d66a3c002a3779ec6746a6b9fa558514653de0cd333b1796a8a5eb9e3eefc92a402af1f080447ae498a93a1923a7bd8853", 16)
	if !ok {
		panic("Failed to initialize n")
	}

	e, ok = new(big.Int).SetString("10001", 16)
	if !ok {
		panic("Failed to initialize e")
	}

	// start server
	r := http.NewServeMux()

	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("CM2-User")

		fmt.Fprint(w, userId)
	})

	http.ListenAndServe(":8000", VerifyAuthentication(r))
}