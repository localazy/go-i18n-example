# go-i18-example
![Golang Starter in Spanish](./src/assets/react-i18next-example.png)

This is an example repository showing how to use Localazy with [go-i18n](https://github.com/nicksnyder/go-i18n) for localization management.
To learn more, check out the [blog post](https://localazy.com/blog/how-to-localize-golang-app-with-go-i18n-and-localazy) describing the set up in detail.

## Sandbox
https://replit.com/@DanCharvt/go-i18n-example#main.go

## Testing out this repo
This repository contains translations from an actual [Localazy Golang i18next example](https://localazy.com/p/go-i18n-example) app. You can try to run `localazy download` to see how it works. However, this app is in read-only mode, meaning that `localazy upload` will fail. 

Run the application 
```
go run main.go
```

## Adjusting for your own project

- sign up for [Localazy](https://localazy.com/register),
- [create an app](https://localazy.com/my/create). Although English is recommended, you may use any language as source 
- select *Golang* integration option and install Localazy [CLI](https://testing.localazy.com/docs/cli/installation),
- Retrieve your _writeKey_ and _readKey_ from the integration guide page and change them in `localazy.keys.json`,
- it is recommended to add `localazy.keys.json` to _.gitignore_
- remove every locale file except for your source locale (presumably en.json)
- fill in your source phrases,
- run `localazy upload`,
- when you have accepted strings in other languages, run `localazy download` and check locales folder for the new translations,
-  run the app `go run main.go`
