# Basic DNS Resolver in Go

This project implements a simple command-line DNS resolver in Go. It takes a domain name as input and prints its associated IP addresses.

## Features

*   Resolves a domain name to its corresponding IP addresses (IPv4 and IPv6).
*   Uses Go's built-in `net` package for DNS resolution.
*   Handles and displays errors gracefully.

## Prerequisites

*   Go (version 1.x or higher recommended) - [Installation Instructions](https://go.dev/doc/install)

## Installation

1.  Clone the repository:

    ```bash
    git clone <repository-url> 
    ```

2.  Navigate to the project directory:

    ```bash
    cd <project-directory>
    ```

3.  Build the project (optional, but useful for checking for errors):

    ```bash
    go build
    ```

## Usage

Run the program from the command line, providing a domain name as an argument:

```bash
go run main.go <domain_name>
