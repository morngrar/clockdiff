// Copyright 2022 timediff Authors
// SPDX-License-Identifier: Apache-2.0

// Package timediff provides the Bonzai command branch of the same name.
package timediff

import (
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
)

func init() {
	Z.Conf.SoftInit()
	Z.Vars.SoftInit()
}

// Cmd provides a Bonzai branch command that can be composed into Bonzai
// trees or used as a standalone with light wrapper (see cmd/).
var Cmd = &Z.Cmd{

	Name:      `timediff`,
	Summary:   `a tool for aiding in time reporting`,
	Version:   `v0.2.4`,
	Copyright: `Copyright 2023 Svein-Kåre Bjørnsen`,
	License:   `Apache-2.0`,
	Source:    `git@github.com:morngrar/timediff.git`,
	Issues:    `github.com/morngrar/timediff/issues`,

	Commands: []*Z.Cmd{

		// standard external branch imports (see rwxrob/{help,conf,vars})
		help.Cmd, conf.Cmd, vars.Cmd,

		// local commands (in this module)
		Sub,
	},

	// WARNING: The Description will be dedented using the exact
	// runes from the beginning of the first line to the first
	// non-whitespace character. This means that mixing spaces and tabs in
	// your indentation will created unwanted truncation. Make sure each
	// line has the same indentation exactly.

	Description: `
        {{aka}} is a toolchain for simple calculation of clock time differences.
        see the help of each subcommand separately for usage info.
		`,
}

var Sub = &Z.Cmd{
	Name:    `subtract`,
	Usage:   `[START_TIME END_TIME | COMMAND]`,
	Aliases: []string{"s"},
	Summary: `subtracts two times and prints the number of hours between them`,
	Description: `
		The {{aka}} command either takes a subcommand or two arguments and
		executes its primary function. The only subcommand that is currently
		supported is this help.

		In regards to the primary function, the the first of the arguments is
		START_TIME and the second is END_TIME. The command prints the
		real-valued difference in hours between them.

		The format of the times is expected to be in 24-hour format, where the
		time 8 AM can be written in the following ways:

		    08:00 
		    8:00 
		    8.00 
		    8

		*NOTE*: If START_TIME is **higher** than END_TIME the command assumes
		that the START_TIME was during the previous day. Thus the difference
		will never be negative.  
	`,

	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, args ...string) error {
		if len(args) < 2 || len(args) > 2 {
			return x.UsageError()
		}

		var err error

		startTime := args[0]
		endTime := args[1]

		startTimeHours, err := ParseTime(startTime)
		if err != nil {
			return fmt.Errorf("Error parsing START_TIME: '%w'", err)
		}
		endTimeHours, err := ParseTime(endTime)
		if err != nil {
			return fmt.Errorf("Error parsing END_TIME: '%w'", err)
		}

		diff := endTimeHours - startTimeHours
		if diff < 0 {
			firstDay := 24 - startTimeHours
			diff = firstDay + endTimeHours
		}

		fmt.Printf("%.2f hours\n", diff)

		return nil
	},
}
