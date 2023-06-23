package pandora

import (
	"encoding/json"
	"fmt"
	logger "goPandoraAdmin-Server/internal/log"
	"goPandoraAdmin-Server/model"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// GetShareTokenByFakeopen 从FakeOpen获取share token
func GetShareTokenByFakeopen(shareTokenStruct model.FakeOpenShareTokenRequest) (model.FakeOpenShareTokenRespond, error) {
	logger.Info("begin to get share token by fakeopen...")
	const target = "https://ai.fakeopen.com/token/register"
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(GetProxyURL()),
		},
	}
	var shareToken model.FakeOpenShareTokenRespond
	dataFrom := url.Values{}
	dataFrom.Add("unique_name", shareTokenStruct.UniqueName)
	dataFrom.Add("access_token", shareTokenStruct.AccessToken)
	dataFrom.Add("expires_in", fmt.Sprintf("%v", shareTokenStruct.ExpiresIn))
	if shareTokenStruct.SiteLimit != nil {
		dataFrom.Add("site_limit", *shareTokenStruct.SiteLimit)
	}
	dataFrom.Add("show_conversations", fmt.Sprintf("%v", shareTokenStruct.ShowConversations))
	dataFrom.Add("show_userinfo", fmt.Sprintf("%v", shareTokenStruct.ShowUserInfo))
	body := strings.NewReader(dataFrom.Encode())

	req, err := http.NewRequest(http.MethodPost, target, body)
	if err != nil {
		return shareToken, fmt.Errorf("http.NewRequest failed: %w", err)
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return shareToken, fmt.Errorf("client.Do failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return shareToken, fmt.Errorf("resp.StatusCode != http.StatusOK")
	}
	rawJSON, err := io.ReadAll(resp.Body)
	if err != nil {
		return shareToken, fmt.Errorf("io.ReadAll failed: %w", err)
	}
	err = json.Unmarshal(rawJSON, &shareToken)
	if err != nil {
		return shareToken, fmt.Errorf("json.Unmarshal failed: %w", err)
	}
	return shareToken, nil
}
