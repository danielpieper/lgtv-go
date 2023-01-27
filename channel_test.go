package lgtv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChannelList(t *testing.T) {
	for _, tc := range []struct {
		name           string
		responseStatus int
		responseBody   string
		errorAssertion assert.ErrorAssertionFunc
		expected       []TVChannel
	}{
		{
			name:           "happy path",
			responseStatus: 200,
			responseBody: getXMLEnvelopeOK(`
 <data>
  <chtype>cable</chtype>
  <sourceIndex>3</sourceIndex>
  <physicalNum>1</physicalNum>
  <major>1</major>
  <displayMajor>1</displayMajor>
  <minor>0</minor>
  <displayMinor>-1</displayMinor>
  <chname>Eurosp360 1</chname>
 </data>
 <data>
  <chtype>cable</chtype>
  <sourceIndex>3</sourceIndex>
  <physicalNum>1</physicalNum>
  <major>2</major>
  <displayMajor>2</displayMajor>
  <minor>0</minor>
  <displayMinor>-1</displayMinor>
  <chname>Sky Sport 3 HD</chname>
 </data>
 <data>
  <chtype>cable</chtype>
  <sourceIndex>3</sourceIndex>
  <physicalNum>2</physicalNum>
  <major>3</major>
  <displayMajor>3</displayMajor>
  <minor>0</minor>
  <displayMinor>-1</displayMinor>
  <chname>Sky Sport 4 HD</chname>
 </data>`),
			errorAssertion: assert.NoError,
			expected: []TVChannel{
				{
					ChannelType:  "cable",
					ChannelName:  "Eurosp360 1",
					SourceIndex:  3,
					PhysicalNum:  1,
					Major:        1,
					DisplayMajor: 1,
					Minor:        0,
					DisplayMinor: -1,
				},
				{
					ChannelType:  "cable",
					ChannelName:  "Sky Sport 3 HD",
					SourceIndex:  3,
					PhysicalNum:  1,
					Major:        2,
					DisplayMajor: 2,
					Minor:        0,
					DisplayMinor: -1,
				},
				{
					ChannelType:  "cable",
					ChannelName:  "Sky Sport 4 HD",
					SourceIndex:  3,
					PhysicalNum:  2,
					Major:        3,
					DisplayMajor: 3,
					Minor:        0,
					DisplayMinor: -1,
				},
			},
		},
		{
			name:           "query error",
			responseStatus: 500,
			responseBody:   "invalid response",
			errorAssertion: assert.Error,
			expected:       []TVChannel{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tv := newMockedLgTV(t, tc.responseStatus, tc.responseBody)
			actual, err := tv.GetChannelList()

			tc.errorAssertion(t, err)
			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
}

func TestGetCurrentChannel(t *testing.T) {
	for _, tc := range []struct {
		name           string
		responseStatus int
		responseBody   string
		errorAssertion assert.ErrorAssertionFunc
		expected       TVChannel
	}{
		{
			name:           "happy path",
			responseStatus: 200,
			responseBody: getXMLEnvelopeOK(`
 <data>
  <chtype>terrestrial</chtype>
  <sourceIndex>0</sourceIndex>
  <physicalNum>0</physicalNum>
  <major>0</major>
  <displayMajor>0</displayMajor>
  <minor>0</minor>
  <displayMinor>-1</displayMinor>
  <chname></chname>
  <progName></progName>
  <audioCh>0</audioCh>
  <inputSourceName></inputSourceName>
  <inputSourceType>0</inputSourceType>
  <labelName></labelName>
  <inputSourceIdx>0</inputSourceIdx>
 </data>
`),
			errorAssertion: assert.NoError,
			expected: TVChannel{
				ChannelType:  "terrestrial",
				ChannelName:  "",
				SourceIndex:  0,
				PhysicalNum:  0,
				Major:        0,
				DisplayMajor: 0,
				Minor:        0,
				DisplayMinor: -1,
			},
		},
		{
			name:           "query error",
			responseStatus: 500,
			responseBody:   "invalid response",
			errorAssertion: assert.Error,
			expected:       TVChannel{},
		},
		{
			name:           "no result error",
			responseStatus: 200,
			responseBody:   `<?xml version="1.0" encoding="utf-8"?><envelope><ROAPError>200</ROAPError><ROAPErrorDetail>OK</ROAPErrorDetail></envelope>`,
			errorAssertion: assert.Error,
			expected:       TVChannel{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tv := newMockedLgTV(t, tc.responseStatus, tc.responseBody)
			actual, err := tv.GetCurrentChannel()

			tc.errorAssertion(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
