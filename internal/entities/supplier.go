package entities

type Supplier struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Contact string `json:"contact"`
    Email   string `json:"email"`
	Address string `json:"address"`
}
