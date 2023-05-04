package testing_advanced_tdd

import "errors"

func Level3Data(url string, d TodoData) ([]byte, error) {
	if url == "error" {
		return nil, errors.New("errror url is not allowed")
	}
	return d.GetData(url)
}
