package main

import "io/ioutil"

func main() {
	var message = []byte("AYAYAAA")
	ioutil.WriteFile("./../../token", message, 755)
}
