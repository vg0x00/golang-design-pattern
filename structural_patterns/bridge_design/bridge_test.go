package bridge

import (
	"testing"
)

func TestGetPrinterAPI1(t *testing.T) {
	normalPrinter := &PrinterAPI1{}
	if err := normalPrinter.PrintMessage("hello"); err != nil {
		t.Fatalf("PrinterAPI1 implement failed: %s\n", err.Error())
	}
}

func TestGetPrinterAPI2(t *testing.T) {
	// packetPrinter := PrinterAPI2{}
	// if err := packetPrinter.PrintMessage("hello"); err != nil {
	// 	expectedErrorMessge := "you need to pass an io.Writer impl"
	// 	if !strings.Contains(err.Error(), expectedErrorMessge) {
	// 		t.Fatalf("without io.Writer impl error message wrong\nActual: %s\nExpect: %s\n", err.Error(), expectedErrorMessge)
	// 	}
	// }

	// pass in an io.Writer impl
	testWriter := TestWriter{}
	packetPrinter := PrinterAPI2{
		Writer: &testWriter,
	}

	expectedMessage := "hello"
	err := packetPrinter.PrintMessage(expectedMessage)
	if err != nil {
		t.Errorf("PrinterAPI2 Impl failed with io.Writer impl\n %s\n", err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("API2 writer got wrong Msg field:\n actual: %s\n expect: %s\n", testWriter.Msg, expectedMessage)
	}
}

func TestNormalPrinter_Print(t *testing.T) {
	expectedMsg := "hello, io.Writer"
	normal := NormalPrinter{
		Msg:     expectedMsg,
		Printer: &PrinterAPI1{},
	}
	err := normal.Print()
	if err != nil {
		t.Errorf("normal printer Print method not implmented yet: %s\n", err.Error())
	}

	testWriter := TestWriter{}
	normal = NormalPrinter{
		Msg: expectedMsg,
		Printer: &PrinterAPI2{
			Writer: &testWriter,
		},
	}

	err = normal.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expectedMsg {
		t.Errorf("the expected msg on io.Writer wrong:\n actual: %s\n expected%s\n",
			testWriter.Msg, expectedMsg)
	}
}

func TestPacketPrinter_Print(t *testing.T) {
	paramMsg := "hello, io.Writer"
	expectedMsg := "Message from Packet: hello, io.Writer"
	packetPrinter := PacketPrinter{
		Msg:     paramMsg,
		Printer: &PrinterAPI1{},
	}

	err := packetPrinter.Print()
	if err != nil {
		t.Error(err.Error())
	}

	testWriter := TestWriter{}
	packetPrinter = PacketPrinter{
		Msg: paramMsg,
		Printer: &PrinterAPI2{
			Writer: &testWriter,
		},
	}

	err = packetPrinter.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expectedMsg {
		t.Errorf("the expected message for io.Writer wrong: \n actual: %s\n expect: %s\n",
			testWriter.Msg, expectedMsg)
	}
}
