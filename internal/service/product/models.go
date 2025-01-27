package product

var allProducts = []Product{
	{Title: "one"},
	{Title: "two"},
	{Title: "free"},
	{Title: "four"},
	{Title: "five"},
}

type Product struct {
	Title string
}
