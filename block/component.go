package debug

import (
	"math"
	"time"

	"github.com/go-compcont/compcont-core"
)

const TypeID compcont.ComponentTypeID = "std.block"

type Config struct {
	Duration time.Duration `ccf:"duration"`
}

var factory compcont.IComponentFactory = &compcont.TypedSimpleComponentFactory[Config, any]{
	TypeID: TypeID,
	CreateInstanceFunc: func(ctx compcont.BuildContext, cfg Config) (instance any, err error) {
		if cfg.Duration == 0 {
			cfg.Duration = math.MaxInt64
		}
		time.Sleep(cfg.Duration)
		return
	},
}

func MustRegister(registry compcont.IFactoryRegistry) {
	compcont.MustRegister(registry, factory)
}

func init() {
	MustRegister(compcont.DefaultFactoryRegistry)
}
