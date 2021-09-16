package util

import (
	"pool_backend/src/global"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

// GenerateID generate Twitter Snowflake ID
func GenerateID(id int64) snowflake.ID {

	node, err := snowflake.NewNode(id)
	if err != nil {
		global.Logger.Error("生成雪花算法报错:", err)
	}
	return node.Generate()

}
