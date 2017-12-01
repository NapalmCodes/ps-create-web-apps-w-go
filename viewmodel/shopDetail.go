package viewmodel

type ShopDetail struct {
	Title    string
	Active   string
	Products []Product
}

func NewShopDetail() ShopDetail {
	var result ShopDetail
	result.Active = "shop"
	result.Title = "Lemonade Stand Supply - Juice Shop"
	lemonJuice := MakeLemonJuiceProduct()
	appleJuice := MakeAppleJuiceProduct()
	watermelonJuice := MakeWatermelonJuiceProduct()
	kiwiJuice := MakeKiwiJuiceProduct()
	mangosteenJuice := MakeMangosteenJuiceProduct()
	orangeJuice := MakeOrangeJuiceProduct()
	pineappleJuice := MakePineappleJuiceProduct()
	strawberryJuice := MakeStrawberryJuiceProduct()
	result.Products = []Product{lemonJuice, appleJuice, watermelonJuice,
		kiwiJuice, mangosteenJuice, orangeJuice, pineappleJuice, strawberryJuice}
	return result
}
