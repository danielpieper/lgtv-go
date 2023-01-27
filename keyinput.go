package lgtv

type (
	tvKey string
)

const (
	tvKeyPower                  tvKey = "1"
	tvKeyNumber0                tvKey = "2"
	tvKeyNumber1                tvKey = "3"
	tvKeyNumber2                tvKey = "4"
	tvKeyNumber3                tvKey = "5"
	tvKeyNumber4                tvKey = "6"
	tvKeyNumber5                tvKey = "7"
	tvKeyNumber6                tvKey = "8"
	tvKeyNumber7                tvKey = "9"
	tvKeyNumber8                tvKey = "10"
	tvKeyNumber9                tvKey = "11"
	tvKeyUp                     tvKey = "12"
	tvKeyDown                   tvKey = "13"
	tvKeyLeft                   tvKey = "14"
	tvKeyRight                  tvKey = "15"
	tvKeyOk                     tvKey = "20"
	tvKeyHomeMenu               tvKey = "21"
	tvKeyBack                   tvKey = "23"
	tvKeyVolumeUp               tvKey = "24"
	tvKeyVolumeDown             tvKey = "25"
	tvKeyMuteToggle             tvKey = "26"
	tvKeyChannelUp              tvKey = "27"
	tvKeyChannelDown            tvKey = "28"
	tvKeyBlue                   tvKey = "29"
	tvKeyGreen                  tvKey = "30"
	tvKeyRed                    tvKey = "31"
	tvKeyYellow                 tvKey = "32"
	tvKeyPlay                   tvKey = "33"
	tvKeyPause                  tvKey = "34"
	tvKeyStop                   tvKey = "35"
	tvKeyFastForward            tvKey = "36"
	tvKeyRewind                 tvKey = "37"
	tvKeySkipForward            tvKey = "38"
	tvKeySkipBackward           tvKey = "39"
	tvKeyRecord                 tvKey = "40"
	tvKeyRecordingList          tvKey = "41"
	tvKeyRepeat                 tvKey = "42"
	tvKeyLiveTv                 tvKey = "43"
	tvKeyEpg                    tvKey = "44"
	tvKeyProgramInformation     tvKey = "45"
	tvKeyAspectRatio            tvKey = "46"
	tvKeyExternalInput          tvKey = "47"
	tvKeyPipSecondaryVideo      tvKey = "48"
	tvKeyShowSubtitle           tvKey = "49"
	tvKeyProgramList            tvKey = "50"
	tvKeyTeleText               tvKey = "51"
	tvKeyMark                   tvKey = "52"
	tvKey3dVideo                tvKey = "400"
	tvKey3dLr                   tvKey = "401"
	tvKeyDash                   tvKey = "402"
	tvKeyPreviousChannel        tvKey = "403"
	tvKeyFavoriteChannel        tvKey = "404"
	tvKeyQuickMenu              tvKey = "405"
	tvKeyTextOption             tvKey = "406"
	tvKeyAudioDescription       tvKey = "407"
	tvKeyEnergySaving           tvKey = "409"
	tvKeyAvMode                 tvKey = "410"
	tvKeySimplink               tvKey = "411"
	tvKeyExit                   tvKey = "412"
	tvKeyReservationProgramList tvKey = "413"
	tvKeyPipChannelUp           tvKey = "414"
	tvKeyPipChannelDown         tvKey = "415"
	tvKeySwitchVideo            tvKey = "416"
	tvKeyApps                   tvKey = "417"
)

func (l *LgTV) keyInputCommand(key tvKey) error {
	return l.command(command{Name: tvCmdKeyInput, Value: key})
}

func (l *LgTV) KeyPowerOff() error           { return l.keyInputCommand(tvKeyPower) }
func (l *LgTV) KeyNumber0() error            { return l.keyInputCommand(tvKeyNumber0) }
func (l *LgTV) KeyNumber1() error            { return l.keyInputCommand(tvKeyNumber1) }
func (l *LgTV) KeyNumber2() error            { return l.keyInputCommand(tvKeyNumber2) }
func (l *LgTV) KeyNumber3() error            { return l.keyInputCommand(tvKeyNumber3) }
func (l *LgTV) KeyNumber4() error            { return l.keyInputCommand(tvKeyNumber4) }
func (l *LgTV) KeyNumber5() error            { return l.keyInputCommand(tvKeyNumber5) }
func (l *LgTV) KeyNumber6() error            { return l.keyInputCommand(tvKeyNumber6) }
func (l *LgTV) KeyNumber7() error            { return l.keyInputCommand(tvKeyNumber7) }
func (l *LgTV) KeyNumber8() error            { return l.keyInputCommand(tvKeyNumber8) }
func (l *LgTV) KeyNumber9() error            { return l.keyInputCommand(tvKeyNumber9) }
func (l *LgTV) KeyUp() error                 { return l.keyInputCommand(tvKeyUp) }
func (l *LgTV) KeyDown() error               { return l.keyInputCommand(tvKeyDown) }
func (l *LgTV) KeyLeft() error               { return l.keyInputCommand(tvKeyLeft) }
func (l *LgTV) KeyRight() error              { return l.keyInputCommand(tvKeyRight) }
func (l *LgTV) KeyOK() error                 { return l.keyInputCommand(tvKeyOk) }
func (l *LgTV) KeyHomeMenu() error           { return l.keyInputCommand(tvKeyHomeMenu) }
func (l *LgTV) KeyBack() error               { return l.keyInputCommand(tvKeyBack) }
func (l *LgTV) KeyVolumeUp() error           { return l.keyInputCommand(tvKeyVolumeUp) }
func (l *LgTV) KeyVolumeDown() error         { return l.keyInputCommand(tvKeyVolumeDown) }
func (l *LgTV) KeyMuteToggle() error         { return l.keyInputCommand(tvKeyMuteToggle) }
func (l *LgTV) KeyChannelUp() error          { return l.keyInputCommand(tvKeyChannelUp) }
func (l *LgTV) KeyChannelDown() error        { return l.keyInputCommand(tvKeyChannelDown) }
func (l *LgTV) KeyBlue() error               { return l.keyInputCommand(tvKeyBlue) }
func (l *LgTV) KeyGreen() error              { return l.keyInputCommand(tvKeyGreen) }
func (l *LgTV) KeyRed() error                { return l.keyInputCommand(tvKeyRed) }
func (l *LgTV) KeyYellow() error             { return l.keyInputCommand(tvKeyYellow) }
func (l *LgTV) KeyPlay() error               { return l.keyInputCommand(tvKeyPlay) }
func (l *LgTV) KeyPause() error              { return l.keyInputCommand(tvKeyPause) }
func (l *LgTV) KeyStop() error               { return l.keyInputCommand(tvKeyStop) }
func (l *LgTV) KeyFastForward() error        { return l.keyInputCommand(tvKeyFastForward) }
func (l *LgTV) KeyRewind() error             { return l.keyInputCommand(tvKeyRewind) }
func (l *LgTV) KeySkipForward() error        { return l.keyInputCommand(tvKeySkipForward) }
func (l *LgTV) KeySkipBackward() error       { return l.keyInputCommand(tvKeySkipBackward) }
func (l *LgTV) KeyRecord() error             { return l.keyInputCommand(tvKeyRecord) }
func (l *LgTV) KeyRecordingList() error      { return l.keyInputCommand(tvKeyRecordingList) }
func (l *LgTV) KeyRepeat() error             { return l.keyInputCommand(tvKeyRepeat) }
func (l *LgTV) KeyLiveTv() error             { return l.keyInputCommand(tvKeyLiveTv) }
func (l *LgTV) KeyEpg() error                { return l.keyInputCommand(tvKeyEpg) }
func (l *LgTV) KeyProgramInformation() error { return l.keyInputCommand(tvKeyProgramInformation) }
func (l *LgTV) KeyAspectRatio() error        { return l.keyInputCommand(tvKeyAspectRatio) }
func (l *LgTV) KeyExternalInput() error      { return l.keyInputCommand(tvKeyExternalInput) }
func (l *LgTV) KeyPipSecondaryVideo() error  { return l.keyInputCommand(tvKeyPipSecondaryVideo) }
func (l *LgTV) KeyShowSubtitle() error       { return l.keyInputCommand(tvKeyShowSubtitle) }
func (l *LgTV) KeyProgramList() error        { return l.keyInputCommand(tvKeyProgramList) }
func (l *LgTV) KeyTeleText() error           { return l.keyInputCommand(tvKeyTeleText) }
func (l *LgTV) KeyMark() error               { return l.keyInputCommand(tvKeyMark) }
func (l *LgTV) Key3dVideo() error            { return l.keyInputCommand(tvKey3dVideo) }
func (l *LgTV) Key3dLr() error               { return l.keyInputCommand(tvKey3dLr) }
func (l *LgTV) KeyDash() error               { return l.keyInputCommand(tvKeyDash) }
func (l *LgTV) KeyPreviousChannel() error    { return l.keyInputCommand(tvKeyPreviousChannel) }
func (l *LgTV) KeyFavoriteChannel() error    { return l.keyInputCommand(tvKeyFavoriteChannel) }
func (l *LgTV) KeyQuickMenu() error          { return l.keyInputCommand(tvKeyQuickMenu) }
func (l *LgTV) KeyTextOption() error         { return l.keyInputCommand(tvKeyTextOption) }
func (l *LgTV) KeyAudioDescription() error   { return l.keyInputCommand(tvKeyAudioDescription) }
func (l *LgTV) KeyEnergySaving() error       { return l.keyInputCommand(tvKeyEnergySaving) }
func (l *LgTV) KeyAvMode() error             { return l.keyInputCommand(tvKeyAvMode) }
func (l *LgTV) KeySimplink() error           { return l.keyInputCommand(tvKeySimplink) }
func (l *LgTV) KeyExit() error               { return l.keyInputCommand(tvKeyExit) }
func (l *LgTV) KeyReservationProgramList() error {
	return l.keyInputCommand(tvKeyReservationProgramList)
}
func (l *LgTV) KeyPipChannelUp() error   { return l.keyInputCommand(tvKeyPipChannelUp) }
func (l *LgTV) KeyPipChannelDown() error { return l.keyInputCommand(tvKeyPipChannelDown) }
func (l *LgTV) KeySwitchVideo() error    { return l.keyInputCommand(tvKeySwitchVideo) }
func (l *LgTV) KeyApps() error           { return l.keyInputCommand(tvKeyApps) }
