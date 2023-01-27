package lgtv

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeyInput(t *testing.T) {
	const (
		testBaseURL = "http://example.com:8080"
		testSession = 1234567890
	)

	okHttpClient := func(key tvKey) *http.Client {
		return NewTestClient(func(req *http.Request) *http.Response {
			defer req.Body.Close()
			reqBytes, err := io.ReadAll(req.Body)
			require.NoError(t, err)

			assert.Equal(t, "POST", req.Method)
			assert.Equal(t, testBaseURL+"/roap/api/command", req.URL.String())
			assert.Equal(
				t,
				fmt.Sprintf("%s<command><name>HandleKeyInput</name><value>%s</value></command>", xml.Header, key),
				string(reqBytes),
			)

			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(getXMLEnvelopeOK(""))),
				Header:     make(http.Header),
			}
		})
	}

	errHttpClient := func() *http.Client {
		return NewTestClient(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: 404,
				Body:       ioutil.NopCloser(bytes.NewBufferString(getXMLEnvelope(404, "Not Found", ""))),
				Header:     make(http.Header),
			}
		})
	}

	for _, tc := range []struct {
		name           string
		method         string
		httpClient     *http.Client
		errorAssertion assert.ErrorAssertionFunc
	}{
		{"KeyPowerOff happy path", "KeyPowerOff", okHttpClient(tvKeyPower), assert.NoError},
		{"KeyNumber0 happy path", "KeyNumber0", okHttpClient(tvKeyNumber0), assert.NoError},
		{"KeyNumber1 happy path", "KeyNumber1", okHttpClient(tvKeyNumber1), assert.NoError},
		{"KeyNumber2 happy path", "KeyNumber2", okHttpClient(tvKeyNumber2), assert.NoError},
		{"KeyNumber3 happy path", "KeyNumber3", okHttpClient(tvKeyNumber3), assert.NoError},
		{"KeyNumber4 happy path", "KeyNumber4", okHttpClient(tvKeyNumber4), assert.NoError},
		{"KeyNumber5 happy path", "KeyNumber5", okHttpClient(tvKeyNumber5), assert.NoError},
		{"KeyNumber6 happy path", "KeyNumber6", okHttpClient(tvKeyNumber6), assert.NoError},
		{"KeyNumber7 happy path", "KeyNumber7", okHttpClient(tvKeyNumber7), assert.NoError},
		{"KeyNumber8 happy path", "KeyNumber8", okHttpClient(tvKeyNumber8), assert.NoError},
		{"KeyNumber9 happy path", "KeyNumber9", okHttpClient(tvKeyNumber9), assert.NoError},
		{"KeyUp happy path", "KeyUp", okHttpClient(tvKeyUp), assert.NoError},
		{"KeyDown happy path", "KeyDown", okHttpClient(tvKeyDown), assert.NoError},
		{"KeyLeft happy path", "KeyLeft", okHttpClient(tvKeyLeft), assert.NoError},
		{"KeyRight happy path", "KeyRight", okHttpClient(tvKeyRight), assert.NoError},
		{"KeyOK happy path", "KeyOK", okHttpClient(tvKeyOk), assert.NoError},
		{"KeyHomeMenu happy path", "KeyHomeMenu", okHttpClient(tvKeyHomeMenu), assert.NoError},
		{"KeyBack happy path", "KeyBack", okHttpClient(tvKeyBack), assert.NoError},
		{"KeyVolumeUp happy path", "KeyVolumeUp", okHttpClient(tvKeyVolumeUp), assert.NoError},
		{"KeyVolumeDown happy path", "KeyVolumeDown", okHttpClient(tvKeyVolumeDown), assert.NoError},
		{"KeyMuteToggle happy path", "KeyMuteToggle", okHttpClient(tvKeyMuteToggle), assert.NoError},
		{"KeyChannelUp happy path", "KeyChannelUp", okHttpClient(tvKeyChannelUp), assert.NoError},
		{"KeyChannelDown happy path", "KeyChannelDown", okHttpClient(tvKeyChannelDown), assert.NoError},
		{"KeyBlue happy path", "KeyBlue", okHttpClient(tvKeyBlue), assert.NoError},
		{"KeyGreen happy path", "KeyGreen", okHttpClient(tvKeyGreen), assert.NoError},
		{"KeyRed happy path", "KeyRed", okHttpClient(tvKeyRed), assert.NoError},
		{"KeyYellow happy path", "KeyYellow", okHttpClient(tvKeyYellow), assert.NoError},
		{"KeyPlay happy path", "KeyPlay", okHttpClient(tvKeyPlay), assert.NoError},
		{"KeyPause happy path", "KeyPause", okHttpClient(tvKeyPause), assert.NoError},
		{"KeyStop happy path", "KeyStop", okHttpClient(tvKeyStop), assert.NoError},
		{"KeyFastForward happy path", "KeyFastForward", okHttpClient(tvKeyFastForward), assert.NoError},
		{"KeyRewind happy path", "KeyRewind", okHttpClient(tvKeyRewind), assert.NoError},
		{"KeySkipForward happy path", "KeySkipForward", okHttpClient(tvKeySkipForward), assert.NoError},
		{"KeySkipBackward happy path", "KeySkipBackward", okHttpClient(tvKeySkipBackward), assert.NoError},
		{"KeyRecord happy path", "KeyRecord", okHttpClient(tvKeyRecord), assert.NoError},
		{"KeyRecordingList happy path", "KeyRecordingList", okHttpClient(tvKeyRecordingList), assert.NoError},
		{"KeyRepeat happy path", "KeyRepeat", okHttpClient(tvKeyRepeat), assert.NoError},
		{"KeyLiveTv happy path", "KeyLiveTv", okHttpClient(tvKeyLiveTv), assert.NoError},
		{"KeyEpg happy path", "KeyEpg", okHttpClient(tvKeyEpg), assert.NoError},
		{"KeyProgramInformation happy path", "KeyProgramInformation", okHttpClient(tvKeyProgramInformation), assert.NoError},
		{"KeyAspectRatio happy path", "KeyAspectRatio", okHttpClient(tvKeyAspectRatio), assert.NoError},
		{"KeyExternalInput happy path", "KeyExternalInput", okHttpClient(tvKeyExternalInput), assert.NoError},
		{"KeyPipSecondaryVideo happy path", "KeyPipSecondaryVideo", okHttpClient(tvKeyPipSecondaryVideo), assert.NoError},
		{"KeyShowSubtitle happy path", "KeyShowSubtitle", okHttpClient(tvKeyShowSubtitle), assert.NoError},
		{"KeyProgramList happy path", "KeyProgramList", okHttpClient(tvKeyProgramList), assert.NoError},
		{"KeyTeleText happy path", "KeyTeleText", okHttpClient(tvKeyTeleText), assert.NoError},
		{"KeyMark happy path", "KeyMark", okHttpClient(tvKeyMark), assert.NoError},
		{"Key3dVideo happy path", "Key3dVideo", okHttpClient(tvKey3dVideo), assert.NoError},
		{"Key3dLr happy path", "Key3dLr", okHttpClient(tvKey3dLr), assert.NoError},
		{"KeyDash happy path", "KeyDash", okHttpClient(tvKeyDash), assert.NoError},
		{"KeyPreviousChannel happy path", "KeyPreviousChannel", okHttpClient(tvKeyPreviousChannel), assert.NoError},
		{"KeyFavoriteChannel happy path", "KeyFavoriteChannel", okHttpClient(tvKeyFavoriteChannel), assert.NoError},
		{"KeyQuickMenu happy path", "KeyQuickMenu", okHttpClient(tvKeyQuickMenu), assert.NoError},
		{"KeyTextOption happy path", "KeyTextOption", okHttpClient(tvKeyTextOption), assert.NoError},
		{"KeyAudioDescription happy path", "KeyAudioDescription", okHttpClient(tvKeyAudioDescription), assert.NoError},
		{"KeyEnergySaving happy path", "KeyEnergySaving", okHttpClient(tvKeyEnergySaving), assert.NoError},
		{"KeyAvMode happy path", "KeyAvMode", okHttpClient(tvKeyAvMode), assert.NoError},
		{"KeySimplink happy path", "KeySimplink", okHttpClient(tvKeySimplink), assert.NoError},
		{"KeyExit happy path", "KeyExit", okHttpClient(tvKeyExit), assert.NoError},
		{"KeyReservationProgramList happy path", "KeyReservationProgramList", okHttpClient(tvKeyReservationProgramList), assert.NoError},
		{"KeyPipChannelUp happy path", "KeyPipChannelUp", okHttpClient(tvKeyPipChannelUp), assert.NoError},
		{"KeyPipChannelDown happy path", "KeyPipChannelDown", okHttpClient(tvKeyPipChannelDown), assert.NoError},
		{"KeySwitchVideo happy path", "KeySwitchVideo", okHttpClient(tvKeySwitchVideo), assert.NoError},
		{"KeyApps happy path", "KeyApps", okHttpClient(tvKeyApps), assert.NoError},

		{"KeyPowerOff error", "KeyPowerOff", errHttpClient(), assert.Error},
		{"KeyNumber0 error", "KeyNumber0", errHttpClient(), assert.Error},
		{"KeyNumber1 error", "KeyNumber1", errHttpClient(), assert.Error},
		{"KeyNumber2 error", "KeyNumber2", errHttpClient(), assert.Error},
		{"KeyNumber3 error", "KeyNumber3", errHttpClient(), assert.Error},
		{"KeyNumber4 error", "KeyNumber4", errHttpClient(), assert.Error},
		{"KeyNumber5 error", "KeyNumber5", errHttpClient(), assert.Error},
		{"KeyNumber6 error", "KeyNumber6", errHttpClient(), assert.Error},
		{"KeyNumber7 error", "KeyNumber7", errHttpClient(), assert.Error},
		{"KeyNumber8 error", "KeyNumber8", errHttpClient(), assert.Error},
		{"KeyNumber9 error", "KeyNumber9", errHttpClient(), assert.Error},
		{"KeyUp error", "KeyUp", errHttpClient(), assert.Error},
		{"KeyDown error", "KeyDown", errHttpClient(), assert.Error},
		{"KeyLeft error", "KeyLeft", errHttpClient(), assert.Error},
		{"KeyRight error", "KeyRight", errHttpClient(), assert.Error},
		{"KeyOK error", "KeyOK", errHttpClient(), assert.Error},
		{"KeyHomeMenu error", "KeyHomeMenu", errHttpClient(), assert.Error},
		{"KeyBack error", "KeyBack", errHttpClient(), assert.Error},
		{"KeyVolumeUp error", "KeyVolumeUp", errHttpClient(), assert.Error},
		{"KeyVolumeDown error", "KeyVolumeDown", errHttpClient(), assert.Error},
		{"KeyMuteToggle error", "KeyMuteToggle", errHttpClient(), assert.Error},
		{"KeyChannelUp error", "KeyChannelUp", errHttpClient(), assert.Error},
		{"KeyChannelDown error", "KeyChannelDown", errHttpClient(), assert.Error},
		{"KeyBlue error", "KeyBlue", errHttpClient(), assert.Error},
		{"KeyGreen error", "KeyGreen", errHttpClient(), assert.Error},
		{"KeyRed error", "KeyRed", errHttpClient(), assert.Error},
		{"KeyYellow error", "KeyYellow", errHttpClient(), assert.Error},
		{"KeyPlay error", "KeyPlay", errHttpClient(), assert.Error},
		{"KeyPause error", "KeyPause", errHttpClient(), assert.Error},
		{"KeyStop error", "KeyStop", errHttpClient(), assert.Error},
		{"KeyFastForward error", "KeyFastForward", errHttpClient(), assert.Error},
		{"KeyRewind error", "KeyRewind", errHttpClient(), assert.Error},
		{"KeySkipForward error", "KeySkipForward", errHttpClient(), assert.Error},
		{"KeySkipBackward error", "KeySkipBackward", errHttpClient(), assert.Error},
		{"KeyRecord error", "KeyRecord", errHttpClient(), assert.Error},
		{"KeyRecordingList error", "KeyRecordingList", errHttpClient(), assert.Error},
		{"KeyRepeat error", "KeyRepeat", errHttpClient(), assert.Error},
		{"KeyLiveTv error", "KeyLiveTv", errHttpClient(), assert.Error},
		{"KeyEpg error", "KeyEpg", errHttpClient(), assert.Error},
		{"KeyProgramInformation error", "KeyProgramInformation", errHttpClient(), assert.Error},
		{"KeyAspectRatio error", "KeyAspectRatio", errHttpClient(), assert.Error},
		{"KeyExternalInput error", "KeyExternalInput", errHttpClient(), assert.Error},
		{"KeyPipSecondaryVideo error", "KeyPipSecondaryVideo", errHttpClient(), assert.Error},
		{"KeyShowSubtitle error", "KeyShowSubtitle", errHttpClient(), assert.Error},
		{"KeyProgramList error", "KeyProgramList", errHttpClient(), assert.Error},
		{"KeyTeleText error", "KeyTeleText", errHttpClient(), assert.Error},
		{"KeyMark error", "KeyMark", errHttpClient(), assert.Error},
		{"Key3dVideo error", "Key3dVideo", errHttpClient(), assert.Error},
		{"Key3dLr error", "Key3dLr", errHttpClient(), assert.Error},
		{"KeyDash error", "KeyDash", errHttpClient(), assert.Error},
		{"KeyPreviousChannel error", "KeyPreviousChannel", errHttpClient(), assert.Error},
		{"KeyFavoriteChannel error", "KeyFavoriteChannel", errHttpClient(), assert.Error},
		{"KeyQuickMenu error", "KeyQuickMenu", errHttpClient(), assert.Error},
		{"KeyTextOption error", "KeyTextOption", errHttpClient(), assert.Error},
		{"KeyAudioDescription error", "KeyAudioDescription", errHttpClient(), assert.Error},
		{"KeyEnergySaving error", "KeyEnergySaving", errHttpClient(), assert.Error},
		{"KeyAvMode error", "KeyAvMode", errHttpClient(), assert.Error},
		{"KeySimplink error", "KeySimplink", errHttpClient(), assert.Error},
		{"KeyExit error", "KeyExit", errHttpClient(), assert.Error},
		{"KeyReservationProgramList error", "KeyReservationProgramList", errHttpClient(), assert.Error},
		{"KeyPipChannelUp error", "KeyPipChannelUp", errHttpClient(), assert.Error},
		{"KeyPipChannelDown error", "KeyPipChannelDown", errHttpClient(), assert.Error},
		{"KeySwitchVideo error", "KeySwitchVideo", errHttpClient(), assert.Error},
		{"KeyApps error", "KeyApps", errHttpClient(), assert.Error},
	} {
		t.Run(tc.name, func(t *testing.T) {
			url, err := url.Parse(testBaseURL)
			require.NoError(t, err)

			tv := LgTV{
				baseURL: url,
				client:  tc.httpClient,
				session: testSession,
			}

			err = callKeyInputByName(t, tv, tc.method)

			tc.errorAssertion(t, err)
		})
	}
}

func callKeyInputByName(t *testing.T, tv LgTV, methodName string) error {
	actual := reflect.ValueOf(&tv).
		MethodByName(methodName).
		Call(nil)

	if v := actual[0].Interface(); v != nil {
		err, ok := v.(error)
		require.True(t, ok)

		return err
	}

	return nil
}
