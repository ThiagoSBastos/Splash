// Copyright © 2022 Thiago Sousa Bastos <thiagosbastos@live.com>

package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Backlog struct {
	TargetStoryPoints int      `json:"targetStoryPoints"`
	Tasks             []string `json:"tasks"`
	StoryPoints       []int    `json:"storyPoints"`
	Priorities        []int    `json:"priorities"`
}

var selectFromCmd = &cobra.Command{
	Use:   "selectFrom",
	Short: "Select tasks from the tasklist",
	Long: `Select tasks from the tasklist that optimize the value of the sprint
		   while meeting the target story-points. For example: <putAnExampleHere>.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Read from file
		fileName := args[0]
		var fileContents, err = os.ReadFile(fileName)
		if err != nil {
			fmt.Print(err.Error())
			return
		}

		// 2. Fill Backlog with JSON data
		var backlog Backlog
		json.Unmarshal([]byte(fileContents), &backlog)
		fmt.Printf(
			"TargetStoryPoints: %d\nTasks: %s\nStoryPoints: %d\nPriorities: %d\n",
			backlog.TargetStoryPoints, backlog.Tasks, backlog.StoryPoints, backlog.Priorities,
		)

		// 3. Validate JSON
		lenTasks := len(backlog.Tasks)
		lenStoryPoints := len(backlog.StoryPoints)
		lenPriorities := len(backlog.Priorities)
		if lenTasks != lenStoryPoints && lenStoryPoints != lenPriorities {
			fmt.Println("The tasks, story-points and pririties do not have the same size.")
			fmt.Printf("Length of Tasks: %d\n", lenTasks)
			fmt.Printf("Length of StoryPoints: %d\n", lenStoryPoints)
			fmt.Printf("Length of Priorities: %d\n", lenPriorities)
			return
		}

		// TODO: 4. Analyze

		// TODO: 5. Give feedback upon finishing analysis and having results
	},
}

func init() {
	rootCmd.AddCommand(selectFromCmd)
}
