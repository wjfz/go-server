package rlogging

import (
	"fmt"
	"io"
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
)

// RotateType 轮转类型
type RotateType time.Duration

const (
	// RotateWeek 按周轮转
	RotateWeek = RotateType(time.Hour * 24 * 7)
	// RotateDay 按日轮转
	RotateDay = RotateType(time.Hour * 24)
	// RotateHour 按小时轮转
	RotateHour = RotateType(time.Hour)
	// RotateMinute 按分钟轮转
	RotateMinute = RotateType(time.Minute)
)

// NewRotateWriter 轮转日志写入
func NewRotateWriter(name string, rt RotateType) io.Writer {
	var filename string
	var rotationTime = time.Duration(rt)

	switch rt {
	case RotateWeek:
		filename = name + ".%Y%W.log"
	case RotateDay:
		filename = name + ".%Y%m%d.log"
	case RotateHour:
		filename = name + ".%Y%m%d%H.log"
	case RotateMinute:
		filename = name + ".%Y%m%d%H%M.log"
	default:
		panic("日志轮转类型错误")
	}

	linkname := name + ".log"

	logs, err := rotatelogs.New(
		filename,
		rotatelogs.WithLinkName(linkname),
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithClock(rotatelogs.UTC),
		rotatelogs.WithRotationCount(30),
	)

	if err != nil {
		fmt.Println("日志服务出错:", err)
	}

	return logs
}
