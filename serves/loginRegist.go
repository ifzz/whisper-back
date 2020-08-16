package serves

import (
	"crypto/md5"
	"encoding/base64"
	"whisper/sqls"
)

//Login 传入mail和密码，验证密码是否正确，不正确返回false
func Login(mail string, pw string) (result bool) {
	//hashed := Myhash(pw)
	pwFromDb := sqls.Login(mail)
	if pw == pwFromDb {
		println("用户登录成功", mail)
		return true
	}
	return false

}

//Myhash 计算密码的哈希值
func Myhash(pw string) string {
	afterHash := md5.New().Sum([]byte(pw))
	after64 := base64.StdEncoding.EncodeToString(afterHash)
	return after64
}
