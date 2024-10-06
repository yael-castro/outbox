package command

import (
	"context"
)

const (
	successExitCode        = 0
	fatalExitCode          = 1
	incorrectUsageExitCode = 2
)

// Command alias for "func(context.Context) int"
type Command = func(context.Context) int
