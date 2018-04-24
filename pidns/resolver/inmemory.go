package resolver

import (
	"net"
	"sync"
)

type InMemoryResolver struct {
	hosts map[string]string

	ThreadProtection bool
	mutex            sync.RWMutex
}

func NewInMemoryResolver() *InMemoryResolver {
	return &InMemoryResolver{
		hosts:            make(map[string]string),
		ThreadProtection: true,
	}
}

func (r *InMemoryResolver) Hosts() map[string]string {

	if r.ThreadProtection {
		r.mutex.Lock()
		defer r.mutex.Unlock()
	}

	return r.hosts
}

func (r *InMemoryResolver) AddRule(name, host string) {

	if r.ThreadProtection {
		r.mutex.Lock()
		defer r.mutex.Unlock()
	}

	r.hosts[name] = host
}

func (r *InMemoryResolver) Resolve(name string) *[]string {

	name = name[:len(name)-1] // . at end

	if r.ThreadProtection {
		r.mutex.RLock()
		defer r.mutex.RUnlock()
	}

	if host, ok := r.hosts[name]; ok {
		return &[]string{host}
	}

	h, err := net.LookupHost(name)
	if err != nil {
		return nil
	}
	return &h

}
