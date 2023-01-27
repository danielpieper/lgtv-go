package lgtv

import (
	"encoding/xml"
	"fmt"
)

type command struct {
	XMLName xml.Name `xml:"command"`
	Name    tvCmd    `xml:"name"`
	Value   tvKey    `xml:"value,omitempty"`
	X       int      `xml:"x,omitempty"`
	Y       int      `xml:"y,omitempty"`
}

func (a command) isCommand() {}

type (
	tvCmd string
)

const (
	tvCmdKeyInput      tvCmd = "HandleKeyInput"
	tvCmdMouseMove     tvCmd = "HandleTouchMove"
	tvCmdMouseClick    tvCmd = "HandleTouchClick"
	tvCmdTouchWheel    tvCmd = "HandleTouchWheel"
	tvCmdChangeChannel tvCmd = "HandleChannelChange"
	tvCmdScrollUp      tvCmd = "up"
	tvCmdScrollDown    tvCmd = "down"

	tvLaunchApp = "AppExecute"
)

func (l *LgTV) command(cmd command) error {
	if !l.IsAuthenticated() {
		return fmt.Errorf("unable to execute command: unauthenticated")
	}

	req, err := l.createCommandRequest("/roap/api/command", &cmd)
	if err != nil {
		return fmt.Errorf("unable to create command request: %w", err)
	}

	err = l.sendRequest(req, nil)
	if err != nil {
		return fmt.Errorf("unable to execute command: %w", err)
	}

	return nil
}
