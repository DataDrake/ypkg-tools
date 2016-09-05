//
// Copyright © 2016 Ikey Doherty
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"flag"
	"github.com/ikeydoherty/ypkg-tools/ytools/cmd"
	"github.com/ikeydoherty/ypkg-tools/ytools/cmd/yauto"
	"os"
)

var cmds = make(map[string]*cmd.CMD)

func usage() {
	println("USAGE: ytools <tool> <tool_args...>")
	flag.PrintDefaults()
}

func registerCmds() {
	cmds["auto"] = yauto.GetCmd()
}

func handleArgs() *cmd.CMD {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}
	c := cmds[os.Args[1]]
	if c == nil {
		usage()
		os.Exit(1)
	}
	err := c.Flags.Parse(os.Args[2:])
	if err != nil {
		c.Flags.Usage()
		os.Exit(1)
	}
	return c
}

func run() {
	c := cmds[os.Args[1]]
	c.Run()
}

func main() {
	flag.Usage = usage
	registerCmds()
	handleArgs()
	run()
}
