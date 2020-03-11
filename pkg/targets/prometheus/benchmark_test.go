package prometheus

import (
	"github.com/prometheus/prometheus/prompb"
	"github.com/timescale/tsbs/cmd/tsbs_load_prometheus/adapter/noop"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestPrometheusLoader(t *testing.T) {
	adapter := noop.Adapter{}
	server := httptest.NewServer(http.HandlerFunc(adapter.Handler))
	serverURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	pb := Benchmark{AdapterWriteUrl: serverURL.String()}
	pp := pb.GetProcessor().(*Processor)
	batch := &Batch{series: []prompb.TimeSeries{{}}}
	samples, _ := pp.ProcessBatch(batch, true)
	if samples != 1 {
		t.Error("wrong number of samples")
	}
	if adapter.SampleCounter != samples {
		t.Error("wrong number of samples processed")
	}
}