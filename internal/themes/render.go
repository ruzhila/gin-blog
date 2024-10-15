package themes

import (
	"fmt"

	"github.com/gin-gonic/gin/render"
	"github.com/ruzhila/gin-blog/internal/common"
	"gitlab.com/go-box/pongo2gin/v6"
)

func NewRender() (render.HTMLRender, error) {
	templatesDir, hint := common.HintResouce("templates")
	if !hint {
		return nil, fmt.Errorf("templates path %s not found", "templates")
	}
	return pongo2gin.New(pongo2gin.RenderOptions{
		TemplateDir: templatesDir,
		TemplateSet: nil,
		ContentType: "text/html; charset=utf-8",
	}), nil
}
