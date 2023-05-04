package testing_basic

func NumberDescrptionOne(n int) string {
	if n < 0 {
		return "negative"
	}
	if n >= 0 && n <= 9 {
		return "single digit"
	}
	if n >= 10 && n <= 99 {
		return "double digit"
	}
	if n >= 100 && n <= 999 {
		return "double digit"
	}

	return "others"
	//return ""
}
