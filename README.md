# Sequencer

sequencer is a project which exposes an endpoint `v1/sequence` which returns a unique sequence number that is thread safe.

## Getting Started Guide
1. Project Go version: [1.21.0](https://go.dev/doc/go1.21)
2. [Docker](https://www.docker.com/get-started/)

### Build the project
run the following command in the project root.
```text
make build
```

### Run the project
run the following command in the project root.
```
make run
```

This will open up 2 ports with the paths:
- 3000 - sequence service
    - `/v1/sequence` used for getting the sequence number
- 3001 - debug service
    - `/debug/readiness` used for determining if the service is ready to start receiving requests
    - `/debug/liveness` used for probing the service to see if it's still alive.
    - `/debug/pprof/` used for visualising and analysing profile data.
    - `/debug/pprof/cmdline` used for getting the running program's command line, with arguments separated by NUL bytes.
    - `/debug/pprof/profile` used for getting pprof-formatted cpu profile.
    - `/debug/pprof/symbol` used for getting program counters listed in the request, responding with a table mapping program counters to function names.
    - `/debug/pprof/trace` used for getting the execution trace in binary form.

### Test sequence endpoint
```
curl localhost:3000/v1/sequence
```

### Project Layout
**/cmd** - The main application for the project.

**/internal** - Private application code.
```text
sequencer
│   README.md
│   Dockerfile
│   makefile
│   go.mod
│   go.sum
│
└───cmd
│   │
│   └───app
│       │   main.go
│   
└───internal
│   │
│   └───atomicSequence
│   │   │   atomicSequence.go
│   │   │   
│   │   └───test
│   │       │   atomicSequence_test.go
│   │
│   └───config
│   │   │   config.go
│   │   │   
│   │   └───test
│   │       │   config_test.go
│   │
│   └───handlers
│   │   │   handlers.go
│   │   │   
│   │   └───v1
│   │   │   │   v1.go
│   │   │   
│   │   └───health
│   │       │   health.go
```

### Setting custom ports
You can change the ports in which the service runs on by specifying `-e API_PORT=<port>` and/or `-e DEBUG_PORT=<port>`
during the docker run stage. Note you will also have to update the exposed ports using `-p` if you want to use the new ports. example:
```
docker run -e API_PORT=4000 -e DEBUG_PORT=4001 -p 4000:4000 -p 4001:4001 sequencer
```





