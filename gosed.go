package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: kubectl logs -f [pod_name] | ./gosed")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	for scanner.Scan() {
		fmt.Println(strings.ReplaceAll(scanner.Text(), "\u2016", "\n"))
	}

}
