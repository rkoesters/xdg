package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.Usage = usage
	countFlag.Usage = countUsage
	emptyFlag.Usage = emptyUsage
	eraseFlag.Usage = eraseUsage
	infoFlag.Usage = infoUsage
	lsFlag.Usage = lsUsage
	restoreFlag.Usage = restoreUsage
	rmFlag.Usage = rmUsage
}

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
	fmt.Fprintln(os.Stderr, "\trestore")
	fmt.Fprintln(os.Stderr, "\t\trestore a file from the trash")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "\trm")
	fmt.Fprintln(os.Stderr, "\t\tmove a file to the trash")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "For command help, run:", os.Args[0], "COMMAND -help")
}

func countUsage() {
	fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], countName, "[FLAGS]")
	fmt.Fprintln(os.Stderr, "FLAGS:")
	countFlag.PrintDefaults()
}

func emptyUsage() {
	fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], emptyName)
}

func eraseUsage() {
	fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], eraseName, "[FLAGS] FILE...")
	fmt.Fprintln(os.Stderr, "FLAGS:")
	eraseFlag.PrintDefaults()
}

func infoUsage() {
	fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], infoName, "[FLAGS] FILE...")
	fmt.Fprintln(os.Stderr, "FLAGS:")
	infoFlag.PrintDefaults()
}

func lsUsage() {
	fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], lsName)
}

func restoreUsage() {
	fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], restoreName, "FILE [DESTINATION]")
	fmt.Fprintln(os.Stderr, "      ", os.Args[0], restoreName, "FILE... DIRECTORY")
}

func rmUsage() {
	fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], rmName, "PATH...")
}
