package types

import (
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

type Lease4 struct {
	Address net.IP `gorm:"column:address;"`
	Hwaddr net.HardwareAddr `gorm:"column:hwaddr"`
	ClientID []byte `gorm:"column:client_id"`
	// lease length in seconds
	ValidLifetime int `gorm:"column:valid_lifetime"`
	Expire time.Time `gorm:"column:expire"`
	// kea's subnet ID.
	SubnetId int `gorm:"column:subnet_id"`
	// Has forward DNS update been performed by a server
	FqdnFwd bool `gorm:"column:fqdn_fwd"`
	// Has reverse DNS update been performed by a server
	FwdnRev bool `gorm:"column:fqdn_rev"`
	Hostname string `gorm:"column:hostname;type:varchar(255)"`
	State int `gorm:"column:state;DEFAULT:0"`
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
	HostId int `gorm:"column:host_id;not null"`
	DHCPIdentifier []byte `gorm:"column:dhcp_identifier;not null"`
	DHCPIdentifierType int16 `gorm:"column:dhcp_identifier_type;not null;type:smallint"`
	DHCP4SubnetId int `gorm:"column:dhcp4_subnet_id"`
	DHCP6SubnetId int `gorm:"column:dhcp6_subnet_id"`
	IPV4Address net.HardwareAddr `gorm:"column:ipv4_address;type:bigint"`
	Hostname string `gorm:"column:hostname;type:varchar(255)"`
	DHCP4ClientClasses string `gorm:"column:dhcp4_client_classes;type:varchar(255)"`
	DHCP6ClientClasses string `gorm:"column:dhcp6_client_classes;type:varchar(255)"`
	DHCP4NextServer int `gorm:"column:dhcp4_next_server"`
	DHCP4ServerHostname string `gorm:"column:dhcp4_server_hostname;type:varchar(64)"`
	DHCP4BootFileName string `gorm:"column:dhcp4_boot_file_name;type:varchar(128)"`
}