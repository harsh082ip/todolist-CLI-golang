package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DeleteTodo = &cobra.Command{
	Use:   "del",
	Short: "delete a todo by it's id",
	Run:   DeleteTodos,
}

func DeleteTodos(cmd *cobra.Command, args []string) {
	id, _ := cmd.Flags().GetString("id")
	fmt.Printf("Deleted %v todo \n", id)
}

func init() {
	DeleteTodo.Flags().String("id", "", "id of todo you want to delete")
	_ = DeleteTodo.MarkFlagRequired("id")
}
