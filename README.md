# Shopcloud API`s

Interface definitions of Talk-Point APIs.

## Building

To build the client code in you application use the `buf` tool. You can install it by following the instructions on the [buf repository](https://github.com/bufbuild/buf), the best way is to use docker.

```bash
$ buf generate https://github.com/Talk-Point/shopcloud.git --path proto/shopcloud/eventhub
```

## Repository Structure

The repository is structured as follows:
- `proto` - Contains the proto files for the API`s.
- `client` - Contains go code to connect to the API`

