package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// 启动服务器监听
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("服务器启动失败:", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器已启动，等待客户端连接...")

	// 接收两个客户端连接
	connA, err := listener.Accept()
	if err != nil {
		fmt.Println("无法接收客户端 A:", err)
		return
	}
	defer connA.Close()
	fmt.Println("客户端 A 已连接")

	connB, err := listener.Accept()
	if err != nil {
		fmt.Println("无法接收客户端 B:", err)
		return
	}
	defer connB.Close()
	fmt.Println("客户端 B 已连接")

	// 游戏循环
	for {
		// 获取玩家 A 的选择
		moveA := getPlayerMove(connA, "A")

		// 获取玩家 B 的选择
		moveB := getPlayerMove(connB, "B")

		// 比较并宣布结果
		result := getResult(moveA, moveB)
		fmt.Fprintln(connA, result)
		fmt.Fprintln(connB, result)
	}
}

// 获取玩家的输入
func getPlayerMove(conn net.Conn, player string) string {
	var move string
	for {
		fmt.Fprintf(conn, "玩家 %s，请输入 石头（shitou）、剪刀（jiandao） 或 布（bu）：\n", player)
		reader := bufio.NewReader(conn)
		move, _ = reader.ReadString('\n')
		move = strings.TrimSpace(move) // 移除空格和换行符
		move = strings.ToLower(move)   // 转换为小写

		// 支持汉字和拼音输入
		if move == "石头" || move == "shitou" || move == "剪刀" || move == "jiandao" || move == "布" || move == "bu" {
			break
		}
		fmt.Fprintln(conn, "无效输入，请重新输入 石头（shitou）、剪刀（jiandao） 或 布（bu）")
	}

	// 将拼音转换为对应的汉字
	switch move {
	case "shitou":
		move = "石头"
	case "jiandao":
		move = "剪刀"
	case "bu":
		move = "布"
	}

	return move
}

// 比较两位玩家的选择并返回结果
func getResult(moveA string, moveB string) string {
	if moveA == moveB {
		return "平局！"
	}

	if (moveA == "石头" && moveB == "剪刀") ||
		(moveA == "剪刀" && moveB == "布") ||
		(moveA == "布" && moveB == "石头") {
		return "玩家 A 胜利！"
	}

	return "玩家 B 胜利！"
}
