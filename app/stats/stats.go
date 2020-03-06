// +build !confonly

package stats

//go:generate errorgen

import (
	"bytes"
	"context"
	"net"
	"sync"
	"sync/atomic"

	"v2ray.com/core/features/stats"
)

// Counter is an implementation of stats.Counter.
type Counter struct {
	value int64
}

// Value implements stats.Counter.
func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

// Set implements stats.Counter.
func (c *Counter) Set(newValue int64) int64 {
	return atomic.SwapInt64(&c.value, newValue)
}

// Add implements stats.Counter.
func (c *Counter) Add(delta int64) int64 {
	return atomic.AddInt64(&c.value, delta)
}

type IPStorager struct {
	access sync.RWMutex
	ips []net.IP
}

func (s *IPStorager) Add(ip net.IP) bool {
	s.access.Lock()
	defer s.access.Unlock()

	for _, _ip := range s.ips {
		if bytes.Equal(_ip, ip) {
			return false
		}
	}

	s.ips = append(s.ips, ip)

	return true
}

func (s *IPStorager) Empty() {
	s.access.Lock()
	defer s.access.Unlock()

	s.ips = s.ips[:0]
}

func (s *IPStorager) Remove(removeIP net.IP) bool {
	s.access.Lock()
	defer s.access.Unlock()

	for i, ip := range s.ips {
		if bytes.Equal(ip, removeIP) {
			s.ips = append(s.ips[:i], s.ips[i+1:]...)
			return true
		}
	}

	return false
}

func (s *IPStorager) All() []net.IP {
	s.access.RLock()
	defer s.access.RUnlock()

	newIPs := make([]net.IP, len(s.ips))
	copy(newIPs, s.ips)

	return newIPs
}


// Manager is an implementation of stats.Manager.
type Manager struct {
	access   sync.RWMutex
	counters map[string]*Counter
	ipStoragers map[string]*IPStorager
}

func NewManager(ctx context.Context, config *Config) (*Manager, error) {
	m := &Manager{
		counters: make(map[string]*Counter),
	}
	m.ipStoragers = make(map[string]*IPStorager)
	return m, nil
}

func (*Manager) Type() interface{} {
	return stats.ManagerType()
}

func (m *Manager) RegisterCounter(name string) (stats.Counter, error) {
	m.access.Lock()
	defer m.access.Unlock()

	if _, found := m.counters[name]; found {
		return nil, newError("Counter ", name, " already registered.")
	}
	newError("create new counter ", name).AtDebug().WriteToLog()
	c := new(Counter)
	m.counters[name] = c
	return c, nil
}

func (m *Manager) GetCounter(name string) stats.Counter {
	m.access.RLock()
	defer m.access.RUnlock()

	if c, found := m.counters[name]; found {
		return c
	}
	return nil
}

func (m *Manager) RegisterIPStorager(name string) (stats.IPStorager, error) {
	if m.ipStoragers == nil {
		return nil, newError("IPStorager is disabled")
	}

	m.access.Lock()
	defer m.access.Unlock()

	if _, found := m.ipStoragers[name]; found {
		return nil, newError("IPStorager ", name, " already registered.")
	}
	newError("create new IPStorager ", name).AtDebug().WriteToLog()
	s := new(IPStorager)
	m.ipStoragers[name] = s
	return s, nil
}

func (m *Manager) GetIPStorager(name string) stats.IPStorager {
	if m.ipStoragers == nil {
		return nil
	}

	m.access.RLock()
	defer m.access.RUnlock()

	if s, found := m.ipStoragers[name]; found {
		return s
	}
	return nil
}

func (m *Manager) VisitCounters(visitor func(string, stats.Counter) bool) {
	m.access.RLock()
	defer m.access.RUnlock()

	for name, c := range m.counters {
		if !visitor(name, c) {
			break
		}
	}
}

func (m *Manager) VisitIPStoragers(visitor func(string, stats.IPStorager) bool) {
	m.access.RLock()
	defer m.access.RUnlock()

	for name, c := range m.ipStoragers {
		if !visitor(name, c) {
			break
		}
	}
}

// Start implements common.Runnable.
func (m *Manager) Start() error {
	return nil
}

// Close implement common.Closable.
func (m *Manager) Close() error {
	return nil
}
