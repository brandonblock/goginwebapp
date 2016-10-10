//model.articles.go

package main

import (
	"errors"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Storing the article list in memory instead of DB, just 'cause
var articleList = []article{
	article{ID: 1, Title: "Love in the Time of the Apocalypse", Content: "A Homeric jaunt across the (u/dis)topic volleyball courts of our future."},
	article{ID: 2, Title: "The Land of Magical Thinking", Content: "Imps, mermaids, FDR, dryfly... All in a day's work for this government employee!"},
}

//Return a list of all articles
func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

func createNewArticle(title, content string) (*article, error) {
	a := article{ID: len(articleList) + 1, Title: title, Content: content}

	articleList = append(articleList, a)

	return &a, nil
}
