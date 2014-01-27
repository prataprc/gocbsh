package commands

import (
    "github.com/prataprc/cbsh/api"
)

var knownCommands = map[string]api.CommandHandler{}

func Allcommands() map[string]api.CommandHandler {
    return knownCommands
}

