package lgtv

import (
	"fmt"
)

type tvQuery string

const (
	tvQueryCurrentChannel tvQuery = "cur_channel"
	tvQueryChannelList    tvQuery = "channel_list"
	tvQueryContextUi      tvQuery = "context_ui"
	tvQueryVolume         tvQuery = "volume_info"
	tvQueryScreen         tvQuery = "screen_image"
	tvQuery3D             tvQuery = "is_3d"
)

func (l *LgTV) query(query tvQuery, resp any) error {
	if !l.IsAuthenticated() {
		return fmt.Errorf("unable to query info: unauthenticated")
	}

	req, err := l.createInfoRequest(string(query))
	if err != nil {
		return fmt.Errorf("unable to create info request: %w", err)
	}

	err = l.sendRequest(req, resp)
	if err != nil {
		return fmt.Errorf("unable to query info: %w", err)
	}

	return nil
}
