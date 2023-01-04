package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (r *RootCmd) AddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a task to make today",
		Long:  `You can add add a new task in your to do list`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("<add task>")
		},
	}
}
