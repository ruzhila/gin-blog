package models

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Envs struct {
	Prefix         string `env:"PREFIX" comment:"Prefix for all routes"`
	ConsolePrefix  string `env:"CONSOLE" comment:"Prefix for console"`
	AuthPrefix     string `env:"AUTH_PREFIX" comment:"Prefix for auth"`
	Static         string `env:"STATIC" comment:"Prefix for static files"`
	ThemePath      string `env:"THEME_PATH" comment:"Path to the theme"`
	AllowRegistion bool   `env:"ALLOW_REGISTION" comment:"Allow registion"`
}

func GetEnvs() *Envs {
	return &Envs{
		Prefix:         "/",
		Static:         "/static",
		ConsolePrefix:  "/console",
		AuthPrefix:     "/auth",
		ThemePath:      "themes/default",
		AllowRegistion: true,
	}
}

type Config struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Key       string    `json:"key" gorm:"size:200;unique"`
	Value     string    `json:"value"`
	Desc      string    `json:"desc" gorm:"size:2000"`
}

func CheckConfigValue(db *gorm.DB, key, value, desc string) {
	key = strings.ToTitle(key)
	c := Config{Key: key, Value: value, Desc: desc}
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoNothing: true,
	}).Create(&c)
}

func GetConfigValue(db *gorm.DB, key string) (string, bool) {
	key = strings.ToTitle(key)
	var c Config
	if err := db.Where("key", key).Take(&c).Error; err != nil {
		return "", false
	}
	return c.Value, true
}

func SetConfigValue(db *gorm.DB, key, value string) error {
	key = strings.ToTitle(key)
	c := Config{Key: key, Value: value}
	r := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(&c)
	return r.Error
}

func GetConfigValues(db *gorm.DB) []Config {
	var cs []Config
	db.Find(&cs)
	return cs
}

func CheckDefaultConfigValues(db *gorm.DB) {
	CheckConfigValue(db, Key_SiteName, "Gin blog", T_("console.site_name"))
	CheckConfigValue(db, Key_SiteLogo, "/logo.png", T_("console.site_logo"))
	CheckConfigValue(db, Key_SiteIcon, "/favicon.ico", T_("console.site_icon"))
	CheckConfigValue(db, Key_SiteUrl, "https://blog.ruzhil.cn", T_("console.site_url"))
	CheckConfigValue(db, Key_SiteAdmin, "Kui", T_("console.site_admin"))
	CheckConfigValue(db, Key_SiteAdminEmail, "kui@ruzhila.cn", T_("console.site_admin_email"))
	CheckConfigValue(db, Key_SiteTheme, "default", T_("console.site_theme"))
	CheckConfigValue(db, Key_SiteLang, "zh-CN", T_("console.site_lang"))
	CheckConfigValue(db, Key_SiteKeywords, "gin, blog", T_("console.site_keywords"))
	CheckConfigValue(db, Key_SiteDescription, "A blog system powered by gin+gorm, by ruzhila.cn", T_("console.site_description"))
	CheckConfigValue(db, Key_SiteCopyRight, "© 2024 ruzhila.cn", T_("console.site_copy_right"))
	CheckConfigValue(db, Key_SiteGA, "", T_("console.site_ga"))
	CheckConfigValue(db, Key_SiteICP, "", T_("console.site_icp"))
}

func HintResouce(p string) (string, bool) {
	for _, d := range []string{".", "..", "../.."} {
		d = filepath.Join(d, p)
		if _, err := os.Stat(d); err == nil {
			return d, true
		}
	}
	return p, false
}
