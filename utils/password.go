package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/astaxie/beego/logs"
	"math/rand"
	"reflect"
	"time"
)

const letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GetRandomBytes 生成随机字符串
func GetRandomBytes(n int) []byte {
	result := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return result
}

// GetRandomString 生成随机字符串
func GetRandomString(n int) string {
	return string(GetRandomBytes(n))
}

// EncryptPassword 密码加密，
// 加密结果 = 1字节版本号 + 1字节盐值长度 + 盐值 + 加密后的数据
func EncryptPassword(str string) string {
	lenSalt := 10                       // 盐值长度
	dataSalt := GetRandomBytes(lenSalt) // 随机一串字符作为盐值
	dataStr := []byte(str)
	hash := sha256.New()           // 加密算法
	hash.Write(dataSalt)           // 盐值
	hash.Write(dataStr)            // 密码
	dataEncrypted := hash.Sum(nil) // Sha256哈希后的值（加密后的数据）

	// 拼接数据，组装结果
	// 加密结果=1字节版本号+1字节盐值长度+盐值+加密后的数据
	dataAll := make([]byte, 2)
	dataAll[0] = 1             // 此加密函数的版本（以后若更改算法可以通过此字段兼容多个版本）
	dataAll[1] = byte(lenSalt) // 长度
	dataAll = append(dataAll, dataSalt...)
	dataAll = append(dataAll, dataEncrypted...)
	return base64.StdEncoding.EncodeToString(dataAll) // Base64编码
}

// VerifyPassword 密码验证
func VerifyPassword(pwdOrig, pwdEncrypted string) bool {
	// Base64解码
	dataOrig := []byte(pwdOrig)
	dataDecode, err := base64.StdEncoding.DecodeString(pwdEncrypted)
	if err != nil {
		logs.Error(err)
		return false
	}

	// 校验数据长度
	nEncrypted := len(dataDecode)
	nOrig := len(dataOrig)
	if nEncrypted <= 2+nOrig {
		logs.Error("Length error:", nEncrypted, nOrig)
		return false
	}

	// 校验数据前两个字节
	if dataDecode[0] != 1 || 34+int(dataDecode[1]) != nEncrypted {
		logs.Error("Header error:", dataDecode[0], dataDecode[1])
		return false
	}

	// 数据分解
	nType := dataDecode[0]
	nSalt := dataDecode[1]
	dataSalt := dataDecode[2 : nSalt+2]
	dataEncrypted := dataDecode[nSalt+2:]

	// 结果判定
	result := false
	if nType == 1 {
		hash := sha256.New()
		hash.Write(dataSalt)
		hash.Write(dataOrig)
		dataTmp := hash.Sum(nil)
		if reflect.DeepEqual(dataTmp, dataEncrypted) {
			result = true
		}
	}

	return result
}
