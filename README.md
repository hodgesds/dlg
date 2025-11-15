# DLG - Da Load Generator
`dlg` is a universal load generator that can be used to generate load for various protocols and services.

## Supported Protocols

DLG supports the following protocols:

- **HTTP/HTTPS** - HTTP/HTTPS load testing with customizable headers, methods, and payloads
- **gRPC** - gRPC service load testing with metadata and connection options
- **WebSocket** - WebSocket connection testing with read/write operations
- **GraphQL** - GraphQL query execution with variables and custom headers
- **MongoDB** - MongoDB database operations (insert, find, update, delete, count, aggregate)
- **Redis** - Redis commands and operations
- **SQL** - MySQL, PostgreSQL, and ClickHouse database testing
- **MQTT** - MQTT publish/subscribe messaging with QoS levels
- **ETCD** - ETCD key-value operations
- **DNS** - DNS query load testing
- **LDAP** - LDAP authentication and query testing
- **Memcache** - Memcache get/set/delete operations
- **SNMP** - SNMP protocol testing
- **SSH** - SSH command execution
- **UDP** - UDP protocol testing
- **DHCP4** - DHCP4 protocol testing

## Features

- **Multi-protocol support** - Test multiple protocols in a single plan
- **Staged execution** - Organize tests into stages with configurable repeats and durations
- **Concurrent execution** - Run stages concurrently for maximum load
- **Prometheus metrics** - Built-in Prometheus metrics for monitoring
- **Rate limiting** - Control load with configurable rate limiters
- **Flexible configuration** - YAML-based configuration with extensive options
