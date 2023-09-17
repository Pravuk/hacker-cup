package common

type CaseReader struct {
	Count        int
	LinesPerCase int
}

type runCase func()
type input func(lines []string)

func NewCaseReader() *CaseReader {
	return nil
}
