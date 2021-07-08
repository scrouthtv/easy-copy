package flags

type Settings interface {
	ParseLine()
	LoadConfig() error

	Sources() []string
	Target() string

	Mode() Mode
	Verbosity() Verbose
	OnConflict() Conflict
	OnSymlink() Symlink
	DoLSColors() bool
	Dryrun() bool
}

var Current Settings

type Mode uint8

const (
	ModeCopy Mode = iota

	ModeMove
)

type Verbose uint8

const (
	// VerbQuiet indicates that no output should be written at all.
	VerbQuiet Verbose = iota

	// VerbCrit indicates that only critical messages should be written.
	VerbCrit

	// VerbNotice indicates that critical and helpful messages should be written.
	VerbNotice

	// VerbInfo indicates that additional info should be written.
	VerbInfo

	// VerbDebug should only be used for debugging.
	VerbDebug
)

type Conflict uint8

const (
	// ConflictSkip indicates to silently skip any files that already exist.
	ConflictSkip Conflict = iota

	// ConflictAsk indicates to ask the user what to do with conflicting files.
	ConflictAsk

	// ConflictOverwrite indicates to always overwrite existing files.
	ConflictOverwrite
)

type Symlink uint8

const (
	SymlinkIgnore Symlink = iota
	SymlinkLink
	SymlinkDeref
)
