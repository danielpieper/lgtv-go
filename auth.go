package lgtv

import (
	"encoding/xml"
	"fmt"
)

type authCommand struct {
	XMLName xml.Name `xml:"auth"`
	Type    string   `xml:"type"`
	Value   string   `xml:"value,omitempty"`
}

func (a authCommand) isCommand() {}

type authResponse struct {
	Session int `xml:"session"`
}

func (l *LgTV) RequestPairingKey() error {
	req, err := l.createCommandRequest(
		"/roap/api/auth",
		authCommand{Type: "AuthKeyReq"},
	)
	if err != nil {
		return fmt.Errorf("unable to create request pairing key request: %w", err)
	}

	if err := l.sendRequest(req, nil); err != nil {
		return fmt.Errorf("unable to request pairing key: %w", err)
	}

	return nil
}

func (l *LgTV) Authenticate(pairingKey string) error {
	req, err := l.createCommandRequest(
		"/roap/api/auth",
		authCommand{Type: "AuthReq", Value: pairingKey},
	)
	if err != nil {
		return fmt.Errorf("unable to create authentication request: %w", err)
	}

	resp := authResponse{}
	err = l.sendRequest(req, &resp)
	if err != nil {
		return fmt.Errorf("unable to send authentication request: %w", err)
	}

	if resp.Session == 0 {
		return fmt.Errorf("unable to authenticate: session was not provided")
	}

	l.session = resp.Session

	return nil
}

func (l *LgTV) IsAuthenticated() bool {
	return l.session != 0
}
