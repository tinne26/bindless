package lang

// TODO: use a package to autodetect language on init based on locale

// yet another package level hack
var globLang Language = EN
func Current() Language { return globLang }
func Set(language Language) { globLang = language }

type Text struct {
	en string
	es string
	ca string
}
func NewText(en, es, ca string) *Text {
	if en == "" || es == "" || ca == "" {
		panic("can't create empty lang.Text")
	}
	return &Text{ en, es, ca }
}
func (self *Text) Get() string {
	switch globLang.Code() {
	case "en": return self.en
	case "es": return self.es
	case "ca": return self.ca
	default:
		panic(globLang)
	}
}
func (self *Text) English() string { return self.en }

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
	switch globLang.Code() {
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
		panic(globLang)
	}
}
