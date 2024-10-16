package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruzhila/gin-blog/internal/common"
	"github.com/ruzhila/gin-blog/internal/models"
	"github.com/ruzhila/gin-blog/internal/themes"
	"gorm.io/gorm"
)

type Handlers struct {
	db    *gorm.DB
	Theme string
}

func NewHandlers(db *gorm.DB) *Handlers {
	return &Handlers{
		db:    db,
		Theme: "default",
	}
}

func (h *Handlers) Register(engine *gin.Engine) error {
	envs := models.GetEnvs()
	r := engine.Group(envs.Prefix)

	if envs.ConsolePrefix != "" {
		h.registerConsole(r, envs.ConsolePrefix)
	}

	if envs.AuthPrefix != "" {
		h.registerAuth(r, envs.AuthPrefix)
	}

	r.GET("/", h.handleIndexPage)
	r.GET("/post/:slug", h.handlePost)
	r.GET("/tags", h.handleTags)
	r.GET("/tag/:tag", h.handleTag)
	r.GET("/tag/:tag/:slug", h.handlePostWithTag)
	r.GET("/categories", h.handleCategories)
	r.GET("/category/:category", h.handleCategorie)
	r.GET("/category/:category/:slug", h.handlePostWithCategory)
	r.POST("/comment/:slug", h.handleComment)
	r.GET("/sitemap.xml", h.handleSitemap)

	templatesDir, _ := common.HintResouce("templates")
	r.StaticFS(envs.Static, themes.NewThemeFileSystem(templatesDir))
	return nil
}

func (h *Handlers) registerAuth(parent *gin.RouterGroup, path string) {
	r := parent.Group(path)
	r.GET("/signin", h.handleSignInPage)
	r.POST("/signin", h.handleSignIn)
	r.GET("/logout", h.handleLogout)
	r.GET("/signup", h.handleSignUpPage)
	r.POST("/signup", h.handleSignUp)
}

func (h *Handlers) registerConsole(parent *gin.RouterGroup, path string) {
	r := parent.Group(path, h.RequiredAuth)

	r.GET("/", h.handleConsoleIndexPage)
	r.GET("/setup", h.handleConsoleSetupPage)
	r.POST("/setup", h.handleConsoleSetup)

	// user, post, tags, categories CRUD
	r.GET("/users", h.handleConsoleUsers)
	r.PUT("/user/:id", h.handleConsoleCreateUser)
	r.PATCH("/user/:id", h.handleConsoleUpdateUser)
	r.DELETE("/user/:id", h.handleConsoleDeleteUser)

	r.GET("/posts", h.handleConsolePosts)
	r.PUT("/post/:id", h.handleConsoleCreatePost)
	r.PATCH("/post/:id", h.handleConsoleUpdatePost)
	r.DELETE("/post/:id", h.handleConsoleDeletePost)

	r.GET("/tags", h.handleConsoleTags)
	r.PUT("/tag/:id", h.handleConsoleCreateTag)
	r.PATCH("/tag/:id", h.handleConsoleUpdateTag)
	r.DELETE("/tag/:id", h.handleConsoleDeleteTag)

	r.GET("/categories", h.handleConsoleCategories)
	r.PUT("/category/:id", h.handleConsoleCreateCategory)
	r.PATCH("/category/:id", h.handleConsoleUpdateCategory)
	r.DELETE("/category/:id", h.handleConsoleDeleteCategory)

	r.GET("/comments", h.handleConsoleComments)
	r.PUT("/comment/:id", h.handleConsoleCreateComment)
	r.PATCH("/comment/:id", h.handleConsoleUpdateComment)
	r.DELETE("/comment/:id", h.handleConsoleDeleteComment)

	assetsPath, _ := common.HintResouce("console")
	r.StaticFS("/assets", http.Dir(assetsPath))
}
