package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
	//	"runtime"
)

var allocSize int

func writeToFile(path, text string) {
	var file, _ = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0755)
	file.WriteString("\n")
	file.WriteString("m" + text)
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

func playWithFile(path string, sec int) {
	board := createSimpleBoard()
	writeToFile(path, convertToTxt(board))
	convertFromTxt(readFromFile(path)[len(readFromFile(path))-1]).show()
	time.Sleep(time.Duration(sec) * time.Second)
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

		now, ocena = now.evaluateAlfaBeta(21, now.player)
		now.player = !now.player
		now.show()
		fmt.Println(ocena)
		now, ocena = now.evaluateAlfaBeta(21, now.player)
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
	playWithFile("./dane", 0)
}
