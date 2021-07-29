package product

type Product struct {
	ID          string  `json:"id"`
	Sku         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type ProductRepository interface {
	Get(string) (*Product, error)
	GetBySku(string) (*Product, error)
}

type ProductService interface {
	Get(string) (*Product, error)
	GetBySku(string) (*Product, error)
}
