package viewmodel

type Home struct {
	Title  string
	Active string
}

func NewHome() Home {
	result := Home{
		Active: "home",
		Title:  "Lemonade Stand Supply",
	}
	return result
}
