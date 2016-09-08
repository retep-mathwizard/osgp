package commands

import (
	"github.com/spf13/cobra"
	"os"
)

const description string = `
`

var root = &cobra.Command{
	Use:   "sks",
	Short: "Do all the skilstak admin things.",
	//Long:  description,
	Run: func(cmd *cobra.Command, args []string) {
		println(args[0])
	},
}

// Execute ties everything to the Cobra root command.
// We do our own Find()ing so we can infer a command
// from the arguments is a suitable sub command is not
// found to allow shortcuts like `sks robmuh`
func Execute() {
	addCommands()
	cmd, args, err := root.Find(os.Args[1:])
	if err != nil {
		args = inferFromArgs(args)
	}
	cmd.SetArgs(args)
	if err := cmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

// Add commands implemented in their own <command>.go files here
// Keep the file and function name same as command itself.
// Pick simple names that don't conflict with stuff in `commands` package.
func addCommands() {
	root.AddCommand(students)
}

func inferFromArgs(a []string) []string {
	println("Would infer from " + a[0])
	// TODO infer it
	return a
}
