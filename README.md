# kaeya-ctl
Command line tool for communicating with kaeya-server

## Usage
```shell
Usage:
  kaeyactl [flags]

Flags:
  -a, --address string   address of kaeya-server
  -h, --help             help for kaeyactl
  -i, --interactive      interactive mode or not
  -v, --version          version for kaeyactl

```
### Interactive mode
Add `-i` flag to enter interactive mode, support commands:
```shell
set
Set kv pair. Usage: set KEY VALUE

get
Get value of the key. Usage: get KEY

exit
Exit interactive mode. Usage: exit

```
Example
```shell
>> get abc
abc -> 10
>> set abc 100
success
>> get abc
abc -> 100
>> exit
Kaeya-ctl interaction mode exit.
```
## Build
Use `make all` for test and build binary executable file for the following OS:
* Windows
* Linux
* MacOS

For more detail, see the Makefile

## Meaning of the project name
**Kaeya Alberich**, a  character in the game **Genshin Impact**.

<img alt="Kaeya Alberich" height="951" src="./doc/kaeya.jpeg" width="375"/>