package main

import (
	"fmt"
	"os"
	"strings"
)

func handleUsers() {
	user := strings.Replace(user, "=", "", -1)
	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
	for _, u := range users {
		result := getUsers(u)
		if result.Login == "" {
			fmt.Println("No user found with the given username. Check spelling")
			os.Exit(1)
		}
		print("Username:		", result.Login, -1)

		if result.Name != "" {
			print(`Name:			`, result.Name, -1)
		}
		if result.Email != "" {
			print(`Email:			`, result.Email, -1)
		}
		if result.Bio != "" {
			print(`Bio:			`, result.Bio, -1)
		}
		if result.HTMLURL != "" {
			print(`Url:			`, result.HTMLURL, -1)
		}
		if result.Location != "" {
			print(`Location:		`, result.Location, -1)
		}
		print(`Public gists		`, "", result.PublicGists)
		print(`Public repositories	`, "", result.PublicRepos)
		print(`Followers		`, "", result.Followers)
		print(`Following		`, "", result.Following)

		fmt.Println("")
	}
}
