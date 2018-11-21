package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "COMMAND [FLAGS]")
	fmt.Fprintln(os.Stderr, "COMMANDS:")
	fmt.Fprintln(os.Stderr, "\tcount")
	fmt.Fprintln(os.Stderr, "\t\tprint the number of items in the trash")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "\tempty")
	fmt.Fprintln(os.Stderr, "\t\tempty the trash")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "\terase")
	fmt.Fprintln(os.Stderr, "\t\tdelete an item from the trash")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "\tinfo")
	fmt.Fprintln(os.Stderr, "\t\tshow information about an item in the trash")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "\tls")
	fmt.Fprintln(os.Stderr, "\t\tlist the items in the trash")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "\trm")
	fmt.Fprintln(os.Stderr, "\t\tmove a file to the trash")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "For command help, run:", os.Args[0], "COMMAND -help")
}

func countUsage() {
	fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], countName, "[FLAGS]")
	fmt.Fprintln(os.Stderr, "FLAGS:")
	countCommand.PrintDefaults()
}

func emptyUsage() {
	fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], emptyName, "[FLAGS]")
	fmt.Fprintln(os.Stderr, "FLAGS:")
	emptyCommand.PrintDefaults()
}

func eraseUsage() {
	fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], eraseName, "[FLAGS] FILE...")
	fmt.Fprintln(os.Stderr, "FLAGS:")
	eraseCommand.PrintDefaults()
}

func infoUsage() {
	fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], infoName, "[FLAGS] FILE...")
	fmt.Fprintln(os.Stderr, "FLAGS:")
	infoCommand.PrintDefaults()
}

func lsUsage() {
	fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], lsName, "[FLAGS]")
	fmt.Fprintln(os.Stderr, "FLAGS:")
	lsCommand.PrintDefaults()
}

func rmUsage() {
	fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], rmName, "[FLAGS] PATH...")
	fmt.Fprintln(os.Stderr, "FLAGS:")
	rmCommand.PrintDefaults()
}