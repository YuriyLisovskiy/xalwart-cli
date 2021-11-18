### Xalwart CLI

- [Download](#download)
  - [Using Go tools](#using-go-tools)
  - [Manual Download](#manual-download)
- [Usage](#usage)
- [Commands](#commands)
- [License](#license)

## Download

### Using Go tools
TODO:

### Manual Download
To download, go to the [releases](https://github.com/YuriyLisovskiy/xalwart-cli/releases).
The examples in the documentation assume you have put this into your PATH and
renamed to `xalwart` (or symlinked as such).

**Note**: If using macOS and downloading manually, you may need to adjust the permissions
of the file to allow for execution.

To do so, please run: `chmod 755 <filename>` where the filename is the name of the downloaded binary.

## Usage
The CLI follows a standard format:
```sh
xalwart [command] [flags]
```
The commands are described below.

## Commands
- [add](docs/add.md)
  - [command](docs/add.md#command)
  - [controller](docs/add.md#controller)
  - [middleware](docs/add.md#middleware)
  - [migration](docs/add.md#migration)
  - [model](docs/add.md#model)
  - [module](docs/add.md#module)
- [project](docs/project.md)
- [version](docs/version.md)

## License
This library is licensed under the Apache 2.0 License.
