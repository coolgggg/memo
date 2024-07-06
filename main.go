package main

import "fmt"

func main() {
	info := GetOptions(SetClientID(10086), SetCountryCode("cn"))
	fmt.Println(info)
}
