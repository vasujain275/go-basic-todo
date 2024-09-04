package cmd

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vasujain275/go-basic-todo/utils"
)

var Priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "It adds the tasks to tasklist with priority",
	Long: `This command add a specific task with its prioority to the tasklsit.
	For example:
		tasks add {taskname} -p 4
		tasks add {taskname} --priority 4
		`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := utils.GetFile()
		if err != nil {
			panic(err)
		}
		writer := csv.NewWriter(file)
		writer.Flush()
		name := strings.Join(args, " ")
		priority := cmd.LocalFlags().Lookup("priority").Value.String()
		fmt.Println(name, priority)
		row := []string{name, priority}
		writer.Write(row)
		writer.Flush()
		file.Close()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&Priority, "priority", "p", 1, "Priority of the task")
}
