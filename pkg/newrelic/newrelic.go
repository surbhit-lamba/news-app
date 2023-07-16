package newrelic

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/grpc/metadata"
)

var (
	// agent is a pointer to the newrelic app
	agent *newrelic.Application = nil
)

// Options will be exposed to initialize new relic
type Options struct {
	Name                               string
	License                            string
	Enabled                            bool
	CrossApplicationTracer             bool
	DistributedTracer                  bool
	MaxSamplesStored                   int
	TransactionTracerThresholdDuration time.Duration
}

// Attribute struct for adding attributes to transaction
type Attribute struct {
	Name  string
	Value interface{}
}

// Initialize newrelic
func Initialize(op *Options) *newrelic.Application {
	if !op.Enabled {
		return nil
	}
	var err error
	agent, err = newrelic.NewApplication(
		newrelic.ConfigAppName(op.Name),
		newrelic.ConfigLicense(op.License),
		newrelic.ConfigDistributedTracerEnabled(op.DistributedTracer),
		newrelic.ConfigEnabled(op.Enabled),
		func(c *newrelic.Config) {
			// set defaults
			if op.MaxSamplesStored > 0 {
				c.TransactionEvents.MaxSamplesStored = op.MaxSamplesStored
			}
			c.CrossApplicationTracer.Enabled = op.CrossApplicationTracer

			if op.TransactionTracerThresholdDuration != 0 {
				c.TransactionTracer.Threshold.IsApdexFailing = false
				c.TransactionTracer.Threshold.Duration = op.TransactionTracerThresholdDuration
			}
		},
	)
	if err != nil {
		panic("Could not initialize newrelic: " + err.Error())
	}

	return agent
}

// StartSegmentWithContext takes  ctx and segment name
func StartSegmentWithContext(ctx context.Context, name string) *newrelic.Segment {
	txn := newrelic.FromContext(ctx)
	return txn.StartSegment(name)
}

// StartSegment takes segment name and txn
func StartSegment(name string, txn *newrelic.Transaction) *newrelic.Segment {
	return txn.StartSegment(name)
}

func StartNonWebTransaction(name string) *newrelic.Transaction {
	return agent.StartTransaction(name)
}

// AddAttributeWithContext adds attributes to transaction available in context
func AddAttributeWithContext(ctx context.Context, attrs ...Attribute) {
	txn := newrelic.FromContext(ctx)
	// In case of disabled, don't do anything
	if agent == nil || txn == nil {
		return
	}

	for _, attr := range attrs {
		txn.AddAttribute(attr.Name, attr.Value)
	}
}

// StartTransaction to register on NR
func StartTransaction(name string, w http.ResponseWriter, r *http.Request) *newrelic.Transaction {
	txn := agent.StartTransaction(name)
	txn.SetWebRequestHTTP(r)
	txn.SetWebResponse(w)
	return txn
}

func getURL(method, target string) *url.URL {
	var host string
	// target can be anything from
	// https://github.com/grpc/grpc/blob/master/doc/naming.md
	// see https://godoc.org/google.golang.org/grpc#DialContext
	if strings.HasPrefix(target, "unix:") {
		host = "localhost"
	} else {
		host = strings.TrimPrefix(target, "dns:///")
	}
	return &url.URL{
		Scheme: "grpc",
		Host:   host,
		Path:   method,
	}
}

func grpcStartTransaction(ctx context.Context, app *newrelic.Application, fullMethod, name string) *newrelic.Transaction {
	method := strings.TrimPrefix(fullMethod, "/")

	var hdrs http.Header
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		hdrs = make(http.Header, len(md))
		for k, vs := range md {
			for _, v := range vs {
				hdrs.Add(k, v)
			}
		}
	}

	target := hdrs.Get(":authority")
	url := getURL(method, target)

	webReq := newrelic.WebRequest{
		Header:    hdrs,
		URL:       url,
		Method:    method,
		Transport: newrelic.TransportHTTP,
	}
	txn := app.StartTransaction(name)
	txn.SetWebRequest(webReq)

	return txn
}

// NoticeError emits an notice which can traced on new-relic console
func NoticeError(ctx context.Context, err error) {
	txn := newrelic.FromContext(ctx)
	if txn == nil {
		txn := StartNonWebTransaction("NilTxn/NoticeError")
		txn.NoticeError(Wrap(err))
		return
	}

	txn.NoticeError(Wrap(err))
}

// NoticeExpectedError records an error that was expected to occur. Errors recorded with this
// method will not trigger any error alerts or count towards your error metrics.
func NoticeExpectedError(ctx context.Context, err error) {
	txn := newrelic.FromContext(ctx)
	if txn == nil {
		txn := StartNonWebTransaction("NilTxn/NoticeExpectedError")
		txn.NoticeExpectedError(Wrap(err))
		return
	}

	txn.NoticeExpectedError(Wrap(err))
}

// RecordEvent records events which can queried from new-relic insights console
func RecordEvent(event string, properties map[string]interface{}) {
	agent.RecordCustomEvent(event, properties)
}

// StartCustomDataSegment starts a new custom data store segment (for databases which are not listed in newrelic datastore.go)
// eg: FirebaseRealtimeDB
func StartCustomDataSegment(ctx context.Context, product string, operation string) *newrelic.DatastoreSegment {
	productName := newrelic.DatastoreProduct(product)
	txn := newrelic.FromContext(ctx)
	return &newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		Product:   productName,
		Operation: operation,
	}
}

// NewContext returns a new Context that carries the provided transaction
func NewContext(ctx context.Context, txn *newrelic.Transaction) context.Context {
	return newrelic.NewContext(ctx, txn)
}

// FromContext returns transaction from context
func FromContext(ctx context.Context) *newrelic.Transaction {
	return newrelic.FromContext(ctx)
}

// GetContext returns new context with given transaction
func GetContext(txn *newrelic.Transaction) context.Context {
	return newrelic.NewContext(context.Background(), txn)
}
