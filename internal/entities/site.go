package entities

type Site struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Location string `json:"location"`
	Address  string `json:"address"`
}