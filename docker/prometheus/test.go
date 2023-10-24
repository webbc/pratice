package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	//使用GaugeVec类型可以为监控指标设置标签，这里为监控指标增加一个标签"device"
	speed = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "disk_available_bytes",
		Help: "Disk space available in bytes",
	}, []string{"device"})

	tasksTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "test_tasks_total",
		Help: "Total number of test tasks",
	})

	taskDuration = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "task_duration_seconds",
		Help: "Duration of task in seconds",
		//Summary类型的监控指标需要提供分位点
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})

	cpuTemperature = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "cpu_temperature",
		Help: "The temperature of cpu",
		//Histogram类型的监控指标需要提供Bucket
		Buckets: []float64{20, 50, 70, 80},
	})
)

func init() {
	//注册监控指标
	prometheus.MustRegister(speed)
	prometheus.MustRegister(tasksTotal)
	prometheus.MustRegister(taskDuration)
	prometheus.MustRegister(cpuTemperature)
}

func main() {
	//模拟采集监控数据
	fakeData()

	//使用prometheus提供的promhttp.Handler()暴露监控样本数据
	//prometheus默认从"/metrics"接口拉取监控样本数据
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func fakeData() {
	tasksTotal.Inc()
	//设置该条样本数据的"device"标签值为"/dev/sda"
	speed.With(prometheus.Labels{"device": "/dev/sda"}).Set(82115880)

	taskDuration.Observe(10)
	taskDuration.Observe(20)
	taskDuration.Observe(30)
	taskDuration.Observe(45)
	taskDuration.Observe(56)
	taskDuration.Observe(80)

	cpuTemperature.Observe(30)
	cpuTemperature.Observe(43)
	cpuTemperature.Observe(56)
	cpuTemperature.Observe(58)
	cpuTemperature.Observe(65)
	cpuTemperature.Observe(70)
}
