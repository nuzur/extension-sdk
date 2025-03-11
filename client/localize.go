package client

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (c *Client) Localize(key string, lang string, defaultValue string) string {
	localizer := i18n.NewLocalizer(c.i18nbundle, lang)
	localized, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: key,
		DefaultMessage: &i18n.Message{
			ID:    key,
			Other: defaultValue,
		},
	})

	if err != nil {
		fmt.Printf("Error localizing key %s: %s\n", key, err)
		return defaultValue
	}
	return localized
}

func (c *Client) LocalizeWithVariables(key string, lang string, variables map[string]string, defaultValue string) string {
	localizer := i18n.NewLocalizer(c.i18nbundle, lang)
	localized, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: variables,
		DefaultMessage: &i18n.Message{
			ID:    key,
			Other: defaultValue,
		},
	})

	if err != nil {
		fmt.Printf("Error localizing key %s: %s\n", key, err)
		return defaultValue
	}
	return localized
}
