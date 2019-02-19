package internal

import (
	"net"
)

func CIDRMask(ones int, bits int) net.IPMask {
	return net.CIDRMask(ones, bits)
}
func Dial(network string, address string) (net.Conn, error) {
	return net.Dial(network, address)
}
func DialIP(network string, laddr, raddr *net.IPAddr) (*net.IPConn, error) {
	return net.DialIP(network, laddr, raddr)
}
