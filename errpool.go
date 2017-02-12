package errpool

import "sync"

// Pool is a leightweight structure for synchronizing and reporting
// the errors of mutliple concurrent units of work.
type Pool struct {
	wg sync.WaitGroup

	errMux sync.Mutex
	errs   ErrList
}

// NewPool instantiates a new empty Pool.
func NewPool() *Pool {
	return &Pool{
		wg:     sync.WaitGroup{},
		errMux: sync.Mutex{},
		errs:   ErrList{},
	}
}

// Add adds a unit of work to Pool's underlying sync.Waitgroup.
// This value may be either positive or negative.
// See the sync.Waitgroup docs for more clarity.
func (p *Pool) Add(delta int) {
	p.wg.Add(delta)
}

// Done is wrapper around sync.Waitgroup's Done method.
func (p *Pool) Done() {
	p.wg.Done()
}

// Wait blocks until all of the Pool's tasks have completed and returns
// a singular error.
func (p *Pool) Wait() error {
	p.wg.Wait()

	p.errMux.Lock()
	defer p.errMux.Unlock()
	errs := p.errs
	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		return errs
	}
}

// Error adds an error to the Pool's errorlist.
func (p *Pool) Error(err error) {
	if err == nil {
		return
	}

	p.errMux.Lock()
	p.errs = append(p.errs, err)
	p.errMux.Unlock()
}
