package stats

//go:generate errorgen

import (
	"net"
	"v2ray.com/core/features"
)

// Counter is the interface for stats counters.
//
// v2ray:api:stable
type Counter interface {
	// Value is the current value of the counter.
	Value() int64
	// Set sets a new value to the counter, and returns the previous one.
	Set(int64) int64
	// Add adds a value to the current counter value, and returns the previous value.
	Add(int64) int64
}

type IPStorager interface {
	Add(net.IP) bool
	Empty()
	Remove(net.IP) bool
	All() []net.IP
}

// Manager is the interface for stats manager.
//
// v2ray:api:stable
type Manager interface {
	features.Feature

	// RegisterCounter registers a new counter to the manager. The identifier string must not be emtpy, and unique among other counters.
	RegisterCounter(string) (Counter, error)
	// GetCounter returns a counter by its identifier.
	GetCounter(string) Counter

	RegisterIPStorager(string) (IPStorager, error)
	GetIPStorager(string) IPStorager

}

// GetOrRegisterCounter tries to get the StatCounter first. If not exist, it then tries to create a new counter.
func GetOrRegisterCounter(m Manager, name string) (Counter, error) {
	counter := m.GetCounter(name)
	if counter != nil {
		return counter, nil
	}

	return m.RegisterCounter(name)
}

func GetOrRegisterIPStorager(m Manager, name string) (IPStorager, error) {
	ipStorager := m.GetIPStorager(name)
	if ipStorager != nil {
		return ipStorager, nil
	}

	return m.RegisterIPStorager(name)
}

// ManagerType returns the type of Manager interface. Can be used to implement common.HasType.
//
// v2ray:api:stable
func ManagerType() interface{} {
	return (*Manager)(nil)
}

// NoopManager is an implementation of Manager, which doesn't has actual functionalities.
type NoopManager struct{}

// Type implements common.HasType.
func (NoopManager) Type() interface{} {
	return ManagerType()
}

// RegisterCounter implements Manager.
func (NoopManager) RegisterCounter(string) (Counter, error) {
	return nil, newError("not implemented")
}

// GetCounter implements Manager.
func (NoopManager) GetCounter(string) Counter {
	return nil
}

func (NoopManager) RegisterIPStorager(string) (IPStorager, error) {
	return nil, newError("not implemented")
}

func (NoopManager) GetIPStorager(string) IPStorager {
	return nil
}

// Start implements common.Runnable.
func (NoopManager) Start() error { return nil }

// Close implements common.Closable.
func (NoopManager) Close() error { return nil }
