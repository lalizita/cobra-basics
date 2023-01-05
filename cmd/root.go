package cmd

import (
	"fmt"
	"os"

	"github.com/lalizita/cobra-basics/internal/todo"

	"github.com/spf13/cobra"
)

type RootCmd struct {
	*cobra.Command
}

func Execute() {
	todos := &todo.Todos{}
	todos.ReadTodos()

	var rooCmd *RootCmd
	rCmd := &cobra.Command{Use: "app"}

	rCmd.AddCommand(rooCmd.ListCmd(todos))
	rCmd.AddCommand(rooCmd.AddCmd())

	if err := rCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
