package main

import (
	"fmt"
	"net"
)

func main() {

	conn, _ := net.Dial("tcp", "aj-test.plexer.rcs-debug.ccloud-poc.shoot.canary.k8s-hana.ondemand.com:2379")

	defer conn.Close()

	reply := make([]byte, 1024)
	conn.Read(reply)

	fmt.Println(string(reply))

}
