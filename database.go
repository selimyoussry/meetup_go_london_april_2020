package mtp

import "strings"

func isGoArticle(article Article) bool {
	return strings.Contains(article.Title, "Go")
}
