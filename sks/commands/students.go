package commands

import (
	"github.com/spf13/cobra"
)

var students = &cobra.Command{
	Use:   "students",
	Short: "List student logins",
	Run: func(cmd *cobra.Command, args []string) {
		println("would list students")
	},
}
