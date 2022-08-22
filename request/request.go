package request

import (
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ParseBodyJson[T any](request *http.Request, data T) error {
	defer request.Body.Close()
	byteBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(byteBody, &data); err != nil {
		return err
	}

	return nil
}
