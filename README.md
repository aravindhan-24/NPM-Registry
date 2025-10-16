# NPM-Registry

A Go-based implementation of a private NPM registry.

## Overview

The **NPM-Registry** project is designed to provide a private registry solution for managing and hosting your own NPM packages. Built using Go, this registry allows you to have more control over your packages and their distribution.

## Features

- **Private Package Hosting**: Host your own NPM packages securely.
- **Go Implementation**: Leverages Go for performance and scalability.
- **Easy Setup**: Simple configuration and deployment.

## Installation

### Prerequisites

- Go 1.18 or higher
- Node.js and npm

### Steps

1. Clone the repository:

   git clone https://github.com/aravindhan-24/NPM-Registry.git
   cd NPM-Registry


2. Build the Go application:

   go build -o npm-registry

3. Run the registry:

   ./npm-registry

4. Configure your NPM client to use the local registry:

   npm set registry http://localhost:2402/

## Usage

* To publish a package:

  npm publish --registry http://localhost:2402/

* To install a package:

  npm install <package-name> --registry http://localhost:2402/

## Configuration

The registry can be configured through environment variables or command-line flags. For detailed configuration options, refer to the Go application's documentation.
