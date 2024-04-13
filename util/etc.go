package util

import (
	"bufio"
	"fmt"
	"os"
)

// 가만히 있다가 엔터키가 들어오면 다음으로 넘어간다.
func PressKey(msg string) {
	kbReader := bufio.NewReader(os.Stdin)

	fmt.Println(msg)
	kbReader.ReadString('\n')
}
