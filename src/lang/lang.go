package lang

// TODO: use a package to autodetect language on init based on locale

// yet another package level hack
var lang Language = EN
func Current() Language { return lang }
func Set(language Language) { lang = language }

type Language string
const (
	EN Language = "en" // english
	ES Language = "es" // spanish / español
	CA Language = "ca" // catalan / català
)

func (self Language) Code() string { return string(self) }
func (self Language) Name() string {
	switch self {
	case EN: return "english"
	case ES: return "español"
	case CA: return "català"
	default:
		panic(self)
	}
}

// translate text to current language, or panic
// if text fragment can't be translated
func Tr(text string) string {
	switch lang.Code() {
	case "en":
		return text
	case "es":
		switch text {
		case "Back"       : return "Atrás"
		case "Start Game" : return "Empezar"
		case "Language"   : return "Idioma"
		case "Fullscreen" : return "Resolución"
		default:
			panic(text)
		}
	case "ca":
		switch text {
		case "Back"       : return "Enrere"
		case "Start Game" : return "Començar"
		case "Language"   : return "Idioma"
		case "Fullscreen" : return "Resolució"
		default:
			panic(text)
		}
	default:
		panic(lang)
	}
}
