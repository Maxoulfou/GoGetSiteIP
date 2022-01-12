package main

import (
	"fmt"
	"getsiteip/command"
	"os"
)

func main() {
	Site := os.Args[1]
	Ip := command.GetSiteIp(Site)

	fmt.Println("Target site: " + Site)
	fmt.Println("IP: " + Ip)

	return
}
