// Copyright © 2022 Thiago Sousa Bastos <thiagosbastos@live.com>

package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

type Backlog struct {
	TargetStoryPoints int
	Tasks             []string
	StoryPoints       []int
	Priorities        []int
}

type Sprint struct {
	TotalStoryPoints int
	Tasks            []string
	StoryPoints      []int
	Priorities       []int
}

var selectFromCmd = &cobra.Command{
	Use:   "selectFrom <path_to_json_file>",
	Short: "Select tasks from the tasklist",
	Long: `Select tasks from the tasklist that optimize the value of the sprint
while meeting the target story-points.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fileName := args[0]
		var fileContents, err = os.ReadFile(fileName)
		if err != nil {
			fmt.Print(err.Error())
			return
		}

		var backlog Backlog
		json.Unmarshal([]byte(fileContents), &backlog)

		lenTasks := len(backlog.Tasks)
		lenStoryPoints := len(backlog.StoryPoints)
		lenPriorities := len(backlog.Priorities)
		if lenTasks != lenStoryPoints && lenStoryPoints != lenPriorities {
			fmt.Println("The tasks, story-points and priorities do not have the same size.")
			fmt.Fprintf(cmd.OutOrStdout(), "Length of Tasks: %d\n", lenTasks)
			fmt.Fprintf(cmd.OutOrStdout(), "Length of StoryPoints: %d\n", lenStoryPoints)
			fmt.Fprintf(cmd.OutOrStdout(), "Length of Priorities: %d\n", lenPriorities)
			return
		}

		sprint := findOptimalSolution(backlog)

		fmt.Fprintf(
			cmd.OutOrStdout(),
			"Sprint: \n TotalStoryPoints: %d\n Tasks: %s\n StoryPoints: %d\n Priorities: %d\n",
			sprint.TotalStoryPoints, sprint.Tasks, sprint.StoryPoints, sprint.Priorities,
		)
	},
}

func init() {
	rootCmd.AddCommand(selectFromCmd)
}

// Finds the optimal combination of tasks within a sprint by modeling the
// problem as a 0/1 knapsack problem. The procedure for finding the solution
// uses the dynamic programming version of the algorithm.
func findOptimalSolution(backlog Backlog) Sprint {
	numberOfTasks := len(backlog.Tasks)
	maxStoryPoints := backlog.TargetStoryPoints

	values := make([][]int, numberOfTasks+1)
	areKept := make([][]bool, numberOfTasks+1)
	for i := 0; i < numberOfTasks+1; i++ {
		values[i] = make([]int, maxStoryPoints+1)
		areKept[i] = make([]bool, maxStoryPoints+1)
	}

	// Base case: no tasks in the sprint
	for j := 0; j < maxStoryPoints+1; j++ {
		values[0][j] = 0
		areKept[0][j] = false
	}
	// Base case: sprint with no story-points
	for i := 0; i < numberOfTasks+1; i++ {
		values[i][0] = 0
		areKept[i][0] = false
	}

	// Optimization procedure
	for i := 1; i <= numberOfTasks; i++ {
		for j := 1; j <= maxStoryPoints; j++ {

			taskFits := (backlog.StoryPoints[i-1] <= j)
			if !taskFits {
				continue
			}

			currMaxValue := backlog.Priorities[i-1] + values[i-1][j-backlog.StoryPoints[i-1]]
			prevMaxValue := values[i-1][j]

			if currMaxValue > prevMaxValue {
				values[i][j] = currMaxValue
				areKept[i][j] = true
			} else {
				values[i][j] = prevMaxValue
				areKept[i][j] = false
			}
		}
	}

	i := numberOfTasks
	j := maxStoryPoints
	var indices []int
	for i > 0 {
		if areKept[i][j] {
			indices = append(indices, i-1)
			j -= backlog.StoryPoints[i-1]
		}
		i--
	}

	var sprint Sprint
	sprint.TotalStoryPoints = 0

	sort.Ints(indices)
	for i, idx := range indices {
		sprint.Tasks = append(sprint.Tasks, backlog.Tasks[idx])
		sprint.StoryPoints = append(sprint.StoryPoints, backlog.StoryPoints[idx])
		sprint.Priorities = append(sprint.Priorities, backlog.Priorities[idx])
		sprint.TotalStoryPoints += sprint.StoryPoints[i]
	}

	return sprint
}
