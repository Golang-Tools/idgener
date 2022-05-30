// 用于获取MachineID,MachineID指当前机器的id,使用机器第一张网卡的可用ip地址构造
// 提供全局变量
// `MachineIDStr string`用于保存MachineID的字符串形式
// `MachineID uint16`用于保存MachineID原始的整型数形式
// 这两个全局变量都会在模块初始化时计算获得,不需要额外调用函数获得
package machineid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getMachineID(t *testing.T) {
	assert.NotEqual(t, MachineID, 0)
	assert.NotEqual(t, MachineIDStr, "")
	// assert.FailNow(t, "get MachineID", MachineID)
}
