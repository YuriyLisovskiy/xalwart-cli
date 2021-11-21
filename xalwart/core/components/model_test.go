package components

import "testing"

var modeTestModelComponent = ModelComponent{
	class: ClassComponent{
		common: CommonComponent{
			name: "MyTest",
		},
	},
	isJsonSerializable: true,
	customTableName:    "",
}

func TestModelComponent_FileName_Default(t *testing.T) {
	actual := modeTestModelComponent.FileName()
	expected := "my_test"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestModelComponent_FileName_Custom(t *testing.T) {
	modeTestModelComponent.class.customFileName = "custom_file_name"
	actual := modeTestModelComponent.FileName()
	expected := modeTestModelComponent.class.customFileName
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}

	modeTestModelComponent.class.customFileName = ""
}

func TestModelComponent_IsJsonSerializable(t *testing.T) {
	actual := modeTestModelComponent.IsJsonSerializable()
	if !actual {
		t.Errorf("Expected true, received %v", actual)
	}
}

func TestModelComponent_TableName_Default(t *testing.T) {
	actual := modeTestModelComponent.TableName()
	expected := "MyTests"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestModelComponent_TableName_Custom(t *testing.T) {
	modeTestModelComponent.customTableName = "custom_table_name"
	actual := modeTestModelComponent.TableName()
	expected := modeTestModelComponent.customTableName
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}

	modeTestModelComponent.customTableName = ""
}
