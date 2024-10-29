package main

import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func NewCommands() *commands {
	handlers := make(map[string]func(*state, command) error)

	return &commands{
		handlers: handlers,
	}
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return errors.New("command is not exist")
	}

	if err := handler(s, cmd); err != nil {
		return err
	}

	return nil
}
