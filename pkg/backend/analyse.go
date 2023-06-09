package backend

import (
	"bytes"
	"fmt"
	"net/http"
)

type HeaderTokenKeyType string

const (
	HeaderTokenSnyk HeaderTokenKeyType = "crda_snyk_token"
)

// AnalyzeDependencyTree is used to create the stack report against the backend
// will return the response body or an error
func AnalyzeDependencyTree(
	backendHost, ecosystem, cliClient, contentType string,
	content []byte,
	tokens map[HeaderTokenKeyType]string,
	jsonOut bool,
) (*http.Response, error) {
	apiUrl := fmt.Sprintf("%s/api/v3/dependency-analysis/%s", backendHost, ecosystem)

	request, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewReader(content))
	if err != nil {
		return nil, err
	}

	accept := "multipart/mixed"
	if jsonOut {
		accept = "application/json"
	}

	request.Header.Add("Client", cliClient)
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept", accept)
	// include token headers
	for k, t := range tokens {
		request.Header.Add(string(k), t)
	}

	httpClient := &http.Client{}
	return httpClient.Do(request)
}
