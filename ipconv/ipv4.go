package ipconv // import "fixup.cc/go/util/ipconv"

import (
	"net"
	"strconv"
	"strings"
)

type IP4Convertor struct{}

func NewIP4Convertor() *IP4Convertor {
	return &IP4Convertor{}
}

// 网络字节序大端整形 转换为x.x.x.x格式IP
func (this *IP4Convertor) Inetntoa(ip_int int) string {
	return inet_ntoa(uint32(ip_int))
}

// x.x.x.x格式IP 转换为网络字节序大端整形
func (this *IP4Convertor) Inetaton(ip_str string) int {
	return int(inet_aton(ip_str))
}

//////////////////////////////////////////////////////////////////////////////
// 向C致敬
func inet_ntoa(ipnr uint32) string {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0]).String()
}

func inet_aton(ipnr string) uint32 {
	bits := strings.Split(ipnr, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum uint32

	sum += uint32(b0) << 24
	sum += uint32(b1) << 16
	sum += uint32(b2) << 8
	sum += uint32(b3)

	return sum
}

// Miling
func InetNtoa(ipnr int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

func InetAton(ipnr string) int64 {
	bits := strings.Split(strings.Split(ipnr, ":")[0], ".")
	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64
	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}
