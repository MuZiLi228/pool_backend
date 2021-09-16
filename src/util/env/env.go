package env

import (
	"fmt"
	"os"
)

var (
	active Environment
	dev    Environment = &environment{value: "dev"}
	fat    Environment = &environment{value: "fat"}
	pro    Environment = &environment{value: "pro"}
)

var _ Environment = (*environment)(nil)

// Environment 环境配置
type Environment interface {
	Value() string
	IsDev() bool
	IsFat() bool
	IsPro() bool
	t()
}

type environment struct {
	value string
}

func (e *environment) Value() string {
	return e.value
}

func (e *environment) IsDev() bool {
	return e.value == "dev"
}

func (e *environment) IsFat() bool {
	return e.value == "fat"
}

func (e *environment) IsPro() bool {
	return e.value == "pro"
}

func (e *environment) t() {}

func init() {

	//设置开发环境   export GIN_MODE=debug|release
	env1 := string(os.Getenv("GIN_MODE"))

	switch env1 {
	case "debug":
		active = dev
	case "release":
		active = pro
	default:
		active = fat
		fmt.Println("Warning: '-env' cannot be found, or it is illegal. The default 'fat' will be used.")
	}
}

// Active 当前配置的env
func Active() Environment {
	return active
}
