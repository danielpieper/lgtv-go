package lgtv

import (
	"encoding/xml"
	"fmt"
)

type TVVolume struct {
	Mute     bool `xml:"mute"`
	MinLevel int  `xml:"minLevel"`
	MaxLevel int  `xml:"maxLevel"`
	Level    int  `xml:"level"`
}

type volumeResponse struct {
	XMLName xml.Name `xml:"envelope"`
	Data    TVVolume `xml:"data"`
}

func (l *LgTV) GetVolume() (TVVolume, error) {
	resp := volumeResponse{}
	err := l.query(tvQueryVolume, &resp)
	if err != nil {
		return TVVolume{}, fmt.Errorf("unable to retrieve volume information: %w", err)
	}

	return resp.Data, nil
}

func (l *LgTV) IsMuted() (bool, error) {
	vol, err := l.GetVolume()
	if err != nil {
		return false, fmt.Errorf("unable to retrieve mute status: %w", err)
	}

	return vol.Mute, nil
}
