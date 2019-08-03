package lib

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)


func Crypto(str string) string{
	md5HashInBytes := md5.Sum([]byte(str))
	md5HashInString := hex.EncodeToString(md5HashInBytes[:])
	fmt.Println(md5HashInString)
	return md5HashInString
}

func MakeShortUrl(url string) ShortUrl{
	short := Crypto(url)
	return ShortUrl{url, short}
}