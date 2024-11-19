package client

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (c *Client) Localize(key string, lang string) (string, error) {
	localizer := i18n.NewLocalizer(c.i18nbundle, lang)
	return localizer.Localize(&i18n.LocalizeConfig{
		MessageID: key,
	})
}
