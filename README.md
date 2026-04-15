# Pinger | Network Latency Checker

<p align="center">
  <img src="https://img.shields.io/badge/Network-Latency%20Check-00ADD8?style=for-the-badge" alt=""/>
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go" alt=""/>
</p>

---

### Fast, concurrent network latency measurement

Pinger sends ICMP echo requests to multiple hosts simultaneously and reports round-trip times with statistics.

```bash
./pinger google.com 8.8.8.8 github.com
```

**Features:**
- Concurrent pings for speed
- Statistics: min, max, avg, stddev
- JSON output for monitoring systems
- Configurable count and interval

---

## Quick Start

```bash
# Ping multiple hosts
./pinger 1.1.1.1 8.8.8.8 google.com

# With custom settings
./pinger -count 10 -interval 200ms google.com

# JSON output
./pinger -json hosts.txt
```

---

## Output Example

```
PING google.com (142.250.80.46)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

64 bytes from 142.250.80.46: icmp_seq=1 ttl=117 time=14.2ms
64 bytes from 142.250.80.46: icmp_seq=2 ttl=117 time=13.8ms
64 bytes from 142.250.80.46: icmp_seq=3 ttl=117 time=14.1ms

--- google.com ping statistics ---
3 packets transmitted, 3 received, 0% loss
rtt min/avg/max = 13.8/14.0/14.2 ms
```

---

MIT © 2024 [simplestar-992](https://github.com/simplestar-992)
