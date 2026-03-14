# TwinFlow

> HTTP traffic capture and replay for detecting breaking changes between service versions.

---

## Overview

TwinFlow lets you record live HTTP traffic from a running service and replay it against a new version — surfacing response diffs before a deployment goes wrong. It's built for teams that want a safety net during refactors, dependency upgrades, or API migrations.

---

## Features

- **HTTP Traffic Capture** — Transparently proxy and record requests/responses from any HTTP service
- **Request Replay** — Replay captured traffic against a new service version with accurate timing and headers
- **Response Diff Detection** — Automatically compare old vs. new responses and highlight deviations
- **CLI Interface** — Simple, scriptable commands with no UI overhead
- **Docker Support** — Run in any containerized environment with zero host dependencies

---

## Installation

```bash
# Using Go
go install github.com/your-org/twinflow@latest

# Using Docker
docker pull your-org/twinflow:latest
```

---

## Quick Start

### 1. Record traffic

Start TwinFlow as a proxy in front of your existing service. All traffic passing through will be captured to disk.

```bash
twinflow record --listen :8080 --target http://service:9000
```

| Flag       | Description                                              |
| ---------- | -------------------------------------------------------- |
| `--listen` | Address TwinFlow listens on                              |
| `--target` | Upstream service to proxy requests to                    |
| `--output` | Directory to write capture files (default: `./captures`) |

### 2. Replay against a new version

Point TwinFlow at the capture directory and a new service target. It will replay every recorded request and report any response differences.

```bash
twinflow replay --capture ./captures --target http://new-service:9000
```

| Flag        | Description                                                |
| ----------- | ---------------------------------------------------------- |
| `--capture` | Directory containing recorded traffic                      |
| `--target`  | New service version to replay against                      |
| `--report`  | Output format: `text`, `json`, or `html` (default: `text`) |

---

## How It Works

```
                   ┌─────────────────────────────────────┐
  Incoming         │             TwinFlow                 │
  Requests  ──────▶│  (proxy + capture / replay + diff)  │──────▶  Service
                   └─────────────────────────────────────┘
```

1. **Record mode** — TwinFlow acts as a transparent reverse proxy, forwarding requests to the target service and saving each request/response pair to a capture file.
2. **Replay mode** — TwinFlow reads the capture files, sends each request to the new service, and compares the responses against the originals.
3. **Diff detection** — Differences in status codes, headers, or response bodies are reported, helping you catch breaking changes before they reach production.

---

## Docker

```bash
# Record mode
docker run --rm -v $(pwd)/captures:/captures \
  your-org/twinflow record \
  --listen :8080 \
  --target http://service:9000 \
  --output /captures

# Replay mode
docker run --rm -v $(pwd)/captures:/captures \
  your-org/twinflow replay \
  --capture /captures \
  --target http://new-service:9000
```

---

## Roadmap

- [ ] gRPC support
- [ ] Request filtering and sampling
- [ ] Noise reduction (ignore known-volatile fields like timestamps)
- [ ] Web UI for diff visualization
- [ ] CI/CD integration (GitHub Actions, GitLab CI)

---

## Contributing

Contributions are welcome. Please open an issue before submitting a pull request for significant changes.

1. Fork the repository
2. Create a feature branch (`git checkout -b feat/my-feature`)
3. Commit your changes (`git commit -m 'feat: add my feature'`)
4. Push and open a pull request

---

## License

[MIT](LICENSE)
