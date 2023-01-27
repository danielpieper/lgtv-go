package lgtv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContextUI(t *testing.T) {
	const testBaseURL = "http://example.com:8080"
	for _, tc := range []struct {
		name           string
		responseStatus int
		responseBody   string
		errorAssertion assert.ErrorAssertionFunc
		expected       string
	}{
		{
			name:           "happy path TouchPad",
			responseStatus: 200,
			responseBody: getXMLEnvelopeOK(`
 <data>
  <mode>TouchPad</mode>
 </data>
`),
			errorAssertion: assert.NoError,
			expected:       "TouchPad",
		},
		{
			name:           "happy path VolCh",
			responseStatus: 200,
			responseBody: getXMLEnvelopeOK(`
 <data>
  <mode>VolCh</mode>
 </data>
`),
			errorAssertion: assert.NoError,
			expected:       "VolCh",
		},
		{
			name:           "query error",
			responseStatus: 500,
			responseBody:   "invalid response",
			errorAssertion: assert.Error,
			expected:       "",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tv := newMockedLgTV(t, tc.responseStatus, tc.responseBody)
			actual, err := tv.GetContextUI()

			tc.errorAssertion(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
