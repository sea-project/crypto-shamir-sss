package sss

import (
	"fmt"
	stdbytes "github.com/sea-project/stdlib-bytes"
	"strings"
	"testing"
)

func TestCombine(t *testing.T) {
	secret := "0ddb327ad1059662da1f02f1b8521bf0f69cf5cecc09a4d8fc7f928fc9726818" // our secret
	n := byte(5)                                                                 // create 5 shares
	k := byte(2)		                                                      	// require 2 of them to combine
	bytes := stdbytes.FromHex(secret)
	shares, err := Split(n, k, bytes) // split into 5 shares
	if err != nil {
		fmt.Println(err)
		return
	}
	// 打印切片后私钥
	fmt.Println("print slice prikey begin")
	for x, y := range shares {
		fmt.Println(x, (stdbytes.Bytes2Hex(y)))
	}
	fmt.Println("print slice prikey end")
	// select a random subset of the total shares
	subset := make(map[byte][]byte, k)
	for x, y := range shares { // just iterate since maps are randomized
		subset[x] = y
		fmt.Println(x, (stdbytes.Bytes2Hex(y)))
		if len(subset) == int(k) {
			break
		}
	}
	combine := Combine(subset)
	fmt.Println(stdbytes.Bytes2Hex(combine))
	if strings.Compare(stdbytes.Bytes2Hex(combine), secret) == 0 {
		t.Logf("success")
	} else {
		t.Error("fail")
	}
}
