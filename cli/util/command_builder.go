package util

import (
	"errors"
	"fmt"
	"log"

	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/spf13/cobra"
)

type CommandBuilder struct {
	name                     string
	shortDescription         string
	longDescription          string
	validateName             func()
	componentBuilder         func() (core.Component, error)
	postCreateMessageBuilder func(core.Component) string
}

func (cb *CommandBuilder) SetName(name string) *CommandBuilder {
	cb.name = name
	return cb
}

func (cb *CommandBuilder) SetShortDescription(text string) *CommandBuilder {
	cb.shortDescription = text
	return cb
}

func (cb *CommandBuilder) SetLongDescription(text string) *CommandBuilder {
	cb.longDescription = text
	return cb
}

func (cb *CommandBuilder) SetNameValidator(validator func()) *CommandBuilder {
	cb.validateName = validator
	return cb
}

func (cb *CommandBuilder) SetComponentBuilder(builder func() (core.Component, error)) *CommandBuilder {
	cb.componentBuilder = builder
	return cb
}

func (cb *CommandBuilder) SetPostCreateMessageBuilder(builder func(core.Component) string) *CommandBuilder {
	cb.postCreateMessageBuilder = builder
	return cb
}

func (cb *CommandBuilder) Command(overwrite *bool) *cobra.Command {
	return &cobra.Command{
		Use:   cb.name,
		Short: cb.shortDescription,
		Long:  cb.longDescription,
		Run: func(cmd *cobra.Command, args []string) {
			if cb.validateName != nil {
				cb.validateName()
			}

			if cb.componentBuilder == nil {
				log.Fatal(errors.New("component builder is nil"))
			}

			component, err := cb.componentBuilder()
			if err != nil {
				log.Fatal(err)
			}

			must(core.Generate(component, *overwrite))
			if cb.postCreateMessageBuilder != nil {
				fmt.Println(cb.postCreateMessageBuilder(component))
			}
		},
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
