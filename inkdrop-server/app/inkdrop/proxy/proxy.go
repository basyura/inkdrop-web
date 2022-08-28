package proxy

import (
	"fmt"
	"inkdrop-server/app/config"
	"io/ioutil"
	"net/http"
	"net/url"
)

var root string
var userID string
var password string

func init() {
	conf := config.Instance()
	root = conf.Inkdrop.Root
	userID = conf.Inkdrop.UserID
	password = conf.Inkdrop.Password
}

func Get(path string) ([]byte, error) {

	endpoint, err := url.JoinPath(root, path)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to new request : %w", err)
	}
	req.SetBasicAuth(userID, password)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to Do request : %w", err)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(b))
	return b, err
}
