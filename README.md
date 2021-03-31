# Go assignment
Simple CLI made for VirtusLab internship assessment.

To run it clone this repository and build the project using
```
go build *.go
```

After that type 
```
./main help
```
for more detailed instructions.

## Methods
CLI has 2 main methods
 * `loadConfig()` that returns Config structure loaded from `config.json` file.
 * `startServer(filename string)` that starts listening on config.port with a file provided in the argument. It handles http requests via `net/http` package