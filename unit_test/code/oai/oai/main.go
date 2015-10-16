package main

import (
	"fmt"

	"github.com/nathj07/talks/unit_test/code/oai"
)

func main() {
	oai := oai.OAI{
		OAIFetcher: oai.OAIFetcher{},
	}
	data, err := oai.FetchMetadataFormats("http://www.worldsciencepublisher.org/journals/index.php/ACSA/oai")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Retrieved Data:\n%+v", data)
}
