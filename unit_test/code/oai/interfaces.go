package oai

type Fetcher interface {
	Fetch(url string) ([]byte, error)
}
