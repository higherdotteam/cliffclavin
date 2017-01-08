package main

import "fmt"
import "net"
import "os"

import "bufio"

func main() {
	fmt.Println("cliffclavin helps with mail.")
	conn, err := net.Dial("tcp", os.Getenv("CLIFFCLAVIN")+":587")
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	r, _ := reader.ReadString('\n')
	fmt.Println(r)
	writer.WriteString("EHLO " + os.Getenv("CLIFFCLAVIN") + "\r\n")
	writer.Flush()
	r, _ = reader.ReadString('\n')
	fmt.Println(r)

}
