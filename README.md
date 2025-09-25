<p align="center">
<img width="512" height="512" alt="bloom-filter-go logo" src="https://github.com/user-attachments/assets/32a55c46-c427-4509-85ef-8bc264bb5de9">
</p>
<h1 align="center">bloom-filter-go</h1>
<p align="center">
A custom implementation of a high-performance bloom filter in Go, inspired by Redis/Dragonfly implementation, with exclusive communication via Unix Domain Sockets.
</p>
<p align="center">
<img src="https://img.shields.io/badge/build-passing-green" alt="Build Status">
<a href="https://golang.org/"><img src="https://img.shields.io/badge/Go-1.21+-blue.svg" alt="Go Version"></a>
<a href="https://www.google.com/search?q=LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License"></a>
</p>

## About the Project

bloom-filter-go is an in-memory bloom filter server designed to be extremely fast and efficient in environments where the bloom filter and application run on the same machine.

The main characteristic that sets it apart is the use of Unix Domain Sockets (UDS) for inter-process communication (IPC), eliminating the TCP/IP network stack overhead and resulting in drastically lower latency and higher throughput.

## Key Features

- **Extreme Performance**: Communication via Unix Sockets for near-instantaneous IPC.
- **Simple API**: Intuitive client interface for operations like ADD and EXISTS
- **Lightweight and Minimal**: Written in pure Go, with no external dependencies.

## Why use Unix Sockets?

Unlike TCP/IP, which is designed for communication between different machines on a network, Unix Sockets are an IPC method that operates directly through the operating system kernel. This means:

* **Lower Latency**: No handshakes, packets, or network routes. Communication is direct.
* **Higher Throughput**: Data transfer is faster as it avoids the complexity of the network stack.
* **More Security**: Socket file access can be restricted using standard Unix file permissions (chmod), ensuring only authorized processes can connect to the bloom filter.

bloom-filter-go is ideal for microservice architectures or monolithic applications running on the same host that need a shared bloom-filter with ultra-high speed.

## Roadmap (TODO)
- [ ] **Implement the MEXISTS and MADD on USD**
- [ ] **Stress Testing**: Perform benchmarks and load tests to validate performance and stability.
- [ ] **Test Coverage**: Write unit and integration tests to ensure code reliability.
- [ ] **Database Integration**: Test bloom-filter-go for exists validation for a real database.
- [ ] **Bloom Filters Implementation**:
  - [ ] Build Bloom Filter variations optimized for different data types.
  - [ ] Create a generic and flexible interface for using Bloom Filters.
- [ ]  **Complete Documentation**: Detail the API and project architecture.

## Contributions

Contributions are very welcome! If you have ideas for improvements or found a bug, feel free to open an Issue or a Pull Request.

### 📝 License

This project is licensed under GNUv3. See the LICENSE file for more details.
