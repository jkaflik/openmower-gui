//autogenerated:yes
//nolint:revive,lll
package xbot_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
)

type MapOverlay struct {
	msg.Package `ros:"xbot_msgs"`
	Polygons    []MapOverlayPolygon
}
