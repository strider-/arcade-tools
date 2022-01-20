# arcade-tools

Custom API / CLI tooling for controlling my modified Arcade1Up cabinet. Run `make help` to list all the available make commands for the project.
In order to run `make install` or `make update` locally on the Pi, you may need to ensure that `$PATH` contains the go binary location (If you're unsure where it's been installed to, run `which go`). To do this, edit `/etc/sudoers` and ensure that `Defaults secure_path` contains the go binary location.

Before you run `make update`, sure you have latest by running `git pull origin master`.

The `GET /api/v1/now-playing` endpoint relies on the following line being somewhere in `/opt/retroarch/configs/all/runcommand-onstart.sh`:
```bash
printf "$1\n$2\n$3\n`basename ${3%.*}`\n$4" > /tmp/now-playing
```
...and this line in `/opt/retroarch/configs/all/runcommand-onend.sh`:
```bash
rm -rf /tmp/now-playing
```

## Project structure
```
├── README.md # this document
├── go.mod
├── go.sum
├── cmd
│   ├── api   # json api project
│   └── cli   # command line interface project
├── internal  # private library code
└── scripts   # project scripts/service definitions
              # to be ran/installed on the RPi via make
```

## Expected RPi GPIO pin usage
NOTE: The [relay](https://www.amazon.com/dp/B0057OC6D8?psc=1) in question is active low, all code reflects this
| Phys. Pin | GPIO # | Usage |
|----------:|---------:|-------|
|     2     |     5V   | Power |
|     9     |     GND  | Ground|
|    11     |     17   | Relay 2|
|    12     |     18   | Relay 1|


## Command Line Interface Arguments
```
  -relay int
        Relay to change (1 or 2)
  -off
        Turn the relay off (default)
  -on
        Turn the relay on
```
Examples:
```
# turn on relay 1
$ arcade-tools -relay 1 -on

# turn off relay 2
$ arcade-tools -relay 2 -off
```

## JSON API endpoints

```
GET /api/v1/relays     # get info on both relays
GET /api/v1/relay/:id  # get info on specific relay

POST /api/v1/relay/:id # change relay state, id must be 1 or 2
# { "state": 1 } to turn a relay on
# { "state": 0 } to turn a relay off
```