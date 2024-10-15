package locales

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocaleLookup(t *testing.T) {
	{
		v := TR("site_name")
		assert.Equal(t, "site_name", v)
	}
	{
		v := TR("console.site_name")
		assert.Equal(t, "Site Name", v)
	}
	{
		SetDefaultLang("zh-CN")
		v := TR("console.site_name")
		assert.Equal(t, "网站名称", v)
		defaultLocales["console.test"] = "Test Name"
		v = TR("console.test")
		assert.Equal(t, "Test Name", v)
	}
}
