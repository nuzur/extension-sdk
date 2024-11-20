package client

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (c *Client) Localize(key string, lang string, defaultValue string) string {
	localizer := i18n.NewLocalizer(c.i18nbundle, lang)
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: key,
		DefaultMessage: &i18n.Message{
			ID:    key,
			Other: defaultValue,
		},
	})
}

func (c *Client) LocalizeWithVariables(key string, lang string, variables map[string]string, defaultValue string) string {
	localizer := i18n.NewLocalizer(c.i18nbundle, lang)
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: variables,
		DefaultMessage: &i18n.Message{
			ID:    key,
			Other: defaultValue,
		},
	})
}
