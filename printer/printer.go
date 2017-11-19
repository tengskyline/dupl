package printer

import "github.com/mibk/dupl/syntax"

type FileReader interface {
	ReadFile(filename string) ([]byte, error)
}

type Printer interface {
	PrintHeader() error
	PrintClones(dups [][]*syntax.Node) error
	PrintFooter() error
}
