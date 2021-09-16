package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"pool_backend/src/global"
	"time"

	"github.com/tidwall/gjson"
	"golang.org/x/crypto/bcrypt"
)

//EncryptPwd 密码加密
func EncryptPwd(userPassword string) ([]byte, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		global.Logger.Warn("密码加密方法报错!", err.Error())
		return nil, errors.New("密码加密错误！")
	}
	return pwd, nil
}

//VerifyPwd 密码校验
func VerifyPwd(hashed, userPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		global.Logger.Warn("密码解密方法报错!", err.Error())
		return false
	}
	return true
}

//GetJSON 获取远程的Json并校验
func GetJSON(url string) (jsonString string, err error) {
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	jsonString = string(data)

	if !gjson.Valid(jsonString) {
		err = errors.New(url + "\n" + "Json Invalid ！！！")
	}

	return jsonString, err
}

//GenRandomNumber 生成6位随机字符串
func GenRandomNumber() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}


//Decimal float64保留2位小数
func Decimal(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}
