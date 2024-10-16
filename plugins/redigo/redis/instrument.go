package redis

import (
	"embed"
	"github.com/apache/skywalking-go/plugins/core/instrument"
)

//go:embed *
var fs embed.FS

//skywalking:nocopy
type Instrument struct {
}

func NewInstrument() *Instrument {
	return &Instrument{}
}

func (i *Instrument) Name() string {
	return "redigo"
}

func (i *Instrument) BasePackage() string {
	return "github.com/gomodule/redigo"
}

func (i *Instrument) VersionChecker(version string) bool {
	return true
}

func (i *Instrument) Points() []*instrument.Point {
	return []*instrument.Point{
		{
			PackagePath: "",
			At:          instrument.NewStructEnhance("activeConn"),
		},
		{
			PackagePath: "",
			At:          instrument.NewStructEnhance("errorConn"),
		},
		{
			PackagePath: "",
			At: instrument.NewMethodEnhance("*activeConn", "Do",
				instrument.WithArgType(0, "string"),
				instrument.WithResultCount(2),
				instrument.WithResultType(0, "interface{}"),
				instrument.WithResultType(1, "error")),
			Interceptor: "DoInterceptor",
		},
		{
			PackagePath: "",
			At: instrument.NewMethodEnhance("*errorConn", "Do",
				instrument.WithArgType(0, "string"),
				instrument.WithResultCount(2),
				instrument.WithResultType(1, "error")),
			Interceptor: "DoInterceptor",
		},
		{
			PackagePath: "",
			At: instrument.NewStaticMethodEnhance("Dial",
				instrument.WithArgType(0, "string"),
				instrument.WithArgType(1, "string"),
				instrument.WithResultCount(2),
				instrument.WithResultType(0, "Conn"),
				instrument.WithResultType(1, "error"),
			),
			Interceptor: "DialInterceptor",
		},
	}
}

func (i *Instrument) PluginSourceCodePath() string {
	return "redis"
}

func (i *Instrument) FS() *embed.FS {
	return &fs
}
