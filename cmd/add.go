/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"
	"time"

	"github.com/lalizita/cobra-basics/internal/todo"
	"github.com/spf13/cobra"
)

func (r *RootCmd) AddCmd(t *todo.Todos) *cobra.Command {
	return &cobra.Command{
		Use:   "add [task]",
		Short: "Describe your task",
		Long:  "Write a simple description about your task.",
		Run: func(cmd *cobra.Command, args []string) {
			task := todo.TaskBody{
				Task:      strings.Join(args[:], " "),
				Checked:   false,
				CreatedAt: time.Now(),
			}

			*t = append(*t, task)

			_ = t.Store()

		},
	}
}
