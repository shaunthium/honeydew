package deploy

import (
	"fmt"

	"github.com/shaunthium/honeydew/ssh"
)

func Deploy() error {
	fmt.Println("in deploy")
	ssh.SSHIntoServer()
	return nil
}
