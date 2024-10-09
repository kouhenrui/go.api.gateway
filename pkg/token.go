package pkg

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	secretKey = []byte("your-secret-key") // 你的密钥，应该存储在安全的地方
)

type TokenClaims struct {
	User UserClaim
	jwt.RegisteredClaims
}

// UserClaim 表示用户的身份声明
type UserClaim struct {
	ID    uint   `json:"id"`    // 用户ID
	Phone string `json:"phone"` // 用户手机号码
	Role  uint   `json:"role"`  // 用户角色
	Name  string `json:"name"`  // 用户姓名
}

func NewTokenClaim(user UserClaim) *TokenClaims {

	return &TokenClaims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 设置过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     //签发时间
			Issuer:    "api-gateway",                                      // 签名颁发者
			Subject:   "sign",                                             //签名主题
		},
	}
}

// GenerateToken 生成 JWT
func GenerateToken(claims *TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("生成 token 失败: %w", err)
	}
	return signedToken, nil
}

// ValidateToken 验证 JWT
func ValidateToken(tokenString string) (*UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("验证 token 失败: %w", err)
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return &claims.User, nil
	}
	return nil, fmt.Errorf("无效的 token")
}
