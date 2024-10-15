package handlers

import (
	"github.com/flosch/pongo2/v6"
	"github.com/ruzhila/gin-blog/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func WithFuncs(db *gorm.DB, c pongo2.Context) pongo2.Context {
	c["posts"] = map[string]any{
		"query": func(offset, limit int) []models.Post {
			r, err := models.GetPosts(db, offset, limit)
			if err != nil {
				logrus.WithField("offset", offset).WithField("limit", limit).WithError(err).Error("handlers: get posts failed")
			}
			return r
		},
		"total": func() int64 {
			return models.CountPosts(db)
		},
	}
	c["categories"] = map[string]any{
		"query": func(parent *uint) []models.Category {
			r, err := models.GetCategories(db, parent)
			if err != nil {
				logrus.WithField("parent", parent).WithError(err).Error("handlers: get categories failed")
			}
			return r
		},
		"total": func(parent *uint) int64 {
			return models.CountCategories(db, parent)
		},
	}
	return c
}
