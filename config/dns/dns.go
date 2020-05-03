package dns

// Config is used for configuring a DNS load test. See:
// https://godoc.org/github.com/miekg/dns#ClientConfig
// https://godoc.org/github.com/miekg/dns#Client.Exchange
type Config struct {
	Endpoints       []string `yaml:"endpoints"`
	ResourceRecords []string `yaml:"resourceRecords"`
	TimeoutSec      int      `yaml:"timeoutSec,omitempty"`
}
