# Groundcover Resource Provider

The Groundcover Resource Provider lets you manage [groundcover](https://groundcover.com) observability resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumiverse/groundcover
```

or `yarn`:

```bash
yarn add @pulumiverse/groundcover
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumiverse_groundcover
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumiverse/pulumi-groundcover/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumiverse.Groundcover
```

## Configuration

The following configuration options are available for the `groundcover` provider:

- `groundcover:apiKey` (environment: `GROUNDCOVER_API_KEY`) - groundcover API Key
- `groundcover:backendId` (environment: `GROUNDCOVER_BACKEND_ID`) - groundcover Backend ID
- `groundcover:apiUrl` (environment: `GROUNDCOVER_API_URL`) - groundcover API URL (defaults to `https://api.groundcover.com`)

## Resources

The provider supports the following resources:

- `groundcover.Apikey` - API Key management
- `groundcover.Dashboard` - Dashboard management
- `groundcover.Dataintegration` - Data integration configuration
- `groundcover.Ingestionkey` - Ingestion key management
- `groundcover.Logspipeline` - Logs pipeline configuration
- `groundcover.Metricsaggregation` - Metrics aggregation rules
- `groundcover.Monitor` - Monitor/alert configuration
- `groundcover.Policy` - Policy management
- `groundcover.Secret` - Secret management
- `groundcover.Serviceaccount` - Service account management

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/groundcover/api-docs/).
