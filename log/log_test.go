package log

import "testing"

func Test(t *testing.T) {
	Logger.Error()
	t.Log(123)
}
