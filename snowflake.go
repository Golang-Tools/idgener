package idgener

import (
	"strconv"

	"github.com/Golang-Tools/optparams"
	"github.com/bwmarrin/snowflake"
)

type SnowflakeGen struct {
	generator *snowflake.Node
	Opt       SnowflakeOpts
}

func NewSnowflakeGen(opts ...optparams.Option[SnowflakeOpts]) (*SnowflakeGen, error) {
	g := new(SnowflakeGen)
	g.Opt = DefaultSnowflakeOpt
	optparams.GetOption(&g.Opt, opts...)
	node, err := snowflake.NewNode(g.Opt.NodeID)
	if err != nil {
		return nil, err
	}
	g.generator = node
	return g, nil
}

//Next 随机生成key
func (g *SnowflakeGen) Next() (string, error) {
	id := g.generator.Generate().Int64()
	return strconv.FormatUint(uint64(id), 32), nil
}

func (g *SnowflakeGen) String() string {
	return "snowflake"
}
