package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gui0923/qdrant-golang-client/model/point"
)

type QDrantHttpClient struct {
	Scheme   string
	HostName string
	Port     int
}

func CreateHttpclient(schema string, hostname string, port int) *QDrantHttpClient {
	return &QDrantHttpClient{
		Scheme:   schema,
		HostName: hostname,
		Port:     port,
	}
}

func (client *QDrantHttpClient) GetPoints(collectionName string, request *point.PointsStringGetRequest) (point.PointsGetResponse, error) {
	bytesData, _ := json.Marshal(request)
	url := client.Scheme + "://" + client.HostName + ":" + strconv.Itoa(client.Port) + "/collections/" + collectionName + "/points"
	resp, err := http.Post(url, "application/json", bytes.NewReader(bytesData))
	if err != nil {
		return point.PointsGetResponse{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return point.PointsGetResponse{}, err
	}
	response := &point.PointsGetResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return point.PointsGetResponse{}, err
	}
	return *response, nil
}

func (client *QDrantHttpClient) UpsertPoints(collectionName string, request *point.PointsListUpsertRequest) (point.UpdateResultResponse, error) {
	bytesData, err := json.Marshal(request)
	if err != nil {
		return point.UpdateResultResponse{}, err
	}
	url := client.Scheme + "://" + client.HostName + ":" + strconv.Itoa(client.Port) + "/collections/" + collectionName + "/points"
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bytesData))
	if err != nil {
		return point.UpdateResultResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return point.UpdateResultResponse{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return point.UpdateResultResponse{}, err
	}
	defer resp.Body.Close()
	response := &point.UpdateResultResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return point.UpdateResultResponse{}, err
	}
	return *response, nil
}
