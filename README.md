<p align="center">
<img width="512" height="512" alt="cache-go logo" src="https://github.com/user-attachments/assets/6547d51f-70b1-4525-aed4-9bd236594f8e">
</p>
<h1 align="center">cache-go</h1>

<p align="center">
Uma implementação customizada de um cache de alta performance em Go, inspirado no Redis/Dragonfly, com comunicação exclusiva via Unix Domain Sockets.
</p>

<p align="center">
<img src="https://img.shields.io/badge/build-passing-green" alt="Build Status">
<a href="https://golang.org/"><img src="https://img.shields.io/badge/Go-1.21+-blue.svg" alt="Go Version"></a>
<a href="https://www.google.com/search?q=LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License"></a>
</p>

## ⚡️ Sobre o Projeto

cache-go é um servidor de cache de em memória, projetado para ser extremamente rápido e eficiente em ambientes onde o cache e a aplicação rodam na mesma máquina.

A principal característica que o diferencia é o uso de Unix Domain Sockets (UDS) para comunicação entre processos (IPC), eliminando o overhead da pilha de rede TCP/IP e resultando em latência drasticamente menor e maior throughput.

## ✨ Features Principais

- Performance Extrema: Comunicação via Unix Sockets para IPC quase instantâneo.

- API Simples: Interface de cliente intuitiva para operações como SET, GET e DELETE.

- Leve e Mínimo: Escrito em Go puro, sem dependências externas.

## 🤔 Por que usar Unix Sockets?

Ao contrário do TCP/IP, que é projetado para comunicação entre diferentes máquinas em uma rede, os Unix Sockets são um método de IPC que opera diretamente através do kernel do sistema operacional. Isso significa:

* Menor Latência: Não há handshakes, pacotes, ou rotas de rede. A comunicação é direta.

* Maior Throughput: A transferência de dados é mais rápida, pois evita a complexidade da pilha de rede.

* Mais Segurança: O acesso ao arquivo do socket pode ser restrito usando as permissões padrão de arquivos do Unix (chmod), garantindo que apenas processos autorizados possam se conectar ao cache.

cache-go é ideal para arquiteturas de microsserviços ou aplicações monolíticas que rodam no mesmo host e precisam de um cache compartilhado de altíssima velocidade.

## 🗺️ Roadmap (TODO)

- [ ] 🧪 Testes de Stress: Realizar benchmarks e testes de carga para validar a performance e a estabilidade.

- [ ] ✅ Cobertura de Testes: Escrever testes unitários e de integração para garantir a confiabilidade do código.

- [ ] 🐘 🐬 Integração com Banco de Dados: Testar o cache-go como uma camada de cache para um banco de dados real.

- [ ] 🌺 Implementação de Bloom Filters:

- [ ] Construir variações de Bloom Filters otimizadas para diferentes tipos de dados.

- [ ] Criar uma interface genérica e flexível para o uso de Bloom Filters no cache.

- [ ] 📄 Documentação Completa: Detalhar a API e a arquitetura do projeto.

## 🤝 Contribuições
Contribuições são muito bem-vindas! Se você tem ideias para melhorias ou encontrou um bug, sinta-se à vontade para abrir uma Issue ou um Pull Request.

### 📝 Licença
Este projeto está sob a licença GNUv3. Veja o arquivo LICENSE para mais detalhes.
