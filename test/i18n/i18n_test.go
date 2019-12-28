package i18n

import (
	"fmt"
	"src/plugins/i18n"
	"testing"
)

func TestI18n(t *testing.T) {
	i18n.Instance.PushMessage("zh-cn", "k", "dada")
	m := i18n.Instance.GetMessage("zh-cn", "k")
	m += "s"
	fmt.Println(m)
}
