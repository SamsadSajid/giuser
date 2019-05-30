package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	reposEndPoint = "/repos"
)

type Params struct {
	Name        string `json:"name"`
	HTMLURL     string `json:"html_url"`
	Description string `json:"description"`
	GITURL      string `json:"git_url"`
	SSHURL      string `json:"ssh_url"`
	CLONEURL    string `json:"clone_url"`
	SVNURL      string `json:"svn_url"`
	CREATEDAT   string `json:"created_at"`
	SIZE        int    `json:"size"`
}

type Repos struct {
	Param []Params
}

func getAllRepos(name string) Repos {
	resp, err := http.Get(githubAPIURL + userEndpoint + name + reposEndPoint)
	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}

	repos := Repos{}
	json.Unmarshal(body, &repos)

	fmt.Println(repos)

	// for k, v := range repos {
	// 	for _, m := range v {
	// 		fmt.Println("", v[m])
	// 	}
	// }

	return repos
}
