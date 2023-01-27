package lgtv

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type tvCommand interface {
	isCommand()
}

type LgTV struct {
	client  *http.Client
	baseURL *url.URL
	session int
}

type baseResponse struct {
	XMLName     xml.Name `xml:"envelope"`
	ErrorCode   int      `xml:"ROAPError"`
	ErrorDetail string   `xml:"ROAPErrorDetail"`
}

func New(baseURL string) (LgTV, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return LgTV{}, fmt.Errorf("unable to parse baseURL: %w", err)
	}
	return LgTV{
		client:  &http.Client{Timeout: time.Duration(1) * time.Second},
		baseURL: url,
	}, nil
}

func (l *LgTV) createCommandRequest(path string, cmd tvCommand) (*http.Request, error) {
	xmlReq, err := xml.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal command: %w", err)
	}

	xmlbody := xml.Header + string(xmlReq)

	req, err := http.NewRequest("POST", l.baseURL.JoinPath(path).String(), strings.NewReader(xmlbody))
	if err != nil {
		return nil, fmt.Errorf("unable to create post request: %w", err)
	}

	req.Header.Set("Content-type", `application/atom+xml`)

	return req, nil
}

func (l *LgTV) createInfoRequest(target string) (*http.Request, error) {
	const path = "/roap/api/data"

	url := l.baseURL.JoinPath(path)
	qry := url.Query()
	qry.Set("target", target)
	url.RawQuery = qry.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create info request: %w", err)
	}

	req.Header.Set("Content-type", `application/atom+xml`)

	return req, nil
}

func (l *LgTV) sendRequest(req *http.Request, tvResp any) error {
	resp, err := l.client.Do(req)
	if err != nil {
		return fmt.Errorf("unable to send request: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response: %w", err)
	}

	xmlResp := struct {
		XMLName     xml.Name `xml:"envelope"`
		ErrorCode   int      `xml:"ROAPError"`
		ErrorDetail string   `xml:"ROAPErrorDetail"`
	}{}
	if err := xml.Unmarshal(respBytes, &xmlResp); err != nil {
		return fmt.Errorf("unable to xml unmarshal base response: %w", err)
	}

	if xmlResp.ErrorCode != 200 {
		return fmt.Errorf("lg tv returned error %d: %s", xmlResp.ErrorCode, xmlResp.ErrorDetail)
	}
	if tvResp != nil {
		if err := xml.Unmarshal(respBytes, tvResp); err != nil {
			return fmt.Errorf("unable to xml unmarshal tv response: %w", err)
		}
	}

	return nil
}
