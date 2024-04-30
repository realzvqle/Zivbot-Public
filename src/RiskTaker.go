package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
)

func WritetoFile(money int) {
	ioutil.WriteFile("money", []byte(strconv.Itoa(money)), 0644)

}
func RiskTaker(money int) (string, int) {
	//response := "wait"
	num := rand.Int()
	if num%7 != 0 {
		money += 1
		if num%6 == 0 {
			money += 3
			return "Lucky You!, You Got Some Extra Money. You Now Have: " + fmt.Sprint(money), money
		}
		//ioutil.WriteFile("money", []byte(fmt.Sprintf(money)), 0644)
		return "You Are Safe, For Now, Take Some Money. You Now Have: " + fmt.Sprint(money), money
	} else {
		money = 0
		return "SHUT UP, YOU LOST YOUR MONEY, You Now Have: " + fmt.Sprint(money), money
	}

	//return response
}
