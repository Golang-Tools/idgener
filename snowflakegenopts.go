package idgener

import (
	"github.com/Golang-Tools/idgener/machineid"
	"github.com/Golang-Tools/optparams"
)

type SnowflakeOpts struct {
	NodeID int64
}

var DefaultSnowflakeOpt = SnowflakeOpts{
	NodeID: int64(machineid.MachineID),
}

//WithNodeID 设置节点id
func WithNodeID(nodeid int64) optparams.Option[SnowflakeOpts] {
	return optparams.NewFuncOption(func(o *SnowflakeOpts) {
		o.NodeID = nodeid
	})
}
