package main

// ./exec4_10 repo:golang/go is:open json decoder

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

type categoryType int

const (
	LessThanOneMonth categoryType = iota
	LessThanOneYear
	MoreThanOneYear

	categoryTypeCount
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	now := time.Now()
	oneMonth := now.AddDate(0, -1, 0)
	oneYear := now.AddDate(-1, 0, 0)

	categories := map[categoryType][]*github.Issue{
		LessThanOneMonth: make([]*github.Issue, 0),
		LessThanOneYear:  make([]*github.Issue, 0),
		MoreThanOneYear:  make([]*github.Issue, 0),
	}

	for _, item := range result.Items {
		if item.CreatedAt.After(oneMonth) {
			categories[LessThanOneMonth] = append(categories[LessThanOneMonth], item)
		}
		if item.CreatedAt.After(oneYear) {
			categories[LessThanOneYear] = append(categories[LessThanOneYear], item)
		}
		if item.CreatedAt.Before(oneYear) {
			categories[MoreThanOneYear] = append(categories[MoreThanOneYear], item)
		}
	}

	for cateType := LessThanOneMonth; cateType < categoryTypeCount; cateType++ {
		fmt.Println("=======================================================")
		for _, item := range categories[cateType] {
			output(item)
		}
	}
}

func output(item *github.Issue) {
	fmt.Printf("#%-10d %-20s %-80s %s\n",
		item.Number, item.User.Login, item.Title, item.CreatedAt.String())
}
