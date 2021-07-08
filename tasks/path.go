package tasks

import "path/filepath"

type Path struct {
	Base string
	Sub  string
}

func (p *Path) AsAbs() string {
	return filepath.Join(p.Base, p.Sub)
}
