package stub

import "easy-copy/flags"

type StubImpl struct {
	v flags.Verbose
}

func New() *StubImpl {
	return &StubImpl{v: flags.VerbDebug}
}

func (s *StubImpl) ParseLine(line string) {}

func (s *StubImpl) LoadConfig(args []string) error {
	return nil
}

func (s *StubImpl) Sources() []string {
	return nil
}

func (s *StubImpl) Target() string {
	return ""
}

func (s *StubImpl) Mode() flags.Mode {
	return flags.ModeCopy
}

func (s *StubImpl) Verbosity() flags.Verbose {
	return s.v
}

func (s *StubImpl) SetVerbosity(v flags.Verbose) {
	s.v = v
}

func (s *StubImpl) OnConflict() flags.Conflict {
	return flags.ConflictAsk
}

func (s *StubImpl) OnSymlink() flags.Symlink {
	return flags.SymlinkIgnore
}

func (s *StubImpl) DoLSColors() bool {
	return true
}

func (s *StubImpl) Dryrun() bool {
	return true
}

func (s *StubImpl) Parallel() bool {
	return true
}

func (s *StubImpl) SetOnConflict(c flags.Conflict) {}
