package lib

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Crypto(str string){
	md5HashInBytes := md5.Sum([]byte(str))
	md5HashInString := hex.EncodeToString(md5HashInBytes[:])
	fmt.Println(md5HashInString)
}