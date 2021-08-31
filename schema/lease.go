package schema

import (
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// this is in accordance with Kea's schema so do not change field names willy-nilly
//CREATE TABLE public.lease4 (
//    address bigint NOT NULL,
//    hwaddr bytea,
//    client_id bytea,
//    valid_lifetime bigint,
//    expire timestamp with time zone,
//    subnet_id bigint,
//    fqdn_fwd boolean,
//    fqdn_rev boolean,
//    hostname character varying(255),
//    state bigint DEFAULT 0
//);

// wrapper around net.IP that has required SQL conversions
type IPv4 struct {
	net.IP
}

func (ip *IPv4) Scan(v interface{}) error {
	switch ip_int := v.(type) {
	case int64:
		result := make(net.IP, 4)
		ii := uint32(ip_int)
		binary.BigEndian.PutUint32(result, ii)
		ip.IP = result
	default:
		return fmt.Errorf("scan got type %T: %v, only int64 is supported\n", v, v)
	}
	return nil
}

func (i *IPv4) Value() (driver.Value, error) {
	ii := i.To4()
	if ii == nil {
		return nil, fmt.Errorf("can't convert IP: %v to integer", i)
	}
	ipInt := binary.BigEndian.Uint32(ii)
	return int64(ipInt), nil
}

type Lease4 struct {
	Address *IPv4 `gorm:"column:address"`
	//Address *IPv4 `gorm:"column:address;primary_key"`
	Hwaddr   []byte `gorm:"column:hwaddr;primary_key"`
	ClientID []byte `gorm:"column:client_id"`
	// lease length in seconds
	ValidLifetime int       `gorm:"column:valid_lifetime"`
	Expire        time.Time `gorm:"column:expire"`
	// kea's subnet ID.
	SubnetId int `gorm:"column:subnet_id"`
	// Has forward DNS update been performed by a server
	FqdnFwd bool `gorm:"column:fqdn_fwd"`
	// Has reverse DNS update been performed by a server
	FwdnRev  bool   `gorm:"column:fqdn_rev"`
	Hostname string `gorm:"column:hostname;type:varchar(255)"`
	State    int    `gorm:"column:state;DEFAULT:0"`
}

func (h *Lease4) Macaddr() net.HardwareAddr {
	return h.Hwaddr
}

//CREATE TABLE public.hosts (
//    host_id integer NOT NULL,
//    dhcp_identifier bytea NOT NULL,
//    dhcp_identifier_type smallint NOT NULL,
//    dhcp4_subnet_id integer,
//    dhcp6_subnet_id integer,
//    ipv4_address bigint,
//    hostname character varying(255) DEFAULT NULL::character varying,
//    dhcp4_client_classes character varying(255) DEFAULT NULL::character varying,
//    dhcp6_client_classes character varying(255) DEFAULT NULL::character varying,
//    dhcp4_next_server bigint,
//    dhcp4_server_hostname character varying(64) DEFAULT NULL::character varying,
//    dhcp4_boot_file_name character varying(128) DEFAULT NULL::character varying
//);

type Host struct {
	HostId int `gorm:"column:host_id;not null,autoIncrement"`
	// MAC
	DHCPIdentifier []byte `gorm:"column:dhcp_identifier;not null"`
	// 0 for DHCP
	DHCPIdentifierType  int16  `gorm:"column:dhcp_identifier_type;not null;type:smallint"`
	DHCP4SubnetId       int    `gorm:"column:dhcp4_subnet_id"`
	DHCP6SubnetId       int    `gorm:"column:dhcp6_subnet_id"`
	IPV4Address         uint32 `gorm:"column:ipv4_address;type:bigint"`
	Hostname            string `gorm:"column:hostname;type:varchar(255)"`
	DHCP4ClientClasses  string `gorm:"column:dhcp4_client_classes;type:varchar(255)"`
	DHCP6ClientClasses  string `gorm:"column:dhcp6_client_classes;type:varchar(255)"`
	DHCP4NextServer     int    `gorm:"column:dhcp4_next_server"`
	DHCP4ServerHostname string `gorm:"column:dhcp4_server_hostname;type:varchar(64)"`
	DHCP4BootFileName   string `gorm:"column:dhcp4_boot_file_name;type:varchar(128)"`
}

func (h *Host) Macaddr() net.HardwareAddr {
	return h.DHCPIdentifier
}
