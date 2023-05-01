package main

import (
	"fmt"

	"github.com/Rezking/cryptit_with_go/decrypt"
	"github.com/Rezking/cryptit_with_go/encrypt"
)

func main() {
	// fmt.Println(encrypt.Nimbus("Kodekloud"))
	encrypted := encrypt.Nimbus("Kodekloud")
	fmt.Println(encrypted)

	fmt.Println("The decrypted value is", decrypt.Nimbus(encrypted))
}
