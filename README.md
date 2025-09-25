<p align="center">
<img width="512" height="512" alt="cache-go logo" src="https://github.com/user-attachments/assets/6547d51f-70b1-4525-aed4-9bd236594f8e">
</p>
<h1 align="center">cache-go</h1>
<p align="center">
A custom implementation of a high-performance cache in Go, inspired by Redis/Dragonfly, with exclusive communication via Unix Domain Sockets.
</p>
<p align="center">
<img src="https://img.shields.io/badge/build-passing-green" alt="Build Status">
<a href="https://golang.org/"><img src="https://img.shields.io/badge/Go-1.21+-blue.svg" alt="Go Version"></a>
<a href="https://www.google.com/search?q=LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License"></a>
</p>

## About the Project

cache-go is an in-memory cache server designed to be extremely fast and efficient in environments where the cache and application run on the same machine.

The main characteristic that sets it apart is the use of Unix Domain Sockets (UDS) for inter-process communication (IPC), eliminating the TCP/IP network stack overhead and resulting in drastically lower latency and higher throughput.

## Key Features

- **Extreme Performance**: Communication via Unix Sockets for near-instantaneous IPC.
- **Simple API**: Intuitive client interface for operations like SET, GET, and DELETE.
- **Lightweight and Minimal**: Written in pure Go, with no external dependencies.

## Why use Unix Sockets?

Unlike TCP/IP, which is designed for communication between different machines on a network, Unix Sockets are an IPC method that operates directly through the operating system kernel. This means:

* **Lower Latency**: No handshakes, packets, or network routes. Communication is direct.
* **Higher Throughput**: Data transfer is faster as it avoids the complexity of the network stack.
* **More Security**: Socket file access can be restricted using standard Unix file permissions (chmod), ensuring only authorized processes can connect to the cache.

cache-go is ideal for microservice architectures or monolithic applications running on the same host that need a shared cache with ultra-high speed.

## Roadmap (TODO)

- [ ] **Stress Testing**: Perform benchmarks and load tests to validate performance and stability.
- [ ] **Test Coverage**: Write unit and integration tests to ensure code reliability.
- [ ] **Database Integration**: Test cache-go as a cache layer for a real database.
- [ ] **Bloom Filters Implementation**:
  - [ ] Build Bloom Filter variations optimized for different data types.
  - [ ] Create a generic and flexible interface for using Bloom Filters in the cache.
- [ ]  **Complete Documentation**: Detail the API and project architecture.

## Contributions

Contributions are very welcome! If you have ideas for improvements or found a bug, feel free to open an Issue or a Pull Request.

### 📝 License

This project is licensed under GNUv3. See the LICENSE file for more details.
