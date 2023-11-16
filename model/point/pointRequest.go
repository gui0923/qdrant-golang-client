package point

import "github.com/gui0923/qdrant-golang-client/model"

type Payload struct {
	Json map[string]interface{} `json:"json"`
}

type PointStruct struct {
	Id      string                 `json:"id"`
	Vector  map[string][]float32   `json:"vector"`
	Payload map[string]interface{} `json:"payload"`
}

type PointsListUpsertRequest struct {
	Points []PointStruct `json:"points"`
}

type PointsStringGetRequest struct {
	Ids         []string `json:"ids"`
	WithPayload bool     `json:"with_payload"`
	WithVector  bool     `json:"with_vector"`
}

type PointsGetResponse struct {
	model.AbstractResponse
	Result []PointStruct `json:"result"`
}

type UpdateStatus struct {
	Name string `json:"name"`
}

type UpdateResult struct {
	Status      string `json:"status"`
	OperationId int64  `json:"operation_id"`
}

type UpdateResultResponse struct {
	model.AbstractResponse
	Result UpdateResult `json:"result"`
}

type PointsListDeleteRequest struct {
	Points []string `json:"points"`
}
