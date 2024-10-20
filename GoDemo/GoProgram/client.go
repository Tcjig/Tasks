package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.18.130:8080")
	if err != nil {
		fmt.Println("无法连接服务器:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		// 读取服务器消息
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(message)

		// 输入玩家的选择
		playerMove, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, playerMove)

		// 显示结果
		result, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(result)
	}
}
