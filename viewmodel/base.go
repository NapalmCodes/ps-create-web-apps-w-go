package viewmodel

type Base struct {
	Title string
}

func NewBase() Base {
	return Base{
		Title: "Lemonade Stand Supply",
	}
}
