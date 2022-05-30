// 用于获取MachineID,MachineID指当前机器的id,使用机器第一张网卡的可用ip地址构造
// 提供全局变量
// `MachineIDStr string`用于保存MachineID的字符串形式
// `MachineID uint16`用于保存MachineID原始的整型数形式
// 这两个全局变量都会在模块初始化时计算获得,不需要额外调用函数获得
package machineid

import (
	"errors"
	"net"
	"strconv"
)

func privateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}
	return nil, errors.New("no private ip address")
}
func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

func lower16BitPrivateIP() (uint16, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}

//MachineID 当前机器的id
var MachineID uint16 = 0

//MachineIDStr 当前机器id的字符串形式
var MachineIDStr string

//getMachineID 获取生成器的MachineID
func getMachineID() string {
	if MachineID == 0 {
		mID, err := lower16BitPrivateIP()
		if err != nil {
			return strconv.FormatUint(uint64(MachineID), 16)
		}
		MachineID = mID
	}
	return strconv.FormatUint(uint64(MachineID), 16)
}

func init() {
	MachineIDStr = getMachineID()
}
