//autogenerated:yes
//nolint:revive,lll
package mower_msgs

import (
    "github.com/bluenviron/goroslib/v2/pkg/msg"
)


type Perimeter struct {
    msg.Package `ros:"mower_msgs"`
    Left float32
    Center float32
    Right float32
}

