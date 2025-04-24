package request

import "net/http"

type HttpClientInterface interface {
	Get(url string, headers map[string]string) (*http.Response, error)
	Post(url string, body []byte, headers map[string]string) (*http.Response, error)
}

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() HttpClientInterface {
	return &HttpClient{
		client: &http.Client{},
	}
}

func (httpClient *HttpClient) Get(url string, headers map[string]string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	return httpClient.client.Do(request)
}

func (httpClient *HttpClient) Post(url string, body []byte, headers map[string]string) (*http.Response, error) {
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	request.Body = http.NoBody

	return httpClient.client.Do(request)
}
