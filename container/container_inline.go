package container

import "github.com/go-compcont/compcont-core"

const InlineContainerType compcont.ComponentTypeID = "std.container-inline"

type ContainerInlineConfig struct {
	Components []compcont.ComponentConfig `ccf:"components"`
}

var inlineFactory compcont.IComponentFactory = &compcont.TypedSimpleComponentFactory[ContainerInlineConfig, compcont.IComponentContainer]{
	TypeID: InlineContainerType,
	CreateInstanceFunc: func(ctx compcont.BuildContext, config ContainerInlineConfig) (instance compcont.IComponentContainer, err error) {
		instance = compcont.NewComponentContainer(
			compcont.WithParentContainer(ctx.Container),
			compcont.WithFactoryRegistry(ctx.Container.FactoryRegistry()),
			compcont.WithContext(ctx),
		)
		err = instance.LoadNamedComponents(config.Components)
		return
	},
}

func MustRegisterContainerInline(r compcont.IFactoryRegistry) {
	compcont.MustRegister(r, inlineFactory)
}

func init() {
	MustRegisterContainerInline(compcont.DefaultFactoryRegistry)
}
