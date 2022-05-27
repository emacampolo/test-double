package order

type Product int

const (
	Book Product = iota
)

type Warehouse interface {
	HasInventory(p Product) bool
	Remove(p Product, qty int) error
}

type Order struct {
	Product  Product
	Quantity int

	isFilled bool
}

func (o *Order) IsFilled() bool {
	return o.isFilled
}

func (o *Order) Fill(w Warehouse) {
	if !w.HasInventory(o.Product) {
		return
	}

	err := w.Remove(o.Product, o.Quantity)
	if err == nil {
		o.isFilled = true
	}
}
