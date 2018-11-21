// xdg-trash is a command line interface for the trash package for the
// purposes of debugging.
package main

import (
	"flag"
	"fmt"
	"github.com/rkoesters/xdg/trash"
	"log"
	"os"
	"time"
)

const (
	countName   = "count"
	emptyName   = "empty"
	eraseName   = "erase"
	infoName    = "info"
	lsName      = "ls"
	restoreName = "restore"
	rmName      = "rm"
)

var (
	countFlag   = flag.NewFlagSet(countName, flag.ExitOnError)
	emptyFlag   = flag.NewFlagSet(emptyName, flag.ExitOnError)
	eraseFlag   = flag.NewFlagSet(eraseName, flag.ExitOnError)
	infoFlag    = flag.NewFlagSet(infoName, flag.ExitOnError)
	lsFlag      = flag.NewFlagSet(lsName, flag.ExitOnError)
	restoreFlag = flag.NewFlagSet(restoreName, flag.ExitOnError)
	rmFlag      = flag.NewFlagSet(rmName, flag.ExitOnError)
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(os.Args[0] + ": ")

	flag.Parse()

	if flag.NArg() < 1 {
		log.Println("command required")
		flag.Usage()
		os.Exit(1)
	}

	log.SetPrefix(os.Args[0] + " " + flag.Arg(0) + ": ")

	switch flag.Arg(0) {
	case countName:
		countFlag.Parse(flag.Args()[1:])
		countMain()
	case emptyName:
		emptyFlag.Parse(flag.Args()[1:])
		emptyMain()
	case eraseName:
		eraseFlag.Parse(flag.Args()[1:])
		eraseMain()
	case infoName:
		infoFlag.Parse(flag.Args()[1:])
		infoMain()
	case lsName:
		lsFlag.Parse(flag.Args()[1:])
		lsMain()
	case restoreName:
		restoreFlag.Parse(flag.Args()[1:])
		restoreMain()
	case rmName:
		rmFlag.Parse(flag.Args()[1:])
		rmMain()
	default:
		log.SetPrefix(os.Args[0] + ": ")
		log.Printf("unknown command '%v'\n", flag.Arg(0))
		flag.Usage()
		os.Exit(1)
	}
}

var (
	countQuiet = countFlag.Bool("q", false, "suppress output, set exit status to count")
)

func countMain() {
	if countFlag.NArg() != 0 {
		log.Print("count does not accept arguments")
		countFlag.Usage()
		os.Exit(1)
	}

	files, err := trash.Files()
	if err != nil {
		log.Fatal(err)
	}

	if *countQuiet {
		os.Exit(len(files))
	} else {
		fmt.Println(len(files))
	}
}

func emptyMain() {
	if emptyFlag.NArg() != 0 {
		log.Print("empty does not accept arguments")
		emptyFlag.Usage()
		os.Exit(1)
	}

	err := trash.Empty()
	if err != nil {
		log.Fatal(err)
	}
}

var (
	eraseRecursive = eraseFlag.Bool("r", false, "recursively erase")
)

func eraseMain() {
	if eraseFlag.NArg() == 0 {
		eraseFlag.Usage()
		os.Exit(1)
	}

	for _, file := range eraseFlag.Args() {
		var err error

		if *eraseRecursive {
			err = trash.EraseAll(file)
		} else {
			err = trash.Erase(file)
		}

		if err != nil {
			log.Fatal(err)
		}
	}
}

var (
	infoCompact = infoFlag.Bool("1", false, "one file info per line")
)

func infoMain() {
	if infoFlag.NArg() == 0 {
		infoFlag.Usage()
		os.Exit(1)
	}

	for _, file := range infoFlag.Args() {
		info, err := trash.Stat(file)
		if err != nil {
			log.Fatal(err)
		}

		if *infoCompact {
			fmt.Printf("%v:%v:%v\n", file, info.Path, info.DeletionDate.Format(time.RFC3339))
		} else {
			fmt.Println("File:", file)
			fmt.Println("Original Path:", info.Path)
			fmt.Println("Deletion Date:", info.DeletionDate.Format(time.RFC822))
		}
	}
}

func lsMain() {
	if lsFlag.NArg() != 0 {
		log.Print("ls does not accept arguments")
		lsFlag.Usage()
		os.Exit(1)
	}

	files, err := trash.Files()
	if err != nil {
		log.Fatal(err)
	}

	for _, fname := range files {
		fmt.Println(fname)
	}
}

func restoreMain() {
	if restoreFlag.NArg() == 0 {
		restoreFlag.Usage()
		os.Exit(1)
	}

	for _, file := range restoreFlag.Args() {
		err := trash.Restore(file)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func rmMain() {
	if rmFlag.NArg() == 0 {
		rmFlag.Usage()
		os.Exit(1)
	}

	for _, path := range rmFlag.Args() {
		err := trash.Trash(path)
		if err != nil {
			log.Fatal(err)
		}
	}
}
