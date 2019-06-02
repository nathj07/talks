package httptesting

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dankinder/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRetryFetcherHTTPTEST(t *testing.T) {
	attempts := 0
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts += 1
		w.WriteHeader(http.StatusNotFound)
	}))

	defer s.Close()

	statusCode, err := retryFetcher(s.URL, []int{404, 400}, 5)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if statusCode != http.StatusNotFound {
		t.Errorf("Unexpected status code: %d", statusCode)
	}
	if attempts != 5 {
		t.Errorf("Unepxceted number of attempts: %d", attempts)
	}
}

func TestRetryFetcherHTTPMOCK(t *testing.T) {
	mh := &httpmock.MockHandler{}
	mh.On("Handle", http.MethodGet, "/", mock.Anything).Return(httpmock.Response{
		Status: http.StatusNotFound,
	}).Times(5)
	s := httpmock.NewServer(mh)
	defer s.Close()
	statusCode, err := retryFetcher(s.URL(), []int{404}, 5)
	require.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)
	mh.AssertExpectations(t)
}
