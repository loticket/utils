package verify

import "testing"

func TestVerify(t *testing.T) {
	t.Log("-----------")
	var date string = "2021-06-28 11:36:60"
	t.Log(VerifyTimeDate(date))
	return
}
