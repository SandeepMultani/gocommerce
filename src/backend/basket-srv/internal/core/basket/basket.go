package basket

type Basket struct {
	ID         string                `json:"id"`
	Items      map[string]BasketItem `json:"items"`
	Total      float32               `json:"total"`
	UpsertedAt int64                 `json:"upserted_at"`
}

type BasketItem struct {
	ProductID   string  `json:"product_id"`
	Sku         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
	Total       float32 `json:"total"`
}

type Product struct {
	ProductID   string  `json:"product_id"`
	Sku         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type BasketService interface {
	Get(string) (*Basket, error)
	Create(string) (*Basket, error)
	Delete(string) error
	AddItem(string, *Product) (*Basket, error)
	RemoveItem(string, string) (*Basket, error)
}

type BasketRepository interface {
	Get(string) (*Basket, error)
	Create(*Basket) error
	Delete(string) error
	Update(*Basket) error
}
