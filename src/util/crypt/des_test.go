package crypt

import (
	"testing"
)

func TestEncryptDesByEcb(t *testing.T) {
	//表格测试
	tests := []struct{ key, src string }{
		{"12345678", "https://www.baidu.com/hello.m3u8"},
		{"12436587", "https://www.baidu.com/hello.m3u8"},
		{"11335678", "https://www.baidu.com/hello.m3u8"},
		{"12347777", "https://www.baidu.com/hello.m3u8"},
	}
	for _, tt := range tests {
		if  _, err := EncryptECB(tt.src, tt.key); err != nil {
			t.Errorf("des by ecb 加密算法报错:%s", err.Error())
		}
		  
		// ciphertext, err := EncryptECB(tt.src, tt.key); 
		// if err != nil {
		// 	t.Errorf("des by ecb 加密算法报错:%s", err.Error())
		// }
		// t.Log("des by ecb 加密测试成功:", ciphertext)
	}


}

func TestDecryptDesByEcb(t *testing.T) {
	// ciphertext := "a23e37cf73da4bb3b91cdf08bdeef2949398b03aa469d40d60680864760beb44"
	// key := "20210812"
	key := "29081452"
	ciphertext := "A25E9AFE5509EE001C59656D1BB1505C857C7D478B2F829C48E15B741FB9B2956C05BE90DCB9F492941DFA44A64D88C3B747B929668ADD5900E6B665EC684F405FA3D0C3C65BCB515678E94471625E99F1147389E2DBC1484DECC6DA96F055E1424ED96DA4E539CDF4B8DF5BBF268792E482AB35B53AE6836C9D15A0814EA5F61F8C93CC23308B3A3C5B809257FBE5C9"
	ciphertext, err := DecryptECB(ciphertext, key)
	if err != nil {
		t.Errorf("des by ecb 解密算法报错:%s", err.Error())
	}
	t.Log("des by ecb 解密测试成功:", ciphertext)

}

func BenchmarkEncryptDesByEcb(b *testing.B) {
	aesKey := "1234567890"
	originData := "https://www.baidu.com/hello.m3u8"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncryptECB(originData, aesKey)
	}
}
