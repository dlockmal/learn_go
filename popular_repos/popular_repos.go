package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"
)

// Owner is the repository owner
type Owner struct {
	Login string
}

// Item is the single repository data structure
type Item struct {
	ID              int
	Name            string
	FullName        string `json:"full_name"`
	Owner           Owner
	Description     string
	CreatedAt       string `json:"created_at"`
	StargazersCount int    `json:"stargazers_count"`
}

// JSONData contains the GitHub API response
type JSONData struct {
	Count int `json:"total_count"`
	Items []Item
}

func main() {
	res, err := http.Get("https://api.github.com/search/repositories?q=stars:>=10000+language:go&sort=stars&order=desc")

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatal("Unexpected status code", res.StatusCode)
	}

	data := JSONData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
	printData(data)
}

func printData(data JSONData) {
	log.Printf("Repositories found: %d", data.Count)
	const format = "%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Repository", "Stars", "Created at", "Description")
	fmt.Fprintf(tw, format, "----------", "-----", "----------", "-----------")
	for _, i := range data.Items {
		desc := i.Description
		if len(desc) > 75 {
			desc = string(desc[:75]) + "..."
		}
		t, err := time.Parse(time.RFC3339, i.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(tw, format, i.FullName, i.StargazersCount, t.Year(), desc)
	}
	tw.Flush()
}
