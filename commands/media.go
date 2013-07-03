package commands

import (
	"fmt"
	"github.com/gedex/ginsta/utils"
	"github.com/gedex/go-instagram/instagram"
	"os"
)

var (
	cmdMediaInfo = &Command{
		Callback: mediaInfo,
		Usage:    "media_info [USER_ID]",
		Short:    "Get information about a media object.",
		Long:     `Get information about a media object.`,
	}
	cmdSearchMedia = &Command{
		Callback: searchMedia,
		Usage:    "search_media [-lat LAT] [-lng LNG] [-d DISTANCE] [-min_ts MIN_TIMESTAMP] [-max_ts MAX_TIMESTAMP]",
		Short:    "Search for media in a given area.",
		Long: `Search for media in a given area. -lat and -lnt must be provided.
In conjunction with TIMESTAMP, the default time span is set to 5 days.
The time span must not exceed 7 days. Default DISTANCE is 1km (-d 1000),
max distance is 5km.`,
	}
	cmdPopularMedia = &Command{
		Callback: popularMedia,
		Usage:    "popular_media",
		Short:    "Get the list of what media is most popular at the moment.",
		Long:     `Get the list of what media is most popular at the moment.`,
	}

	flagMediaLat, flagMediaLng, flagMediaDistance float64
	flagMediaMinTS, flagMediaMaxTS                int64
)

func init() {

	cmdSearchMedia.Flag.Float64Var(&flagMediaLat, "lat", 0, "LAT")
	cmdSearchMedia.Flag.Float64Var(&flagMediaLng, "lng", 0, "LNG")
	cmdSearchMedia.Flag.Float64Var(&flagMediaDistance, "d", 0, "DISTANCE")

	cmdSearchMedia.Flag.Int64Var(&flagMediaMinTS, "min_ts", 0, "MIN_TS")
	cmdSearchMedia.Flag.Int64Var(&flagMediaMaxTS, "max_ts", 0, "MAX_TS")
}

func mediaInfo(r *Runner, c *Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", c.Usage)
		os.Exit(1)
	}
	mediaId := args[0]

	inst := r.Client.Instagram
	m, err := inst.Media.Get(mediaId)
	utils.Check(err)

	utils.MediaPrinter(*m)
	os.Exit(0)
}

func searchMedia(r *Runner, c *Command, args []string) {
	inst := r.Client.Instagram
	params := new(instagram.Parameters)

	if flagMediaLat != 0 {
		params.Lat = flagMediaLat
	}
	if flagMediaLng != 0 {
		params.Lng = flagMediaLng
	}
	if flagMediaDistance != 0 {
		params.Distance = flagMediaDistance
	}
	if flagMediaMinTS != 0 {
		params.MinTimestamp = flagMediaMinTS
	}
	if flagMediaMaxTS != 0 {
		params.MaxTimestamp = flagMediaMaxTS
	}

	media, next, err := inst.Media.Search(params)
	utils.Check(err)

	utils.MediaSlicePrinter(media, next, "-max-id")
	os.Exit(0)
}

func popularMedia(r *Runner, c *Command, args []string) {
	inst := r.Client.Instagram
	media, next, err := inst.Media.Popular()
	utils.Check(err)

	utils.MediaSlicePrinter(media, next, "-max-id")
	os.Exit(0)
}
