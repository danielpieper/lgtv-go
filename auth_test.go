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

func TestRequestPairingKey(t *testing.T) {
	const testBaseURL = "http://example.com:8080"
	for _, tc := range []struct {
		name           string
		httpClient     *http.Client
		errorAssertion assert.ErrorAssertionFunc
	}{
		{
			name: "happy path",
			httpClient: NewTestClient(func(req *http.Request) *http.Response {
				defer req.Body.Close()
				reqBytes, err := io.ReadAll(req.Body)
				require.NoError(t, err)

				assert.Equal(t, "POST", req.Method)
				assert.Equal(t, testBaseURL+"/roap/api/auth", req.URL.String())
				assert.Equal(t, xml.Header+"<auth><type>AuthKeyReq</type></auth>", string(reqBytes))

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
			}
			err = tv.RequestPairingKey()

			tc.errorAssertion(t, err)
		})
	}
}

func TestAuthenticate(t *testing.T) {
	const (
		testBaseURL    = "http://example.com:8080"
		testPairingKey = "123456"
	)
	for _, tc := range []struct {
		name           string
		httpClient     *http.Client
		errorAssertion assert.ErrorAssertionFunc
	}{
		{
			name: "happy path",
			httpClient: NewTestClient(func(req *http.Request) *http.Response {
				defer req.Body.Close()
				reqBytes, err := io.ReadAll(req.Body)
				require.NoError(t, err)

				assert.Equal(t, "POST", req.Method)
				assert.Equal(t, testBaseURL+"/roap/api/auth", req.URL.String())
				assert.Equal(
					t,
					fmt.Sprintf("%s<auth><type>AuthReq</type><value>%s</value></auth>", xml.Header, testPairingKey),
					string(reqBytes),
				)

				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(getXMLEnvelopeOK(`<session>1234567890</session>`))),
					Header:     make(http.Header),
				}
			}),
			errorAssertion: assert.NoError,
		},
		{
			name: "session not provided",
			httpClient: NewTestClient(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(getXMLEnvelopeOK(""))),
					Header:     make(http.Header),
				}
			}),
			errorAssertion: assert.Error,
		},
		{
			name: "http client error",
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
			}
			err = tv.Authenticate(testPairingKey)

			tc.errorAssertion(t, err)
		})
	}
}

func TestIsAuthenticated(t *testing.T) {
	const testSession = 1234567890
	for _, tc := range []struct {
		name     string
		session  int
		expected bool
	}{
		{
			name:     "authenticated",
			session:  testSession,
			expected: true,
		},
		{
			name:     "not authenticated",
			session:  0,
			expected: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tv := LgTV{
				session: tc.session,
			}
			actual := tv.IsAuthenticated()

			assert.Equal(t, tc.expected, actual)
		})
	}
}
