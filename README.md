# 🌳 timediff - a tool for calculating clock time differences

[![GoDoc](https://godoc.org/github.com/morngrar/timediff?status.svg)](https://godoc.org/github.com/morngrar/timediff)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

## Usage

```
$ timediff
NAME
       timediff - a tool for aiding in time reporting

SYNOPSIS
       timediff COMMAND

COMMANDS
       help       - display help similar to man page format
       conf       - manage conf in /home/sk/.config/timediff/config.yaml
       var        - cache variables in /home/sk/.cache/timediff/vars
       s|subtract - subtracts two times and prints the number of hours between them

DESCRIPTION
       timediff is a toolchain for simple calculation of clock time differences. see the help of each
       subcommand separately for usage info.

CONTACT
       Source: git@github.com:morngrar/timediff.git
       Issues: github.com/morngrar/timediff/issues

LEGAL
       timediff (v0.2.2) Copyright 2023 Svein-Kåre Bjørnsen
       License Apache-2.0
```

```
$ timediff s help
NAME
       subtract - subtracts two times and prints the number of hours between them

ALIASES
       subtract (s)

SYNOPSIS
       subtract [COMMAND]

COMMANDS
       help - display help similar to man page format

DESCRIPTION
       The subtract (s) command takes two arguments, the first of which is START_TIME and the second is
       END_TIME. The command prints the real-valued difference in hours between them.

       The format of the times is expected to be in 24-hour format, where the time 8 AM can be written in
       the following ways:

           08:00 
           8:00 
           8.00 
           8

       NOTE: If START_TIME is higher than END_TIME the command assumes that the START_TIME was
       during the previous day. Thus the difference will never be negative.
```

## Install

### Prebuilt executable

Download the latest release for your platform from the [releases page](https://github.com/morngrar/timediff/releases), put it somewhere in your PATH. If you're running bash as your shell, you can also enable command tab completion by following the instruction in the [tab completion section](#tab-completion).

Every release is signed with gpg, the public key of which can be found [here](https://keyserver.ubuntu.com/pks/lookup?search=49564b7270e13347a20deee9327bdc9bf2b00d67&fingerprint=on&op=index).

### Bonzai style.

This command can be installed as a standalone program or composed into a
Bonzai command tree.

Standalone

```
go install github.com/morngrar/timediff/cmd/timediff@latest
```

Composed

```go
package z

import (
	Z "github.com/rwxrob/bonzai/z"
	timediff "github.com/morngrar/timediff"
)

var Cmd = &Z.Cmd{
	Name:     `z`,
	Commands: []*Z.Cmd{help.Cmd, timediff.Cmd},
}
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C timediff timediff
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.


## Other Examples

* <https://github.com/rwxrob/z> - the one that started it all
