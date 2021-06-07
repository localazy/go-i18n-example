package main

import (
	"encoding/json"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func initLocalizer(langs ...string) *i18n.Localizer {
	// Create a new i18n bundle with English as default language.
	bundle := i18n.NewBundle(language.English)

	// Register a json unmarshal function for i18n bundle.
	// This is to enable usage of json format
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Load source language
	bundle.LoadMessageFile("./locales/en.json")
	bundle.LoadMessageFile("./locales/es-419.json")
	bundle.LoadMessageFile("./locales/es.json")

	// Initialize localizer which will look for phrase keys in passed languages
	// in a strict order (first language is searched first)
	// When no key in any of the languages is found, it fallbacks to default - English language
	localizer := i18n.NewLocalizer(bundle, langs...)

	return localizer
}

func main() {
	localizer := initLocalizer(
		language.LatinAmericanSpanish.String(),
		language.Spanish.String(),
	)

	simpleMessage, _ := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "listen", // source key identifier
	})

	pluralMessage, _ := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{ID: "person"}, // another source key identifier
		PluralCount:    2,
	})

	fmt.Println(simpleMessage + ", " + pluralMessage)
}
