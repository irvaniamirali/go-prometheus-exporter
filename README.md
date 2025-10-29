# Go Custom Prometheus Exporter

A **production-ready** Prometheus exporter in Go that collects **system** and **application** metrics.

- **System**: CPU, Memory, Disk, Network  
- **App**: HTTP request count, latency (histogram)  
- Fully containerized with **Docker + Prometheus + Grafana**

---

## Live Dashboard

**Real-time monitoring**  
[View Live Grafana Snapshot](https://snapshots.raintank.io/dashboard/snapshot/REwZSBloPRPw02Jm5ugfC3nO2kP4D0ov)

---

## Features

- `prometheus/client_golang` integration
- `gopsutil` for accurate system metrics
- Simulated HTTP traffic for demo
- **Multi-stage Docker build** (~15MB final image)
- Ready-to-import Grafana dashboard
- `/health` endpoint
- Clean, testable, GitHub-ready code

---

## Quick Start

```bash
git clone https://github.com/irvaniamirali/go-prometheus-exporter.git
cd go-prometheus-exporter
docker-compose up -d
```

## Access

| Service    | URL                                                            |
| ---------- | -------------------------------------------------------------- |
| Exporter   | [http://localhost:9091/metrics](http://localhost:9091/metrics) |
| Prometheus | [http://localhost:9090](http://localhost:9090)                 |
| Grafana    | [http://localhost:3000](http://localhost:3000) (admin/admin)   |

**Grafana:** Import `grafana/dashboard.json` → **+** → **Import**

---

## Architecture

```
[Go Exporter] → /metrics → [Prometheus] → [Grafana Dashboard]
```

---

## Collector Pattern (Describe/Collect)

* **Counter** → Request count
* **Gauge** → System stats
* **Histogram** → Latency distribution

---

## Project Structure

```
.
├── main.go
├── metrics/
│   ├── system.go     # CPU, RAM, Disk, Network
│   └── app.go        # HTTP requests, latency
├── Dockerfile
├── docker-compose.yml
├── prometheus/prometheus.yml
└── grafana/dashboard.json
```

---

## Contributing

1. Fork the repo
2. Create branch: `git checkout -b feat/new-feature`
3. Commit: `git commit -m 'feat: add new feature'`
4. Push & open PR
