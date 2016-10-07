//models.articles_test.go

package main

import "testing"

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// Check that the length of fetched articles is
	// equal to the global var keeping track of articles
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}

func TestCreateNewArticle(t *testing.T) {
	originalLength := len(getAllArticles())

	a, err := createNewArticle("New test title", "New test content")

	allArticles := getAllArticles()
	newLength := len(allArticles)

	if err != nil || newLength != originalLength+1 || a.Title != "New test title" || a.Content != "New test content" {
		t.Fail()
	}
}
