package httprequest

type CreateBasketRequest struct {
	BasketId string `json:"basket_id"`
}

type RemoveItemRequest struct {
	ProductID string `json:"product_id"`
}
