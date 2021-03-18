package discovergy

import (
	"github.com/BirknerAlex/discovergy_exporter/pkg/prometheus"
	"log"
	"time"
)

func RunUpdater(email string, password string) error {
	c := &Client{
		email:    email,
		password: password,
	}

	meters, err := c.GetMeters()
	if err != nil {
		return err
	}

	for {
		for _, meter := range meters {
			log.Printf("Requesting last reading for meter %v", meter.ID)
			lastReading, err := c.LastReading(meter.ID)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			label := map[string]string{"meterId": meter.ID}

			prometheus.MilliWattTotalGauge.With(label).Set(float64(lastReading.Values.MilliWattTotalPower))
			prometheus.MilliWattPhase1Gauge.With(label).Set(float64(lastReading.Values.MilliWattPhase1Power))
			prometheus.MilliWattPhase2Gauge.With(label).Set(float64(lastReading.Values.MilliWattPhase2Power))
			prometheus.MilliWattPhase3Gauge.With(label).Set(float64(lastReading.Values.MilliWattPhase3Power))

			prometheus.VoltagePhase1Gauge.With(label).Set(float64(lastReading.Values.VoltagePhase1))
			prometheus.VoltagePhase2Gauge.With(label).Set(float64(lastReading.Values.VoltagePhase2))
			prometheus.VoltagePhase3Gauge.With(label).Set(float64(lastReading.Values.VoltagePhase3))
		}

		time.Sleep(5 * time.Second)
	}
}
