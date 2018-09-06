//Author : dmc
//
//Date: 2018/9/6 上午10:58
//
//Description: 
package generator

import (
	"testing"
)

func TestNewGeneratePkg(t *testing.T) {
	g := NewGeneratePkg("hello")
	g.Generate()
}

func init()  {
	setDefaults()
}
