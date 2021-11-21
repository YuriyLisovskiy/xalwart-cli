package components

import "testing"

var migrationTestMigrationComponent = MigrationComponent{
	class:         ClassComponent{},
	isInitial:     true,
	migrationName: "001_TestMigration",
	className:     "Migration001_TestMigration",
}

func TestMigrationComponent_ClassName(t *testing.T) {
	actual := migrationTestMigrationComponent.ClassName()
	expected := "Migration001_TestMigration"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestMigrationComponent_FileName_Default(t *testing.T) {
	actual := migrationTestMigrationComponent.FileName()
	expected := "001_test_migration"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestMigrationComponent_FileName_Custom(t *testing.T) {
	migrationTestMigrationComponent.class.customFileName = "custom_file_name"
	actual := migrationTestMigrationComponent.FileName()
	expected := migrationTestMigrationComponent.class.customFileName
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}

	migrationTestMigrationComponent.class.customFileName = ""
}

func TestMigrationComponent_IsInitial(t *testing.T) {
	actual := migrationTestMigrationComponent.IsInitial()
	if !actual {
		t.Errorf("Expected true, received %v", actual)
	}
}

func TestMigrationComponent_MigrationName(t *testing.T) {
	actual := migrationTestMigrationComponent.MigrationName()
	expected := migrationTestMigrationComponent.migrationName
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestNewMigrationComponent_GenericClassName(t *testing.T) {
	migration, _ := NewMigrationComponent(
		HeaderComponent{},
		templateBoxMock{Templates: map[string]string{}},
		"001_TestMigration",
		"/tmp",
		"",
		true,
	)
	actual := migration.ClassName()
	expected := "Migration001_TestMigration"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestNewMigrationComponent_CustomClassName(t *testing.T) {
	migration, _ := NewMigrationComponent(
		HeaderComponent{},
		templateBoxMock{Templates: map[string]string{}},
		"TestMigration_001",
		"/tmp",
		"",
		true,
	)
	actual := migration.ClassName()
	expected := "TestMigration_001"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}
