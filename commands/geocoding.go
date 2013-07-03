package commands

import (
	"fmt"
	"github.com/gedex/ginsta/utils"
	"os"
	"strconv"
)

var (
	cmdGeocoding = &Command{
		Callback: runGeocoding,
		Usage:    "geocoding [ADDRESS]",
		Short:    "Converting addresses into latitude and longitude representation.",
		Long:     `Converting addresses (like '1600 Amphitheatre Parkway, Mountain View, CA') into geographic coordinates (like latitude 37.423021 and longitude -122.083739).`,
	}
	cmdReverseGeocoding = &Command{
		Callback: runReverseGeocoding,
		Usage:    "reverse_geocoding [LAT] [LNG]",
		Short:    "Converting geographic coordinates into a human-readable address.",
		Long:     `Converting geographic coordinates (like '37.423021, -122.083739') into human-readable address (like '1600 Amphitheatre Parkway, Mountain View, CA').`,
	}
)

func runGeocoding(r *Runner, cmd *Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %v\n", cmd.Usage)
		os.Exit(2)
	}
	location := args[0]
	results, err := utils.Geocoding(location, 0, 0)
	utils.Check(err)

	for _, result := range results.Results {
		fmt.Println(result.FormattedAddress, ":", result.Geometry.Location.String())
	}
	os.Exit(0)
}

func runReverseGeocoding(r *Runner, cmd *Command, args []string) {
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %v\n", cmd.Usage)
		os.Exit(2)
	}
	lat, err := strconv.ParseFloat(args[0], 64)
	utils.Check(err)
	lng, err := strconv.ParseFloat(args[1], 64)
	utils.Check(err)
	results, err := utils.Geocoding("", lat, lng)
	utils.Check(err)

	for _, result := range results.Results {
		fmt.Println(result.FormattedAddress)
	}
	os.Exit(0)
}
