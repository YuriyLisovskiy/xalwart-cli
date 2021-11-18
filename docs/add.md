# Add
- [command](#command)
- [controller](#controller)
- [middleware](#middleware)
- [migration](#migration)
- [model](#model)
- [module](#module)

### Description
Used to create a new component for existing project.

**Flags**

| Flag     | Shorthand | Description     | Example | Required? (Y/N) |
|----------|-----------|-----------------|---------|-----------------|
| `--help` | `-h`      | Help for `add`. | `-h`    | N               |

**Examples**
```sh
xalwart add --help
```

## Command
Used to create a new command component.
Command files will have snake case value of `name` flag as names by default.

**Flags**

| Flag          | Shorthand | Description                      | Example          | Required? (Y/N) |
|---------------|-----------|----------------------------------|------------------|-----------------|
| `--file`      | `-f`      | Custom file name of new command. | `-f dump`        | N               |
| `--help`      | `-h`      | Help for `command`.              | `-h`             | N               |
| `--name`      | `-n`      | Name of a new command.           | `-n DumpModules` | Y               |
| `--overwrite` | `-o`      | Overwrite files if exist.        | `-o`             | N               |
| `--root`      | `-r`      | Root path for a new command.     | `-r ./commands`  | N               |

**Examples**
```sh
xalwart add command --name DumpModules
xalwart add command -n RunServer -f run -o
```

## Controller
Used to create a new controller component.
Controller files will have snake case `{name_flag}_controller` as names by default.

**Flags**

| Flag          | Shorthand | Description                         | Example            | Required? (Y/N) |
|---------------|-----------|-------------------------------------|--------------------|-----------------|
| `--file`      | `-f`      | Custom file name of new controller. | `-f index`         | N               |
| `--help`      | `-h`      | Help for `controller`.              | `-h`               | N               |
| `--name`      | `-n`      | Name of a new controller.           | `-n Products`      | Y               |
| `--overwrite` | `-o`      | Overwrite files if exist.           | `-o`               | N               |
| `--root`      | `-r`      | Root path for a new controller.     | `-r ./controllers` | N               |

**Examples**
```sh
xalwart add controller --name products
xalwart add controller -n Index -f index -o
```

## Middleware
Used to create a new middleware component.
Middleware files will have snake case value of `name` flag as names by default.

**Flags**

| Flag            | Shorthand | Description                                           | Example                     | Required? (Y/N) |
|-----------------|-----------|-------------------------------------------------------|-----------------------------|-----------------|
| `--class-based` | `-c`      | Add class-based middleware instead of function-based. | `-c`                        | N               |
| `--file`        | `-f`      | Custom file name of new middleware.                   | `-f test`                   | N               |
| `--help`        | `-h`      | Help for `middleware`.                                | `-h`                        | N               |
| `--name`        | `-n`      | Name of a new middleware.                             | `-n CustomExceptionHandler` | Y               |
| `--overwrite`   | `-o`      | Overwrite files if exist.                             | `-o`                        | N               |
| `--root`        | `-r`      | Root path for a new middleware.                       | `-r ./middleware`           | N               |

**Examples**
```sh
xalwart add middleware --name CustomExceptionHandler -c
xalwart add middleware -n Test -f index -o
```

## Migration
Used to create a new migration component.
Migration files will have snake case value of `name` flag as names by default.

Recommended migration name structure is `{number}_{ShortDescription}`, for example: `001_Initial`.
In this case, migration class will have `Migration{number}_{ShortDescription}` name.

**Flags**

| Flag          | Shorthand | Description                        | Example            | Required? (Y/N) |
|---------------|-----------|------------------------------------|--------------------|-----------------|
| `--file`      | `-f`      | Custom file name of new migration. | `-f migration_001` | N               |
| `--help`      | `-h`      | Help for `migration`.              | `-h`               | N               |
| `--initial`   | `-i`      | Mark migration as initial one.     | `-i`               | N               |
| `--name`      | `-n`      | Name of a new migration.           | `-n 001_Initial`   | Y               |
| `--overwrite` | `-o`      | Overwrite files if exist.          | `-o`               | N               |
| `--root`      | `-r`      | Root path for a new migration.     | `-r ./migrations`  | N               |

**Examples**
```sh
xalwart add migration --name 001_CreatedUser
xalwart add migration -n 001_Initial -f migration_001 -o -i
```

## Model
Used to create a new model component.
Model files will have snake case value of `name` flag as names by default.

**Flags**

| Flag                  | Shorthand | Description                                                                      | Example            | Required? (Y/N) |
|-----------------------|-----------|----------------------------------------------------------------------------------|--------------------|-----------------|
| `--file`              | `-f`      | Custom file name of new model.                                                   | `-f client_model`  | N               |
| `--help`              | `-h`      | Help for `model`.                                                                | `-h`               | N               |
| `--json-serializable` | `-j`      | Inherit from 'xw::IJsonSerializable' interface and implement 'to_json()' method. | `-j`               | N               |
| `--name`              | `-n`      | Name of a new model.                                                             | `-n Client`        | Y               |
| `--overwrite`         | `-o`      | Overwrite files if exist.                                                        | `-o`               | N               |
| `--root`              | `-r`      | Root path for a new model.                                                       | `-r ./models`      | N               |
| `--table`             | `-t`      | Custom name which will be used for table in SQL database.                        | `-t client_models` | N               |

**Examples**
```sh
xalwart add model --name Client
xalwart add model -n Product -f product_model -o -j
```

## Module
Used to create a new module component.
Module files will have `module` names by default and will be placed in the directory with snake case
name from value of `name` flag.

**Flags**

| Flag          | Shorthand | Description                     | Example                 | Required? (Y/N) |
|---------------|-----------|---------------------------------|-------------------------|-----------------|
| `--file`      | `-f`      | Custom file name of new module. | `-f module_config`      | N               |
| `--help`      | `-h`      | Help for `module`.              | `-h`                    | N               |
| `--name`      | `-n`      | Name of a new module.           | `-n api`                | Y               |
| `--overwrite` | `-o`      | Overwrite files if exist.       | `-o`                    | N               |
| `--root`      | `-r`      | Root path for a new module.     | `-r ./SuperProject/src` | N               |

**Examples**
```sh
xalwart add module --name Api
xalwart add module -n core -f config -o
```
