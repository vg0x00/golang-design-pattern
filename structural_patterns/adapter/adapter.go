package adapter

import "fmt"

type OldPrinter interface {
	Print(msg string) string
}

type NewPrinter interface {
	PrintSaved() string
}

type OldPrinterImpl struct {
}

func (p *OldPrinterImpl) Print(msg string) string {
	return fmt.Sprintf("old printer: %s\n", msg)
}

type NewPrinterImpl struct {
	OldPrinter OldPrinter
	Msg        string
}

// route through this method in new interface impl
func (p *NewPrinterImpl) PrintSaved() string {
	returnMsg := ""
	if p.OldPrinter != nil {
		returnMsg = p.OldPrinter.Print(p.Msg)
	} else {
		returnMsg = p.Msg
	}

	return returnMsg
}
