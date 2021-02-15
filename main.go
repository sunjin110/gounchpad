package main

import (
	"bufio"
	"fmt"
	"gounchpad/pkg/sound"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
)

// 標準出力をcleanして見やすいようにするやつ
var clearCmdMap map[string]func()

func init() {
	log.Println("init")

	// 1文字ずつ入力を読み込むようにする
	// stty raw
	cmd := exec.Command("stty", "raw")
	cmd.Stdin = os.Stdin
	cmd.Run()

	clearCmdMap = map[string]func(){}

	clearCmdMap["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdin
		cmd.Run()
	}

	// TODO windows clean, mac clean
}

func main() {
	log.Println("hello")

	defer func() {
		// 入力モードをもとに戻す
		cmd := exec.Command("stty", "-raw")
		cmd.Stdin = os.Stdin
		cmd.Run()
	}()

	reader := bufio.NewReader(os.Stdin)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt) // 停止signalの検知

	go func() {
		defer close(sig)
		for {
			r, _, err := reader.ReadRune()
			if err != nil {
				if err == io.EOF {
					return
				}
				panic(err)
			}
			clear()
			fmt.Println(string(r))

			// sound
			go sound.Sound(r)

			if r == 'q' {
				clear()
				break
			}

		}
	}()

	<-sig
	fmt.Println("bye...")

}

// clear cuiの画面をclearする
func clear() {
	clearFunc, exists := clearCmdMap[runtime.GOOS]
	if exists {
		clearFunc()
	} else {
		panic("clear funcがありません")
	}
}
