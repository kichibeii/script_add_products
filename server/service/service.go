package services

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/script_add_products/server/commons"
	"github.com/script_add_products/server/domain/repositories"
	"github.com/script_add_products/server/domain/thirdparties"
)

type ISyncProductService interface {
	Run(ctx context.Context) error
}

type SyncProductService struct {
	opt        commons.Options
	repo       repositories.IRepository
	thirdParty thirdparties.IThirdParty
}

func NewSyncProductService(opt commons.Options, repo repositories.IRepository, thirdParty thirdparties.IThirdParty) ISyncProductService {
	return &SyncProductService{
		opt:        opt,
		repo:       repo,
		thirdParty: thirdParty,
	}
}

func (svc *SyncProductService) Run(ctx context.Context) error {
	fmt.Println("get all products from databases")
	// get data product that dont have shopify_id
	products, err := svc.repo.GetAllData(ctx)
	if err != nil {
		return err
	}

	// preparing data
	// create a file
	fmt.Println("preparing data")
	pathFile, mapDataProducts, err := preparingData(products)
	if err != nil {
		return err
	}

	fmt.Println("process to insert to shopify")
	// create a bulkd
	err = svc.ProcessBulkInsert(ctx, pathFile)
	if err != nil {
		return err
	}

	url := ""
	status := ""

	for {
		time.Sleep(10 * time.Second)

		fmt.Println("checking loop")
		status, url, err = svc.thirdParty.Webhook(ctx)
		if err != nil {
			return err
		}

		if strings.ToLower(status) == "completed" {
			fmt.Println("webhook is ready")
			break
		}
	}

	fmt.Println("processing data from url response")
	mappingInsertedData, err := getDataFromLink(url)
	if err != nil {
		return err
	}

	fmt.Println("updating products")
	for key, title := range mappingInsertedData {
		if val, ok := mapDataProducts[title]; ok {
			svc.repo.UpdateProduct(ctx, val, key)
		}
	}

	return nil
}

func getDataFromLink(url string) (map[string]string, error) {
	mapProducts := make(map[string]string)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		line := scanner.Text()

		var dataUrl DataFromURL
		err = json.Unmarshal([]byte(line), &dataUrl)
		if err != nil {
			return nil, err
		}

		mapProducts[dataUrl.Data.ProductCreate.Product.Id] = dataUrl.Data.ProductCreate.Product.Title
	}

	return mapProducts, nil
}

func (svc *SyncProductService) ProcessBulkInsert(ctx context.Context, pathFile string) error {
	fmt.Println("reserve a link to upload file")
	// reserve a link
	staggedTarget, err := svc.thirdParty.StagedUploadCreate(ctx)
	if err != nil {
		return err
	}

	fmt.Println("upload data")
	err = svc.thirdParty.UploadFileJsonl(ctx, staggedTarget, pathFile)
	if err != nil {
		return err
	}

	// manual filling key for debugging
	key := ""
	for _, data := range staggedTarget.Parameters {
		if data.Name == "key" {
			key = data.Value
		}
	}

	fmt.Println("creating bulk operation", key)
	err = svc.thirdParty.CreateBulk(ctx, key)
	if err != nil {
		return err
	}

	return nil
}

func preparingData(produts []repositories.ProducstEntity) (string, map[string]int, error) {
	// preparing map
	mapDataProducts := make(map[string]int)

	// Specify the file path
	filePath := "data.json"

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	// Create a bufio.Writer
	writer := bufio.NewWriter(file)

	for _, data := range produts {
		// mappping data into file
		dataFile := InputData{}
		dataFile.Input.Title = data.Title
		variants := []Variant{}
		for _, variant := range data.Variants {
			variants = append(variants, Variant{
				Price:  fmt.Sprintf("%d", variant.Price),
				Option: []string{variant.Title},
			})
		}
		dataFile.Input.Variants = variants

		images := []Media{}
		for _, image := range data.Images {
			images = append(images, Media{
				MediaContentType: "IMAGE",
				OriginalSource:   image.Link,
			})
		}
		dataFile.Media = images

		marshalData, err := json.Marshal(dataFile)
		if err != nil {
			return "", nil, err
		}

		_, err = writer.WriteString(string(marshalData) + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return "", nil, err
		}

		// Flush the buffer
		err = writer.Flush()
		if err != nil {
			fmt.Println("Error flushing buffer:", err)
			return "", nil, err
		}

		// mapping data products to mapDataProducts
		mapDataProducts[data.Title] = data.ID
	}

	return filePath, mapDataProducts, nil
}
