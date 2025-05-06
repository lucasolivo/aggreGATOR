package main

import (
	"fmt"
	"log"
	"github.com/lucasolivo/aggreGATOR/internal/config"
	"github.com/lucasolivo/aggreGATOR/internal/cli"
	"os"
	"errors"
)

func main() {
    cfg, err := config.Read()
    if err != nil {
        log.Fatalf("Failed to read config: %v", err)
    }
    
    s := &cli.State{
		CfgPoint: &cfg,
	}

	handlers := &cli.Commands{
		CommandNames: make(map[string]func(*cli.State, cli.Command) error),
	}

	handlers.Register("login", cli.HandlerLogin)

	CLArgs := os.Args

	if len(CLArgs) < 2 {
		fmt.Println(errors.New("Not enough arguments provided"))
		os.Exit(1)
	}

	calledCommand := cli.Command{
		Name: CLArgs[1],
		Args: CLArgs[2:],
	}

	err = handlers.Run(s, calledCommand)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


}