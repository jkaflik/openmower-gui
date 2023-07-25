//autogenerated:yes
//nolint:revive,lll
package xbot_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/geometry_msgs"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/std_msgs"
)

const (
	AbsolutePose_SOURCE_GPS                              uint8  = 1
	AbsolutePose_SOURCE_LIGHTHOUSE                       uint8  = 2
	AbsolutePose_SOURCE_SENSOR_FUSION                    uint8  = 100
	AbsolutePose_FLAG_GPS_RTK                            uint16 = 1
	AbsolutePose_FLAG_GPS_RTK_FIXED                      uint16 = 2
	AbsolutePose_FLAG_GPS_RTK_FLOAT                      uint16 = 4
	AbsolutePose_FLAG_GPS_DEAD_RECKONING                 uint16 = 8
	AbsolutePose_FLAG_SENSOR_FUSION_RECENT_ABSOLUTE_POSE uint16 = 1
	AbsolutePose_FLAG_SENSOR_FUSION_DEAD_RECKONING       uint16 = 8
)

type AbsolutePose struct {
	msg.Package         `ros:"xbot_msgs"`
	msg.Definitions     `ros:"uint8 SOURCE_GPS=1,uint8 SOURCE_LIGHTHOUSE=2,uint8 SOURCE_SENSOR_FUSION=100,uint16 FLAG_GPS_RTK=1,uint16 FLAG_GPS_RTK_FIXED=2,uint16 FLAG_GPS_RTK_FLOAT=4,uint16 FLAG_GPS_DEAD_RECKONING=8,uint16 FLAG_SENSOR_FUSION_RECENT_ABSOLUTE_POSE=1,uint16 FLAG_SENSOR_FUSION_DEAD_RECKONING=8"`
	Header              std_msgs.Header
	SensorStamp         uint32
	ReceivedStamp       uint32
	Source              uint8
	Flags               uint16
	OrientationValid    uint8
	MotionVectorValid   uint8
	PositionAccuracy    float32
	OrientationAccuracy float32
	Pose                geometry_msgs.PoseWithCovariance
	MotionVector        geometry_msgs.Vector3
	VehicleHeading      float64
	MotionHeading       float64
}