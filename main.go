package main

import "fmt"
import "net"
import "os"

import "bufio"

func main() {
	fmt.Println("cliffclavin helps with mail.")
	conn, err := net.Dial("tcp", os.Getenv("SERVER")+":587")
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	r, _ := reader.ReadString('\n')
	fmt.Println(r)
	writer.WriteString("EHLO " + os.Getenv("SERVER") + "\r\n")
	writer.Flush()
	r, _ = reader.ReadString('\n')
	fmt.Println(r)

	writer.WriteString("AUTH LOGIN\r\n")
	writer.Flush()
	r, _ = reader.ReadString('\n')
	fmt.Println(r)

	writer.WriteString("MAIL FROM:<" + os.Getenv("FROM") + ">\r\n")
	writer.Flush()
	r, _ = reader.ReadString('\n')
	fmt.Println(r)

	writer.WriteString("RCPT TO:<" + os.Getenv("TO") + ">\r\n")
	writer.Flush()
	r, _ = reader.ReadString('\n')
	fmt.Println(r)

	writer.WriteString("DATA\r\nhi\r\n.\r\n")
	writer.Flush()
	r, _ = reader.ReadString('\n')
	fmt.Println(r)

	writer.WriteString("QUIT\r\n")
	writer.Flush()
	r, _ = reader.ReadString('\n')
	fmt.Println(r)

}
