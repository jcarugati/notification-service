# Go Notification Service
This repository contains a Go-based implementation of a notification service that provides rate limiting and leverages a mock gateway for sending notifications.

### Features
- **Local Rate Limiter**: Prevents spamming by limiting the number of notifications sent in a given time frame.
- **Local Gateway**: Mocks the functionality of sending a notification. This can be replaced with a real-world gateway for actual notifications.
- **Flexible Rules**: Define custom rules for different types of notifications.
- **API Routes**: Provides an HTTP endpoint for sending notifications.

### Table of Contents
1. [Installation](#installation)
2. [Usage](#usage)
3. [Structure](#structure)

### Installation

Before you begin, make sure you have Go installed on your machine.

1. Clone the repository:
    ```bash
    git clone https://github.com/jcarugati/notification-service.git
    ```
2. Navigate to the repository folder:
    ```bash
    cd notification-service
    ```
3. Install the necessary dependencies (Assuming you are using Go Modules):
    ```bash
    go mod download
    ```
4. Compile and run:
    ```bash
   go build main.go
   ```
   ```bash
   go run main.go
   ```

### Usage
#### Sending Notifications:

Make a **POST** request to `/api/v1/notification` with the following JSON structure:
```json
{
  "user_id": "user123",
  "message": "Hello from Go Notification Service!",
  "type": "status"
}
```

The service will handle the rate-limiting and mock notification sending process.

The supported `type` values are declared in the manifest and can be modified along with their respective rules. By
default, the following types are supported:
- `status`
- `news`
- `marketing`

If the `type` sent is not supported the service will return a `400 Bad Request` response.
If the rate limit is exceeded for a user and a notification type, the service will return a `429 Too Many Requests` response.

#### Health Check:
Make a **GET** request to `/health` to check the health status of the service.

### Structure
- **LocalRateLimiter**: Provides rate limiting functionality using an in-memory cache.
- **LocalGateway**: A mock gateway to simulate the process of sending notifications.
- **NotificationService**: The main service responsible for sending notifications.
- **RateLimiter**: Interface for the rate limiter.
- **Gateway**: Interface for the notification gateway.
- **Manifest**: Represents the notification rules loaded from a YAML file.
