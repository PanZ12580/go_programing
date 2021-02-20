package main

import (
	"fmt"
	"go_programing/chapter4/json/github"
	"log"
)

func main() {
	res, err := github.SearchIssues([]string{"repo:golang/go", "is:open", "json", "decoder"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("total issues: %d\n", res.TotalCount)
	for _, item := range res.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	classify := github.Classify(res)
	for k, v := range classify {
		fmt.Printf("%s: \n", k)
		for _, item := range v {
			fmt.Printf("%.55s\n", item.Title)
		}
	}
}
