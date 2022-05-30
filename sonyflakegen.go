package idgener

import (
	"strconv"

	"github.com/Golang-Tools/optparams"
	"github.com/sony/sonyflake"
)

type SonyflakeGen struct {
	generator *sonyflake.Sonyflake
	Opt       sonyflake.Settings
}

func NewSonyflakeGen(opts ...optparams.Option[sonyflake.Settings]) (*SonyflakeGen, error) {
	g := new(SonyflakeGen)
	g.Opt = DefaultSonyflakeOpt
	optparams.GetOption(&g.Opt, opts...)
	g.generator = sonyflake.NewSonyflake(g.Opt)
	return g, nil
}

//Next 随机生成key
func (g *SonyflakeGen) Next() (string, error) {
	id, err := g.generator.NextID()
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(id, 32), nil
}

func (g *SonyflakeGen) String() string {
	return "sonyflake"
}
