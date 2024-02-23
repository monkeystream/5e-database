package main

import (
	"context"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

type translator struct {
	c *translate.Client
}

func newTranslator(key string) *translator {
	c, err := translate.NewClient(context.Background(), option.WithAPIKey(key))
	if err != nil {
		panic(err)
	}

	return &translator{
		c: c,
	}
}

func (t *translator) Translate(s []string) []string {
	if len(s) == 0 {
		return nil
	} else if len(s) == 1 && s[0] == "" {
		return []string{""}
	}

	res, err := t.c.Translate(context.Background(), s, language.Russian, &translate.Options{
		Source: language.English,
	})
	if err != nil {
		panic(err)
	}

	for i, text := range res {
		s[i] = text.Text
	}

	return s
}

func (t *translator) close() {
	t.c.Close()
}
