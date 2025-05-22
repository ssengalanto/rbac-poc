# RBAC Proof of Concept (rbac-poc)

## Overview

This repository demonstrates a Role-Based Access Control (RBAC) implementation using Go (Golang). It showcases how to manage access permissions based on user roles within a system, ensuring that users can only perform actions that their roles permit.

## Features

- **Role Management**: Define and manage various user roles.
- **Policy Enforcement**: Implement policies that dictate what actions are allowed for each role.
- **Middleware Integration**: Utilize middleware to enforce RBAC policies in your application.

## Project Structure

```
rbac-poc/
├── .idea/              # IDE configuration files
├── middleware/         # RBAC middleware implementation
├── go.mod              # Go module definition
├── go.sum              # Go module dependencies
├── main.go             # Entry point of the application
├── model.conf          # Model configuration for RBAC
└── policy.csv          # RBAC policies defining permissions
```

## Technologies Used

- **Go (Golang)**: The primary programming language used for backend development.
- **Casbin**: Authorization library used to define and enforce role-based access control (RBAC) policies via middleware.
- **Go-Chi**: Lightweight HTTP router and middleware framework for defining routes and handling requests.

## Setup Instructions

### Prerequisites

- Go 1.18 or higher installed on your machine.
- A code editor like Visual Studio Code or GoLand for development.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ssengalanto/rbac-poc.git
   cd rbac-poc
   ```

2. Initialize Go modules:

   ```bash
   go mod tidy
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

   The application will start, and you can begin testing the RBAC implementation.

## Sample Requests
   ```bash
curl -H "X-User-Role: admin" http://localhost:8080/workspaces
curl -H "X-User-Role: account_owner" http://localhost:8080/users
curl -H "X-User-Role: event_owner" http://localhost:8080/guests
curl -H "X-User-Role: co_host" http://localhost:8080/workspaces
   ```

## Usage

- Define roles and permissions in the `policy.csv` file.
- Configure models and policies in `model.conf`.
- Implement the RBAC middleware in your Go application to enforce access control based on the defined policies.

## Contributing

Contributions are welcome! Please fork the repository, create a new branch, and submit a pull request with your proposed changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
