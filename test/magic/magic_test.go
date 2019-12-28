package magic

import (
	"fmt"
	"src/instruct"
	"testing"
)

func TestMagic(t *testing.T) {
	m := instruct.NewMagic('t', 'h', 'r', 'd')
	fmt.Println(m)
	//	fmt.Println(util.RandomUUID())
}
