//go:build !linux

package device

import (
	"github.com/liloew/wireguard-go/conn"
	"github.com/liloew/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
