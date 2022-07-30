package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetLatestVideoList get latest video list
func GetLatestVideoList(mid int64) ([]VideoRecord, error) {
	responseBody, err := HTTPGetRequest(RequestURL, [][]string{
		{"mid", fmt.Sprint(mid)},
		{"order", "pubdate"},
		{"tid", "0"},
		{"pn", "1"},
		{"ps", "5"},
	})
	if err != nil {
		return []VideoRecord{}, err
	}
	var response biliResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return []VideoRecord{}, err
	}
	var records []VideoRecord
	for _, vData := range response.Data.List.VList {
		record := VideoRecord{
			URL:     URLPrefix + vData.BVID,
			Pic:     vData.Pic,
			Title:   vData.Title,
			Author:  vData.Author,
			Created: vData.Created,
		}
		records = append(records, record)
	}

	return records, nil
}

// HTTPGetRequest send a http get request
func HTTPGetRequest(url string, queryList [][]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for _, queryItem := range queryList {
		if len(queryItem) != 2 {
			return nil, fmt.Errorf("%v is not a valid query item", queryItem)
		}
		q.Add(queryItem[0], queryItem[1])
	}
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
