package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/motivq/baresip-go/binding"
	"github.com/motivq/baresip-go/ctrltcp"
)

func main() {
	go testDial()
	go testTCP()

	go binding.Start()

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())

}

func testDial() {
	time.Sleep(3 * time.Second)
	binding.UAConnect("*61")
}

func testTCP() {
	time.Sleep(2 * time.Second)
	ctrltcp.GetConn()
	ctrltcp.EventHandle(func(info ctrltcp.EventInfo) {
		fmt.Printf("实时事件 %+v", info)
	})
}
