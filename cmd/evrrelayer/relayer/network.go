package relayer

import (
	"fmt"
	"github.com/Evrynetlabs/evrynet-node/evrclient"
	"net/url"
	"strings"
)

// IsWebsocketURL returns true if the given URL is a websocket URL
func IsWebsocketURL(rawurl string) bool {
	u, err := url.Parse(rawurl)
	if err != nil {
		return false
	}
	return u.Scheme == "ws" || u.Scheme == "wss"
}

// SetupWebsocketEthClient returns boolean indicating if a URL is valid websocket evrclient
func SetupWebsocketEvrClient(ethURL string) (*evrclient.Client, error) {
	if strings.TrimSpace(ethURL) == "" {
		return nil, nil
	}

	if !IsWebsocketURL(ethURL) {
		return nil, fmt.Errorf("invalid websocket eth client URL: %s", ethURL)
	}

	client, err := evrclient.Dial(ethURL)
	if err != nil {
		return nil, err
	}

	return client, nil
}

