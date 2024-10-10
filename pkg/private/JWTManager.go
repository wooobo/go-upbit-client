package private

import (
	"crypto/sha512"
	"encoding/hex"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/url"
)

type JWTManager struct {
	AccessKey string
	SecretKey string
}

func NewJWT(accessKey, secretKey string) *JWTManager {
	return &JWTManager{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}

func (j *JWTManager) getSHA512Hash(query string) string {
	hash := sha512.Sum512([]byte(query))
	return hex.EncodeToString(hash[:])
}

func (j *JWTManager) CreateTokenWithQuery(values url.Values) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"access_key":     j.AccessKey,
		"nonce":          uuid.New().String(),
		"query":          values.Encode(),
		"query_hash":     j.getSHA512Hash(values.Encode()),
		"query_hash_alg": "SHA512",
	})

	signedToken, _ := token.SignedString([]byte(j.SecretKey))
	return "Bearer " + signedToken
}
