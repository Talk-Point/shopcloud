# Welcome to ShopCloud API Documentation - BETA

This documentation provides information about the ShopCloud API, which consists of various gRPC microservices. The API can be accessed using gRPC or gRPC-Web, and it also supports REST-style calls using curl.

## Overview

The ShopCloud API is composed of the following microservices:

## Overview

The ShopCloud API is composed of the following microservices:

| Microservice | Endpoint | Proto File | Web URL |
|--------------|----------|------------|---------|
| ContentHub | [ContentHub Documentation](contenthub.md) | [contenthub.proto](contenthub.proto) | [contenthub.talk-point.de](https://contenthub.talk-point.de) |
| SupplierHub | [SupplierHub Documentation](supplierhub.md) | [supplierhub.proto](supplierhub.proto) | [supplierhub.talk-point.de](https://supplierhub.talk-point.de) |
| HeartbeatHub | [HeartbeatHub Documentation](heartbeathub.md) | [heartbeathub.proto](heartbeathub.proto) | [heartbeathub.talk-point.de](https://heartbeathub.talk-point.de) |
| ImageHub | [ImageHub Documentation](imagehub.md) | [imagehub.proto](imagehub.proto) | [imagehub.talk-point.de](https://imagehub.talk-point.de) |
| Barcode API | [Barcode API Documentation](barcode-api.md) | [api.json](api.json) | [barcode.talk-point.de](https://barcode.talk-point.de) |
| MetaHub | [MetaHub Documentation](metahub.md) | [metahub.proto](metahub.proto) | [metahub.talk-point.de](https://metahub.talk-point.de) |
| MailHub | [MailHub Documentation](mailhub.md) | [mailhub.proto](mailhub.proto) | [mailhub.talk-point.de](https://mailhub.talk-point.de) |
| ConsoleHub | [ConsoleHub Documentation](consolehub.md) | [consolehub.proto](consolehub.proto) | [consolehub.talk-point.de](https://consolehub.talk-point.de) |
| EventHub | [EventHub Documentation](eventhub.md) | [eventhub.proto](eventhub.proto) | [eventhub.talk-point.de](https://eventhub.talk-point.de) |
| SageHub | [SageHub Documentation](sagehub.md) | [sagehub.proto](sagehub.proto) | [sagehub.talk-point.de](https://sagehub.talk-point.de) |
| NotificationHub | [NotificationHub Documentation](notificationhub.md) | [notificationhub.proto](notificationhub.proto) | [notificationhub.talk-point.de](https://notificationhub.talk-point.de) |
| WarehouseBrainHub | [WarehouseBrainHub Documentation](warehousebrainhub.md) | [warehousebrainhub.proto](warehousebrainhub.proto) | [warehousebrainhub.talk-point.de](https://warehousebrainhub.talk-point.de) |

Each microservice has its own set of proto files that define the available services, messages, and enums.

## Getting Started

To get started with the ShopCloud API, follow these steps:

1. Set up your development environment with the necessary tools and dependencies.
2. Obtain the required authentication credentials (API keys, tokens, etc.).
3. Choose your preferred method of interacting with the API (gRPC, gRPC-Web, or REST).
4. Refer to the documentation for each microservice to understand the available endpoints, request/response formats, and usage examples.

## Authentication

Authentication is required to access the ShopCloud API. Depending on the microservice, you may need to include API keys, tokens, or other authentication mechanisms in your requests. Refer to the individual microservice documentation for specific authentication details.

## Error Handling

The ShopCloud API returns standard gRPC error codes and messages in case of failures or invalid requests. Make sure to handle errors appropriately in your client application.

## Microservices

### ContentHub

The ContentHub microservice provides functionality related to managing content, attributes, tags, and more.

- [ContentHub Documentation](contenthub.md)
- [ContentHub Proto File](contenthub.proto)

### SupplierHub

The SupplierHub microservice allows interaction with supplier-related data and operations.

- [SupplierHub Documentation](supplierhub.md)
- [SupplierHub Proto File](supplierhub.proto)

### HeartbeatHub

The HeartbeatHub microservice enables monitoring and heartbeat functionality for cloud services.

- [HeartbeatHub Documentation](heartebeathub.md)
- [HeartbeatHub Proto File](heartebeathub.proto)

... (continue with other microservices)

## Additional Resources

- [gRPC Documentation](https://grpc.io/docs/)
- [gRPC-Web Documentation](https://github.com/grpc/grpc-web)
- [Protocol Buffers Documentation](https://developers.google.com/protocol-buffers)
- [MkDocs Documentation](https://www.mkdocs.org)

## Support

If you encounter any issues or have questions regarding the ShopCloud API, please contact our support team at support@talk-point.de.