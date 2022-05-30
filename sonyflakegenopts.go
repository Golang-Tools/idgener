package idgener

import (
	"time"

	"github.com/Golang-Tools/idgener/machineid"
	"github.com/Golang-Tools/optparams"
	"github.com/sony/sonyflake"
)

var DefaultSonyflakeOpt = sonyflake.Settings{
	StartTime: time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC),
	MachineID: func() (uint16, error) {
		return machineid.MachineID, nil
	},
}

//WithStartTime 设置起始时间
func WithStartTime(starttime time.Time) optparams.Option[sonyflake.Settings] {
	return optparams.NewFuncOption(func(o *sonyflake.Settings) {
		o.StartTime = starttime
	})
}

//WithMachineID 设置机器id
func WithMachineID(machineid uint16) optparams.Option[sonyflake.Settings] {
	return optparams.NewFuncOption(func(o *sonyflake.Settings) {
		o.MachineID = func() (uint16, error) {
			return machineid, nil
		}
	})
}

//WithCheckMachineID 设置校验机器id的函数
func WithCheckMachineID(checkfunc func(uint16) bool) optparams.Option[sonyflake.Settings] {
	return optparams.NewFuncOption(func(o *sonyflake.Settings) {
		o.CheckMachineID = checkfunc
	})
}
