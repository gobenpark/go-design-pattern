package with_command_facade

type originator struct {
	Command Command
}

func (o *originator) NewMemento() Memento {
	return Memento{o.Command}
}

func (o *originator) ExtractAndStoreCommand(m Memento) {
	o.Command = m.memento
}
