package auth

import (
	"github.com/xxtea/xxtea-go/xxtea"
)

const (
	KEY = "78sk89w8dh09sll892oa9n8hs9s892jb"
)

func Encrypt(debug bool, src []byte) (dst []byte) {
	if debug {
		dst = src
	} else {
		dst = xxtea.Encrypt(src, []byte(KEY))
	}
	return
}

///解密
func Decrypt(debug bool, src []byte) (dst []byte) {
	if debug {
		dst = src
	} else {
		dst = xxtea.Decrypt(src, []byte(KEY))
	}
	return
}
