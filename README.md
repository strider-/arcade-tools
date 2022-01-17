# arcade-tools

Custom API / CLI tooling for controlling my modified Arcade1Up cabinet. Run `make help` to list all the available make commands for the project.

## Project structure
```
├── README.md # this document
├── go.mod
├── go.sum
├── cmd
│   ├── api   # json api project
│   └── cli   # command line interface project
└── internal  # private library code
```

## Expected RPi GPIO pin usage
NOTE: The relay in question is assumed active low
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