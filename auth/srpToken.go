package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

type ecb struct {
	b         cipher.Block
	blockSize int
}

type ecbEncrypter ecb

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

func newECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

type newSRPTokenOptions struct {
	Timestamp *time.Time
}

// PKCS7Padding adds padding to data
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func newSrpToken(serviceName string, opts ...*newSRPTokenOptions) string {
	now := time.Now()
	timestamp := &now

	for _, opt := range opts {
		if opt.Timestamp != nil {
			timestamp = opt.Timestamp
		}
	}

	key := "ssoisno12345678987654321"
	src := fmt.Sprintf("%v-%v", timestamp.Unix(), serviceName)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()
	data := PKCS7Padding([]byte(src), blockSize)

	encryptData := make([]byte, len(data))

	ecb := newECBEncrypter(block)
	ecb.CryptBlocks(encryptData, data)

	token := base64.URLEncoding.EncodeToString(encryptData)
	token = strings.ReplaceAll(token, "=", "")
	return token
}
