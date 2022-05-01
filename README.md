# Splash

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/ThiagoSBastos/Splash/Go?style=plastic)

> A Sprint Planning CLI Application

The motivation of this application is to automate the selection of tasks within a Sprint in an Agile team.


Splash reads JSON file containing the target story points for a sprint and tasks in the backlog, which consists of a list of names, priorities, and story points. By calling the `selectFrom` command, Splash outputs the Sprint that maximizes the priorities while keeping the total story points lesser or equal to target story points.

## Usage
```
Splash [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  selectFrom  Select tasks from the tasklist

Flags:
  -h, --help   help for Splash

Use "Splash [command] --help" for more information about a command.
```

## Example
Imagine having the following backlog in a JSON file:
```JSON
{
    "targetStoryPoints": 5,
    "tasks": [
        "TaskA",
        "TaskB",
        "TaskC",
        "TaskD",
        "TaskE",
        "TaskF"
    ],
    "storyPoints": [1, 1, 8, 3, 5, 0],
    "priorities": [1, 2, 5, 10, 2, 4]
}
```
By typing:
```
Splash selectFrom path/to/JSONFile
```
The output is:
```
Sprint:
 TotalStoryPoints: 5
 Tasks: [TaskA TaskB TaskD TaskF]
 StoryPoints: [1 1 3 0]
 Priorities: [1 2 10 4]
```
And that's it! Your sprint is automatically planned for you and your team!

## Notes
This application assumes that the tasks don't have any dependencies. Therefore, if your backlog has any kind of dependency, I advise you to filter the tasks beforehand to have a feasible output when using Splash.