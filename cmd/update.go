package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var UpdateTodo = &cobra.Command{
	Use:   "update",
	Short: "update a todo using it's Id",
	Run:   UpdateTodos,
}

func UpdateTodos(cmd *cobra.Command, args []string) {

	id, _ := cmd.Flags().GetString("id")
	fmt.Printf("Updated %v todo \n", id)
}

func init() {
	UpdateTodo.Flags().String("id", "", "id of the Todo you want to update")
	_ = UpdateTodo.MarkFlagRequired("id")
}
