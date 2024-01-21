package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListTodo = &cobra.Command{
	Use:   "list",
	Short: "list all todos",
	Run:   ListTodos,
}

func ListTodos(cmd *cobra.Command, args []string) {
	id, _ := cmd.Flags().GetString("id")
	if id != "" {
		fmt.Printf("Fetched Todo by id: %v \n", id)
		return
	}
	fmt.Println("here is a list of all your todos")
}

func init() {
	ListTodo.Flags().String("id", "", "optional: fetch specific todos by ID")
}
