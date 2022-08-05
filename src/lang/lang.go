package lang

import "os"
import "log"
import "strings"

import "github.com/jeandeaual/go-locale"

// yet another package level hack
var globLang Language
func init() {
	// program args for language have preference over locale auto-detection
	globLang = UNSET
	for _, arg := range os.Args {
		switch arg {
		case "--en": globLang = EN
		case "--es": globLang = ES
		case "--ca": globLang = CA
		}
	}
	if globLang != UNSET { return }

	// detect locales
	locales, err := locale.GetLocales()
	if err != nil {
		log.Printf("Error while retrieving locales: %s", err.Error())
		globLang = ES
		return
	}

	// process locales and find best match
	for _, locale := range locales {
		langCode := strings.SplitN(locale, "-", 2)[0]
		switch langCode {
		case "en":
			globLang = EN
			return
		case "es":
			globLang = ES
			return
		case "ca":
			globLang = CA
			return
		}
	}

	globLang = EN
}

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
	UNSET Language = "unset"
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
