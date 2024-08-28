# Shopcloud API`s

Interface definitions of Talk-Point APIs.

## Services

- ContentHub - API for managing product content
- EventHub - API for managing events
- HeartbeatHub - API for managing heartbeats and KPI`s
- ImageHub - API for managing images and create the cdn links
- MailHub - API for sending emails and receving mails
- NotificationHub - API for notifications
- SageHub - Connector API for the ERP System
- SuppliuerHub - API for managing supplier product and orders
- WarehousebrainHub - API automatic trading system
- BarcodeHub - API for managing barcodes

## Building

To build the client code in you application use the `buf` tool. You can install it by following the instructions on the [buf repository](https://github.com/bufbuild/buf), the best way is to use docker.

```bash
$ buf generate https://github.com/Talk-Point/shopcloud.git --path proto/shopcloud/eventhub
```

## Repository Structure

The repository is structured as follows:
- `proto` - Contains the proto files for the API`s.
- `client` - Contains go code to connect to the API`

