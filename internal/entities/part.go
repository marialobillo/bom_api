package entities

type Part struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Supplier_id string     `json:"supplier_id"`
	Price       float64    `json:"price"`
	Available   bool       `json:"available"`
	Description string     `json:"description"`
	Quantity    int        `json:"quantity"`
	Created_at   string     `json:"created_at"`
	Updated_at   string     `json:"updated_at"`
}