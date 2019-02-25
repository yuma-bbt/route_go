package main

import (
	"fmt"
	"strconv"
)

type edge struct {
	leaf_node string
}

type edge_box []*edge

func do_input_name(n string) (r *edge) {
	r = new(edge)
	r.leaf_node = n
	return r
}

func ip_maker(i int) (ips string) {
	ips = "192.168." + strconv.Itoa(i) + ".0/24"
	return ips
}

func main() {
	var leaf1 edge_box

	var tmp []string
	var addr_range1_g = 9

	for x := 0; x < addr_range1_g; x++ {
		tmp = append(tmp, ip_maker(x))
	}
	//fmt.Println(tmp)
	for _, tmp := range tmp {
		//		fmt.Println(tmp)
		leaf1 = append(leaf1, do_input_name(tmp))
	}

	fmt.Println(leaf1[2].leaf_node)
	//leaf1 = append(leaf1, do_input_name("192.168."+x+".0/24"))

	//for _, i := range leaf1 {
	//fmt.Println(i.leaf_node)
	//}
	//fmt.Println(leaf[0].leaf_node)
}
