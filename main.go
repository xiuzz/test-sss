package main

import (
	"encoding/json"
	"fmt"

	"sss/sss"
)

func main() {
	// fmt.Println("hello world")

	secret := byte(126)
	secrets := sss.SplitSecret(secret, 10, 2)
	data, _ := json.Marshal(secrets)
	fmt.Println(string(data))
	value := sss.Recover(secrets[:2], 2)
	fmt.Println(value)
}
