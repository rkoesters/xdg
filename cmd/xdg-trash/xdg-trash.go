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
	countName = "count"
	emptyName = "empty"
	eraseName = "erase"
	infoName  = "info"
	lsName    = "ls"
	rmName    = "rm"
)

var (
	countCommand = flag.NewFlagSet(countName, flag.ExitOnError)
	emptyCommand = flag.NewFlagSet(emptyName, flag.ExitOnError)
	eraseCommand = flag.NewFlagSet(eraseName, flag.ExitOnError)
	infoCommand  = flag.NewFlagSet(infoName, flag.ExitOnError)
	lsCommand    = flag.NewFlagSet(lsName, flag.ExitOnError)
	rmCommand    = flag.NewFlagSet(rmName, flag.ExitOnError)
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

	// Parse the command.
	switch flag.Arg(0) {
	case countName:
		countCommand.Parse(flag.Args()[1:])
	case emptyName:
		emptyCommand.Parse(flag.Args()[1:])
	case eraseName:
		eraseCommand.Parse(flag.Args()[1:])
	case infoName:
		infoCommand.Parse(flag.Args()[1:])
	case lsName:
		lsCommand.Parse(flag.Args()[1:])
	case rmName:
		rmCommand.Parse(flag.Args()[1:])
	default:
		log.Printf("unknown command '%v'\n", flag.Arg(0))
		flag.Usage()
		os.Exit(1)
	}
	log.SetPrefix(os.Args[0] + " " + flag.Arg(0) + ": ")

	// Run the command.
	switch {
	case countCommand.Parsed():
		if countCommand.NArg() != 0 {
			log.Print("count does not accept arguments")
			countCommand.Usage()
			os.Exit(1)
		}

		files, err := trash.Files()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(len(files))
	case emptyCommand.Parsed():
		if emptyCommand.NArg() != 0 {
			log.Print("empty does not accept arguments")
			emptyCommand.Usage()
			os.Exit(1)
		}

		err := trash.Empty()
		if err != nil {
			log.Fatal(err)
		}
	case eraseCommand.Parsed():
		if eraseCommand.NArg() == 0 {
			eraseCommand.Usage()
			os.Exit(1)
		}

		for _, file := range eraseCommand.Args() {
			err := trash.Erase(file)
			if err != nil {
				log.Fatal(err)
			}
		}
	case infoCommand.Parsed():
		if infoCommand.NArg() == 0 {
			infoCommand.Usage()
			os.Exit(1)
		}

		for _, file := range infoCommand.Args() {
			info, err := trash.Stat(file)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("File:", file)
			fmt.Println("Original Path:", info.Path)
			fmt.Println("Deletion Date:", info.DeletionDate.Format(time.RFC822))
		}
	case lsCommand.Parsed():
		if lsCommand.NArg() != 0 {
			log.Print("ls does not accept arguments")
			lsCommand.Usage()
			os.Exit(1)
		}

		files, err := trash.Files()
		if err != nil {
			log.Fatal(err)
		}

		for _, fname := range files {
			fmt.Println(fname)
		}
	case rmCommand.Parsed():
		if rmCommand.NArg() == 0 {
			rmCommand.Usage()
			os.Exit(1)
		}

		for _, path := range rmCommand.Args() {
			err := trash.Trash(path)
			if err != nil {
				log.Fatal(err)
			}
		}
	default:
		panic("command parsed but not run")
	}
}
