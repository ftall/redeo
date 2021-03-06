package info

import (
	"strconv"
	"sync/atomic"
)

// Value must be exportable as a string
type Value interface {
	String() string
}

// StringValue is the simplest value type
type StringValue string

func (v StringValue) String() string { return string(v) }

// IntValue converts a static integer into a value
func IntValue(n int) Value { return StringValue(strconv.Itoa(n)) }

// Callback function
type Callback func() string

func (c Callback) String() string { return c() }

// Counter is a numeric counter value
type Counter struct{ v int64 }

// NewCounter return a Counter
func NewCounter() *Counter { return &Counter{} }

// Inc atomically increase the Counter by delta
func (c *Counter) Inc(delta int64) int64 { return atomic.AddInt64(&c.v, delta) }

// Set atomically set the value of Counter to v
func (c *Counter) Set(v int64) { atomic.StoreInt64(&c.v, v) }

// Value atomically get the value stored in Counter
func (c *Counter) Value() int64 { return atomic.LoadInt64(&c.v) }

// String return the value of Counter as string
func (c *Counter) String() string { return strconv.FormatInt(c.Value(), 10) }
