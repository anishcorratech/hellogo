package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Name: ")
	input, _ := reader.ReadString('\n')
	read_line := strings.TrimSuffix(input, "\n")
	fmt.Printf("Hello %s nice to meet you\n", read_line)
}
