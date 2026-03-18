<div align="center">

<br/>

```
████████╗██╗    ██╗██╗███╗   ██╗███████╗██╗      ██████╗ ██╗    ██╗
╚══██╔══╝██║    ██║██║████╗  ██║██╔════╝██║     ██╔═══██╗██║    ██║
   ██║   ██║ █╗ ██║██║██╔██╗ ██║█████╗  ██║     ██║   ██║██║ █╗ ██║
   ██║   ██║███╗██║██║██║╚██╗██║██╔══╝  ██║     ██║   ██║██║███╗██║
   ██║   ╚███╔███╔╝██║██║ ╚████║██║     ███████╗╚██████╔╝╚███╔███╔╝
   ╚═╝    ╚══╝╚══╝ ╚═╝╚═╝  ╚═══╝╚═╝     ╚══════╝ ╚═════╝  ╚══╝╚══╝
```

**Traffic replay and API breaking-change detection — before it reaches production.**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat-square&logo=docker)](https://docker.com)
[![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey?style=flat-square)]()

</div>

---

## What is TwinFlow?

TwinFlow records **real production traffic** and replays it against a new version of your service — catching breaking API changes, missing fields, and latency regressions **before** you deploy.

Think of it as a shadow test harness that uses your own live traffic as the test suite.

```
Production Traffic          New Service Version
      │                             │
      ▼                             ▼
  [ Record ]  ──── replay ────► [ Compare ]
      │                             │
  captures/                   ❌ BREAKING CHANGE
  traffic.log                  - Field removed: name
                               - Field added: username
                               STATUS: UNSAFE TO DEPLOY
```

---

## Features

| Feature                             | Description                                              |
| ----------------------------------- | -------------------------------------------------------- |
| 🎙️ **Traffic Recording**            | Capture real HTTP requests and responses from production |
| ▶️ **Traffic Replay**               | Replay captured traffic against any target service       |
| 🔍 **Breaking Change Detection**    | Detect added/removed/modified fields in API responses    |
| ⏱️ **Latency Regression Detection** | Flag performance degradations between versions           |
| 🖥️ **CLI-First Workflow**           | Simple, composable commands that fit any pipeline        |
| 🐳 **Docker Support**               | Zero-dependency deployment via Docker                    |
| 🌍 **Cross-Platform Builds**        | Pre-built binaries for Linux, macOS, and Windows         |

---

## Installation

### Option 1 — Build from Source

> Requires Go 1.21+

```bash
git clone https://github.com/Saad7890-web/Twinflow.git
cd twinflow
go build -o twinflow ./cmd/twinflow
```

### Option 2 — Cross-Platform Binaries

Build for all major platforms in one step:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o dist/twinflow-linux ./cmd/twinflow

# macOS
GOOS=darwin GOARCH=amd64 go build -o dist/twinflow-mac ./cmd/twinflow

# Windows
GOOS=windows GOARCH=amd64 go build -o dist/twinflow.exe ./cmd/twinflow
```

### Option 3 — Docker

```bash
docker build -t twinflow .
```

---

## Quick Start

### Step 1 — Record Production Traffic

Point TwinFlow at your running service to start capturing requests:

```bash
# Without Docker
./twinflow record --target http://localhost:9000

# With Docker
docker run --network=host -v $(pwd)/captures:/captures \
  twinflow record --target http://localhost:9000
```

Captured traffic is saved to the `captures/` directory.

### Step 2 — Deploy Your New Service Version

Bring up your candidate service (the version you want to test) on a different port — for example, `localhost:9001`.

### Step 3 — Replay and Compare

```bash
# Without Docker
./twinflow replay --target http://localhost:9001

# With Docker
docker run --network=host -v $(pwd)/captures:/captures \
  twinflow replay --target http://localhost:9001
```

TwinFlow replays every recorded request against the new service and compares the responses.

---

## Example Output

```bash
$ curl http://localhost:8080/user

Replaying 42 captured requests against http://localhost:9001...

[1/42] GET /health .............. ✅ OK (12ms)
[2/42] GET /user ................. ❌ BREAKING CHANGE
       - Field removed: name
       - Field added: username
[3/42] POST /orders .............. ✅ OK (34ms → 41ms, +20%)
...

────────────────────────────────────────────
SUMMARY
────────────────────────────────────────────
Total Requests   : 42
Passed           : 40
Breaking Changes : 1  ← GET /user
Latency Warnings : 1  ← POST /orders (+20%)

STATUS: ⛔ UNSAFE TO DEPLOY
────────────────────────────────────────────
```

---

## Use Cases

- **Safe Deployment** — Validate a new service version against real-world traffic before going live
- **API Contract Validation** — Ensure your API response shape never silently drifts between versions
- **Regression Detection** — Catch performance regressions using actual production request volumes
- **Microservice Testing** — Verify downstream service compatibility without writing manual test cases

---

## Roadmap

- [ ] CI/CD integration (GitHub Actions, GitLab CI)
- [ ] Kubernetes support
- [ ] HTML/JSON diff reports
- [ ] Request filtering and sampling
- [ ] gRPC support

---

## Contributing

Contributions are welcome! Please open an issue to discuss what you'd like to change, or submit a pull request directly.

1. Fork the repository
2. Create your feature branch: `git checkout -b feature/my-feature`
3. Commit your changes: `git commit -m 'Add my feature'`
4. Push to the branch: `git push origin feature/my-feature`
5. Open a pull request

---

## Author

Built by **Saad Islam Omy**

---

<div align="center">

If TwinFlow saves you from a bad deploy, consider giving it a ⭐

</div>
