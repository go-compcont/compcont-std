package compcontzap

import (
	"github.com/go-compcont/compcont-core"
	"go.uber.org/zap"
)

const TypeID compcont.ComponentTypeID = "std.logger.zap"

var factory compcont.IComponentFactory = &compcont.TypedSimpleComponentFactory[Config, *zap.Logger]{
	TypeID: TypeID,
	CreateInstanceFunc: func(ctx compcont.BuildContext, config Config) (instance *zap.Logger, err error) {
		logger, err := New(config)
		if err != nil {
			return
		}
		instance = logger
		return
	},
}

func MustRegister(registry compcont.IFactoryRegistry) {
	compcont.MustRegister(registry, factory)
}

func init() {
	MustRegister(compcont.DefaultFactoryRegistry)
}
