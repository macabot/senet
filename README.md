# senet

## Setup

The setup requires 2 terminals.

**Terminal 1**: Install [air](https://github.com/cosmtrek/air) and run it in the root directory:

```shell
air
```

`air` will watch the Go files and rebuild the fairytale WASM file whenever you make any changes.

**Terminal 2**: Install the [fairytale cli](https://github.com/macabot/fairytale#cli) and run it in the root directory:

```shell
fairytale serve :8000 cmd/fairytale/main.wasm --watch --assets cmd/client-hypp/public
```

`fairytale` will watch the WASM file and assets. Whenever any changes are made it will reload the web page.
You can visit the fairytale app on http://localhost:8000/.
