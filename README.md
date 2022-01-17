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
├── internal  # private library code
└── pkg       # public code
```

## Expected RPi GPIO pin usage
NOTE: The relay in question is assumed active low
| Phys. Pin | GPIO # | Usage |
|----------:|---------:|-------|
|     2     |     5V   | Power |
|     9     |     GND  | Ground|
|    11     |     17   | Relay 2|
|    12     |     18   | Relay 1|