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
	bytesData, err := json.Marshal(request)
	if err != nil {
		return point.PointsGetResponse{}, err
	}
	url := client.Scheme + "://" + client.HostName + ":" + strconv.Itoa(client.Port) + "/collections/" + collectionName + "/points"
	resp, err := http.Post(url, "application/json", bytes.NewReader(bytesData))
	if err != nil {
		return point.PointsGetResponse{}, err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
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
	return doUpdateRequest(url, "PUT", bytesData)
}

func (client *QDrantHttpClient) DeletePoints(collectionName string, request *point.PointsListDeleteRequest) (point.UpdateResultResponse, error) {
	bytesData, err := json.Marshal(request)
	if err != nil {
		return point.UpdateResultResponse{}, err
	}
	url := client.Scheme + "://" + client.HostName + ":" + strconv.Itoa(client.Port) + "/collections/" + collectionName + "/points/delete"
	return doUpdateRequest(url, "POST", bytesData)
}

func doUpdateRequest(url string, method string, bytesData []byte) (point.UpdateResultResponse, error) {
	body, err := requestHttp(url, method, bytesData)
	if err != nil {
		return point.UpdateResultResponse{}, err
	}
	response := &point.UpdateResultResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return point.UpdateResultResponse{}, err
	}
	return *response, nil
}

func requestHttp(url string, method string, bytesData []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bytesData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return body, nil
}
