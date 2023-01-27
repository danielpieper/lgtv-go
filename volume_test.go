package lgtv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVolume(t *testing.T) {
	for _, tc := range []struct {
		name           string
		responseStatus int
		responseBody   string
		errorAssertion assert.ErrorAssertionFunc
		expected       TVVolume
	}{
		{
			name:           "happy path",
			responseStatus: 200,
			responseBody: getXMLEnvelopeOK(`
 <data>
  <mute>false</mute>
  <minLevel>0</minLevel>
  <maxLevel>100</maxLevel>
  <level>13</level>
 </data>
`),
			errorAssertion: assert.NoError,
			expected: TVVolume{
				Mute:     false,
				MinLevel: 0,
				MaxLevel: 100,
				Level:    13,
			},
		},
		{
			name:           "query error",
			responseStatus: 500,
			responseBody:   "invalid response",
			errorAssertion: assert.Error,
			expected:       TVVolume{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tv := newMockedLgTV(t, tc.responseStatus, tc.responseBody)
			actual, err := tv.GetVolume()

			tc.errorAssertion(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestIsMuted(t *testing.T) {
	for _, tc := range []struct {
		name           string
		responseStatus int
		responseBody   string
		errorAssertion assert.ErrorAssertionFunc
		expected       bool
	}{
		{
			name:           "happy path: not muted",
			responseStatus: 200,
			responseBody: getXMLEnvelopeOK(`
 <data>
  <mute>false</mute>
  <minLevel>0</minLevel>
  <maxLevel>100</maxLevel>
  <level>13</level>
 </data>
`),
			errorAssertion: assert.NoError,
			expected:       false,
		},
		{
			name:           "happy path: muted",
			responseStatus: 200,
			responseBody: getXMLEnvelopeOK(`
 <data>
  <mute>true</mute>
  <minLevel>0</minLevel>
  <maxLevel>100</maxLevel>
  <level>13</level>
 </data>
`),
			errorAssertion: assert.NoError,
			expected:       true,
		},
		{
			name:           "query error",
			responseStatus: 500,
			responseBody:   "invalid response",
			errorAssertion: assert.Error,
			expected:       false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tv := newMockedLgTV(t, tc.responseStatus, tc.responseBody)
			actual, err := tv.IsMuted()

			tc.errorAssertion(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
