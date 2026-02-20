# InfluxDB Resource Provider

The InfluxDB Resource Provider lets you manage [InfluxDB](https://www.influxdata.com/) resources.

## ⚠️ Deprecation Notice

This Pulumi provider for **InfluxDB** is **deprecated** and is no longer actively maintained as of **February 2026**.  
While earlier versions were published, **no further updates, bug fixes, or support will be provided**.

- Existing releases may continue to work, but compatibility with future Pulumi or InfluxDB versions is not guaranteed.  
- Do **not** use this provider for new projects.  
- The repository is archived and remains public for reference purposes only.  
- If you require continued usage, you may **fork the repository** and maintain your own version

## Supported InfluxDB flavours

### v3

* [InfluxDB Cloud Serverless](https://www.influxdata.com/products/influxdb-cloud/serverless/)

### v2

* [InfluxDB Cloud TSM](https://docs.influxdata.com/influxdb/cloud/)
* [InfluxDB OSS](https://docs.influxdata.com/influxdb/v2/)

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @thulasirajkomminar/influxdb
```

or `yarn`:

```bash
yarn add @thulasirajkomminar/influxdb
```

### Python

To use from Python, install using `pip`:

```bash
pip install thulasirajkomminar_influxdb
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/thulasirajkomminar/pulumi-influxdb/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package ThulasirajKomminar.InfluxDB
```

## Configuration

The following configuration points are available for the `influxdb` provider:

- `influxdb:password` (environment: `INFLUXDB_PASSWORD`) - The InfluxDB password
- `influxdb:token` (environment: `INFLUXDB_TOKEN`) - An InfluxDB token string
- `influxdb:url` (environment: `INFLUXDB_URL`) - The InfluxDB server URL
- `influxdb:username` (environment: `INFLUXDB_USERNAME`) - The InfluxDB username

## Authentication

The InfluxDB provider supports two [authentication methods](https://docs.influxdata.com/influxdb/v2/api/v2/#tag/Authentication):

* Token-based authentication (recommended).
* Username and password authentication.

### Authentication Priority

When both authentication methods are provided, **token authentication takes priority**. This means:

- If both `token` and `username`/`password` are configured, the provider will use token authentication
- Token authentication is the recommended method for better security and simplicity
- Username/password authentication is used only when no token is provided

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/influxdb/api-docs/).
