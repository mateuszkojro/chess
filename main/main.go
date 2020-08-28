package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var allocSize int

func writeToFile(path, text string) {
	var file, _ = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0755)
	file.WriteString("\n")
	file.WriteString(text)
	file.Close()
}

func readFromFile(path string) []string {
	var res []string
	var file, _ = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0755)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

// !! Uzywam czystego alfa beta nie tego ok
func playWithFile(path string, sec int) {
	board := createSimpleBoard()
	ocena := 0
	for i := 0; i < 10; i++ {
		fmt.Println("Reading from file...Now")
		boardsText := readFromFile(path)
		for boardsText[len(boardsText)-1][0] == '*' {
			fmt.Println("Another player didnt play yet going to sleep for 10 s")
			time.Sleep(10 * time.Second)
			boardsText = readFromFile(path)
		}
		board = convertFromTxt(boardsText[len(boardsText)-1])
		board, ocena = board.evaluateAlfaBeta(20, board.player)
		fmt.Println("Ruch:", len(boardsText)-1, " Ocena:", ocena)
		writeToFile(path, convertToTxt(board))
		/*output, _:=*/ exec.Command("python", "show.py").Output()
		// /* ruchy czranego
		board.player = !board.player
		board, ocena = board.evaluateAlfaBeta(20, board.player)
		fmt.Println("Ruch:", len(boardsText)-1, " Ocena:", ocena)
		writeToFile(path, convertToTxt(board))
		exec.Command("python", "show.py").Output()
		//koniec ruchy czranego */
		//fmt.Println("Linux said:", string(output))
		fmt.Println("Going to sleep for", sec, "seconds")
		time.Sleep(time.Duration(sec) * time.Second)
	}
}

func playWithYourself() {
	fmt.Println("analyzeBoard: ")
	now := createSimpleBoard()
	now.player = true
	now, ocena := now.evaluateAlfaBeta(101, now.player)
	now.show()
	now.player = !now.player
	now, ocena = now.evaluateAlfaBeta(101, now.player)
	now.show()
	now.player = !now.player

	count := 15

	for i := 0; i < count; i++ {
		progress(i, count)

		now, ocena = now.evaluateAlfaBeta(101, now.player)
		now.player = !now.player
		now.show()
		fmt.Println(ocena)
		now, ocena = now.evaluateAlfaBeta(101, now.player)
		now.player = !now.player
		now.show()
		fmt.Println(ocena)

		fmt.Println("Ruch", i+1)
	}

	//fmt.Println(now.player)
	//now, ocena = now.evaluateAlfaBeta(2)
	//now.player = !now.player
	//now, ocena = now.evaluateAlfaBeta(2)
	now.show()
	fmt.Println("ocena: ", ocena)
}

func main() {
	runtime.GOMAXPROCS(8)
	//runtime.GOMAXPROCS(1)
	playWithFile("./dane", 5)
	//writeToFile("./dane", convertToTxt(createSimpleBoard()))
	//playWithYourself()
}
