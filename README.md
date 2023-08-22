# Discovergy-Exporter

> Archived since Discovergy has discontinued their B2C business.

This is a prometheus exporter for [Discovergy](https://discovergy.com/) smart meters.

Currently, only supports exporting the current values of all connected meters:

* Total power usage in milliwatts
* Phase 1-3 power usage in milliwatts
* Voltage of phase 1-3 in millivolts
* Phase 1-3 current in milliampere

## Usage

Run exporter example:
```bash
go run cmd/exporter.go -email my.email.address@domain.tld -password YourDiscovergyPassword
```

By default the server is listening on `http://localhost:2112/metrics`

Help:
```
go run cmd/exporter.go -help
  -email string
        Discovergy account email
  -listen string
        Listen address (default ":2112")
  -password string
        Discovergy account password
  -sleep int
        Sleep time between refreshing Discovergy readings (default 30)
```
