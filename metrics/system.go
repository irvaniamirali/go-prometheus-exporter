package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

// SystemCollector collects host system metrics
type SystemCollector struct {
	cpuUsage     *prometheus.Desc
	memUsed      *prometheus.Desc
	diskUsage    *prometheus.Desc
	netSentTotal *prometheus.Desc
	netRecvTotal *prometheus.Desc
}

func NewSystemCollector() *SystemCollector {
	return &SystemCollector{
		cpuUsage: prometheus.NewDesc(
			"system_cpu_usage_percent",
			"Current CPU usage in percent",
			nil, nil,
		),
		memUsed: prometheus.NewDesc(
			"system_memory_used_bytes",
			"Currently used memory in bytes",
			nil, nil,
		),
		diskUsage: prometheus.NewDesc(
			"system_disk_usage_percent",
			"Disk usage percentage for root partition",
			[]string{"path"}, nil,
		),
		netSentTotal: prometheus.NewDesc(
			"system_network_sent_bytes_total",
			"Total bytes sent over network",
			[]string{"interface"}, nil,
		),
		netRecvTotal: prometheus.NewDesc(
			"system_network_received_bytes_total",
			"Total bytes received over network",
			[]string{"interface"}, nil,
		),
	}
}

// Describe sends the super-set of all possible descriptors
func (c *SystemCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.cpuUsage
	ch <- c.memUsed
	ch <- c.diskUsage
	ch <- c.netSentTotal
	ch <- c.netRecvTotal
}

// Collect is called by Prometheus when collecting metrics
func (c *SystemCollector) Collect(ch chan<- prometheus.Metric) {
	// CPU
	if percent, err := cpu.Percent(time.Second, false); err == nil && len(percent) > 0 {
		ch <- prometheus.MustNewConstMetric(c.cpuUsage, prometheus.GaugeValue, percent[0])
	}

	// Memory
	if vm, err := mem.VirtualMemory(); err == nil {
		ch <- prometheus.MustNewConstMetric(c.memUsed, prometheus.GaugeValue, float64(vm.Used))
	}

	// Disk (root)
	if usage, err := disk.Usage("/"); err == nil {
		ch <- prometheus.MustNewConstMetric(c.diskUsage, prometheus.GaugeValue, usage.UsedPercent, usage.Path)
	}

	// Network
	if counters, err := net.IOCounters(true); err == nil {
		for _, counter := range counters {
			ch <- prometheus.MustNewConstMetric(c.netSentTotal, prometheus.CounterValue, float64(counter.BytesSent), counter.Name)
			ch <- prometheus.MustNewConstMetric(c.netRecvTotal, prometheus.CounterValue, float64(counter.BytesRecv), counter.Name)
		}
	}
}
