package gra

import (
	"testing"
)

func TestName(t *testing.T) {
	Main()
}
func TestGetRandStr(t *testing.T) {
	s := GetRandStr(4)
	ImgText(200, 100, s)
}
