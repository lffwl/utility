package jwt

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestVerifyToken(t *testing.T) {

	tJwt := NewJwt(Config{
		ExpiryAt: 4 * time.Second,
	})

	token, err := tJwt.CreateToken(111)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("token：", token)

	req, err := http.NewRequest("", "", nil)

	req.Header.Set(DefaultTokenKey, token)

	data, highError := tJwt.VerifyToken(req)
	if highError.Error != nil {
		t.Error(highError.Error)
		return
	}

	fmt.Println("data：", data)

}

func TestRefreshToken(t *testing.T) {
	tJwt := NewJwt(Config{
		ExpiryAt: 1 * time.Second,
	})

	token, err := tJwt.CreateToken(111)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("token：", token)

	req, err := http.NewRequest("", "", nil)

	req.Header.Set(DefaultTokenKey, token)

	time.Sleep(2 * time.Second)

	newToken, err := tJwt.RefreshToken(req)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("newToken：", newToken)
}
