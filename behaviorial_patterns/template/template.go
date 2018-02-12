package template

import "strings"

// interface or protocol
type Messager interface {
	Message() string
}

// first interface implement
type Template struct {
}

func (s *Template) first() string {
	return "first"
}

func (s *Template) third() string {
	return "third"
}

func (s *Template) ExecuteProcess(m Messager) string {
	return strings.Join([]string{s.first(), m.Message(), s.third()}, " ")
}

// second interface implement
type AnonymousTemplate struct {
}

func (s *AnonymousTemplate) first() string {
	return "first"
}

func (s *AnonymousTemplate) third() string {
	return "third"
}

func (s *AnonymousTemplate) ExecuteProcess(f func() string) string {
	return strings.Join([]string{s.first(), f(), s.third()}, " ")
}
