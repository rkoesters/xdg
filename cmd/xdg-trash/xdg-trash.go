package main

import (
	"flag"
	"fmt"
	"github.com/rkoesters/xdg/trash"
	"log"
	"os"
	"time"
)

var (
	countCommand = flag.NewFlagSet("count", flag.ExitOnError)
	emptyCommand = flag.NewFlagSet("empty", flag.ExitOnError)
	eraseCommand = flag.NewFlagSet("erase", flag.ExitOnError)
	infoCommand  = flag.NewFlagSet("info", flag.ExitOnError)
	lsCommand    = flag.NewFlagSet("ls", flag.ExitOnError)
	trashCommand = flag.NewFlagSet("trash", flag.ExitOnError)
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(os.Args[0] + ": ")

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 1 {
		log.Println("command required")
		flag.Usage()
		os.Exit(1)
	}

	// Parse the command.
	switch flag.Arg(0) {
	case countCommand.Name():
		countCommand.Parse(flag.Args()[1:])
	case emptyCommand.Name():
		emptyCommand.Parse(flag.Args()[1:])
	case eraseCommand.Name():
		eraseCommand.Parse(flag.Args()[1:])
	case infoCommand.Name():
		infoCommand.Parse(flag.Args()[1:])
	case lsCommand.Name():
		lsCommand.Parse(flag.Args()[1:])
	case trashCommand.Name():
		trashCommand.Parse(flag.Args()[1:])
	default:
		log.Printf("unknown command '%v'\n", flag.Arg(0))
		flag.Usage()
		os.Exit(1)
	}
	log.SetPrefix(flag.Arg(0) + ": ")

	// Run the command.
	switch {
	case countCommand.Parsed():
		if countCommand.NArg() != 0 {
			log.Fatal("count does not accept arguments")
		}

		files, err := trash.Files()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(len(files))
	case emptyCommand.Parsed():
		if emptyCommand.NArg() != 0 {
			log.Fatal("empty does not accept arguments")
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
			log.Fatal("ls does not accept arguments")
		}

		files, err := trash.Files()
		if err != nil {
			log.Fatal(err)
		}

		for _, fname := range files {
			fmt.Println(fname)
		}
	case trashCommand.Parsed():
		if trashCommand.NArg() == 0 {
			trashCommand.Usage()
			os.Exit(1)
		}

		for _, path := range trashCommand.Args() {
			err := trash.Trash(path)
			if err != nil {
				log.Fatal(err)
			}
		}
	default:
		panic("command parsed but not run")
	}
}
