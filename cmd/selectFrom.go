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
		sprint := findOptimalSolution(backlog)

		// TODO: 5. Give feedback upon finishing analysis and having results
		fmt.Printf(
			"TotalStoryPoints: %d\nTasks: %s\nStoryPoints: %d\nPriorities: %d\n",
			sprint.TotalStoryPoints, sprint.Tasks, sprint.StoryPoints, sprint.Priorities,
		)
	},
}

func init() {
	rootCmd.AddCommand(selectFromCmd)
}

// Solves the knapsack problem using dynamic programming
func findOptimalSolution(backlog Backlog) Sprint {
	numberOfTasks := len(backlog.Tasks)
	maxStoryPoints := backlog.TargetStoryPoints

	// 1. construct a values and a areKept tables
	values := make([][]int, numberOfTasks+1) // row with the 'items' to pick
	areKept := make([][]bool, numberOfTasks+1)
	for i := 0; i < numberOfTasks+1; i++ {
		values[i] = make([]int, maxStoryPoints+1) // column with the 'weights'
		areKept[i] = make([]bool, maxStoryPoints+1)
	}

	// 2a. Base case: no items in the 'knapsack'
	for j := 0; j < maxStoryPoints+1; j++ {
		values[0][j] = 0
		areKept[0][j] = false
	}
	// 2b. Base case: knapsack with no weight
	for i := 0; i < numberOfTasks+1; i++ {
		values[i][0] = 0
		areKept[i][0] = false
	}

	// 3. Fill tables with the values
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

	// 4. Retrieve the indices of the values that give the optimal solution
	n := numberOfTasks
	c := maxStoryPoints
	var indices []int
	for n > 0 {
		if areKept[n][c] {
			indices = append(indices, n-1)
			c -= backlog.StoryPoints[n-1]
		}
		n--
	}

	// 5. Fill the solution data structure
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
