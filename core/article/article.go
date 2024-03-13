package core

import (
	"github.com/0xBradock/go-rw/domain"
)

type Article struct{}

func NewArticle() domain.ArticleService {
	return &Article{}
}

func (a *Article) GetArticlesFeed() ([]domain.Article, error) {
	return nil, nil
}

func (a *Article) GetArticles(tag domain.Tag, author domain.Author, favorited bool) ([]domain.Article, error) {
	return nil, nil
}

func (a *Article) GetArticle(slug string) (*domain.Article, error) {
	return nil, nil
}
