package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if err := initializeTaskFiles(); err != nil {
		log.Fatal(err)
	}

	flag.Parse()
	if err := Validate(flag.Args()); err != nil {
		log.Fatal(err)
	}
	if err := Parse(flag.Args()); err != nil {
		log.Fatal(err)
	}

}

func initializeTaskFiles() error {
	if _, err := os.Stat(tmpDBPath); err == nil {
		if err := os.Remove(tmpDBPath); err != nil {
			return fmt.Errorf("error removing temporary tasks file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("error checking temporary tasks file: %w", err)
	}

	if _, err := os.Stat(dbPath); err == nil {
		return nil
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("error checking tasks file: %w", err)
	}

	if err := os.WriteFile(dbPath, []byte("[]"), 0644); err != nil {
		return fmt.Errorf("error creating tasks file: %w", err)
	}
	return nil
}
