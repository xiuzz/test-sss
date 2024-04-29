package sss

import (
	"fmt"
	"math/rand"
	"sort"
)

const p = 257

type Secret struct {
	Index int  `json:"index"`
	Share byte `json:"share"`
}

func makeRandParameter(t int) []int {
	params := make([]int, t)

	for i := 0; i < t; i++ {
		for {
			params[i] = rand.Intn(p)
			if params[i] != 0 {
				break
			}
		}

	}

	return params
}

type shuffle []int

func (s shuffle) Len() int {
	return len(s)
}

func (s shuffle) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s shuffle) Less(i, j int) bool {
	return rand.Int()%2 == 0
}

func makeIndexes(n int) []int {
	var s shuffle = make([]int, 256)
	for i := 0; i < 256; i++ {
		s[i] = i + 1
	}
	sort.Sort(s)
	return s[:n]
}

func Recover(secrets []Secret, t int) byte {
	if len(secrets) != t {
		panic("aaaa")
	}

	coffees := make([][]int, t)
	values := make([]int, t)
	for i := 0; i < t; i++ {
		coffee := make([]int, t)
		tmp := 1
		for j := 0; j < t; j++ {
			coffee[j] = tmp
			tmp *= secrets[i].Index % p
		}
		coffees[i] = coffee

		values[i] = int(secrets[i].Share)
	}

	m, s := recursion(coffees, values)
	fmt.Println(m, s)

	return byte(s / m) // TODO
}

func recursion(coffees [][]int, values []int) (int, int) {
	if len(coffees) != len(values) {
		panic("11111")
	}
	if len(coffees) == 1 {
		c, m := coffees[0][0], values[0]
		fmt.Println(c, m)
		c = ((c % p) + p) % p
		m = ((m % p) + p) % p
		return c, m
	}

	t := len(coffees)
	ltrt := coffees[t-1][t-1]
	newCoffees := make([][]int, len(coffees)-1)
	newValues := make([]int, len(coffees)-1)
	for i := 0; i < t-1; i++ {
		multi := coffees[i][t-1]

		coffee := make([]int, len(coffees)-1)
		for j := 0; j < t-1; j++ {
			coffee[j] = (coffees[i][j]*ltrt - coffees[t-1][j]*multi) % p
		}
		newCoffees[i] = coffee
		newValues[i] = (values[i]*ltrt - values[t-1]*multi) % p
	}

	return recursion(newCoffees, newValues)
}

func SplitSecret(secret byte, n, t int) []Secret {
	params := makeRandParameter(t)
	params[0] = int(secret)

	fmt.Println(params)
	indexes := makeIndexes(n) // TODO
	secrets := make([]Secret, n)
	for i := 0; i < n; i++ {
		secrets[i] = Secret{
			Share: calculate(params, indexes[i]),
			Index: indexes[i],
		}
	}

	return secrets
}

func calculate(param []int, index int) byte {
	sum := 0
	tmp := 1
	for i := 0; i < len(param); i++ {
		sum += param[i] * tmp % p
		tmp *= index
		//tmp = tmp % p
	}
	return byte(sum)
}
