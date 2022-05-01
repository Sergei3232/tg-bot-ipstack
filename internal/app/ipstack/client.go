package ipstack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ClientIP struct {
	Hostname  string
	AccessKey string
}

type IpInfo struct {
	Ip            string  `json:"ip"`
	Hostname      string  `json:"hostname"`
	Type          string  `json:"type"`
	ContinentCode string  `json:"continent_code"`
	ContinentName string  `json:"continent_name"`
	CountryCode   string  `json:"country_code"`
	CountryName   string  `json:"country_name"`
	RegionCode    string  `json:"region_code"`
	RegionName    string  `json:"region_name"`
	City          string  `json:"city"`
	Zip           string  `json:"zip"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
}

type QueryIP interface {
	GetInfoIp(ip string) (*IpInfo, error)
}

func NewClientIP(hostname, accessKey string) QueryIP {
	return &ClientIP{
		hostname,
		accessKey,
	}
}

func (c *ClientIP) GetInfoIp(ip string) (*IpInfo, error) {
	requestText := c.Hostname + "/" + ip + "/" + "?" + "access_key=" + c.AccessKey
	clietnHttp := http.Client{}
	resp, err := clietnHttp.Get(requestText)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	var data IpInfo
	json.Unmarshal(body, &data)

	return &data, nil
}

func (i *IpInfo) ToString() string {
	result := "ip: " + i.Ip + "\n" +
		"hostname: " + i.Hostname + "\n" +
		"type: " + i.Type + "\n" +
		"continent_code: " + i.ContinentCode + "\n" +
		"continent_name: " + i.ContinentName + "\n" +
		"country_code: " + i.CountryCode + "\n" +
		"country_name: " + i.CountryName + "\n" +
		"region_code: " + i.RegionCode + "\n" +
		"region_name: " + i.RegionName + "\n" +
		"city: " + i.City + "\n" +
		"zip: " + i.Zip + "\n" +
		"latitude: " + fmt.Sprintf("%#v", i.Latitude) + "\n" +
		"longitude: " + fmt.Sprintf("%#v", i.Longitude)

	return result
}
