package pkg

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	secretKey = []byte("your-secret-key") // 你的密钥，应该存储在安全的地方
)

// Claims 自定义声明
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT
func GenerateToken(username string) (string, error) {
	// 设置声明
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 设置过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     //签发时间
			Issuer:    "api-gateway",                                      // 签名颁发者
			Subject:   "sign",                                             //签名主题
		},
	}

	// 生成 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("生成 token 失败: %w", err)
	}

	return signedToken, nil
}

// ValidateToken 验证 JWT
func ValidateToken(tokenString string) (*Claims, error) {
	// 解析 token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证 token 签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("验证 token 失败: %w", err)
	}

	// 验证 claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("无效的 token")
}
