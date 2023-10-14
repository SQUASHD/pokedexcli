package main

import (
	"errors"
	"fmt"
)

func commandMap(c *config, args ...string) error {
	locationRes, err := c.client.ListLocations(c.nextLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = locationRes.Next
	c.previousLocationsURL = locationRes.Previous

	for _, locArea := range locationRes.Results {
		fmt.Printf("%s\n", locArea.Name)
	}

	return nil
}

func commandMapb(c *config, args ...string) error {
	if c.previousLocationsURL == nil {
		return errors.New("no previous locations")
	}
	locationRes, err := c.client.ListLocations(c.previousLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = locationRes.Next
	c.previousLocationsURL = locationRes.Previous

	for _, locArea := range locationRes.Results {
		fmt.Printf("%s\n", locArea.Name)
	}

	return nil
}
