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
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockTvCommand struct {
	XMLName xml.Name `xml:"test"`
	Value   string   `xml:"value"`
}

func (m mockTvCommand) isCommand() {}

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func newMockedLgTV(t *testing.T, responseStatus int, responseBody string) LgTV {
	const (
		testBaseURL = "http://example.com:8080"
		testSession = 12345
	)

	url, err := url.Parse(testBaseURL)
	require.NoError(t, err)

	return LgTV{
		baseURL: url,
		client: NewTestClient(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: responseStatus,
				Body:       ioutil.NopCloser(bytes.NewBufferString(responseBody)),
				Header:     make(http.Header),
			}
		}),
		session: testSession,
	}
}

func getXMLEnvelopeOK(data string) string {
	return getXMLEnvelope(200, "OK", data)
}

func getXMLEnvelope(roapError int, roapErrorDetail, data string) string {
	return fmt.Sprintf(
		`%s<envelope><ROAPError>%d</ROAPError><ROAPErrorDetail>%s</ROAPErrorDetail>%s</envelope>`,
		xml.Header,
		roapError,
		roapErrorDetail,
		data,
	)
}

func TestNew(t *testing.T) {
	const testBaseURL = "http://example.com:8080"

	actual, err := New(testBaseURL)

	assert.NoError(t, err)
	assert.Equal(t, time.Second, actual.client.Timeout)
}

func TestCreateCommandRequest(t *testing.T) {
	const (
		testBaseURL = "http://example.com:8080"
		testPath    = "any/path"
		testValue   = "any value"

		expectedMethod = "POST"
		expectedBody   = `<?xml version="1.0" encoding="UTF-8"?>
<test><value>any value</value></test>`
		expectedContentType = `application/atom+xml`
	)

	url, err := url.Parse(testBaseURL)
	require.NoError(t, err)

	tv := LgTV{baseURL: url}
	cmd := mockTvCommand{Value: testValue}
	actual, err := tv.createCommandRequest(testPath, cmd)

	assert := assert.New(t)
	assert.NoError(err)

	if assert.NotNil(actual) {
		assert.Equal(expectedMethod, actual.Method)
		assert.Equal(expectedContentType, actual.Header.Get("Content-type"))

		actualBody, err := io.ReadAll(actual.Body)
		if err != nil {
			t.Fatalf("unable to read request body: %v", err)
		}
		assert.Equal(expectedBody, string(actualBody))
	}
}

func TestCreateInfoRequest(t *testing.T) {
	const (
		testBaseURL = "http://example.com:8080"
		testPath    = "any/path"
		testValue   = "any value"

		expectedMethod      = "GET"
		expectedContentType = `application/atom+xml`
	)

	url, err := url.Parse(testBaseURL)
	require.NoError(t, err)

	tv := LgTV{baseURL: url}
	actual, err := tv.createInfoRequest(testPath)

	assert := assert.New(t)
	assert.NoError(err)

	if assert.NotNil(actual) {
		assert.Equal(expectedMethod, actual.Method)
		assert.Equal(expectedContentType, actual.Header.Get("Content-type"))
	}
}

func TestSendRequest(t *testing.T) {
	const testBaseURL = "http://example.com:8080"

	for _, tc := range []struct {
		name           string
		responseStatus int
		responseBody   string
		assertion      assert.ErrorAssertionFunc
	}{
		{
			name:           "happy path",
			responseStatus: 200,
			responseBody: `<?xml version="1.0" encoding="UTF-8"?>
<envelope><ROAPError>200</ROAPError><ROAPErrorDetail></ROAPErrorDetail></envelope>`,
			assertion: assert.NoError,
		},
		{
			name:           "roap error",
			responseStatus: 200,
			responseBody: `<?xml version="1.0" encoding="UTF-8"?>
<envelope><ROAPError>404</ROAPError><ROAPErrorDetail>Not Found</ROAPErrorDetail></envelope>`,
			assertion: assert.Error,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", testBaseURL, nil)
			if err != nil {
				t.Fatalf("unable to create request: %v", err)
			}

			tv := newMockedLgTV(t, tc.responseStatus, tc.responseBody)

			err = tv.sendRequest(req, nil)
			tc.assertion(t, err)
		})
	}
}
