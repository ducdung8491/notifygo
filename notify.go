package notifygo

// Notifier ...
type Notifier interface {
	// Notify ...
	Notify(msg string) error
}

// NotifyFunc ...
type NotifyFunc func(msg string) error

func (f NotifyFunc) Notify(msg string) error {
	return f(msg)
}

type pipe struct {
	ns []Notifier
}

// Pipe ...
func Pipe(ns ...Notifier) *pipe {
	return &pipe{ns}
}

func (p *pipe) Notify(msg string) error {
	var err error
	for _, n := range p.ns {
		err = n.Notify(msg)
	}
	return err
}

type group struct {
	ncs []Notifier
}

// Group ...
func Group(ns ...Notifier) *group {
	return &group{ns}
}

func (g *group) Notify(msg string) error {
	var err error
	return err
}
