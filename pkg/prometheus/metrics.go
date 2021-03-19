package prometheus

import "github.com/prometheus/client_golang/prometheus"

var MilliWattTotalGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "discovergy_milliwatt_total",
		Help: "Current power usage total in milliwatt",
	},
	[]string{
		"meterId",
	},
)

var MilliWattPhase1Gauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "discovergy_milliwatt_phase1",
		Help: "Phase 1 power usage in milliwatt",
	},
	[]string{
		"meterId",
	},
)

var MilliWattPhase2Gauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "discovergy_milliwatt_phase2",
		Help: "Phase 2 power usage in milliwatt",
	},
	[]string{
		"meterId",
	},
)

var MilliWattPhase3Gauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "discovergy_milliwatt_phase3",
		Help: "Phase 3 power usage in milliwatt",
	},
	[]string{
		"meterId",
	},
)

var VoltagePhase1Gauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "discovergy_voltage_phase1",
		Help: "Phase 1 voltage in millivolt",
	},
	[]string{
		"meterId",
	},
)

var VoltagePhase2Gauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "discovergy_voltage_phase2",
		Help: "Phase 2 voltage in millivolt",
	},
	[]string{
		"meterId",
	},
)

var VoltagePhase3Gauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "discovergy_voltage_phase3",
		Help: "Phase 3 voltage in millivolt",
	},
	[]string{
		"meterId",
	},
)

var AmperePhase1Gauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "discovergy_ampere_phase1",
		Help: "Phase 1 ampere in milliampere",
	},
	[]string{
		"meterId",
	},
)

var AmperePhase2Gauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "discovergy_ampere_phase2",
		Help: "Phase 2 ampere in milliampere",
	},
	[]string{
		"meterId",
	},
)

var AmperePhase3Gauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "discovergy_ampere_phase3",
		Help: "Phase 3 ampere in milliampere",
	},
	[]string{
		"meterId",
	},
)

func init() {
	prometheus.MustRegister(MilliWattTotalGauge)
	prometheus.MustRegister(MilliWattPhase1Gauge)
	prometheus.MustRegister(MilliWattPhase2Gauge)
	prometheus.MustRegister(MilliWattPhase3Gauge)

	prometheus.MustRegister(VoltagePhase1Gauge)
	prometheus.MustRegister(VoltagePhase2Gauge)
	prometheus.MustRegister(VoltagePhase3Gauge)

	prometheus.MustRegister(AmperePhase1Gauge)
	prometheus.MustRegister(AmperePhase2Gauge)
	prometheus.MustRegister(AmperePhase3Gauge)
}
