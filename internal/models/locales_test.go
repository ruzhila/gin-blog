package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocaleLookup(t *testing.T) {
	{
		v := T_("site_name")
		assert.Equal(t, "site_name", v)
	}
	{
		v := T_("console.site_name")
		assert.Equal(t, "Site Name", v)
	}
	{
		SetDefaultLang("zh-CN")
		v := T_("console.site_name")
		assert.Equal(t, "网站名称", v)
		defaultLocales["console.test"] = "Test Name"
		v = T_("console.test")
		assert.Equal(t, "Test Name", v)
	}
}
