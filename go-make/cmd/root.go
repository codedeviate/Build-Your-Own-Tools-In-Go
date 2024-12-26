package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/username/go-make/executor"
	"github.com/username/go-make/vars"

	"github.com/spf13/cobra"
	"github.com/username/go-make/make_file"
)

var rootCmd = &cobra.Command{
	Use:   "gomake",
	Short: "GoMake is a lightweight alternative to GNU Make",
	Long:  `GoMake parses a simple Makefile and executes the tasks defined in it.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please specify a target task")
			os.Exit(1)
		}
		taskName := ""
		for _, arg := range args {
			if strings.Contains(arg, "=") {
				parts := strings.Split(arg, "=")
				vars.VarList[parts[0]] = vars.EvalVar(parts[1])
			} else if len(taskName) == 0 {
				taskName = arg
			} else {
				log.Fatal("Too many arguments")
				os.Exit(1)
			}
		}
		makeFileName, _ := cmd.Flags().GetString("file")
		if len(makeFileName) == 0 {
			makeFileName = "Makefile"
		}
		tasks, err := make_file.ParseMakefile(makeFileName)
		if err != nil {
			fmt.Printf("Error parsing Makefile: %v\n", err)
			os.Exit(1)
		}

		if err := executor.RunTask(taskName, tasks); err != nil {
			fmt.Printf("Error executing task: %v\n", err)
			os.Exit(1)
		}
	},
}

func Execute() error {
	rootCmd.Flags().StringP("file", "f", "Makefile", "Specify the Makefile to use")
	return rootCmd.Execute()
}
