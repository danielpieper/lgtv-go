package lgtv

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMouseMove(t *testing.T) {
	const (
		testBaseURL = "http://example.com:8080"
		testSession = 1234567890
		testX       = 123
		testY       = 456
	)

	for _, tc := range []struct {
		name           string
		httpClient     *http.Client
		x              int
		y              int
		errorAssertion assert.ErrorAssertionFunc
	}{
		{
			name: "happy path",
			x:    testX,
			y:    testY,
			httpClient: NewTestClient(func(req *http.Request) *http.Response {
				defer req.Body.Close()
				reqBytes, err := io.ReadAll(req.Body)
				require.NoError(t, err)

				assert.Equal(t, "POST", req.Method)
				assert.Equal(t, testBaseURL+"/roap/api/command", req.URL.String())
				assert.Equal(
					t,
					fmt.Sprintf("%s<command><name>HandleTouchMove</name><x>%d</x><y>%d</y></command>", xml.Header, testX, testY),
					string(reqBytes),
				)

				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(getXMLEnvelopeOK(""))),
					Header:     make(http.Header),
				}
			}),
			errorAssertion: assert.NoError,
		},
		{
			name: "http client error",
			x:    testX,
			y:    testY,
			httpClient: NewTestClient(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: 404,
					Body:       ioutil.NopCloser(bytes.NewBufferString(getXMLEnvelope(404, "Not Found", ""))),
					Header:     make(http.Header),
				}
			}),
			errorAssertion: assert.Error,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			url, err := url.Parse(testBaseURL)
			require.NoError(t, err)

			tv := LgTV{
				baseURL: url,
				client:  tc.httpClient,
				session: testSession,
			}

			err = tv.MouseMove(tc.x, tc.y)

			tc.errorAssertion(t, err)
		})
	}
}
