package webtools

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/scheiblingco/gofn/errtools"
	"github.com/scheiblingco/gofn/typetools"
)

type RequestMethod string

const (
	GET    RequestMethod = "GET"
	POST   RequestMethod = "POST"
	PUT    RequestMethod = "PUT"
	DELETE RequestMethod = "DELETE"
	PATCH  RequestMethod = "PATCH"
)

type RestRequest struct {
	Url    string
	Method RequestMethod

	Headers map[string]string

	BodyReader io.Reader

	errs []error
}

type RestResponse struct {
	StatusCode int
	Headers    map[string][]string
	Response   *http.Response

	bodyRead bool
}

func (r *RestRequest) addError(err error) {
	r.errs = append(r.errs, err)
}

func (r *RestRequest) WithAuthorization(auth Authorization) *RestRequest {
	return auth.Apply(r)
}

func (r *RestRequest) WithHeader(key, value string) *RestRequest {
	if key == "" {
		r.addError(errtools.InvalidKeyError("header key cannot be empty"))
	}

	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}

	r.Headers[key] = value

	return r
}

func (r *RestRequest) WithHeaders(headers map[string]string) *RestRequest {
	for k, v := range headers {
		r = r.WithHeader(k, v)
	}

	return r
}

func (r *RestRequest) WithBodyBytes(body []byte) *RestRequest {
	if r.Method == GET || r.Method == DELETE {
		r.addError(errtools.BodyNotAcceptedError(strings.ToLower(string(r.Method)) + " requests do not accept a body"))
	}

	r.BodyReader = bytes.NewReader(body)

	return r
}

func (r *RestRequest) WithBodyString(body string) *RestRequest {
	if r.Method == GET || r.Method == DELETE {
		r.addError(errtools.BodyNotAcceptedError(strings.ToLower(string(r.Method)) + " requests do not accept a body"))
	}

	return r.WithBodyBytes([]byte(body))
}

func (r *RestRequest) WithJsonBody(body interface{}, contentType *string) *RestRequest {
	if r.Method == GET || r.Method == DELETE {
		r.addError(errtools.BodyNotAcceptedError(strings.ToLower(string(r.Method)) + " requests do not accept a body"))
	}

	cType := "application/json"

	if contentType != nil {
		cType = *contentType
	}

	if val, ok := body.([]byte); ok {
		return r.WithBodyBytes(val).WithHeader("Content-Type", cType)
	}

	if val, ok := body.(string); ok {
		return r.WithBodyString(val).WithHeader("Content-Type", cType)
	}

	jsonString, err := json.Marshal(body)
	if err != nil {
		r.addError(err)
		return r
	}

	return r.WithBodyBytes(jsonString).WithHeader("Content-Type", cType)
}

func (r *RestRequest) WithXmlBody(body interface{}, contentType *string) *RestRequest {
	if r.Method == GET || r.Method == DELETE {
		r.addError(errtools.BodyNotAcceptedError(strings.ToLower(string(r.Method)) + " requests do not accept a body"))
	}

	cType := "application/xml"

	if contentType != nil {
		cType = *contentType
	}

	if val, ok := body.([]byte); ok {
		return r.WithBodyBytes(val).WithHeader("Content-Type", cType)
	}

	if val, ok := body.(string); ok {
		return r.WithBodyString(val).WithHeader("Content-Type", cType)
	}

	xmlString, err := xml.Marshal(body)
	if err != nil {
		r.addError(err)
		return r
	}

	return r.WithBodyBytes(xmlString).WithHeader("Content-Type", cType)
}

func (r *RestRequest) WithMultipartFormBody(body []MultipartField) *RestRequest {
	if r.Method == GET || r.Method == DELETE {
		r.addError(errtools.BodyNotAcceptedError(strings.ToLower(string(r.Method)) + " requests do not accept a body"))
	}

	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)

	for _, field := range body {
		if err := field.AddToWriter(writer); err != nil {
			r.addError(err)
		}
	}

	if len(r.errs) > 0 {
		return r
	}

	if err := writer.Close(); err != nil {
		r.addError(err)
		return r
	}

	r.BodyReader = buf

	return r.WithHeader("Content-Type", writer.FormDataContentType())
}

func (r *RestRequest) WithUrlencodedFormBody(body interface{}, contentType *string) *RestRequest {
	if r.Method == GET || r.Method == DELETE {
		r.addError(errtools.BodyNotAcceptedError(strings.ToLower(string(r.Method)) + " requests do not accept a body"))
	}

	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}

	cType := "application/x-www-form-urlencoded"

	if contentType != nil {
		cType = *contentType
	}

	data := url.Values{}

	switch bty := body.(type) {
	case map[string]string:
		for k, v := range bty {
			data.Set(k, v)
		}

		return r.WithBodyString(data.Encode()).WithHeader("Content-Type", cType)

	case map[string]interface{}:
		for k, v := range bty {
			if typetools.IsStringlikeType(v) || typetools.IsNumericType(v) {
				data.Set(k, typetools.EnsureString(v))
			} else {
				r.addError(errtools.InvalidTypeError("urlencoded form body values must be string-like or numeric"))
			}
		}

		return r.WithBodyString(data.Encode()).WithHeader("Content-Type", cType)

	case map[string][]string:
		for k, v := range bty {
			for _, val := range v {
				data.Add(k, val)
			}
		}

		return r.WithBodyString(data.Encode()).WithHeader("Content-Type", cType)
	}

	if typetools.IsStringlikeType(body) {
		return r.WithBodyString(typetools.EnsureString(body)).WithHeader("Content-Type", cType)
	}

	r.addError(errtools.InvalidTypeError("urlencoded form body must be a map[string]string, map[string]interface{}, map[string][]string or string-like"))

	return r
}

func (r *RestRequest) WithQueryParams(params map[string]string) *RestRequest {
	if len(params) == 0 {
		return r
	}

	if strings.Contains(r.Url, "?") {
		r.Url += "&"
	} else {
		r.Url += "?"
	}

	for k, v := range params {
		r.Url += k + "=" + v + "&"
	}

	r.Url = strings.TrimRight(r.Url, "&")

	return r
}

func (r *RestRequest) Validate() error {
	if len(r.errs) > 0 {
		return errtools.MultipleErrors(r.errs)
	}

	if len(r.Url) == 0 {
		return errtools.MissingValueError("url")
	}

	if r.Method == "" {
		return errtools.MissingValueError("method")
	}

	return nil
}

func (r *RestRequest) Execute() (*RestResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(string(r.Method), r.Url, r.BodyReader)
	if err != nil {
		return nil, err
	}

	if r.Headers != nil && len(r.Headers) > 0 {
		for k, v := range r.Headers {
			req.Header.Add(k, v)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return &RestResponse{
		Response:   resp,
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
	}, nil
}

func (r *RestResponse) BodyAsBytes() ([]byte, error) {
	if r.bodyRead {
		return nil, errtools.BodyConsumedError("body has already been read")
	}
	r.bodyRead = true
	defer r.Response.Body.Close()
	return io.ReadAll(r.Response.Body)
}

func (r *RestResponse) BodyAsString() (string, error) {
	b, err := r.BodyAsBytes()
	if r.bodyRead {
		return "", errtools.BodyConsumedError("body has already been read")
	}
	r.bodyRead = true
	return string(b), err
}

func (r *RestResponse) UnmarshalJsonBody(v interface{}) error {
	if r.bodyRead {
		return errtools.BodyConsumedError("body has already been read")
	}
	r.bodyRead = true
	return json.NewDecoder(r.Response.Body).Decode(v)
}

func (r *RestResponse) UnmarshalXmlBody(v interface{}) error {
	if r.bodyRead {
		return errtools.BodyConsumedError("body has already been read")
	}
	r.bodyRead = true
	return xml.NewDecoder(r.Response.Body).Decode(v)
}

func (r *RestResponse) Close() {
	r.Response.Body.Close()
}

func NewRequest(method RequestMethod, url string) *RestRequest {
	return &RestRequest{
		Method: method,
		Url:    url,
	}
}

func GetRequest(url string) *RestRequest {
	return NewRequest(GET, url)
}

func PostRequest(url string) *RestRequest {
	return NewRequest(POST, url)
}

func PutRequest(url string) *RestRequest {
	return NewRequest(PUT, url)
}

func DeleteRequest(url string) *RestRequest {
	return NewRequest(DELETE, url)
}

func PatchRequest(url string) *RestRequest {
	return NewRequest(PATCH, url)
}
