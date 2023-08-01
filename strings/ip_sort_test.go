package strings

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	ips := IPS{"192.168.111.10", "192.168.111.4", "1.1.1.1", "192.168.111.1", "192.168.1.10", "192.168.1.11"}
	sort.Sort(ips)
	fmt.Printf("%#v", ips)
}
