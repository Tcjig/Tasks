package main

import "fmt"

var currentPlayer string
var board [3][3]string

func main() {
	// 打印棋盘部分
	board = [3][3]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}

	currentPlayer := "A"

	// 利用无限循环条件来持续进行
	for {
		printBoard(board)
		row, col := playerInput()
		if piecesMove(&board, row, col, currentPlayer) {
			if playerWin(currentPlayer) {
				printBoard(board)
				fmt.Printf("玩家 %s 获胜！\n", currentPlayer)
				break
			}
			if playerDogfall() {
				printBoard(board)
				fmt.Println("两个玩家平局")
				break
			}
			if currentPlayer == "A" {
				currentPlayer = "B"
			} else {
				currentPlayer = "A"
			}
		} else {
			fmt.Println("输入错误，请重新输入")
		}
	}
}

// 打印初始棋盘函数
func printBoard(board [3][3]string) {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			fmt.Printf(" %s ", board[row][col])
			if col < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if row < 2 {
			fmt.Println("---|---|---")
		}
	}
}

// 获取玩家输入函数
func playerInput() (int, int) {
	var row, col int
	for {
		fmt.Print("请输入正确的行和列：")

		//错误处理
		_, err := fmt.Scan(&row, &col)
		if err != nil || row < 0 || col < 0 || row > 2 || col > 2 {
			fmt.Println("输入错误，请重新输入")
			continue
		}
		return row, col
	}
}

// 移动棋子函数
func piecesMove(board *[3][3]string, row int, col int, player string) bool {
	if board[row][col] == " " {
		board[row][col] = player
		return true
	}
	return false
}

// 检查胜利条件函数
func playerWin(player string) bool {
	for i := 0; i < 3; i++ {
		if (board[i][0] == player && board[i][1] == player && board[i][2] == player) ||
			(board[0][i] == player && board[1][i] == player && board[2][i] == player) {
			return true
		}
	}
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}

// 检查平局条件函数
func playerDogfall() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}
