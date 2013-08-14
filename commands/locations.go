package commands

import (
	"fmt"
	"github.com/gedex/ginsta/utils"
	"github.com/gedex/go-instagram/instagram"
	"os"
)

var (
	cmdLocationInfo = &Command{
		Callback: locationInfo,
		Usage:    "location_info [LOCATION_ID]",
		Short:    "Get information about a location object.",
		Long:     `Get information about a location object.`,
	}
	cmdRecentMediaByLocation = &Command{
		Callback: recentMediaByLocation,
		Usage:    "recent_media_by_location [-min_id MIN_ID] [-max_id MAX_ID] [-min_ts MIN_TIMESTAMP] [-max_ts MAX_TIMESTAMP] [LOCATION_ID]",
		Short:    "Get the list recent media in a given LOCATION_ID.",
		Long:     `Get the list recent media in a given LOCATION_ID.`,
	}
	cmdSearchLocation = &Command{
		Callback: searchLocation,
		Usage:    "search_location [-lat LAT] [-lng LNG] [-d DISTANCE]",
		Short:    "Search for location in a given area.",
		Long: `Search for media in a given area. -lat and -lnt must be provided.
In conjunction with TIMESTAMP, the default time span is set to 5 days.
The time span must not exceed 7 days. Default DISTANCE is 1km (-d 1000),
max distance is 5km.`,
	}

	flagLocationLat, flagLocationLng, flagLocationDistance float64
	flagLocationMediaMinID, flagLocationMediaMaxID         string
	flagLocationMediaMinTS, flagLocationMediaMaxTS         int64
)

func init() {
	cmdSearchLocation.Flag.Float64Var(&flagLocationLat, "lat", 0, "LAT")
	cmdSearchLocation.Flag.Float64Var(&flagLocationLng, "lng", 0, "LNG")
	cmdSearchLocation.Flag.Float64Var(&flagLocationDistance, "d", 0, "DISTANCE")

	cmdRecentMediaByLocation.Flag.StringVar(&flagLocationMediaMinID, "min_id", "", "MIN_ID")
	cmdRecentMediaByLocation.Flag.StringVar(&flagLocationMediaMaxID, "max_id", "", "MAX_ID")

	cmdRecentMediaByLocation.Flag.Int64Var(&flagLocationMediaMinTS, "min_ts", 0, "MIN_TS")
	cmdRecentMediaByLocation.Flag.Int64Var(&flagLocationMediaMaxTS, "max_ts", 0, "MAX_TS")
}

func locationInfo(r *Runner, c *Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", c.Usage)
		os.Exit(1)
	}
	locationId := args[0]

	inst := r.Client.Instagram
	location, err := inst.Locations.Get(locationId)
	utils.Check(err)

	utils.LocationPrinter(location)
	os.Exit(0)
}

func recentMediaByLocation(r *Runner, c *Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", c.Usage)
		os.Exit(1)
	}
	locationId := args[0]

	inst := r.Client.Instagram
	params := new(instagram.Parameters)

	if flagLocationMediaMinID != "" {
		params.MinID = flagLocationMediaMinID
	}
	if flagLocationMediaMaxID != "" {
		params.MaxID = flagLocationMediaMaxID
	}
	if flagLocationMediaMinTS != 0 {
		params.MinTimestamp = flagLocationMediaMinTS
	}
	if flagLocationMediaMaxTS != 0 {
		params.MaxTimestamp = flagLocationMediaMaxTS
	}

	media, next, err := inst.Locations.RecentMedia(locationId, params)
	utils.Check(err)

	utils.MediaSlicePrinter(media, next, "-max-id")
	os.Exit(0)
}

func searchLocation(r *Runner, c *Command, args []string) {
	if flagLocationLat == 0 || flagLocationLng == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", c.Usage)
		os.Exit(1)
	}

	inst := r.Client.Instagram
	params := new(instagram.Parameters)
	if flagLocationDistance != 0 {
		params.Distance = flagLocationDistance
	}

	locations, err := inst.Locations.Search(flagLocationLat, flagLocationLng, params)
	utils.Check(err)

	utils.LocationSlicePrinter(locations, nil, "-max-id")
	os.Exit(0)
}
