# senet

## Setup
The setup requires 3 terminals.

Terminal 1: Run a file server for the static assets.
This example uses Python.
```shell
cd cmd/client-hypp/public
python3 -m http.server 8001
```

Terminal 2: Build the WASM file.
Run this command every time you make a change to the application.
```shell
cd cmd/fairytale
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

Terminal 3: Serve the fairytale app.
```shell
cd cmd/fairytale
fairytale serve :8000 main.wasm
```

You can visit the app on http://localhost:8000/.
