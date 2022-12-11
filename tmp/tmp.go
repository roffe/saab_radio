package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("eeprom.bin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("var eeprom = []byte{")
	for _, b := range data {
		fmt.Printf("0x%02X, ", b)
	}
	fmt.Println("}")

}
