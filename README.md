# senet

## Setup

The setup requires 3 terminals.

### Terminal 1

Install [air](https://github.com/cosmtrek/air) and run it in the root directory:

```shell
air
```

`air` will watch the Go files and rebuild the fairytale WASM file whenever you make any changes.

### Terminal 2

Install [sass](https://sass-lang.com/) and run it in the root directory:

```shell
sass --watch cmd/client-hypp/scss:cmd/client-hypp/public
```

`sass` will watch the scss files and recompile the css file whenever you make any changes.

### Terminal 3

Install the [fairytale cli](https://github.com/macabot/fairytale#cli) and run it in the root directory:

```shell
fairytale serve :8000 cmd/fairytale/main.wasm --watch --assets cmd/client-hypp/public
```

`fairytale` will watch the WASM file and assets. Whenever any changes are made it will reload the web page.
You can visit the fairytale app on http://localhost:8000/.

## Test

```sh
go test $(go list ./... | grep -vE 'cmd|tale|webrtc')
```

## Build

```sh
./cmd/client-hypp/build
```

## Development

### Package dependency tree

Red nodes directly or indirectly import `syscall/js`.

```mermaid
flowchart TD

subgraph cmd
    client-hypp
    fairytale
end

subgraph internal
    app

    subgraph app-group["app"]
        component
        dispatch
        state
        tale

        subgraph tale-group["tale"]
            control
        end
    end

    subgraph pkg
        clone
        set
        stack
        webrtc
    end
end

client-hypp --> app

fairytale --> state
fairytale --> tale

component --> state
component --> set
component --> dispatch

dispatch --> state

state --> clone
state --> set

control --> state

tale --> component
tale --> state
tale --> control
tale --> dispatch

app --> component
app --> state

classDef syscallJS fill:#f00;
class client-hypp,fairytale,tale,control,webrtc syscallJS;
```
