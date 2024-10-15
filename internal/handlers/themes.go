package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var NotAllowAccessExts = map[string]bool{
	".exe": true,
	".tpl": true,
	".env": true,
}

type ThemeFileSystem struct {
	Root string
}

func NewThemeFileSystem(themePath string) *ThemeFileSystem {
	return &ThemeFileSystem{
		Root: themePath,
	}
}

func (fs *ThemeFileSystem) Open(name string) (http.File, error) {
	ext := strings.ToLower(filepath.Ext(name))
	if NotAllowAccessExts[ext] {
		return nil, os.ErrNotExist
	}
	return http.Dir(fs.Root).Open(name)
}
