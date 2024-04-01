package services

type DataFromURL struct {
	Data DataURL `json:"data"`
}

type DataURL struct {
	ProductCreate ProductCreate `json:"productCreate"`
}

type ProductCreate struct {
	Product Product `json:"product"`
}

type Product struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type InputData struct {
	Input Input   `json:"input"`
	Media []Media `json:"media"`
}

type Input struct {
	Title    string    `json:"title"`
	Variants []Variant `json:"variants"`
}

type Variant struct {
	Price  string   `json:"price"`
	Option []string `json:"options"`
}

type Media struct {
	MediaContentType string `json:"mediaContentType"`
	OriginalSource   string `json:"originalSource"`
}
