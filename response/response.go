package response

import (
	"io"
	"net/http"

	"github.com/henomis/restclientgo"
)

type Response struct {
	Code           int         `json:"code"`
	HttpStatusCode int         `json:"httpStatusCode"`
	Message        string      `json:"message"`
	RawBody        *string     `json:"-"`
	Headers        http.Header `json:"-"`
}

func (r *Response) IsSuccess() bool {
	return (r.Code != 200)
}

func (r *Response) SetStatusCode(code int) error {
	r.HttpStatusCode = code
	return nil
}

func (r *Response) SetBody(body io.Reader) error {

	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	s := string(b)
	r.RawBody = &s

	return nil
}

func (r *Response) SetHeaders(headers restclientgo.Headers) error {
	r.Headers = http.Header(headers)
	return nil
}
