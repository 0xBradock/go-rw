package domain

type ArticleService interface {
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

type Author string

type Tag string

type User struct {
	Email    string
	Token    string
	Username string
	Bio      string
	Image    string
}

type Profile struct {
	Username  string
	Bio       string
	Image     string
	Following bool
}

type Article struct {
	Slug           string
	Title          string
	Description    string
	Body           string
	CreatedAt      string
	UpdatedAt      string
	TagList        []string
	Favorited      bool
	FavoritesCount uint
	Author         Profile
}
