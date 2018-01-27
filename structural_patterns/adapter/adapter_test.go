package adapter

import (
	"testing"
)

func TestOldPrinterAdapter(t *testing.T) {
	adapter := NewPrinterImpl{
		OldPrinter: &OldPrinterImpl{},
		Msg:        "new Adapter msg",
	}
	msg := adapter.PrintSaved()
	if msg != "old printer: new Adapter msg\n" {
		t.Fatalf("adapter with old interface impl got wrong Msg")
	}
}

func TestNewPrinterAdapter(t *testing.T) {
	adapter := NewPrinterImpl{
		OldPrinter: nil,
		Msg:        "new Adapter msg",
	}

	msg := adapter.PrintSaved()
	if msg != "new Adapter msg" {
		t.Fatalf("new interface adapter got wrong return msg: %s\n", msg)
	}

}
