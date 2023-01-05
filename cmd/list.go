package cmd

import (
	"fmt"

	"github.com/lalizita/cobra-basics/internal/todo"

	"github.com/spf13/cobra"
)

func (r *RootCmd) ListCmd(todos *todo.Todos) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all your tasks for today",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("<List tasks>")
			fmt.Printf("%+v\n", todos)
		},
	}
}
