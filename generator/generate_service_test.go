package generator

import (
	"testing"
)

func TestNewGenerateService(t *testing.T) {
	g:=NewGenerateService("hello4","http",false,false,false,[]string{})
	g.Generate()
}