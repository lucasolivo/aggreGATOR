package cli

import (
	"github.com/lucasolivo/aggreGATOR/internal/config"
	"errors"
	"fmt"
)

type State struct {
	CfgPoint *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	CommandNames map[string]func(*State, Command) error
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("username is required")
	}
	username := cmd.Args[0]
	err := s.CfgPoint.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Println("user has been set")
	return nil
}

func (c *Commands) Run(s *State, cmd Command) error {
    handler, ok := c.CommandNames[cmd.Name]
    if !ok {
        return errors.New("command not available")
    }
    err := handler(s, cmd)
    return err
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
    c.CommandNames[name] = f
}