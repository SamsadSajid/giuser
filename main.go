package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/ogier/pflag"
)

var (
	user     string
	emoji    string
	language string
	stars    string
)

func handleReposWithLanguage() {
	language := strings.Replace(language, "=", "", -1)
	stars := strings.Replace(stars, "=", "", -1)
	fmt.Printf("Searching repositories with language: %s with stars %s\n", language, stars)

	searchReposWithLanguage(language, stars)
}

func main() {
	// parse flags
	flag.Parse()

	numberOfFlags := flag.NFlag()

	switch numberOfFlags {
	// if user does not supply flags, print usage
	case 0:
		printLogo()
		printUsage()
	case 1:
		if user != "" {
			handleUsers()
		} else if language != "" {
			handleReposWithLanguage()
		}
	case 2:
		// -l -s
		if language != "" {
			handleReposWithLanguage()
		}
	}
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users with username")
	flag.StringVarP(&language, "language", "l", "", "Search repositories with language")
	flag.StringVarP(&stars, "stars", "s", "50", "Filters repositories with number of stars")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
