//autogenerated:yes
//nolint:revive,lll
package mower_map

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
)

type ConvertToNavigationAreaSrvReq struct {
	msg.Package `ros:"mower_map"`
	Index       uint32
}

type ConvertToNavigationAreaSrvRes struct {
	msg.Package `ros:"mower_map"`
}

type ConvertToNavigationAreaSrv struct {
	msg.Package `ros:"mower_map"`
	ConvertToNavigationAreaSrvReq
	ConvertToNavigationAreaSrvRes
}
