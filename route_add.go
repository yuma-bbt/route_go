package main

import (
	"github.com/vishvananda/netlink"
)

func main() {
	lo, _ := netlink.ParseAddr("169.254.169.254 / 32")
	netlink.AddrAdd(lo, addr)
}
