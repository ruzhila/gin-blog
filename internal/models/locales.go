package models

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var defaultLocales map[string]string
var currentLocales map[string]string

func init() {
	defaultLocales = LoadLocales("en")
	currentLocales = defaultLocales
}

func SetDefaultLang(lang string) {
	newLocales := LoadLocales(lang)
	if newLocales != nil {
		currentLocales = newLocales
	}
}

func LoadLocales(lang string) map[string]string {
	localeFile, hint := HintResouce("locales/" + lang + ".yaml")
	if !hint {
		return nil
	}
	defaultLocales = map[string]string{}

	data, err := os.ReadFile(localeFile)
	if err != nil {
		logrus.WithError(err).WithField("file", localeFile).Error("read locale file failed")
		return nil
	}
	var raw map[string]any
	if err := yaml.Unmarshal(data, &raw); err != nil {
		logrus.WithError(err).WithField("file", localeFile).Error("parse locale file failed")
		return nil
	}
	flattenMap("", raw, defaultLocales)
	return defaultLocales
}

func flattenMap(prefix string, nestedMap map[string]any, result map[string]string) {
	for key, value := range nestedMap {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}
		switch v := value.(type) {
		case map[any]any:
			stringMap := make(map[string]any)
			for k, val := range v {
				stringMap[k.(string)] = val
			}
			flattenMap(fullKey, stringMap, result)
		case map[string]any:
			flattenMap(fullKey, v, result)
		default:
			result[fullKey] = fmt.Sprintf("%v", v)
		}
	}
}

// T_ is a simple i18n function
func T_(id string) string {
	if v, ok := currentLocales[id]; ok {
		return v
	}
	if v, ok := defaultLocales[id]; ok {
		return v
	}
	return id
}
