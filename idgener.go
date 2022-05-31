//用于生成jwt的id
package idgener

import (
	"errors"
)

//IDGenInterface 满足该接口的就认为是IDGen
type IDGenInterface interface {
	Next() (string, error)
	String() string
}

type IDGENAlgorithm uint8

const (
	//AckModeAckWhenGet 获取到后确认
	IDGEN_UUIDV4 IDGENAlgorithm = iota
	IDGEN_SNOYFLAKE
	IDGEN_SNOWFLAKE
)

//IDGenNameToIDGen 通过名字获得一个IDGen实例
//@params idgen_name string idgen的名字,目前支持uuid4,sonyflake,snowflake
func IDGenNameToIDGen(idgen_name IDGENAlgorithm) (IDGenInterface, error) {
	switch idgen_name {
	case IDGEN_UUIDV4:
		{
			return NewUUID4Gen()
		}
	case IDGEN_SNOYFLAKE:
		{
			return NewSonyflakeGen()
		}
	case IDGEN_SNOWFLAKE:
		{
			return NewSnowflakeGen()
		}
	default:
		{
			return nil, errors.New("unknown IDGen algo name")
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
func Next(idgen_name IDGENAlgorithm) (string, error) {
	switch idgen_name {
	case IDGEN_UUIDV4:
		{
			return DefaultUUID4.Next()
		}
	case IDGEN_SNOYFLAKE:
		{
			return DefaultSonyflake.Next()
		}
	case IDGEN_SNOWFLAKE:
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
