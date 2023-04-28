package anykeyquit

import (
	"bufio"
	"fmt"
	"os"
)

func PressAnyKeyQuit() {
	for {
		fmt.Println("PressAnyKeyQuit  ...")
		in := bufio.NewReader(os.Stdin)
		in.ReadLine()
		break
	}
}
