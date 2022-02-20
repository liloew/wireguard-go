//go:build !linux

package device

import (
	"github.com/liloew/wireguard/conn"
	"github.com/liloew/wireguard/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
