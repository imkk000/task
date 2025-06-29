package summary

import (
	"strings"

	"github.com/imkk000/task/v3/internal/logger"
	"github.com/imkk000/task/v3/taskfile/ast"
)

func PrintTasks(l *logger.Logger, t *ast.Taskfile, c []string) {
	for i, call := range c {
		PrintSpaceBetweenSummaries(l, i)
		if task, ok := t.Tasks.Get(call); ok {
			PrintTask(l, task)
		}
	}
}

func PrintSpaceBetweenSummaries(l *logger.Logger, i int) {
	spaceRequired := i > 0
	if !spaceRequired {
		return
	}

	l.Outf(logger.Default, "\n")
	l.Outf(logger.Default, "\n")
}

func PrintTask(l *logger.Logger, t *ast.Task) {
	printTaskName(l, t)
	printTaskDescribingText(t, l)
	printTaskDependencies(l, t)
	printTaskAliases(l, t)
	printTaskCommands(l, t)
}

func printTaskDescribingText(t *ast.Task, l *logger.Logger) {
	if hasSummary(t) {
		printTaskSummary(l, t)
	} else if hasDescription(t) {
		printTaskDescription(l, t)
	} else {
		printNoDescriptionOrSummary(l)
	}
}

func hasSummary(t *ast.Task) bool {
	return t.Summary != ""
}

func printTaskSummary(l *logger.Logger, t *ast.Task) {
	lines := strings.Split(t.Summary, "\n")
	for i, line := range lines {
		notLastLine := i+1 < len(lines)
		if notLastLine || line != "" {
			l.Outf(logger.Default, "%s\n", line)
		}
	}
}

func printTaskName(l *logger.Logger, t *ast.Task) {
	l.Outf(logger.Default, "task: ")
	l.Outf(logger.Green, "%s\n", t.Name())
	l.Outf(logger.Default, "\n")
}

func printTaskAliases(l *logger.Logger, t *ast.Task) {
	if len(t.Aliases) == 0 {
		return
	}
	l.Outf(logger.Default, "\n")
	l.Outf(logger.Default, "aliases:\n")
	for _, alias := range t.Aliases {
		l.Outf(logger.Default, " - ")
		l.Outf(logger.Cyan, "%s\n", alias)
	}
}

func hasDescription(t *ast.Task) bool {
	return t.Desc != ""
}

func printTaskDescription(l *logger.Logger, t *ast.Task) {
	l.Outf(logger.Default, "%s\n", t.Desc)
}

func printNoDescriptionOrSummary(l *logger.Logger) {
	l.Outf(logger.Default, "(task does not have description or summary)\n")
}

func printTaskDependencies(l *logger.Logger, t *ast.Task) {
	if len(t.Deps) == 0 {
		return
	}

	l.Outf(logger.Default, "\n")
	l.Outf(logger.Default, "dependencies:\n")

	for _, d := range t.Deps {
		l.Outf(logger.Default, " - %s\n", d.Task)
	}
}

func printTaskCommands(l *logger.Logger, t *ast.Task) {
	if len(t.Cmds) == 0 {
		return
	}

	l.Outf(logger.Default, "\n")
	l.Outf(logger.Default, "commands:\n")
	for _, c := range t.Cmds {
		isCommand := c.Cmd != ""
		l.Outf(logger.Default, " - ")
		if isCommand {
			l.Outf(logger.Yellow, "%s\n", c.Cmd)
		} else {
			l.Outf(logger.Green, "Task: %s\n", c.Task)
		}
	}
}
