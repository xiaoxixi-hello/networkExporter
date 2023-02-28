package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type HTTPServer struct {
	r gin.IRoutes
	s *http.Server
}

func SetUp() (r *gin.Engine) {
	r = gin.New()

	// 创建一个自定义的注册表
	registry := prometheus.NewRegistry()

	// 在自定义的注册表中注册该指标
	for _, v := range a() {
		registry.MustRegister(v)
	}

	r.GET("/metrics", func(c *gin.Context) {
		handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry})
		handler.ServeHTTP(c.Writer, c.Request)
	})
	return
}

type NetCollector struct {
	queueLengthDesc *prometheus.Desc
	labelValues     []string
	v               float64
}

func (c *NetCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.queueLengthDesc
}

func (c *NetCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(c.queueLengthDesc, prometheus.GaugeValue, c.v, c.labelValues...)
}

// 网络入带宽
func inBandwidth(id string, i float64) *NetCollector {
	return &NetCollector{
		queueLengthDesc: prometheus.NewDesc(
			"tx_SRLine_InBandwidth",
			"The number of TencentCloud special railway line InBandwidth",
			// 动态标签的key列表
			[]string{"SRLine_zone", "SRLine_id"},
			// 静态标签
			prometheus.Labels{"module": "http-server"},
		),
		// 动态标签的值, 这里必须与声明的动态标签的key一一对应
		labelValues: []string{"上海金融区", id},
		v:           i,
	}
}

func a() []*NetCollector {
	scoreMap := make(map[string]float64)
	scoreMap["张三"] = 100
	scoreMap["小明"] = 100
	allNetCollector := make([]*NetCollector, 0, 4)
	for k, v := range scoreMap {
		allNetCollector = append(allNetCollector, &NetCollector{
			queueLengthDesc: prometheus.NewDesc(
				"tx_SRLine_InBandwidth",
				"The number of TencentCloud special railway line InBandwidth",
				// 动态标签的key列表
				[]string{"SRLine_zone", "SRLine_idc"},
				// 静态标签
				prometheus.Labels{"SRLine_id": k},
			),
			// 动态标签的值, 这里必须与声明的动态标签的key一一对应
			labelValues: []string{"上海金融区", k},
			v:           v,
		})
	}
	return allNetCollector
}

// OutBandwidth 网络出带宽
func OutBandwidth() *NetCollector {
	return &NetCollector{
		queueLengthDesc: prometheus.NewDesc(
			"tx_SRLine_OutBandwidth",
			"The number of TencentCloud special railway line OutBandwidth",
			// 动态标签的key列表
			[]string{"SRLine_zone", "SRLine_id"},
			// 静态标签
			prometheus.Labels{"module": "http-server"},
		),
		// 动态标签的值, 这里必须与声明的动态标签的key一一对应
		labelValues: []string{"上海金融区", "0001"},
		v:           98,
	}
}
