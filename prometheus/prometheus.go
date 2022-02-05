package prometheus

import (
	"awesomeAssistant/util"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

const BasePath = "http://192.168.0.2:9090"
const ApiQueryPath = "/api/v1/query"

// NowValue struct to in moment data response, not time range
type NowValue struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Instance string `json:"instance"`
				Job      string `json:"job"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

func FreeRam() (float64, string) {
	values := make(map[string]string)
	values["query"] = "100-((node_memory_MemTotal_bytes-node_memory_MemFree_bytes)/node_memory_MemTotal_bytes)*100"
	values["time"] = strconv.FormatInt(time.Now().Unix(), 10)
	err, validUrl := util.UrlToCanonical(BasePath, ApiQueryPath, values)
	if err != nil {
		log.Println(err)
	}

	client := new(http.Client)
	req, err := http.NewRequest("GET", validUrl, nil)
	req.Header.Add("Accept-Encoding", "gzip, deflate")

	resp, err := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Println(err)
		}
		defer func(reader io.ReadCloser) {
			err := reader.Close()
			if err != nil {
				log.Println(err)
			}
		}(reader)
	default:
		reader = resp.Body
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(reader)
	if err != nil {
		log.Println(err)
	}
	var cont NowValue
	err = json.Unmarshal(buf.Bytes(), &cont)
	if err != nil {
		log.Println(err)
	}
	reqTime, okTime := cont.Data.Result[0].Value[0].(float64)
	reqVal, okVal := cont.Data.Result[0].Value[1].(string)
	if okTime && okVal {
		return reqTime, reqVal
	}
	return 0, ""
}
