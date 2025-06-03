# senet

Play at <https://senet.macabot.com>.

## Dependencies

- [Go](https://go.dev/)
- [Sass](https://sass-lang.com/)
- [Air](https://github.com/cosmtrek/air)
- [Fairytale cli](https://github.com/macabot/fairytale#cli)
- [Brotli](https://github.com/google/brotli)
- [wasm-opt](https://github.com/WebAssembly/binaryen)

## Setup

The setup requires the following terminals.

### Terminal 1

Run `air` in the root directory:

```shell
air
```

`air` will watch the files and run the [build script](#build) whenever you make any changes.

### Terminal 2

Run the fairytale cli in the root directory:

```shell
fairytale serve :8000 cmd/fairytale/main.wasm --watch --assets public
```

`fairytale` will watch the WASM file and assets. Whenever any changes are made it will reload the web page.
You can visit the fairytale app on <http://localhost:8000/>.

### Terminal 3

Run the static file server

```shell
go run cmd/server/main.go -d ./public
```

You can visit the Senet app on <http://localhost:8001/>.

## Test

```sh
go test $(GOOS=js GOARCH=wasm go list ./... | grep -v 'cmd')
```

```sh
docker run --rm \
    -v "$(pwd)":/workspace \
    -v "$HOME/go/pkg/mod":/go/pkg/mod \
    -v "$HOME/.cache/go-build":/root/.cache/go-build \
    -w /workspace \
    macabot/senet-builder:0.2.0 go test $(GOOS=js GOARCH=wasm go list ./... | grep -v 'cmd')
```

## Build

Run the build script as follows:

```sh
./build $environment [$public_dir]
```

The `$environment` must be either 'development' or 'production'.
The `$public_dir` is optional and defaults to './public'.

Alternatively, you may use the Docker `senet-builder` container:

```sh
docker run --rm \
    -v "$(pwd)":/workspace \
    -v "$HOME/go/pkg/mod":/go/pkg/mod \
    -v "$HOME/.cache/go-build":/root/.cache/go-build \
    -w /workspace \
    macabot/senet-builder:0.2.0 ./build development
```

## Development

### Creating a new Docker image

Build the Docker image:

```sh
docker build -t macabot/senet-builder:$tag .
```

Push the Docker image to the registry:

```sh
docker push macabot/senet-builder:$tag
```

See <https://hub.docker.com/r/macabot/senet-builder>.

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
class client-hypp,fairytale syscallJS;
```

## Page navigation

```mermaid
graph LR
    GitHub

    subgraph home["Home"]
        home-play["play"]
        home-rules["rules"]
        home-sourceCode["source-code"]
    end

    subgraph rules["Rule"]
        rules-start["start"]
        rules-home["home"]
    end

    subgraph SPA

        subgraph start["Start"]
            start-home["home"]
            start-tutorial["tutorial"]
            start-local["local"]
            start-online["online"]
            start-rules["rules"]
            start-sourceCode["source-code"]
        end

        subgraph online["Online"]
            online-new["new"]
            online-join["join"]
            online-back["back"]
        end
        subgraph NewGame
            NewGame-next["next"]
            NewGame-cancel["cancel"]
        end
        subgraph JoinGame
            JoinGame-next["next"]
            JoinGame-cancel["cancel"]
        end

        subgraph WhoGoesFirst
            WhoGoesFirst-play["play"]
            WhoGoesFirst-cancel["cancel"]
        end

        subgraph game
            game-quit["quit"]
        end

    end

    GitHub --> home

    home-play --> start
    home-rules --> rules
    home-sourceCode --> GitHub

    rules-home --> home
    rules-start --> start

    start-home --> home
    start-rules --> rules
    start-sourceCode --> GitHub
    start-online --> online
    start-tutorial --> game
    start-local --> game

    online-back --> start
    online-new --> NewGame
    online-join --> JoinGame

    NewGame-next --> WhoGoesFirst
    NewGame-cancel --> start

    JoinGame-next --> WhoGoesFirst
    JoinGame-cancel --> start

    WhoGoesFirst-play --> game
    WhoGoesFirst-cancel --> start

    game-quit --> start
```

## Online player vs players

When playing an online game, two players use [WebRTC](https://developer.mozilla.org/en-US/docs/Web/API/WebRTC_API) to connect directly to one another.
This means there is no trusted third party to generate a random throw of the sticks.
Instead a [commitment scheme](https://en.wikipedia.org/wiki/Coin_flipping#Telecommunications) is used:

```mermaid
sequenceDiagram
    Note right of Flipper: Flipper is player whose turn it is
    loop Until end of game
        par Wait for opponent to be ready
            Flipper->>Caller: Send is ready
        and
            Caller->>Flipper: Send is ready
        end
        Flipper-->>Flipper: Generate Flipper secret
        Flipper->>Caller: Send Flipper secret
        Caller-->>Caller: Generate Caller secret
        Caller-->>Caller: Generate Caller predictions
        Caller-->>Caller: Generate commitment
        Caller->>Flipper: Send commitment
        Flipper-->>Flipper: Generate Flipper results
        Flipper->>Caller: Send Flipper results
        Caller->>Flipper: Send Caller secret and predictions
        Flipper-->>Flipper: Verify commitment
        Flipper-->>Flipper: Player throws sticks
        Flipper->>Caller: Send sticks are thrown
        Flipper-->>Flipper: Player moves piece
        Flipper->>Caller: Send piece is moved
    end
    Note right of Flipper: Depending on throw, Flipper and Caller switch roles
```

The throw of the sticks is based on the [XNOR](https://en.wikipedia.org/wiki/XNOR_gate) operation on every prediction and result pair.
E.g.
| Caller predictions | Flipper results | Throw flips |
| ------------------ | --------------- | ----------- |
| 1 | 1 | 1 |
| 0 | 1 | 0 |
| 0 | 1 | 0 |
| 0 | 0 | 1 |
| | Throw | 2 |
