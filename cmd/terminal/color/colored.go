package color

type String string

func (c String) Set(styles ...string) {
	for _, style := range styles {
		c = String(style + c.String())
	}
}

func (c String) String() string {
	return string(c)
}
