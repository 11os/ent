package main

import (
	"fmt"
	"./src/server"
	"./src/utils"
)

func main() {
	server.Start()
	fmt.Print("detail = %s" , utils.Get58list("http://sz.58.com/ershoufang/?PGTID=0d100000-0000-4c2c-7d40-cae3f496c747&ClickID=3"))
}
