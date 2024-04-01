package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/script_add_products/server/commons"
)

type IRepository interface {
	GetAllData(ctx context.Context) ([]ProducstEntity, error)
	UpdateProduct(ctx context.Context, id int, shopifyId string) error
}

type Repository struct {
	opt commons.Options
}

func NewRepository(opt commons.Options) IRepository {
	return &Repository{
		opt: opt,
	}
}

func (r *Repository) GetAllData(ctx context.Context) ([]ProducstEntity, error) {
	// get products
	models := []GetProductsData{}

	if result := r.opt.Database.Raw(`
		SELECT 
		    p.id,
		    p.title,
		    p.shopify_id,
		    CONCAT('[', GROUP_CONCAT(DISTINCT CONCAT('{"link":"', pi.link, '","id":',pi.id, ',"product_id":',pi.product_id , '}') ORDER BY pi.id SEPARATOR ','), ']') AS product_images,
		    CONCAT('[', GROUP_CONCAT(DISTINCT CONCAT('{"product_id":', pv.product_id, ', "title":"', pv.title, '", "price":', pv.price, ', "quantity":', pv.quantity, '}') ORDER BY pv.product_id SEPARATOR ','), ']') AS product_variants
		FROM 
		    products p
		LEFT JOIN 
		    product_images pi ON p.id  = pi.product_id 
		LEFT JOIN 
		    product_variants pv ON p.id = pv.product_id
		WHERE 
			p.shopify_id = ''
		GROUP BY 
		    p.id
		`).Scan(&models); result.Error != nil {
		return []ProducstEntity{}, result.Error
	}

	return mappingDataModelsToEntity(models)
}

func (r *Repository) UpdateProduct(ctx context.Context, id int, shopifyId string) error {
	fmt.Println("the data", id, shopifyId)
	model := ProducstModel{
		ID: id,
	}

	if result := r.opt.Database.Table("products").Model(&model).UpdateColumn("shopify_id", shopifyId); result.Error != nil {
		return result.Error
	}

	return nil
}

func mappingDataModelsToEntity(models []GetProductsData) ([]ProducstEntity, error) {
	productsEntity := []ProducstEntity{}
	for _, model := range models {
		productImages := []ProductImageEntity{}
		err := json.Unmarshal([]byte(model.ProductImages), &productImages)
		if err != nil {
			return []ProducstEntity{}, err
		}

		productVariants := []ProductVariantEntity{}
		err = json.Unmarshal([]byte(model.ProductVariants), &productVariants)
		if err != nil {
			return []ProducstEntity{}, err
		}

		productsEntity = append(productsEntity, ProducstEntity{
			ID:        model.ID,
			Title:     model.Title,
			ShopifyId: model.ShopifyID,
			Images:    productImages,
			Variants:  productVariants,
		})
	}

	return productsEntity, nil
}
