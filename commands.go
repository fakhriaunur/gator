package main

import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCmds map[string]func(*state, command) error
}

func NewCommands() *commands {
	handlers := make(map[string]func(*state, command) error)

	return &commands{
		registeredCmds: handlers,
	}
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.registeredCmds[cmd.name]
	if !ok {
		return errors.New("command is not exist")
	}

	if err := handler(s, cmd); err != nil {
		return err
	}

	return nil
}
