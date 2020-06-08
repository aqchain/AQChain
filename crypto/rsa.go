package crypto

import (
	"AQChain/utils"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/mr-tron/base58/base58"
	"io/ioutil"
	"log"
	"os"
)

func ReadKey(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	block, _ := pem.Decode(bytes)
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	publicKey := privateKey.Public()
	return GetIDFromPublicKey(publicKey)
}

func GetIDFromPublicKey(key crypto.PublicKey) string {
	return SHA256String(GetPublicKeyString(key))
}

func GetPublicKeyString(key crypto.PublicKey) string {
	bytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		log.Println(err)
	}
	return base58.FastBase58Encoding(bytes)
}

// RSA公钥私钥
func GenerateRSAKeyPairs(bits int) (*rsa.PrivateKey, error) {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	// 生成公钥
	publicKey := &privateKey.PublicKey

	PKIX, _ := x509.MarshalPKIXPublicKey(publicKey)
	PKCS1 := x509.MarshalPKCS1PrivateKey(privateKey)

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: PKIX,
	}

	id := GetIDFromPublicKey(publicKey)

	// 创建文件
	file, err := os.Create(utils.CaculatePath() + id + "+P.pem")
	if err != nil {
		return privateKey, err
	}
	if err = pem.Encode(file, block); err != nil {
		return privateKey, err
	}

	block = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: PKCS1,
	}
	file, err = os.Create(utils.CaculatePath() + id + "+S.pem")
	if err != nil {
		return privateKey, err
	}
	if err = pem.Encode(file, block); err != nil {
		return privateKey, err
	}

	return privateKey, nil
}

func RSAVerify(theMsg string, sig []byte, PubKey []byte) bool {
	//var theMsg = "the message you want to encode 你好 世界"
	//fmt.Println("Source:", theMsg)
	//私钥签名
	//sig, _ := RsaSign([]byte(theMsg))

	fmt.Println(string(sig))
	//公钥验证
	if RsaSignVer([]byte(theMsg), sig, PubKey) != nil {
		fmt.Println("验证失败")
		return false
	} else {
		fmt.Println("验证通过")
		return true

	}

	////公钥加密
	// enc, _ := RsaEncrypt([]byte(theMsg))
	//  fmt.Println("Encrypted:", string(enc))
	//  //私钥解密
	//  decstr, _ := RsaDecrypt(enc)
	//  fmt.Println("Decrypted:", string(decstr))
}

//私钥签名
func RsaSign(data []byte, PriKey []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	//获取私钥
	block, _ := pem.Decode(PriKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed)
}

//公钥验证
func RsaSignVer(data []byte, signature []byte, PubKey []byte) error {
	hashed := sha256.Sum256(data)
	block, _ := pem.Decode(PubKey)
	if block == nil {
		return errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//验证签名
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}

// 公钥加密
func RsaEncrypt(data []byte, PubKey []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(PubKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

// 私钥解密
func RsaDecrypt(ciphertext []byte, PriKey []byte) ([]byte, error) {
	//获取私钥
	block, _ := pem.Decode(PriKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
