package bridge

import (
	"errors"
	"fmt"
	"io"
	"log"
)

// two printer interface
// 1: print message to console directly.(normal printer)
// 2: print message to io.Writer with prefix string: Message from Packet (packet printer)

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterAPI1 struct { // default print message to console
}

func (p *PrinterAPI1) PrintMessage(msg string) error {
	// return errors.New("not implemented yet")
	log.Println(msg)
	return nil
}

// second printer impl works with io.Writer

type PrinterAPI2 struct { // delegate print job to it's writer
	Writer io.Writer
}

func (p *PrinterAPI2) PrintMessage(msg string) error { // call Writer.Write
	if p.Writer == nil {
		return errors.New("not implemented yet")
	}
	fmt.Fprintf(p.Writer, msg) // let the inner Writer do the job
	return nil
}

// writer implement
type TestWriter struct {
	Msg string
}

func (w *TestWriter) Write(p []byte) (int, error) {
	n := len(p)
	if n > 0 {
		w.Msg = string(p)
		return n, nil
	}
	err := errors.New("content writer received was empty")
	return 0, err
}

// abstract Pritner interface
// bridge between: different Printer implementtation and PrinterAbstraction obj
// you can change outside abstract obj and inner Printer implementation as you like.
type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (p *NormalPrinter) Print() error {
	// return errors.New("not implmented yet")
	return p.Printer.PrintMessage(p.Msg)
}

type PacketPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (p *PacketPrinter) Print() error {
	// return errors.New("not implemented yet")
	return p.Printer.PrintMessage(fmt.Sprintf("Message from Packet: %s", p.Msg))
}
