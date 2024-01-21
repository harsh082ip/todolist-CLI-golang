package main

import (
	"fmt"
	"os"

	"github.com/harsh082ip/todolist-CLI-golang/cmd"
	// "github.com/spf13/cobra"
)

func main() {

	cmd.RootCmd.AddCommand(cmd.ListTodo)
	// cmd.Root2.AddCommand(cmd.RootCmd)

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
