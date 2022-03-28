// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/metric/nonrecording"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.MeterConfig{}
	_ = syncint64.Counter(nil)
	_ = nonrecording.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	cfg       config
	requests  syncint64.Counter
	errors    syncint64.Counter
	duration  syncint64.Histogram
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...Option) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	c := &Client{
		cfg:       newConfig(opts...),
		serverURL: u,
	}
	if c.requests, err = c.cfg.Meter.SyncInt64().Counter(otelogen.ClientRequestCount); err != nil {
		return nil, err
	}
	if c.errors, err = c.cfg.Meter.SyncInt64().Counter(otelogen.ClientErrorsCount); err != nil {
		return nil, err
	}
	if c.duration, err = c.cfg.Meter.SyncInt64().Histogram(otelogen.ClientDuration); err != nil {
		return nil, err
	}
	return c, nil
}

// AnyContentTypeBinaryStringSchema invokes anyContentTypeBinaryStringSchema operation.
//
// GET /anyContentTypeBinaryStringSchema
func (c *Client) AnyContentTypeBinaryStringSchema(ctx context.Context) (res AnyContentTypeBinaryStringSchemaOK, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("anyContentTypeBinaryStringSchema"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "AnyContentTypeBinaryStringSchema",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/anyContentTypeBinaryStringSchema"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeAnyContentTypeBinaryStringSchemaResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AnyContentTypeBinaryStringSchemaDefault invokes anyContentTypeBinaryStringSchemaDefault operation.
//
// GET /anyContentTypeBinaryStringSchemaDefault
func (c *Client) AnyContentTypeBinaryStringSchemaDefault(ctx context.Context) (res AnyContentTypeBinaryStringSchemaDefaultDefStatusCode, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("anyContentTypeBinaryStringSchemaDefault"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "AnyContentTypeBinaryStringSchemaDefault",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/anyContentTypeBinaryStringSchemaDefault"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeAnyContentTypeBinaryStringSchemaDefaultResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MultipleGenericResponses invokes multipleGenericResponses operation.
//
// GET /multipleGenericResponses
func (c *Client) MultipleGenericResponses(ctx context.Context) (res MultipleGenericResponsesRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("multipleGenericResponses"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "MultipleGenericResponses",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/multipleGenericResponses"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMultipleGenericResponsesResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OctetStreamBinaryStringSchema invokes octetStreamBinaryStringSchema operation.
//
// GET /octetStreamBinaryStringSchema
func (c *Client) OctetStreamBinaryStringSchema(ctx context.Context) (res OctetStreamBinaryStringSchemaOK, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("octetStreamBinaryStringSchema"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "OctetStreamBinaryStringSchema",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/octetStreamBinaryStringSchema"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeOctetStreamBinaryStringSchemaResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OctetStreamEmptySchema invokes octetStreamEmptySchema operation.
//
// GET /octetStreamEmptySchema
func (c *Client) OctetStreamEmptySchema(ctx context.Context) (res OctetStreamEmptySchemaOK, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("octetStreamEmptySchema"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "OctetStreamEmptySchema",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/octetStreamEmptySchema"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeOctetStreamEmptySchemaResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}