package main

import "snippetbox.wes-kibagendi/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
