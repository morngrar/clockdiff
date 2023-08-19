// Copyright 2022 clockdiff Authors
// SPDX-License-Identifier: Apache-2.0

// Package example provides the Bonzai command branch of the same name.
package clockdiff

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

	Name:      `clockdiff`,
	Summary:   `a tool for aiding in time reporting`,
	Version:   `v0.0.1`,
	Copyright: `Copyright 2023 Svein-Kåre Bjørnsen`,
	License:   `Apache-2.0`,
	Source:    `git@github.com:morngrar/clockdiff.git`,
	Issues:    `github.com/morngrar/clockdiff/issues`,

	Commands: []*Z.Cmd{

		// standard external branch imports (see rwxrob/{help,conf,vars})
		help.Cmd, conf.Cmd, vars.Cmd,

		// local commands (in this module)
		sub,
	},

	// WARNING: The Description will be dedented using the exact
	// runes from the beginning of the first line to the first
	// non-whitespace character. This means that mixing spaces and tabs in
	// your indentation will created unwanted truncation. Make sure each
	// line has the same indentation exactly.

	Description: `
		Thing
		`,
}

// private leaf
var sub = &Z.Cmd{
	Name:    `subtract`,
	Aliases: []string{"s"},
	Summary: `subtracts two times and prints the number of hours between them`,
	Description: `
        The {{aka}} command takes two arguments, the first of which is 
        START_TIME and the second is END_TIME. The command prints the
        real-valued difference in hours between them.

        The format of the times is expected to be in 24-hour format, where the
        time 8 AM can be written in the following ways:

            08:00
            8:00
            8.00
            8
        `,

	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, args ...string) error {
		if len(args) < 2 || len(args) > 2 {
			return x.UsageError()
		}
		// TODO: validate arguments

		fmt.Printf("testing test %s", "asldkj")

		return nil
	},
}
