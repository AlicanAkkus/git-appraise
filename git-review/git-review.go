/*
Copyright 2015 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Command git-review manages code reviews stored as git-notes in the source repo.
//
// To install, run:
//
//    $ go get source.developers.google.com/id/0tH0wAQFren.git/src
//
// And for usage information, run:
//
//    $ git-review help
package main

import (
	"fmt"
	"os"
	"sort"
	"source.developers.google.com/id/0tH0wAQFren.git/commands"
	"source.developers.google.com/id/0tH0wAQFren.git/repository"
	"strings"
)

const usageMessageTemplate = `Usage: %s <command>

Where <command> is one of:
  %s

For individual command usage, run:
  %s help <command>
`

func usage() {
	command := os.Args[0]
	var subcommands []string
	for subcommand := range commands.CommandMap {
		subcommands = append(subcommands, subcommand)
	}
	sort.Strings(subcommands)
	fmt.Printf(usageMessageTemplate, command, strings.Join(subcommands, "\n  "), command)
}

func help() {
	if len(os.Args) < 3 {
		usage()
		return
	}
	subcommand, ok := commands.CommandMap[os.Args[2]]
	if !ok {
		fmt.Printf("Unknown command \"%s\"\n", os.Args[2])
		usage()
		return
	}
	subcommand.Usage(os.Args[0])
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	if os.Args[1] == "help" {
		help()
		return
	}
	if !repository.IsGitRepo() {
		fmt.Printf("%s must be run from within a git repo.", os.Args[0])
		return
	}
	subcommand, ok := commands.CommandMap[os.Args[1]]
	if !ok {
		fmt.Printf("Unknown command \"%s\"", os.Args[1])
		usage()
		return
	}
	if err := subcommand.Run(os.Args[2:]); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
