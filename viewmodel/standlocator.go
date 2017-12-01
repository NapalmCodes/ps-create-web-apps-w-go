package viewmodel

type StandLocator struct {
	Title  string
	Active string
}

func NewStandLocator() StandLocator {
	result := StandLocator{
		Active: "standlocator",
		Title:  "Lemonade Stand Supply - Stand Locator",
	}
	return result
}

type StandCoordinate struct {
	Title     string  `json:"title"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lng"`
}
