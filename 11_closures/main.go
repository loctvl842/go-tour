package main

import "fmt"

type service struct {
	host    string
	port    int
	timeout int
}

type Option func(*service)

func WithHost(host string) Option { return func(s *service) { s.host = host } }

func WithPort(port int) Option { return func(s *service) { s.port = port } }

func WithTimeout(timeout int) Option { return func(s *service) { s.timeout = timeout } }

func New(opts ...Option) *service {
	s := service{
		host:    "localhost",
		port:    8080,
		timeout: 1000,
	}

	for _, opt := range opts {
		opt(&s)
	}
	return &s
}

func main() {
	s := New(
		WithHost("localhost"),
		WithPort(3000),
	)
  fmt.Println(s)
	fmt.Println((*s).host)
  fmt.Println(s.host)
  fmt.Println(s.host == (*s).host)
	fmt.Println(s.port)
	fmt.Println(s.timeout)
}
