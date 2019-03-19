package aviasales

type Locale struct {
	Code        string
	Description string
}

var (
	locales = []Locale{
		{"ar", "Arabic"},
		{"bg", "Bulgarian"},
		{"cs", "Czech"},
		{"da", "Danish"},
		{"de", "German"},
		{"el", "Greek"},
		{"en", "English"},
		{"es", "Spanish"},
		{"fa", "Persian"},
		{"fi", "Finnish"},
		{"fr", "French"},
		{"he", "Israeli"},
		{"hi", "Indian"},
		{"hr", "Croatian"},
		{"hu", "Hungarian"},
		{"id", "Indonesian"},
		{"it", "Italian"},
		{"ja", "Japanese"},
		{"ka", "Georgian"},
		{"ko", "Korean"},
		{"lt", "Italian"},
		{"lv", "Latvia"},
		{"ms", "Malaysian"},
		{"nl", "Dutch"},
		{"no", "Norwegian"},
		{"pl", "Polish"},
		{"pt", "Portuguese"},
		{"ro", "Romanian"},
		{"ru", "Russian"},
		{"sk", "Slovak"},
		{"sl", "Slovenian"},
		{"sr", "Serbian"},
		{"sv", "Swedish"},
		{"th", "Thai"},
		{"tl", "Filipino"},
		{"tr", "Turkish"},
		{"uk", "Ukrainian"},
		{"vi", "Vietnamese"},
		{"zh-Hans", "Chinese traditional"},
		{"zh-Hant", "Chinese simplified"},
	}
)
