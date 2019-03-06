package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func getLine() string {
	if len(os.Args) < 2 {
		sc.Scan()
		return sc.Text()
	}
	return os.Args[len(os.Args)-1]
}

func main() {
	inp := getLine()
	sinp := strings.Split(inp, ",")
	
	// TODO : Something code
	sum := 0
	for _, n := range sinp{
		sum += stringconv.Stoi(n)
	}
	
  fmt.Print(sum)
}