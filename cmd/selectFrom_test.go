package cmd

import (
	"bytes"
	"testing"
)

func TestExpectedFile(t *testing.T) {
	regularTest := []string{"selectFrom", "test.json"}
	actualOutput := new(bytes.Buffer)
	rootCmd.SetOut(actualOutput)
	rootCmd.SetArgs(regularTest)
	rootCmd.Execute()

	const expectedOutput = "Sprint: \n TotalStoryPoints: 5\n Tasks: [TaskA TaskB TaskD TaskF]\n StoryPoints: [1 1 3 0]\n Priorities: [1 2 10 4]\n"
	if actualOutput.String() != expectedOutput {
		t.Fatalf("expected \n\"%s\" got \n\"%s\"", expectedOutput, actualOutput.String())
	}
}

func TestEmptyInputFile(t *testing.T) {
	emptyJSON := []string{"selectFrom", "test_empty.json"}
	actualOutput := new(bytes.Buffer)
	rootCmd.SetOut(actualOutput)
	rootCmd.SetArgs(emptyJSON)
	rootCmd.Execute()

	const expectedOutput = "Sprint: \n TotalStoryPoints: 0\n Tasks: []\n StoryPoints: []\n Priorities: []\n"
	if actualOutput.String() != expectedOutput {
		t.Fatalf("expected \n\"%s\" got \n\"%s\"", expectedOutput, actualOutput.String())
	}
}

func TestFileWithArraysWithDifferentLengths(t *testing.T) {
	file_different_lengths := []string{"selectFrom", "test_diff_lengths.json"}
	actualOutput := new(bytes.Buffer)
	rootCmd.SetOut(actualOutput)
	rootCmd.SetArgs(file_different_lengths)
	rootCmd.Execute()

	const expectedOutput = "Length of Tasks: 5\nLength of StoryPoints: 6\nLength of Priorities: 4\n"
	if actualOutput.String() != expectedOutput {
		t.Fatalf("expected \n\"%s\" got \n\"%s\"", expectedOutput, actualOutput.String())
	}
}
