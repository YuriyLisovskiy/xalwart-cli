### xalwart-cli

CLI Application for web-site development using
[Xalwart Framework](https://github.com/YuriyLisovskiy/xalwart)

#### Table of contents
* [Build](#build)
* [Installation](#installation)

#### Build
To build xalwart-cli project you need to have at least Go v1.11.
For further information, read how to
[install](https://golang.org/doc/install) the Go programming
language. 

Use [Git Bash](https://git-scm.com/downloads) terminal on Windows.

```bash
$ git clone https://github.com/YuriyLisovskiy/xalwart-cli.git
$ cd xalwart-cli
$ make build
```

#### Installation

Download binary from
[releases](https://github.com/YuriyLisovskiy/xalwart-cli/releases)
or build the project.

* **Linux**

    If you built the project by yourself, run:
    ```bash
    $ sudo make install
    ```
  
    Otherwise, follow the next steps:
    ```bash
    $ sudo tar xvzf ~/Downalods/xalwart-cli-<os>-<arch>.tar.gz -C /usr/local/bin
    $ sudo chmod a+x /usr/local/bin/xalwart
    ```

* **Windows**

    * Create `C:\\xalwart` directory if not exists.
    * Unpack downloaded archive `xalwart-cli-<os>-<arch>.tar.gz`
      or copy `./bin/xalwart.exe` file (if you built the project
      manually) to `C:\\xalwart`.
    * Run the following command in `cmd` terminal to append
      `xalwart` application to `PATH`:
      ```bash
      pathman /au C:\xalwart
      ```
      Or add `C:\xalwart` to `PATH` manually.
