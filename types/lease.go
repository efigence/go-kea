package types

import (
	"net"
	"time"
)

type LeaseV4 struct {
	Hostname string
	IP net.IP
	Hwaddr net.HardwareAddr
	ClientID []byte
	// lease length in seconds
	Lifetime int
	Expire time.Time
	// kea's subnet ID.
	SubnetId int
	// Has forward DNS update been performed by a server
	DnsUpdatedFwd bool
	// Has reverse DNS update been performed by a server
	DnsUpdatedRev bool
}