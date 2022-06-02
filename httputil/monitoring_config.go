package httputil

// MonitoringConfig structure of data for handling configuration for
// controlling what content is monitored.
type MonitoringConfig struct {
	MonitoredTypes   []string
	MonitoredMethods []string
}
