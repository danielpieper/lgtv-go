package lgtv

import (
	"bytes"
	"image"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetScreen(t *testing.T) {
	const (
		testBaseURL = "http://example.com:8080"
		testSession = 12345
	)

	okImageBytes := func() []byte {
		b, err := os.ReadFile("screen_test.png")
		require.NoError(t, err)
		return b
	}

	okImage := func() image.Image {
		f, err := os.Open("screen_test.png")
		require.NoError(t, err)
		defer f.Close()
		image, _, err := image.Decode(f)
		require.NoError(t, err)
		return image
	}

	for _, tc := range []struct {
		name           string
		responseStatus int
		responseBody   []byte
		session        int
		errorAssertion assert.ErrorAssertionFunc
		expected       image.Image
	}{
		{
			name:           "happy path",
			responseStatus: 200,
			responseBody:   okImageBytes(),
			session:        testSession,
			errorAssertion: assert.NoError,
			expected:       okImage(),
		},
		{
			name:           "unauthenticated",
			responseStatus: 200,
			errorAssertion: assert.Error,
		},
		{
			name:           "invalid image",
			responseStatus: 200,
			responseBody:   []byte("invalid image"),
			session:        testSession,
			errorAssertion: assert.Error,
		},
		{
			name:           "query error",
			responseStatus: 500,
			session:        testSession,
			errorAssertion: assert.Error,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			url, err := url.Parse(testBaseURL)
			require.NoError(t, err)

			tv := LgTV{
				baseURL: url,
				client: NewTestClient(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: tc.responseStatus,
						Body:       ioutil.NopCloser(bytes.NewReader(tc.responseBody)),
						Header:     make(http.Header),
					}
				}),
				session: tc.session,
			}

			actual, err := tv.GetScreen()

			tc.errorAssertion(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
