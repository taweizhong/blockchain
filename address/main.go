package main

import "fmt"

func main() {
	w := NewWalletKeyPair()

	fmt.Printf("%x\n", w.PublicKey)
	fmt.Printf("%s\n", w.PrivateKey)

	address := w.GetAddress()
	fmt.Println(address)
	if IsValidAddress(address) {
		fmt.Println("正确")
	}

}
