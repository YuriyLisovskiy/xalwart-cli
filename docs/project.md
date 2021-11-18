# Project

## Description
Used to create a new project.

**Flags**

| Flag           | Shorthand | Description                                       | Example           | Required? (Y/N) |
|----------------|-----------|---------------------------------------------------|-------------------|-----------------|
| `--help`       | `-h`      | Help for `project`.                               | `-h`              | N               |
| `--key-length` | `-k`      | Length of secret key to be generated in settings. | `-k 50`           | N               |
| `--name`       | `-n`      | Name of a new project.                            | `-n TestProject`  | Y               |
| `--overwrite`  | `-o`      | Overwrite files if exist.                         | `-o`              | N               |
| `--root`       | `-r`      | Root path for a new project.                      | `-r ./middleware` | N               |
| `--use-orm`    | `-O`      | Use standard ORM, provided by framework.          | `-O`              | N               |
| `--use-server` | `-s`      | Use standard web server, provided by framework.   | `-s`              | N               |
| `--`           | `-`       |                                                   | `-`               | N               |

**Examples**

```sh
xalwart project -n TestProject
xalwart project -n SecondProject -k 64 -o -s false
```
