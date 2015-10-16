package oai

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MDResult struct {
	XMLName xml.Name   `xml:"OAI-PMH"`
	Records []MDFormat `xml:"ListMetadataFormats>metadataFormat"`
}
type MDFormat struct {
	Prefix    string `xml:"metadataPrefix"`
	Schema    string `xml:"schema"`
	Namespace string `xml:"metadataNamespace"`
}

// implement the Fetcher Interface
type OAIFetcher struct{}

func (of OAIFetcher) Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET failed: %v", err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// composition - OAI has the Fetcher Interface
type OAI struct {
	OAIFetcher Fetcher // use the interface type to allow for injection
}

func (oai OAI) FetchMetadataFormats(baseURL string) ([]MDFormat, error) {
	// call Fetch
	body, err := oai.OAIFetcher.Fetch(fmt.Sprintf("%s?verb=ListMetadataFormats", baseURL))
	if err != nil {
		return nil, fmt.Errorf("Error fetching metadata formats for %s: %v", baseURL, err)
	}

	return oai.parse(body)
}

func (oai *OAI) parse(body []byte) ([]MDFormat, error) {
	result := &MDResult{}
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result.Records, nil
}
