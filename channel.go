package lgtv

import (
	"encoding/xml"
	"fmt"
)

type TVChannel struct {
	ChannelType  string `xml:"chtype"`
	ChannelName  string `xml:"chname"`
	SourceIndex  int    `xml:"sourceIndex"`
	PhysicalNum  int    `xml:"physicalNum"`
	Major        int    `xml:"major"`
	DisplayMajor int    `xml:"displayMajor"`
	Minor        int    `xml:"minor"`
	DisplayMinor int    `xml:"displayMinor"`
}

type channelListResponse struct {
	XMLName xml.Name    `xml:"envelope"`
	Data    []TVChannel `xml:"data"`
}

func (l *LgTV) GetChannelList() ([]TVChannel, error) {
	resp := channelListResponse{}
	err := l.query(tvQueryChannelList, &resp)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve channel list: %w", err)
	}

	return resp.Data, nil
}

func (l *LgTV) GetCurrentChannel() (TVChannel, error) {
	resp := channelListResponse{}
	err := l.query(tvQueryCurrentChannel, &resp)
	if err != nil {
		return TVChannel{}, fmt.Errorf("unable to retrieve current channel: %w", err)
	}

	if len(resp.Data) != 1 {
		return TVChannel{}, fmt.Errorf("unable to retrieve current channel: no result")
	}

	return resp.Data[0], nil
}
