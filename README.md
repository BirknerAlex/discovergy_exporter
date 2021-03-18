# Discovergy-Exporter

This is a prometheus exporter for [Discovergy](https://discovergy.com/) smart meters.

Currently, only supports exporting the current values of all connected meters:

* Total power usage in milliwatts
* Phase 1-3 power usage in milliwatts
* Voltage of phase 1-3 in millivolts

## Metrics endpoint

`http://localhost:2112/metrics`
