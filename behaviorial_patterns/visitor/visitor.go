package visitor

import (
	"fmt"
	"io"
	"os"
)

type Visitor interface {
	VisitA(v *MessageA)
	VisitB(v *MessageB)
}

type MessageA struct {
	Content string
	Writer  io.Writer
}

func (m *MessageA) Print() {
	if m.Writer == nil {
		m.Writer = os.Stdout
	}
	fmt.Fprintf(m.Writer, "%s", m.Content)
}

func (m *MessageA) Accept(v Visitor) {
	v.VisitA(m)
}

type MessageAVisitor struct {
}

func (m *MessageAVisitor) VisitA(v *MessageA) {
	v.Content = fmt.Sprintf("%s %s", "visitor A", v.Content)
}
func (m *MessageAVisitor) VisitB(v *MessageB) {
	v.Content = fmt.Sprintf("%s %s", "viistor B", v.Content)
}

type MessageB struct {
	Content string
	Writer  io.Writer
}

func (m *MessageB) Print() {
	if m.Writer == nil {
		m.Writer = os.Stdout
	}
	fmt.Fprintf(m.Writer, "%s", m.Content)
}

func (m *MessageB) Accept(v Visitor) {
	v.VisitB(m)
}

type MessageBVisitor struct{}

func (m *MessageBVisitor) VisitA(v *MessageA) {
	v.Content = fmt.Sprintf("%s %s", "visitor A", v.Content)
}
func (m *MessageBVisitor) VisitB(v *MessageB) {
	v.Content = fmt.Sprintf("%s %s", "visitor B", v.Content)
}
