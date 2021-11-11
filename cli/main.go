package main

import "github.com/YuriyLisovskiy/xalwart-cli/generator"

func main() {
	//unit, err := generator.NewCommandUnit("dump", "./generated/commands", "")
	//must(err)
	//must(generator.GenerateUnit(unit))
	//
	//unit, err = generator.NewControllerUnit("index", "./generated/controllers", "")
	//must(err)
	//must(generator.GenerateUnit(unit))
	//
	//unit, err = generator.NewModuleUnit("main", "./generated/modules", "")
	//must(err)
	//must(generator.GenerateUnit(unit))

	unit, err := generator.NewProjectUnit(
		"CustomService",
		"./generated/projects",
		50,
		true,
		true,
	)
	must(err)
	must(generator.GenerateUnit(unit))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
