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
	ipInt, err := ip4ToInt(host.IP)
	if err != nil {
		return err
	}
	res := k.db.Select(
		"DHCPIdentifier",
		"DHCPIdentifierType",
		"DHCP4SubnetId",
		"IPV4Address",
		"Hostname",
	).Create(schema.Host{
		//HostId:              0,
		DHCPIdentifier:     host.Macaddr,
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
	return res.Error
}

func (k *Kea) Delete(ip net.IP, mac []byte) error {
	ipInt, err := ip4ToInt(ip)
	if err != nil {
		return err
	}
	res := k.db.Where(schema.Host{
		DHCPIdentifier: mac,
		IPV4Address:    ipInt,
	}).Delete(schema.Host{})
	return res.Error
}

func (k *Kea) ListBySubnetID(subnetId int) (hosts []Host, err error) {
	h := []schema.Host{}
	result := k.db.Where(&schema.Host{DHCP4SubnetId: subnetId}).Find(&h)
	hosts = make([]Host, len(h))
	for id, host := range h {
		ip := intToIPv4(host.IPV4Address)
		hosts[id] = Host{
			Hostname: host.Hostname,
			Macaddr:  host.DHCPIdentifier,
			IP:       ip,
			SubnetId: host.DHCP4SubnetId,
		}
	}
	return hosts, result.Error
}

//func (k *Kea)ListUsedSubnets() (map[int]int, error) {
//	k.db.Raw("SELECT count(*),dhcp4_subnet_id from hosts group by dhcp_subnet_id", 3).Scan(&result)
//}

func ip4ToInt(ipAddr net.IP) (uint32, error) {
	ip := ipAddr.To4()
	if ip == nil {
		return 0, fmt.Errorf("couldn't convert to IPv4")
	}
	return binary.BigEndian.Uint32(ip), nil
}
func intToIPv4(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
