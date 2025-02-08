package reloading

import (
	"github.com/go-compcont/compcont-core"
	"github.com/go-resty/resty/v2"
)

const TypeID compcont.ComponentTypeID = "std.reloading"

var factory compcont.IComponentFactory = &compcont.TypedSimpleComponentFactory[Config, IReloading]{
	TypeID: TypeID,
	CreateInstanceFunc: func(ctx compcont.BuildContext, cfg Config) (instance IReloading, err error) {
		var restyClient *resty.Client
		if cfg.Resty != nil {
			restyClient = cfg.Resty.MustLoadComponent(ctx.Container).Instance
		}
		instance = NewReloading(cfg, restyClient)
		return
	},
}

func MustRegister(registry compcont.IFactoryRegistry) {
	compcont.MustRegister(registry, factory)
}

func init() {
	MustRegister(compcont.DefaultFactoryRegistry)
}
