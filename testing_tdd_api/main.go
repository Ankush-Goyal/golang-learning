package testing_tdd_api

func Level3Data(url string, d TodoData) ([]byte, error) {
	return d.GetData(url)
}
