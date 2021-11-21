package components

import "testing"

func TestMiddlewareComponent_FullName_ClassBased(t *testing.T) {
	middlewareComponent := MiddlewareComponent{
		class: ClassComponent{
			common: CommonComponent{
				name: "TestName",
			},
			componentType: "test_type",
		},
		isClassBased: true,
	}
	actual := middlewareComponent.FullName()
	expected := "TestNameTestType"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestMiddlewareComponent_FullName_NotClassBased(t *testing.T) {
	middlewareComponent := MiddlewareComponent{
		class: ClassComponent{
			common: CommonComponent{
				name: "TestName",
			},
			componentType: "test_type",
		},
		isClassBased: false,
	}
	actual := middlewareComponent.FullName()
	expected := "test_name_test_type"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}
