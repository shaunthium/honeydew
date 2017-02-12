package setup

import (
	"fmt"

	"github.com/shaunthium/honeydew/ssh"
)

func Setup() error {
	fmt.Println("in setup")
	ssh.SSHIntoServer()
	return nil
}
