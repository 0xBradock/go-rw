package domain

type ArticleRepo interface {
	// GetArticlesFeed return recent articles from users you follow
	//
	// Get most recent articles from users you follow. Use query parameters to limit. Auth is required
	GetArticlesFeed() ([]Article, error)

	// GetArticles return recent articles globally
	//
	// Get most recent articles globally. Use query parameters to filter results. Auth is optional
	GetArticles(tag Tag, author Author, favorited bool) ([]Article, error)

	// GetArticle returns an article
	//
	// Get an article. Auth not required
	GetArticle(slug string) (*Article, error)
}
