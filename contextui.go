package lgtv

import (
	"encoding/xml"
	"fmt"
)

type contextUIResponse struct {
	XMLName xml.Name `xml:"envelope"`
	Mode    string   `xml:"data>mode"`
}

func (l *LgTV) GetContextUI() (string, error) {
	resp := contextUIResponse{}
	err := l.query(tvQueryContextUi, &resp)
	if err != nil {
		return "", fmt.Errorf("unable to retrieve context UI information: %w", err)
	}

	return resp.Mode, nil
}
