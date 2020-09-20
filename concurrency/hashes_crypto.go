// Hashes and Cryptography
package concurrency

import (
	"fmt"
	"hash/crc32"
	"crypto/sha1"
)

func CalculateNonCryptoHash()  {
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println("Non Crypto Hash Value: ",v)
}

func CalculateCryptoHash()  {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println("Crypto Hash Value: ",bs)
}


