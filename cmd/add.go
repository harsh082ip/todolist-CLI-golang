package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AddTodo = &cobra.Command{
	Use:   "add",
	Short: "Add a Todo",
	Run:   AddTodos,
}

func AddTodos(cmd *cobra.Command, args []string) {
	title, _ := cmd.Flags().GetString("title")
	desc, _ := cmd.Flags().GetString("desc")
	if desc != "" {
		fmt.Println("You have Successfully Added a Todooo")
		fmt.Printf("%v --- %v \n", title, desc)
		return
	}
	fmt.Println("You have Successfully Added a Todo")
	fmt.Printf("%v \n", title)
}

func init() {
	AddTodo.Flags().String("title", "", "Title of the Todo")
	AddTodo.Flags().String("desc", "", "Optional: Description of Todo")
	_ = AddTodo.MarkFlagRequired("title")
	// _ = AddTodo.MarkFlagRequired("desc")
}
