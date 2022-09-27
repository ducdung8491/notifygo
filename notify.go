package notifygo

type Notifier interface {
	Notify(msg string) error
}

type NotifyFunc func(msg string) error

func (f NotifyFunc) Notify(msg string) error {
	return f(msg)
}

type pipe struct {
	ns []Notifier
}

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

func Group(ns ...Notifier) *group {
	return &group{ns}
}

func (g *group) Notify(msg string) error {
	var err error
	return err
}
