package reporter

import "github.com/ferpart/hlsvalidator/domain"

func Report(hlsReporters []domain.IHLSReporter) {
	for _, hlsReporter := range hlsReporters {
		hlsReporter.Report()
	}
}
