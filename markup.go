//go:generate go tool peg -inline -switch grammar.peg

package markup

func Parse(s string) (metas map[string]string, err error) {
	p := &Parser{
		Buffer: s,
		Metas: make(map[string]string),
	}
	p.Init()
	if err := p.Parse(); err != nil {
		return metas, err
	}
	p.Execute()
	//p.PrintSyntaxTree()
	return p.Metas, nil
}
