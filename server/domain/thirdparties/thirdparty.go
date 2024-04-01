package thirdparties

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/script_add_products/server/commons"
)

type IThirdParty interface {
	StagedUploadCreate(ctx context.Context) (StagedTarget, error)
	UploadFileJsonl(ctx context.Context, data StagedTarget, filePath string) error
	CreateBulk(ctx context.Context, filePath string) error
	Webhook(ctx context.Context) (string, string, error)
}

type ThirdParty struct {
	opt commons.Options
}

type WebhookResponse struct {
	Status string `json:"status"`
	Url    string `json:"url"`
}

func NewThirdParty(opt commons.Options) IThirdParty {
	return &ThirdParty{
		opt: opt,
	}
}

func (t *ThirdParty) StagedUploadCreate(ctx context.Context) (StagedTarget, error) {
	var response RootStagged

	payload := []byte(`{
		"query": "mutation stagedUploadsCreate($input: [StagedUploadInput!]!) { stagedUploadsCreate(input: $input) { stagedTargets { url resourceUrl parameters { name value } } } }",
		"variables": {
		   "input": [
			 {
			   "filename": "bulk_op_vars",
			   "mimeType": "text/jsonl",
			   "httpMethod": "POST",
			   "resource": "BULK_MUTATION_VARIABLES"
			 }
		   ]
		 }
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, t.opt.Config.Thirdparty.BaseUrl, bytes.NewBuffer(payload))
	if err != nil {
		return StagedTarget{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Shopify-Access-Token", t.opt.Config.Thirdparty.Key)

	res, err := client.Do(req)
	if err != nil {
		return StagedTarget{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return StagedTarget{}, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return StagedTarget{}, err
	}

	return response.Data.StagedUploadsCreate.StaggedTargets[0], nil
}

func (t *ThirdParty) UploadFileJsonl(ctx context.Context, data StagedTarget, filePath string) error {
	url := strings.ReplaceAll(data.URL, "com/", "com")

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return err
	}

	_, err = part.Write(fileContent)
	if err != nil {
		return err
	}

	for _, parameter := range data.Parameters {
		err = writer.WriteField(parameter.Name, parameter.Value)
		if err != nil {
			return err
		}
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Shopify-Access-Token", t.opt.Config.Thirdparty.Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyTarat, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("response upload data", string(bodyTarat))
	return nil
}

func (t *ThirdParty) CreateBulk(ctx context.Context, filePath string) error {
	var response ResponseBulkOperation
	query := fmt.Sprintf(`mutation {
		bulkOperationRunMutation(
			mutation: "mutation populateProduct($input:ProductInput!, $media:[CreateMediaInput!]){ productCreate(input:$input, media:$media){ product{ id title } userErrors{ field message } } }",
			stagedUploadPath: "%s"
		) {
			bulkOperation {
				id
				url
				status
			}
			userErrors {
				message
				field
			}
		}
	}`, filePath)

	data := map[string]string{"query": query}
	payload, _ := json.Marshal(data)

	req, err := http.NewRequest(http.MethodPost, t.opt.Config.Thirdparty.BaseUrl, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("X-Shopify-Access-Token", t.opt.Config.Thirdparty.Key)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	fmt.Println("RESPONSE SIH", response)

	return nil
}

func (t *ThirdParty) Webhook(ctx context.Context) (string, string, error) {
	payload := []byte(`{"query":" query {currentBulkOperation(type:MUTATION){ id status url }}"}`)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, t.opt.Config.Thirdparty.BaseUrl, bytes.NewBuffer(payload))

	if err != nil {
		return "", "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Shopify-Access-Token", t.opt.Config.Thirdparty.Key)

	res, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", "", err
	}

	var response ResponseWebhook

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", "", err
	}

	return response.Data.CurrentBulkOperation.Status, response.Data.CurrentBulkOperation.URL, err

}
