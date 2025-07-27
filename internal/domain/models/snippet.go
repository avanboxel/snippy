package models

type Snippet struct {
	Id       int
	Code     string
	Tags     []string
	Language string
}

func NewSnippet(code string, tags []string, lang string) *Snippet {
	return &Snippet{
		Code:     code,
		Tags:     tags,
		Language: lang,
	}
}
