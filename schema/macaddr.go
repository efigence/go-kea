package schema

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"net"
)

// https://github.com/lib/pq/pull/390
// Macaddr is a wrapper for transferring Macaddr values back and forth easily.
type Macaddr struct {
	net.HardwareAddr
}

func ParseMAC(s string) (mac Macaddr, err error) {
	mac.HardwareAddr, err = net.ParseMAC(s)
	return mac, err
}

// Method mostly for testing without fuss
func MustParseMAC(s string) (mac Macaddr) {
	mac, err := ParseMAC(s)
	if err != nil {
		panic(err)
	}
	return mac
}

// Scan implements the Scanner interface.
func (m *Macaddr) Scan(value interface{}) error {
	m.HardwareAddr = nil
	macaddrAsBytes, ok := value.([]byte)
	if !ok {
		return errors.New("Could not convert scanned value to bytes")
	}
	parsedMacaddr, err := net.ParseMAC(string(macaddrAsBytes))
	if err != nil {
		return err
	}
	m.HardwareAddr = parsedMacaddr
	return nil
}

// Value implements the driver Valuer interface. Note if m.Valid is false
// or m.HardwareAddr is nil the database column value will be set to NULL.
func (m Macaddr) Value() (driver.Value, error) {
	if m.HardwareAddr == nil {
		return nil, fmt.Errorf("hardware address is nil")
	}
	return []byte(m.HardwareAddr.String()), nil
}

func (m Macaddr) MarshalJSON() ([]byte, error) {
	if m.HardwareAddr != nil {
		return []byte(`"` + m.HardwareAddr.String() + `"`), nil
	} else {
		return []byte(`null`), nil
	}

}
func (m *Macaddr) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	m.HardwareAddr, err = net.ParseMAC(s)
	return err
}
func (m *Macaddr) Byte() []byte {
	return m.HardwareAddr
}
