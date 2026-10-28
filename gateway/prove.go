package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Prove asks the identity gateway whether this app is a registered client.
// Boot fails closed if the gateway does not know the app.
func Prove(ctx context.Context, identityURL, appID string) error {
	if identityURL == "" || appID == "" {
		return fmt.Errorf("GATEWAY_UNSET")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, strings.TrimRight(identityURL, "/")+"/v1/identity/clients/"+appID, nil)
	if err != nil {
		return err
	}
	res, err := (&http.Client{Timeout: 5 * time.Second}).Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("CLIENT_UNKNOWN")
	}
	var doc struct {
		ClientID string `json:"clientId"`
	}
	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		return err
	}
	if doc.ClientID != appID {
		return fmt.Errorf("CLIENT_MISMATCH")
	}
	return nil
}
