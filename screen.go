package lgtv

import (
	"encoding/xml"
	"fmt"
	"image"
	_ "image/jpeg"
)

type screenResponse struct {
	XMLName xml.Name `xml:"envelope"`
	Mode    string   `xml:"data>mode"`
}

func (l *LgTV) GetScreen() (image.Image, error) {
	if !l.IsAuthenticated() {
		return nil, fmt.Errorf("unable to query info: unauthenticated")
	}

	req, err := l.createInfoRequest(string(tvQueryScreen))
	if err != nil {
		return nil, fmt.Errorf("unable to create info request: %w", err)
	}

	resp, err := l.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to send request: %w", err)
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response as image: %w", err)
	}

	return img, nil
}
