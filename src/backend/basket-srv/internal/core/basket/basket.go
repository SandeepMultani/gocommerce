package basket

type Basket struct {
	ID    string       `json:"id"`
	Items []BasketItem `json:"items"`
}

type BasketItem struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}

type BasketService interface {
	Get(string) (*Basket, error)
	Create(string) (*Basket, error)
	Delete(string) error
	AddItem(string, string) (*Basket, error)
	RemoveItem(string, string) (*Basket, error)
}

type BasketRepository interface {
	Get(string) (*Basket, error)
	Create(string) (*Basket, error)
	Delete(string) error
	AddItem(string, string) (*Basket, error)
	RemoveItem(string, string) (*Basket, error)
}
