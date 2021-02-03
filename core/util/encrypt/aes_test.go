package encrypt

import (
	"testing"
)

func TestAesDecrypt(t *testing.T) {

	key := ""
	encrypt := ""

	d, e := AesDecrypt(key, encrypt)
	t.Log(e)
	t.Log(d)

}
