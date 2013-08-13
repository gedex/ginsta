package commands

import (
	"fmt"
	"github.com/gedex/ginsta/utils"
	"os"
)

var (
	cmdMediaLikes = &Command{
		Callback: mediaLikes,
		Usage:    "media_likes [MEDIA_ID]",
		Short:    "Get the list of users who like this MEDIA_ID.",
		Long:     `Get the list of users who like this MEDIA_ID.`,
	}
	cmdAddLike = &Command{
		Callback: addLike,
		Usage:    "add_like [MEDIA_ID]",
		Short:    "Like this MEDIA_ID.",
		Long:     `Like this MEDIA_ID.`,
	}
	cmdDelLike = &Command{
		Callback: delLike,
		Usage:    "del_like [MEDIA_ID]",
		Short:    "Unlike this MEDIA_ID.",
		Long:     `Unlike this MEDIA_ID.`,
	}
)

func mediaLikes(r *Runner, c *Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", c.Usage)
		os.Exit(1)
	}
	mediaId := args[0]

	inst := r.Client.Instagram
	users, err := inst.Likes.MediaLikes(mediaId)
	utils.Check(err)

	utils.UserSlicePrinter(users, nil, "-max-id")
	os.Exit(0)
}

func addLike(r *Runner, c *Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", c.Usage)
		os.Exit(1)
	}
	mediaId := args[0]

	inst := r.Client.Instagram
	err := inst.Likes.Like(mediaId)
	utils.Check(err)

	os.Exit(0)
}

func delLike(r *Runner, c *Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", c.Usage)
		os.Exit(1)
	}
	mediaId := args[0]

	inst := r.Client.Instagram
	err := inst.Likes.Unlike(mediaId)
	utils.Check(err)

	os.Exit(0)
}
