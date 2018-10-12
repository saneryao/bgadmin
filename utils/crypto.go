package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

/* 功能：默认私钥解密
 * 参数：strEncrypted表示已加密过的数据转换成的base64编码字符串
 * 返回值：original表示解密成功后的原始明文数据，err表示解密失败时的错误信息
 */
func RsaDecryptByDefaultPrikey(strEncrypted string) (original []byte, err error) {
	var bytesEncrypted []byte
	bytesEncrypted, err = base64.StdEncoding.DecodeString(strEncrypted)
	if err != nil {
		return
	}

	original, err = RsaDecryptByPrikey(bytesEncrypted, defPriKey)
	return
}

/* 功能：默认私钥解密
 * 参数：encrypted表示已加密过的数据，priKey表示私钥数据
 * 返回值：original表示解密成功后的原始明文数据，err表示解密失败时的错误信息
 */
func RsaDecryptByPrikey(encrypted, priKey []byte) (original []byte, err error) {
	block, _ := pem.Decode(priKey) //将密钥解析成实例
	if block == nil {
		err = errors.New("private key error!")
		return
	}

	var priv *rsa.PrivateKey
	priv, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	original, err = rsa.DecryptPKCS1v15(rand.Reader, priv, encrypted) //RSA算法解密
	return
}

// 默认私钥
var defPriKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQC02RFpif6PfAWPGz0lb5ZMEDikUfBy0uUvz2WyT4IL/DSvHOms
BweAmapnD+9z9L9zc8e79PRPHYVDn6XmTlTQi7+B+hluQe9juddFCJzCOyblt5lA
Ytcxokb/TnZYDkeAb8V+Yn8k3WGeV4E2QzhF6M1+bk35rSkwGXSJ//ZXcwIDAQAB
AoGBAJuvKmAynWQkHhz+E/uAPmCorjo3F8mhaA9qeV+xNoDe0vPU7gxb2MeKgblL
bJXkIlRfXZie3RqmjSyT6RuEX1nd2BU3nZ5DkYTnzrVvlurPf7/bwYbz0tDBV3gK
4j8tYZOsrDHzQaht/UxvWpqCmYsghiRiTxl1plOM9E6lFI3BAkEA35MhmJ2dl/nj
LUOt+LXVj+CtnAfJlLzrJKsmHBPMJ3ydCrNlMRyysHaF5b0nvvRBUnjbMkLX4RQI
HexdSv6LIwJBAM8TlBVGrkqCBOUGPU7PJVdc+4PvUybitK4WSJxFqP0ta81A7qkK
sOd+s0pIzI/0Rq5Wbi3qdm/68Rl4rRhqr3ECQQCNAgjwbohUr+BeTg9Ni7GUOwqE
HE1BKB3OVLuGfd4HEYsikp7B6O2yai5tBW4p+3evglYNTydE6BNufAMjJ4OtAkEA
tL2P+OoSYnINt6C0jj77hr9fwI55c21Y6sIEanHax/CHMUXFicINGmFaODJGajd6
IdzKmkUVTzQmkGbmHnOv4QJAUS38+kSd4RWpzl57ZPddOzgpnrJ2Nvo9Pz570irg
VEWrXXKYicJWl4MXA534yQSm3Kv0gkFtUlcaQ+xNVl9WTw==
-----END RSA PRIVATE KEY-----
`)
