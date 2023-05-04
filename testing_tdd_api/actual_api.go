package testing_tdd_api

import (
	"io/ioutil"
	"net/http"
)

type ActualTodoData struct{}

func (atd *ActualTodoData) GetData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
