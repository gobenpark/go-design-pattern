package another_example

import "fmt"

type ProductInfoRetriever interface {
	GetPrice() float32
	GetName() string
}

type Visitor interface {
	Visit(ProductInfoRetriever)
}

type Visitable interface {
	Accept(Visitor)
}

type Product struct {
	Price float32
	Name  string
}

func (p *Product) GetPrice() float32 {
	return p.Price
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) Accept(v Visitor) {
	v.Visit(p)
}

// ==========================
type Rice struct {
	Product
}

type Pasta struct {
	Product
}

type PriceVisitor struct {
	Sum float32
}

func (pv *PriceVisitor) Visit(p ProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

type NamePrinter struct {
	ProductList string
}

func (n *NamePrinter) Visit(p ProductInfoRetriever) {
	n.ProductList = fmt.Sprintf("%s\n%s", p.GetName(), n.ProductList)
}

func main() {
	products := make([]Visitable, 2)
	products[0] = &Rice{
		Product: Product{
			Price: 32.0,
			Name:  "Some rice",
		},
	}
	products[1] = &Pasta{
		Product: Product{
			Price: 40.0,
			Name:  "Some pasta",
		},
	}

	//Print the sum of prices
	priceVisitor := &PriceVisitor{}

	for _, p := range products {
		p.Accept(priceVisitor)
	}

	fmt.Printf("Total: %f\n", priceVisitor.Sum)

	//Print the products list
	nameVisitor := &NamePrinter{}

	for _, p := range products {
		p.Accept(nameVisitor)
	}

	fmt.Printf("\nProduct list:\n-------------\n%s", nameVisitor.ProductList)
}
