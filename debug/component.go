package debug

import (
	"github.com/go-compcont/compcont-core"
	compcontzap "github.com/go-compcont/compcont-std/compcont-zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const TypeID compcont.ComponentTypeID = "std.debug"

type Config struct {
	Logger  *compcont.TypedComponentConfig[any, *zap.Logger] `ccf:"logger"`
	Level   string                                           `ccf:"level"`
	Message string                                           `ccf:"message"`
}

var factory compcont.IComponentFactory = &compcont.TypedSimpleComponentFactory[Config, any]{
	TypeID: TypeID,
	CreateInstanceFunc: func(ctx compcont.BuildContext, cfg Config) (instance any, err error) {
		var logger *zap.Logger
		if cfg.Logger != nil {
			logger = cfg.Logger.MustLoadComponent(ctx.Container).Instance
		} else {
			logger = compcontzap.GetDefault()
		}

		var level zapcore.Level
		if cfg.Level != "" {
			lvl, err1 := zap.ParseAtomicLevel(cfg.Level)
			if err1 != nil {
				logger.Error("parse zap level error", zap.Error(err1))
				return
			}
			level = lvl.Level()
		} else {
			level = zapcore.DebugLevel
		}

		logger.Log(
			level,
			"compcont debug info",
			zap.Stringer("name", ctx.Config.Name),
			zap.Stringer("type", ctx.Config.Type),
			zap.Stringers("deps", ctx.Config.Deps),
			zap.Stringers("absolute_path", ctx.GetAbsolutePath()),
			zap.String("message", cfg.Message),
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
