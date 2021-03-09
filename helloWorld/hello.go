package helloWorld

const englishHelloPrefix = "Hello, "

var languageToGreetingsMap = map[string]string{
	"english": "Hello, ",
	"spanish": "Hola, ",
	"french":  "Bonjour, ",
}

func hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	greeting, ok := languageToGreetingsMap[language]

	if ok {
		return greeting + name
	} else {
		greeting := languageToGreetingsMap["english"]
		return greeting + name
	}
}

