package utility

import (
	"aery-go/internal/consts"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func init() {

}

func EncryptAES(plaintext string) (string, error) {
	key := []byte(consts.Key)

	// 检查密钥长度是否有效
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("invalid key size: key must be 16, 24, or 32 bytes")
	}

	// 创建 AES 加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 创建随机的 IV（初始化向量），长度为 AES 块大小
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 使用 CFB 模式加密明文
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	// 返回 Base64 编码后的密文
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// DecryptAES 使用 AES 解密数据
func DecryptAES(ciphertext string) (string, error) {
	// 解码 Base64 编码的密文
	decodedCiphertext, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// 创建 AES 解密块，key 必须是 16, 24, 或 32 字节长度
	block, err := aes.NewCipher([]byte(consts.Key))
	if err != nil {
		return "", err
	}

	// 检查密文长度是否小于块大小
	if len(decodedCiphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	// 提取 IV
	iv := decodedCiphertext[:aes.BlockSize]
	decodedCiphertext = decodedCiphertext[aes.BlockSize:]

	// 使用 CFB 模式解密
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decodedCiphertext, decodedCiphertext)

	// 返回解密后的明文
	return string(decodedCiphertext), nil
}
