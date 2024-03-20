package core

import (
	"bufio"
	"fmt"
	"os"
)

func Input() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
