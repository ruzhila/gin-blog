package models

import (
	"time"

	"github.com/ruzhila/gin-blog/internal/i18n"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Envs struct {
	Prefix        string `env:"PREFIX" comment:"Prefix for all routes"`
	ConsolePrefix string `env:"CONSOLE" comment:"Prefix for console"`
	AuthPrefix    string `env:"AUTH_PREFIX" comment:"Prefix for auth"`
	Static        string `env:"STATIC" comment:"Prefix for static files"`
}

func GetEnvs() *Envs {
	return &Envs{
		Prefix:        "/",
		Static:        "/static",
		ConsolePrefix: "/console",
		AuthPrefix:    "/auth",
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
	c := Config{Key: key, Value: value, Desc: desc}
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoNothing: true,
	}).Create(&c)
}

func GetConfigValue(db *gorm.DB, key string) (string, bool) {
	var c Config
	if err := db.Where("key", key).Take(&c).Error; err != nil {
		return "", false
	}
	return c.Value, true
}

func SetConfigValue(db *gorm.DB, key, value string) error {
	c := Config{Key: key, Value: value}
	r := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(&c)
	return r.Error
}

func GetConfigValues(db *gorm.DB) map[string]string {
	var cs []Config
	db.Find(&cs)
	m := make(map[string]string)
	for _, c := range cs {
		m[c.Key] = c.Value
	}
	return m
}

func CheckDefaultConfigValues(db *gorm.DB) {
	CheckConfigValue(db, Key_SiteName, "Gin blog", i18n.TR("console.site_name"))
	CheckConfigValue(db, Key_SiteLogo, "/logo.png", i18n.TR("console.site_logo"))
	CheckConfigValue(db, Key_SiteIcon, "/favicon.ico", i18n.TR("console.site_icon"))
	CheckConfigValue(db, Key_SiteUrl, "https://blog.ruzhil.cn", i18n.TR("console.site_url"))
	CheckConfigValue(db, Key_SiteAdmin, "Kui", i18n.TR("console.site_admin"))
	CheckConfigValue(db, Key_SiteAdminEmail, "kui@ruzhila.cn", i18n.TR("console.site_admin_email"))
	CheckConfigValue(db, Key_SiteTheme, "default", i18n.TR("console.site_theme"))
	CheckConfigValue(db, Key_SiteLang, "zh-CN", i18n.TR("console.site_lang"))
	CheckConfigValue(db, Key_SiteKeywords, "gin, blog", i18n.TR("console.site_keywords"))
	CheckConfigValue(db, Key_SiteDescription, "A blog system powered by gin+gorm, by ruzhila.cn", i18n.TR("console.site_description"))
	CheckConfigValue(db, Key_SiteCopyRight, "© 2024 ruzhila.cn", i18n.TR("console.site_copy_right"))
	CheckConfigValue(db, Key_SiteGA, "", i18n.TR("console.site_ga"))
	CheckConfigValue(db, Key_SiteICP, "", i18n.TR("console.site_icp"))
	CheckConfigValue(db, Key_AllowRegistion, "true", i18n.TR("console.allow_registion"))
}
