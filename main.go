package main

import (
	"bufio"
	"log"
	"lsp-minusOne/rpc"
	"os"
)

func main() {
	logger := getLogger("log.txt")
	logger.Println("Hey, I started!")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you did not give me a good file")
	}

	return log.New(logfile, "[lsp-minusOne] ", log.Ldate|log.Ltime|log.Lshortfile)
}
