package crypt

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
	"fmt"
)

//EncryptECB 加密
func EncryptECB(src, key string) (string, error) {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	// //对明文数据进行补码
	data = ZeroPadding(data)
	
	if len(data)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//对明文按照blocksize进行分块加密
		//必要时可以使用go关键字进行并行加密
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return fmt.Sprintf("%X", out), nil
}

//DecryptECB  解密  key必须8位
func DecryptECB(src, key string) (string, error) {
	data, err := hex.DecodeString(src)
	if err != nil {
		return "", err
	}
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = NullUnPadding(out)
	return string(out), nil
}


//ZeroPadding 0填充
func ZeroPadding(in []byte) []byte {
	length := len(in);
	if (length % 8 == 0) {
		return in;
	} else {
		blockCount := length / 8;
		out := make([]byte, (blockCount + 1) * 8)
		var i int;
		for i = 0; i < length; i++ {
			out[i] = in[i];
		}
		return out;
	}
}

// NullUnPadding 0填充减码
func NullUnPadding(in []byte) []byte {
    return bytes.TrimRight(in, string([]byte{0}));
}

//PKCS5Padding 明文补码算法
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//PKCS5UnPadding 明文减码算法
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
