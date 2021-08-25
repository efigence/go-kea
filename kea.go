package kea

import (
	"encoding/binary"
	"fmt"
	"github.com/efigence/go-kea/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net"
)

type Config struct {
	DSN string
}

type Kea struct {
	db *gorm.DB
}

type Host struct {
	Hostname string
	Macaddr  []byte
	IP       net.IP
	SubnetId int
}

func New(cfg Config) (*Kea, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	k := Kea{db: db}
	return &k, nil
}

func (k *Kea) Add(host Host) error {
	dbMac := schema.Macaddr{host.Macaddr}
	ipInt, err := ip4toInt(host.IP)
	if err != nil {
		return err
	}
	k.db.Create(schema.Host{
		//HostId:              0,
		DHCPIdentifier:     dbMac,
		DHCPIdentifierType: 0, // 0 is MAC
		DHCP4SubnetId:      host.SubnetId,
		IPV4Address:        ipInt,
		Hostname:           host.Hostname,
		//DHCP4ClientClasses:  "",
		//DHCP6ClientClasses:  "",
		//DHCP4NextServer:     0,
		//DHCP4ServerHostname: "",
		//DHCP4BootFileName:   "",
	})
	return nil
}

func (k *Kea) ListBySubnetID(subnetId int) (hosts []Host, err error) {
	result := k.db.Where(&schema.Host{DHCP4SubnetId: subnetId}).Find(&hosts)
	return hosts, result.Error
}

//func (k *Kea)ListUsedSubnets() (map[int]int, error) {
//	k.db.Raw("SELECT count(*),dhcp4_subnet_id from hosts group by dhcp_subnet_id", 3).Scan(&result)
//}

func ip4toInt(ipAddr net.IP) (uint32, error) {
	ip := ipAddr.To4()
	if ip == nil {
		return 0, fmt.Errorf("couldn't convert to IPv4")
	}
	return binary.BigEndian.Uint32(ip), nil
}
