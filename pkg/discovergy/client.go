package discovergy

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const baseUrl = "https://api.discovergy.com/public/v1"

type Client struct {
	email    string
	password string
}

type Meter struct {
	ID             string `json:"meterId"`
	ManufacturerID string `json:"manufacturerId"`
	Serial         string `json:"fullSerialNumber"`
}

type MetersResponse []Meter

type LastReading struct {
	Timestamp int64             `json:"time"`
	Values    LastReadingValues `json:"values"`
}

type LastReadingValues struct {
	MilliWattTotalPower  int32 `json:"power"`
	MilliWattPhase1Power int32 `json:"phase1Power"`
	MilliWattPhase2Power int32 `json:"phase2Power"`
	MilliWattPhase3Power int32 `json:"phase3Power"`

	VoltagePhase1 int32 `json:"phase1Voltage"`
	VoltagePhase2 int32 `json:"phase2Voltage"`
	VoltagePhase3 int32 `json:"phase3Voltage"`
}

func (c *Client) LastReading(meterId string) (*LastReading, error) {
	response, err := c.getRequest(fmt.Sprintf("last_reading?meterId=%v", meterId), &LastReading{})
	if err != nil {
		return nil, err
	}

	if readingResponse, ok := response.(*LastReading); ok {
		return readingResponse, nil
	}

	return nil, errors.New("failed to fetch last reading from Discovergy API")
}

func (c *Client) GetMeters() ([]Meter, error) {
	response, err := c.getRequest("meters", &MetersResponse{})
	if err != nil {
		return nil, err
	}

	if metersResponse, ok := response.(*MetersResponse); ok {
		return *metersResponse, nil
	}

	return nil, errors.New("failed to fetch meters from Discovergy API")
}

func (c *Client) getRequest(u string, response interface{}) (interface{}, error) {
	parsedUrl, err := url.Parse(fmt.Sprintf("%v/%v", baseUrl, u))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to parse url: %v", err.Error()))
	}

	req, err := http.NewRequest("GET", parsedUrl.String(), nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to prepare request: %v", err.Error()))
	}

	req.Host = parsedUrl.Host

	req.SetBasicAuth(c.email, c.password)
	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Discovergy api returned wrong status code: %v", r.StatusCode))
	}

	rBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("can not read response from Discovergy api: %v", err.Error()))
	}

	err = json.Unmarshal(rBytes, response)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("can not parse json from Discovergy api: %v", err.Error()))
	}

	return response, nil
}
