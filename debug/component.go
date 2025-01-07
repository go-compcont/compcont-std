package debug

import (
	"github.com/go-compcont/compcont-core"
	"go.uber.org/zap"
)

const TypeID compcont.ComponentTypeID = "std.debug"

type Config struct {
	Logger *compcont.TypedComponentConfig[any, *zap.Logger]
}

var factory compcont.IComponentFactory = &compcont.TypedSimpleComponentFactory[Config, any]{
	TypeID: TypeID,
	CreateInstanceFunc: func(ctx compcont.BuildContext, cfg Config) (instance any, err error) {
		var logger *zap.Logger
		if cfg.Logger != nil {
			loggerComp, err1 := cfg.Logger.LoadComponent(ctx.Container)
			if err1 != nil {
				err = err1
				return
			}
			logger = loggerComp.Instance
		} else {
			logger, err = zap.NewDevelopmentConfig().Build()
			if err != nil {
				return
			}
		}

		logger.Debug(
			"compcont debug info",
			zap.Stringer("name", ctx.Config.Name),
			zap.Stringer("type", ctx.Config.Type),
			zap.Stringers("deps", ctx.Config.Deps),
			zap.Stringers("absolute_path", ctx.GetAbsolutePath()),
		)
		return
	},
}

func MustRegister(registry compcont.IFactoryRegistry) {
	compcont.MustRegister(registry, factory)
}

func init() {
	MustRegister(compcont.DefaultFactoryRegistry)
}
