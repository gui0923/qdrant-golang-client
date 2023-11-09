package client

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (client *QDrantHttpClient) GetPoints(collectionName string, request point.PointsStringGetRequest) point.PointsGetResponse {
	bytesData, _ := json.Marshal(request)
	url := client.Scheme + "://" + client.HostName + ":" + strconv.Itoa(client.Port) + "/collections/" + collectionName + "/points"
	fmt.Printf("url: %v\n", url)
	resp, _ := http.Post(url, "application/json", bytes.NewReader(bytesData))
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("string(body): %v\n", string(body))
	response := &point.PointsGetResponse{}
	json.Unmarshal(body, response)
	return *response
}

func (client *QDrantHttpClient) UpsertPoints(collectionName string, request point.PointsListUpsertRequest) point.UpdateResultResponse {
	bytesData, _ := json.Marshal(request)
	url := client.Scheme + "://" + client.HostName + ":" + strconv.Itoa(client.Port) + "/collections/" + collectionName + "/points"
	fmt.Printf("url: %v\n", url)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(bytesData))
	req.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{}
	resp, _ := httpClient.Do(req)
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("string(body): %v\n", string(body))
	response := &point.UpdateResultResponse{}
	json.Unmarshal(body, response)
	return *response
}