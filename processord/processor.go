package processord

type Exector interface {
	Exec()
	ParsePage(pageNum int)
}

type Processor struct {
	Exec Exector
}

func NewProcessor(exec Exector) *Processor {
	return &Processor{
		Exec: exec,
	}
}

func (p *Processor) StartExec() {
	p.Exec.Exec()
}
