package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/devasherr/lambda/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the lambda language repl (go wild!!)\n", user.Name)
	repl.Start(os.Stdin, os.Stdout)
}
