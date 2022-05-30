//用于生成jwt的id
package idgener

import (
	"errors"
	"strings"
)

//IDGenInterface 满足该接口的就认为是IDGen
type IDGenInterface interface {
	Next() (string, error)
	String() string
}

//IDGenNameToIDGen 通过名字获得一个IDGen实例
//@params idgen_name string idgen的名字,目前支持uuid4,sonyflake,snowflake
func IDGenNameToIDGen(idgen_name string) (IDGenInterface, error) {
	switch strings.ToLower(idgen_name) {
	case "uuid4":
		{
			return NewUUID4Gen()
		}
	case "sonyflake":
		{
			return NewSonyflakeGen()
		}
	case "snowflake":
		{
			return NewSnowflakeGen()
		}
	default:
		{
			return nil, errors.New("unknown IDGen name")
		}
	}
}

//DefaultUUID4 默认uuid4id生成器
var DefaultUUID4 *UUID4Gen

//DefaultSnowflake 默认snowflake生成器,nodeid为本机MachineID
var DefaultSnowflake *SnowflakeGen

//DefaultSonyflake 默认的sonyflake生成器,使用本机MachineID作为MachineID
var DefaultSonyflake *SonyflakeGen

//Next 使用默认id生成器构造id
//@params idgen_name string idgen的名字,目前支持uuid4,sonyflake,snowflake
func Next(idgen_name string) (string, error) {
	switch strings.ToLower(idgen_name) {
	case "uuid4":
		{
			return DefaultUUID4.Next()
		}
	case "sonyflake":
		{
			return DefaultSonyflake.Next()
		}
	case "snowflake":
		{
			if DefaultSnowflake == nil {
				return "", errors.New("default snowflake not init ok")
			}
			return DefaultSnowflake.Next()
		}
	default:
		{
			return "", errors.New("unknown IDGen name")
		}
	}
}

func init() {
	DefaultUUID4, _ = NewUUID4Gen()
	DefaultSonyflake, _ = NewSonyflakeGen()
	DefaultSnowflake, _ = NewSnowflakeGen()
}
