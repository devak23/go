package main

import . "fmt"

type (
	ServerOptions func(options) options
	TransportType int
)

type options struct {
	MaxConnections int
	TransportType  TransportType
	Name           string
}

const (
	UDP TransportType = iota
	TCP
)

type Server struct {
	options

	// The main reason for using the options struct and embedding it inside Server is to use this as a configuration for
	// our server that we want users to provide. We don’t want users to provide data that is not contained in this
	// struct, such as the isAlive flag. This clearly separates concerns, and it will allow us to build the next
	// higher-order functions and partial application layers on top of it.

	isAlive bool
}

// We have embedded options without declaring a new name for the field. This is achieved by simply writing the type of
// struct that you want to embed. When doing so, the Server struct will contain all the fields that the options struct
// has. It’s a way to model object composition.

func MaxConnections(n int) ServerOptions {
	return func(o options) options {
		o.MaxConnections = n
		return o
	}
}

func ServerName(name string) ServerOptions {
	return func(o options) options {
		o.Name = name
		return o
	}
}

func Transport(t TransportType) ServerOptions {
	return func(o options) options {
		o.TransportType = t
		return o
	}
}

func NewServer(os ...ServerOptions) Server {
	opts := options{}
	for _, option := range os {
		opts = option(opts)
	}

	return Server{
		options: opts,
		isAlive: true,
	}
}

func main() {
	server := NewServer(MaxConnections(10), ServerName("MyFirstServer"), Transport(UDP))
	Printf("%+v\n", server)
}
