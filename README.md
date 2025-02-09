# Internet Download Manager (IDM)

## Introduction
The **Internet Download Manager (IDM)** is a secure and role-based file download management system designed for enterprise environments. Unlike traditional download managers, IDM enforces strict access control, ensuring that only authorized users can download specific files. Administrators have full control over file access, while regular users are restricted to pre-approved files, enhancing security and compliance.

## Features
- **Authentication & Authorization:** Secure login with role-based access control (RBAC).
- **Role Management:** Admins can manage user roles and permissions.
- **Restricted Downloads:** Users can only download pre-approved files, while admins have full access.
- **Enterprise Security:** Ensures compliance with organizational policies by limiting unauthorized downloads.
- **Efficient Data Handling:** Integrates caching, message queuing, and distributed storage for scalability.

## Folder Structure
```
├── api            # Contains API definitions and external communication
├── cmd            # Main entry point of the application
├── internal       # Core application logic
│   ├── dataaccess # Handles data access and processing from S3, database
│   │   ├── migrations  # Database migration scripts
│   │   │   ├── mysql   # Migration scripts for MySQL
│   ├── handler    # Outer layer handling HTTP requests
│   ├── logic      # Contains business logic
```

## Technologies Used
- **Programming Language:** Golang
- **Message Broker:** Kafka
- **Cache:** Redis
- **Database:** MySQL
- **Block Storage:** S3 (using MinIO as a replacement)

## Installation and Running
```sh
# Clone the repository
$ git clone <repo_url>
$ cd <project_folder>

# Run the application
$ go run cmd/main.go
```

## Contributions
Contributions are welcome! Please create a pull request or open an issue for discussion.

---
**Contact:** For questions or support, please reach out via email or create an issue on GitHub.

