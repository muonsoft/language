package language

import "golang.org/x/text/language"

type Tag = language.Tag

// Equal compares language tags by base ISO 639 language code.
func Equal(tag1, tag2 Tag) bool {
	base1, _, _ := tag1.Raw()
	base2, _, _ := tag2.Raw()

	return base1 == base2
}

var (
	Und = language.Und

	Afrikaans            = language.Afrikaans
	Amharic              = language.Amharic
	Arabic               = language.Arabic
	ModernStandardArabic = language.ModernStandardArabic
	Azerbaijani          = language.Azerbaijani
	Bulgarian            = language.Bulgarian
	Bengali              = language.Bengali
	Catalan              = language.Catalan
	Czech                = language.Czech
	Danish               = language.Danish
	German               = language.German
	Greek                = language.Greek
	English              = language.English
	AmericanEnglish      = language.AmericanEnglish
	BritishEnglish       = language.BritishEnglish
	Spanish              = language.Spanish
	EuropeanSpanish      = language.EuropeanSpanish
	LatinAmericanSpanish = language.LatinAmericanSpanish
	Estonian             = language.Estonian
	Persian              = language.Persian
	Finnish              = language.Finnish
	Filipino             = language.Filipino
	French               = language.French
	CanadianFrench       = language.CanadianFrench
	Gujarati             = language.Gujarati
	Hebrew               = language.Hebrew
	Hindi                = language.Hindi
	Croatian             = language.Croatian
	Hungarian            = language.Hungarian
	Armenian             = language.Armenian
	Indonesian           = language.Indonesian
	Icelandic            = language.Icelandic
	Italian              = language.Italian
	Japanese             = language.Japanese
	Georgian             = language.Georgian
	Kazakh               = language.Kazakh
	Khmer                = language.Khmer
	Kannada              = language.Kannada
	Korean               = language.Korean
	Kirghiz              = language.Kirghiz
	Lao                  = language.Lao
	Lithuanian           = language.Lithuanian
	Latvian              = language.Latvian
	Macedonian           = language.Macedonian
	Malayalam            = language.Malayalam
	Mongolian            = language.Mongolian
	Marathi              = language.Marathi
	Malay                = language.Malay
	Burmese              = language.Burmese
	Nepali               = language.Nepali
	Dutch                = language.Dutch
	Norwegian            = language.Norwegian
	Punjabi              = language.Punjabi
	Polish               = language.Polish
	Portuguese           = language.Portuguese
	BrazilianPortuguese  = language.BrazilianPortuguese
	EuropeanPortuguese   = language.EuropeanPortuguese
	Romanian             = language.Romanian
	Russian              = language.Russian
	Sinhala              = language.Sinhala
	Slovak               = language.Slovak
	Slovenian            = language.Slovenian
	Albanian             = language.Albanian
	Serbian              = language.Serbian
	SerbianLatin         = language.SerbianLatin
	Swedish              = language.Swedish
	Swahili              = language.Swahili
	Tamil                = language.Tamil
	Telugu               = language.Telugu
	Thai                 = language.Thai
	Turkish              = language.Turkish
	Ukrainian            = language.Ukrainian
	Urdu                 = language.Urdu
	Uzbek                = language.Uzbek
	Vietnamese           = language.Vietnamese
	Chinese              = language.Chinese
	SimplifiedChinese    = language.SimplifiedChinese
	TraditionalChinese   = language.TraditionalChinese
	Zulu                 = language.Zulu
)
