package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

//PageInfo 通用pageInfo
type PageInfo struct {
	Page      uint64  `json:"page"`
	Size      uint64  `json:"size"`
	Total     int64   `json:"total"`
	TotalPage int64   `json:"totalPage"`
}

//Swagger swagger
type Swagger struct {
	Swagger string                 `json:"swagger"`
	Info    map[string]interface{} `json:"info"`
	Paths   map[string]interface{} `json:"paths"`
}



/**
Gorm中自定义Time类型的JSON字段格式
将struct成员类型由time.Time改为JsonTime, 则可实现自定义json序列化后的时间格式
*/

//JSONTime format json time field by myself
type JSONTime struct {
	time.Time
}

//MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (item *JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", item.Local().Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

//Value insert timestamp into mysql need this function.
func (item *JSONTime) Value() (driver.Value, error) {
	//若当前时间字段没有取值
	if item == nil {
		return nil, nil
	}

	var zeroTime time.Time
	if item.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}

	return item.Time, nil
}

//Scan valueof time.Time
func (item *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*item = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

const (
	timeFormart = "2006-01-02 15:04:05"
)

//UnmarshalJSON json反序列化
func (item *JSONTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*item = JSONTime{now}
	return
}
