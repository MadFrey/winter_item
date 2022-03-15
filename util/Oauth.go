package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func GetTokenAuthUrl(code string) string {
	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		"096d20884beb1200e0fe", "aad728aba806256a39b7ecab684989ffbc0fc43f", code,
	)
}

func GetToken(url string) (*Token, error) {

	// 形成请求
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}

	// 将响应体解析为 token，并返回
	var token Token
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}
	return &token, nil
}

func GetUserInfo(token *Token) (map[string]interface{}, error) {
	// 形成请求
	var userInfoUrl = "https://api.github.com/user?access_token=" + token.AccessToken
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}

	var userInfo = make(map[string]interface{})
	fmt.Println(userInfo)
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}
