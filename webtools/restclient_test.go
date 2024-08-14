package webtools_test

import (
	"fmt"
	"maps"
	"testing"

	"github.com/scheiblingco/gofn/webtools"
)

type ResponseBody struct {
	Args    map[string]string `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	URL     string            `json:"url"`
}

func TestGetRequest(t *testing.T) {
	t.Log("Testing GET request")

	expectedHeaders := map[string]string{
		"Accept":     "application/json",
		"User-Agent": "github.com/scheiblingco/gofn",
	}

	expectedArgs := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	resp, err := webtools.GetRequest("https://httpbin.org/get").
		WithHeader("Accept", expectedHeaders["Accept"]).
		WithHeaders(map[string]string{
			"User-Agent": expectedHeaders["User-Agent"],
		}).
		WithQueryParams(expectedArgs).Execute()

	if err != nil {
		t.Error(err)
	}

	bodyObj := &ResponseBody{}

	if err := resp.UnmarshalJsonBody(bodyObj); err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	if len(bodyObj.Args) != len(expectedArgs) || !maps.Equal(bodyObj.Args, expectedArgs) {
		t.Errorf("Expected Args to be map[string]string{\"key1\": \"value1\", \"key2\": \"value2\"}, got %v", bodyObj.Args)
	}

	for k, v := range expectedHeaders {
		if bodyObj.Headers[k] != v {
			t.Errorf("Expected Headers[%s] to be %s, got %s", k, v, bodyObj.Headers[k])
		}
	}

	fmt.Println("Status code:", resp.StatusCode)
	fmt.Println("Headers:", resp.Headers)
	fmt.Println("Body:", bodyObj)

}
