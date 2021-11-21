### Xalwart CLI

- [Download](#download)
  - [Using go install](#using-go-install)
  - [Manual Download](#manual-download)
- [Documentation](#documentation)
- [License](#license)

## Download

### Using go install
Install the xalwart tool with the command:
```
go install github.com/YuriyLisovskiy/xalwart-cli/xalwart@latest
```
Go will automatically install it in your `$GOPATH/bin` directory which should be in your $PATH.

Once installed you should have the `xalwart` command available. Confirm by typing `xalwart` at a
command line.

### Manual Download
To download, go to the [releases](https://github.com/YuriyLisovskiy/xalwart-cli/releases).
The examples in the documentation assume you have put this into your PATH and
renamed to `xalwart` (or symlinked as such).

**Note**: If using macOS and downloading manually, you may need to adjust the permissions
of the file to allow for execution.

To do so, please run: `chmod 755 <filename>` where the filename is the name of the downloaded binary.

## Documentation
Check out [this](https://github.com/YuriyLisovskiy/xalwart-cli/wiki) wiki page.

## License
This library is licensed under the Apache 2.0 License.
