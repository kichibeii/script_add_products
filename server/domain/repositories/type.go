package repositories

// entity
type ProducstEntity struct {
	ID        int                    `json:"id"`
	Title     string                 `json:"title"`
	ShopifyId string                 `json:"shopify_id"`
	Variants  []ProductVariantEntity `json:"variants"`
	Images    []ProductImageEntity   `json:"images"`
}

type ProductVariantEntity struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Title     string `json:"title"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
}

type ProductImageEntity struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Link      string `json:"link"`
}

// end entity

// model
type ProducstModel struct {
	ID        int    `db:"id"`
	Title     string `db:"title"`
	ShopifyId string `db:"shopify_id"`
}

type ProductVariantModel struct {
	ID        int    `db:"id"`
	ProductID int    `db:"product_id"`
	Title     string `db:"title"`
	Price     int    `db:"price"`
	Quantity  int    `db:"quantity"`
}

type ProductImageModel struct {
	ID        int    `db:"id"`
	ProductID int    `db:"product_id"`
	Link      string `db:"link"`
}

type GetProductsData struct {
	ID              int    `db:"id"`
	Title           string `db:"title"`
	ShopifyID       string `db:"shopify_id"`
	ProductImages   string `db:"product_images"`
	ProductVariants string `db:"product_variants"`
}

// end model
