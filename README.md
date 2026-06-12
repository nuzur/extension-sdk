# extension-sdk

The nuzur Extension SDK is a Go library for building [nuzur](https://nuzur.com) extensions. Extensions are gRPC services that integrate with the nuzur platform to add custom functionality to projects.

## Prerequisites

- Go 1.21+
- A nuzur account at [nuzur.com](https://nuzur.com)
- The nuzur CLI — [nuzur.com/cli](https://nuzur.com/cli)

## Getting Started

### 1. Install the Nuzur CLI

The [nuzur-cli](https://github.com/nuzur/nuzur-cli) scaffolds a fully working extension project so you don't have to wire everything up manually.

Follow the installation instructions at [nuzur.com/cli](https://nuzur.com/cli).

### 2. Create your extension in the web app

Before scaffolding, you need to create an extension and an extension version in the [nuzur developer portal](https://app.nuzur.com/dev). 

### 3. Scaffold your extension

Log in and run the scaffold command:

```sh
nuzur-cli extension scaffold
```

The CLI will prompt you to select a project and configure your extension, then generate a ready-to-run project with the following structure:

```
my-extension/
├── main.go              # App entry point
├── config/
│   └── extension.yaml   # Extension metadata & config
├── server/
│   ├── server.go        # gRPC server implementation
│   ├── metadata.go      # GetMetadata handler
│   ├── start_execution.go   # StartExecution handler (your logic goes here)
│   ├── get_execution.go     # GetExecution handler
│   └── submit_step.go       # SubmitExecutionStep handler
├── constants/           # Extension-level constants
├── translations/        # i18n translation files
└── Dockerfile
```

### 4. Implement your extension logic

Open `server/start_execution.go`. The scaffolded handler already:

- Resolves typed config values from the request
- Fetches base dependencies (project schema, team, etc.) via the SDK client
- Creates an execution record on the Nuzur platform
- Returns a structured response

Add your custom logic where the `// TODO your extension logic here` comment is:

```go
func (s *server) StartExecution(ctx context.Context, req *pb.StartExecutionRequest) (*pb.StartExecutionResponse, error) {
    // ... setup (already scaffolded) ...

    // TODO your extension logic here

    // update execution status to succeeded when done
    _, err = s.client.UpdateExecution(ctx, client.UpdateExecutionRequest{
        ExecutionUUID:      uuid.FromStringOrNil(exec.Uuid),
        ProjectUUID:        projectUUID,
        ProjectVersionUUID: projectVersionUUID,
        Status:             pb.ExecutionStatus_EXECUTION_STATUS_SUCCEEDED,
    })
    // ...
}
```

### 5. Run your extension

Start the server locally:

```sh
go run main.go
```

The server listens on port `5555` by default. Set the `PORT` environment variable to override it.

## Configuration

The `config/extension.yaml` file identifies your extension to the Nuzur platform:

```yaml
metadata:
  uuid: <your-extension-uuid>
  version_uuid: <your-extension-version-uuid>
  identifier: my-extension
  display_name: My Extension
  description: What this extension does
  steps: 1
```

These values are populated automatically by the CLI scaffold command.

## SDK Client

The `*client.Client` injected into your server gives you access to Nuzur platform APIs:

| Method | Description |
|---|---|
| `GetBaseDependencies` | Fetch project schema, entities, team, and user info |
| `ResolveConfigValues` | Deserialize typed config values from the execution request |
| `CreateExecution` | Register a new execution record |
| `UpdateExecution` | Update execution status (in-progress → succeeded/failed) |
| `GetExecution` | Fetch an existing execution record |
| `DownloadExecutionResults` | Download previously uploaded results |
| `Localize` | Translate a string using the bundled i18n files |

## Module

If you are wiring the SDK into an existing `fx` application, import the module directly:

```go
import extensionsdk "github.com/nuzur/extension-sdk"

fx.New(
    extensionsdk.Module,
    fx.Provide(server.New),
).Run()
```

`extensionsdk.Module` provides a configured `*client.Client` and starts the gRPC server on lifecycle start.

## Links

- CLI: [nuzur.com/cli](https://nuzur.com/cli)
- CLI source: [github.com/nuzur/nuzur-cli](https://github.com/nuzur/nuzur-cli)
- SDK source: [github.com/nuzur/extension-sdk](https://github.com/nuzur/extension-sdk)
