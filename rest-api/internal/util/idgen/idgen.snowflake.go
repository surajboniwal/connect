package idgen

import (
	"fmt"
	"hash/fnv"
	"os"

	"github.com/bwmarrin/snowflake"
)

type SnowflakeIdGen struct {
}

func NewSnowflakeIdGen() *SnowflakeIdGen {
	return &SnowflakeIdGen{}
}

var node *snowflake.Node

func init() {
	n, err := snowflake.NewNode(1)

	if err != nil {
		logger.E(err)
	}

	node = n
}

func (sf SnowflakeIdGen) New() int64 {
	sfId := node.Generate()
	return sfId.Int64()
}

func getMachineID() int {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return 0
	}

	hash := fnv.New32a()
	_, _ = hash.Write([]byte(hostname))
	return int(hash.Sum32())
}
