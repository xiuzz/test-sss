package main

import (
	"fmt"
	"sss/sss"
)

func main() {
	// fmt.Println("hello world")

	// secret := byte(126)
	// secrets := sss.SplitSecret(secret, 10, 2)
	// data, _ := json.Marshal(secrets)
	// fmt.Println(string(data))
	// value := sss.Recover(secrets[:2], 2)
	// fmt.Println(value)
	secrets := sss.EnCrypto("qfoqofekjxcnqowhfvqhfoqwnfcoaxcahdowqfyhpopwqhdwqpfnnwfcuncqpn", 10, 2)
	fmt.Println(secrets)
	value := sss.DeCrypto(secrets[:2], 2)
	fmt.Println(string(value))
}
