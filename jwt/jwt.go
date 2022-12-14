package jwt

import (
	"errors"
	gjwt "github.com/golang-jwt/jwt/v4"
	uerror "github.com/lffwl/utility/error"
	"net/http"
	"strings"
	"time"
)

var (
	// DefaultExpiryAt default 2 hour
	DefaultExpiryAt       = 2 * time.Hour
	DefaultRefreshTokenAt = 2 * time.Hour
	DefaultTokenKey       = "Access-Token"
	DefaultEncryptKey     = []byte("j9ZREFUDxC0eA02uX6O0pIQ7AnmRuM5v")
	TokenLookupItems      = []string{"header", "query", "cookie"}
)

type Config struct {

	// expiry time
	ExpiryAt time.Duration

	// refresh token at
	RefreshTokenAt time.Duration

	// encrypt key
	EncryptKey []byte

	// token key
	TokenKey string

	// token type get
	TokenLookup string

	// issuer
	Issuer string
}

type Jwt struct {
	conf Config
}

// NewJwt new Jwt
func NewJwt(config Config) *Jwt {

	// default Jwt
	// expiry time
	if config.ExpiryAt == 0 {
		config.ExpiryAt = DefaultExpiryAt
	}

	// refresh token at
	if config.RefreshTokenAt == 0 {
		config.RefreshTokenAt = DefaultRefreshTokenAt
	}

	// token key
	if config.TokenKey == "" {
		config.TokenKey = DefaultTokenKey
	}

	// encrypt key
	if config.EncryptKey == nil {
		config.EncryptKey = DefaultEncryptKey
	}

	// token lookup
	if config.TokenLookup != "" {
		config.TokenLookup = strings.ToLower(config.TokenLookup)
	}

	// return Jwt
	return &Jwt{
		conf: config,
	}

}

type UJwtCustomClaims struct {
	Data interface{} `json:"data"`
	gjwt.RegisteredClaims
}

// CreateToken createToken
func (u *Jwt) CreateToken(data interface{}) (string, error) {

	var (
		tokenString string
		err         error
	)

	// Create the claims
	claims := UJwtCustomClaims{
		data,
		gjwt.RegisteredClaims{
			ExpiresAt: gjwt.NewNumericDate(time.Now().Add(u.conf.ExpiryAt)),
			IssuedAt:  gjwt.NewNumericDate(time.Now()),
			NotBefore: gjwt.NewNumericDate(time.Now()),
			Issuer:    u.conf.Issuer,
		},
	}

	// create token
	token := gjwt.NewWithClaims(gjwt.SigningMethodHS256, claims)

	// create string
	if tokenString, err = token.SignedString(u.conf.EncryptKey); err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken ParseToken
func (u *Jwt) ParseToken(tokenString string) (*gjwt.Token, error) {

	// parser token
	parser := gjwt.NewParser()
	jwtToken, _, err := parser.ParseUnverified(tokenString, &UJwtCustomClaims{})
	if err != nil {
		return nil, err
	}

	return jwtToken, nil

}

// GetData get id
// ?????????????????????
func (u *Jwt) GetData(r *http.Request) (interface{}, error) {
	// token
	token := u.GetToken(r)

	// get token
	jwtToken, err := u.ParseToken(token)
	if err != nil {
		return nil, err
	}

	// get claims
	claims, ok := jwtToken.Claims.(*UJwtCustomClaims)
	if ok {
		return claims.Data, err
	}

	return nil, errors.New("parse claims error")
}

// VerifyToken id
func (u *Jwt) VerifyToken(r *http.Request) (interface{}, uerror.HighError) {

	highError := uerror.HighError{}

	// token
	token := u.GetToken(r)

	// get token
	jwtToken, err := u.ParseToken(token)
	if err != nil {
		highError.Code = uerror.HighErrorNotAuthCode
		highError.Error = err
		return nil, highError
	}

	// get claims
	claims, ok := jwtToken.Claims.(*UJwtCustomClaims)
	if !ok {
		highError.Code = uerror.HighErrorNotAuthCode
		highError.Error = errors.New("token parse error")
		return nil, highError
	}

	if claims.Issuer != u.conf.Issuer {
		highError.Code = uerror.HighErrorNotAuthCode
		highError.Error = errors.New("token verify error")
		return nil, highError
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		highError.Code = uerror.HighErrorAuthFailedCode
		highError.Error = errors.New("token expires")
		return nil, highError
	}

	return claims.Data, highError
}

// GetToken get token
func (u *Jwt) GetToken(r *http.Request) string {

	// not set lookup
	if u.conf.TokenLookup == "" {
		token := ""

		if token = r.Header.Get(u.conf.TokenKey); token != "" {
			return token
		}

		if token = r.FormValue(u.conf.TokenKey); token != "" {
			return token
		}

		if cookie, _ := r.Cookie(u.conf.TokenKey); cookie != nil {
			return cookie.Value
		}

		return token
	}

	// header
	if u.conf.TokenLookup == TokenLookupItems[0] {
		return r.Header.Get(u.conf.TokenKey)
	}

	// query
	if u.conf.TokenLookup == TokenLookupItems[1] {
		return r.FormValue(u.conf.TokenKey)
	}

	// cookie
	if u.conf.TokenLookup == TokenLookupItems[2] {
		cookie, _ := r.Cookie(u.conf.TokenKey)
		return cookie.String()
	}

	return ""
}

func (u *Jwt) RefreshToken(r *http.Request) (string, error) {
	// token
	token := u.GetToken(r)

	// parser token
	jwtToken, err := u.ParseToken(token)
	if err != nil {
		return "", err
	}

	claims, ok := jwtToken.Claims.(*UJwtCustomClaims)
	if !ok {
		return "", errors.New("token error")
	}

	if claims.Issuer != u.conf.Issuer {
		return "", errors.New("token verify error")
	}

	// ???????????????????????????????????????
	if time.Now().Before(claims.ExpiresAt.Time.Add(u.conf.RefreshTokenAt)) {
		return u.CreateToken(claims.Data)
	}

	return "", errors.New("refresh token expire")
}
