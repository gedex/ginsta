package commands

import (
	"fmt"
	"github.com/gedex/ginsta/utils"
	"github.com/gedex/go-instagram/instagram"
	"os"
)

var (
	cmdTagInfo = &Command{
		Callback: tagInfo,
		Usage:    "tag_info [TAG_NAME]",
		Short:    "Get information about a tag [TAG_NAME].",
		Long:     `Get information about a tag [TAG_NAME].`,
	}
	cmdRecentMediaByTag = &Command{
		Callback: recentMediaByTag,
		Usage:    "recent_media_by_tag [-min_id MIN_ID] [-max_id MAX_ID] [TAG_NAME]",
		Short:    "Get the list of media tagged with [TAG_NAME].",
		Long:     `Get the list of media tagged with [TAG_NAME].`,
	}
	cmdSearchTag = &Command{
		Callback: searchTag,
		Usage:    "search_tag [QUERY]",
		Short:    "Search for tag.",
		Long:     `Search for tag by name. This may returns more than one tag.`,
	}

	flagMediaMinID, flagMediaMaxID string
)

func init() {
	cmdRecentMediaByTag.Flag.StringVar(&flagMediaMinID, "min_id", "", "MIN_ID")
	cmdRecentMediaByTag.Flag.StringVar(&flagMediaMaxID, "max_id", "", "MAX_ID")
}

func tagInfo(r *Runner, c *Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", c.Usage)
		os.Exit(1)
	}
	tagName := args[0]

	inst := r.Client.Instagram
	tag, err := inst.Tags.Get(tagName)
	utils.Check(err)

	utils.TagPrinter(tag)
	os.Exit(0)
}

func recentMediaByTag(r *Runner, c *Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", c.Usage)
		os.Exit(1)
	}
	tagName := args[0]

	inst := r.Client.Instagram
	params := new(instagram.Parameters)
	if flagMediaMinID != "" {
		params.MinID = flagMediaMinID
	}
	if flagMediaMaxID != "" {
		params.MaxID = flagMediaMaxID
	}

	media, next, err := inst.Tags.RecentMedia(tagName, params)
	utils.Check(err)

	utils.MediaSlicePrinter(media, next, "-max_id")
	os.Exit(0)
}

func searchTag(r *Runner, c *Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", c.Usage)
		os.Exit(1)
	}
	tagName := args[0]

	inst := r.Client.Instagram
	tags, next, err := inst.Tags.Search(tagName)
	utils.Check(err)

	utils.TagSlicePrinter(tags, next, "-max-id")
	os.Exit(0)
}
