package utils

import (
	"errors"
	"fmt"
	"log"

	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/spf13/cobra"
)

type ComponentCommandBuilder struct {
	name                  string
	shortDescription      string
	longDescription       string
	validator             func() error
	componentBuilder      func() (core.Component, error)
	postRunMessageBuilder func(core.Component) string
}

func (cb *ComponentCommandBuilder) SetName(name string) *ComponentCommandBuilder {
	cb.name = name
	return cb
}

func (cb *ComponentCommandBuilder) SetShortDescription(text string) *ComponentCommandBuilder {
	cb.shortDescription = text
	return cb
}

func (cb *ComponentCommandBuilder) SetLongDescription(text string) *ComponentCommandBuilder {
	cb.longDescription = text
	return cb
}

func (cb *ComponentCommandBuilder) SetNameValidator(validator func() error) *ComponentCommandBuilder {
	cb.validator = validator
	return cb
}

func (cb *ComponentCommandBuilder) SetComponentBuilder(
	builder func() (
		core.Component,
		error,
	),
) *ComponentCommandBuilder {
	cb.componentBuilder = builder
	return cb
}

func (cb *ComponentCommandBuilder) SetPostRunMessageBuilder(builder func(core.Component) string) *ComponentCommandBuilder {
	cb.postRunMessageBuilder = builder
	return cb
}

func (cb *ComponentCommandBuilder) Command(overwrite *bool) *cobra.Command {
	return &cobra.Command{
		Use:   cb.name,
		Short: cb.shortDescription,
		Long:  cb.longDescription,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if cb.validator != nil {
				if err := cb.validator(); err != nil {
					return err
				}
			}

			if cb.componentBuilder == nil {
				return errors.New("component builder is nil")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			component, err := cb.componentBuilder()
			if err != nil {
				log.Fatal(err)
			}

			must(core.Generate(component, *overwrite))
			if cb.postRunMessageBuilder != nil {
				fmt.Println(cb.postRunMessageBuilder(component))
			}
		},
	}
}

func NewComponentCommandBuilder() *ComponentCommandBuilder {
	return &ComponentCommandBuilder{}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
