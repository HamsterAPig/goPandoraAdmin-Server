package pandora

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strings"
)

type AccessTokenPayload struct {
	Profile struct {
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
	} `json:"https://api.openai.com/profile"`
	Auth struct {
		UserID string `json:"user_id"`
	} `json:"https://api.openai.com/auth"`
	Iss   string   `json:"iss"`
	Sub   string   `json:"sub"`
	Aud   []string `json:"aud"`
	Iat   int      `json:"iat"`
	Exp   int      `json:"exp"`
	Azp   string   `json:"azp"`
	Scope string   `json:"scope"`
}

// CheckAccessToken 检查token并且返回payload
func CheckAccessToken(accessToken string) (AccessTokenPayload, error) {
	var ast AccessTokenPayload
	// 从Pandora的源码里面拿到的openai的公钥
	publicKey := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA27rOErDOPvPc3mOADYtQ
BeenQm5NS5VHVaoO/Zmgsf1M0Wa/2WgLm9jX65Ru/K8Az2f4MOdpBxxLL686ZS+K
7eJC/oOnrxCRzFYBqQbYo+JMeqNkrCn34yed4XkX4ttoHi7MwCEpVfb05Qf/ZAmN
I1XjecFYTyZQFrd9LjkX6lr05zY6aM/+MCBNeBWp35pLLKhiq9AieB1wbDPcGnqx
lXuU/bLgIyqUltqLkr9JHsf/2T4VrXXNyNeQyBq5wjYlRkpBQDDDNOcdGpx1buRr
Z2hFyYuXDRrMcR6BQGC0ur9hI5obRYlchDFhlb0ElsJ2bshDDGRk5k3doHqbhj2I
gQIDAQAB
-----END PUBLIC KEY-----`

	// 解析token
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
		if nil != err {
			return nil, fmt.Errorf("failed to parse public key: %v", err)
		}
		return publicKey, nil
	})

	if nil != err {
		return ast, fmt.Errorf("failed to parse token: %v", err)
	}

	// 验证 JWT 的有效性
	if !token.Valid {
		return ast, fmt.Errorf("invalid JWT")
	}

	// 获取 payload
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ast, fmt.Errorf("failed to get JWT claims")
	}
	if _, ok := claims["scope"]; !ok {
		return ast, fmt.Errorf("miss scope")
	}
	scope := claims["scope"]
	if !strings.Contains(scope.(string), "model.read") || !strings.Contains(scope.(string), "model.request") {
		return ast, fmt.Errorf("invalid scope")
	}
	_, ok1 := claims["https://api.openai.com/auth"]
	_, ok2 := claims["https://api.openai.com/profile"]
	if !ok1 || !ok2 {
		return ast, fmt.Errorf("belonging to an unregistered user")
	}

	jsonBytes, err := json.Marshal(claims)
	if err != nil {
		return ast, fmt.Errorf("failed to marshal claims")
	}
	// 将 JSON 转换为具体的结构体
	err = json.Unmarshal(jsonBytes, &ast)
	if err != nil {
		return ast, fmt.Errorf("failed to unmarshal claims")
	}
	return ast, err
}
