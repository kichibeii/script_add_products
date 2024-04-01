package thirdparties

type RootStagged struct {
	Data DataStagged `json:"data"`
}

type DataStagged struct {
	StagedUploadsCreate StagedUpploadCreate `json:"stagedUploadsCreate"`
}

type StagedUpploadCreate struct {
	StaggedTargets []StagedTarget `json:"stagedTargets"`
}

type StagedTarget struct {
	URL         string         `json:"url"`
	ResourceURL string         `json:"resourceUrl"`
	Parameters  []TargetParams `json:"parameters"`
}

type TargetParams struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ResponseWebhook struct {
	Data Data `json:"data"`
}

type Data struct {
	CurrentBulkOperation CurrentBulkOperation `json:"currentBulkOperation"`
}

type CurrentBulkOperation struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

type ResponseBulkOperation struct {
	Data DataBulkOperation `json:"data"`
}

type DataBulkOperation struct {
	BulkOperationRunMutation BulkOperationRunMutation `json:"bulkOperationRunMutation"`
}

type BulkOperationRunMutation struct {
	BulkOperation BulkOperation `json:"bulkOperation"`
	UserErrors    []interface{} `json:"userErrors"`
}

type BulkOperation struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Status string `json:"status"`
}
