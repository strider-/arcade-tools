# arcade-tools

Custom API / CLI tooling for controlling my modified Arcade1Up cabinet. Run `make help` to list all the available make commands for the project.
In order to run `make install` or `make update` locally on the Pi, you may need to ensure that `$PATH` contains the go binary location (If you're unsure where it's been installed to, run `which go`). To do this, edit `/etc/sudoers` and ensure that `Defaults secure_path` contains the go binary location.

The `GET /api/v1/now-playing` endpoint relies on the following line being somewhere in `/opt/retroarch/configs/all/runcommand-onstart.sh`:
```bash
printf "$1\n$2\n$3\n`basename ${3%.*}`\n$4" > ${AC_NOW_PLAYING_FILE:-/tmp/now-playing}
```
...and this line in `/opt/retroarch/configs/all/runcommand-onend.sh`:
```bash
:>| ${AC_NOW_PLAYING_FILE:-/tmp/now-playing}
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
├── scripts   # project scripts/service definitions
│             # to be ran/installed on the RPi via make
└── ui        # front-end for the json api
```

## Expected RPi GPIO pin usage
NOTE: The [relay](https://www.amazon.com/dp/B0057OC6D8?psc=1) in question is active low, all code reflects this
| Phys. Pin | GPIO # | Usage |
|----------:|---------:|-------|
|     2     |     5V   | Power |
|     9     |     GND  | Ground|
|    11     |     17   | Relay 2|
|    12     |     18   | Relay 1|

## JSON API Configuration

| Env. Var | Default | Usage |
|----------|---------|-------|
| `AC_NOW_PLAYING_FILE` | `/tmp/now-playing` | populated with info of ROM in play, empty when no game is being played |
| `AC_RELAY1_NAME` | `lighting` | A friendly name for what is plugged into relay #1 |
| `AC_RELAY2_NAME` | `monitor` | A friendly name for what is plugged into relay #2 |

To specify these values for the `arcade-api.service`, after installation run `systemctl edit arcade-api.service` and populate it like so, 
replacing `<Key>` with one of the above environment variable names and `<Value>` with your new value. After you save the file, run `systemctl restart arcade-api.service` for your changes to take effect.
```
[Service]
Environment=<Key>=<Value>
...
```


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
GET  /api/v1/now-playing # get info on the game being played,
                         # returns a 204 if no game is active

POST /api/v1/kill        # terminates the retroarch process,
                         # use in cases where emulation locks up.
                         # should dump you back to the ES menu

GET  /api/v1/relays      # get info on both relays
GET  /api/v1/relays/:id  # get info on specific relay, id must be 1 or 2

POST /api/v1/relays/:id  # change relay state, id must be 1 or 2
# { "state": 1 } to turn a relay on
# { "state": 0 } to turn a relay off
```

## Front End Configuration
Create a `.env` file based on the `.env.example` provided, populating the `REACT_APP_ARCADE_API_URL` accordingly. After running `npm install`, use `yarn start` to run the project locally. 

## Front End Deployment
There's an incredibly naive Makefile task, `deploy/fe` that will take a `target=<user@host>` argument and build/deploy the front end,
assuming the deployment user/server:
1. Can be accessed by ssh/scp access via public key
1. Has a site setup at `/var/www/arcade-fe`
1. Has read/write permissions in that directory