package i18n

type I18n struct {
	defaultLanguage string
	translations    map[string]map[string]string
}

func NewI18n(defaultLanguage string) *I18n {
	return &I18n{
		defaultLanguage: defaultLanguage,
		translations:    make(map[string]map[string]string),
	}
}

func (i18n *I18n) AddTranslation(language, key, translation string) {
	if _, ok := i18n.translations[language]; !ok {
		i18n.translations[language] = make(map[string]string)
	}
	i18n.translations[language][key] = translation
}

func (i18n *I18n) Translate(key string, language ...string) string {
	lang := ""
	if len(language) > 0 {
		lang = language[0]
	}
	if lang == "" {
		lang = i18n.defaultLanguage
	}
	languageTranslations, ok := i18n.translations[lang]
	if ok {
		translation, ok := languageTranslations[key]
		if ok {
			return translation
		}
	}
	if lang == i18n.defaultLanguage {
		return key
	}
	return i18n.Translate(i18n.defaultLanguage, key)
}
