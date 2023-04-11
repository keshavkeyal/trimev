package main

import (
	"bufio"
	"fmt"
	"os"
	decryption "trimev/decryption"
	encryption "trimev/encryption"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	fmt.Println(input)
	encryption.Encrypt(input)
	////for decryption
	fmt.Println("now lets decrypt")
	var inputs string
	fmt.Scan(&inputs)
	fmt.Println(inputs)
	decryption.Decryption(inputs)
}
