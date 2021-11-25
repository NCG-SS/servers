package main

import (
	"Own_GoInAction1/Activity16-MultipleServers/customserver"
	"fmt"
)

func main() {
	fmt.Println("TCP server")
	customserver.StartJSONQuery()
}
