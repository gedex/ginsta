package commands

import (
	"fmt"
	"github.com/gedex/ginsta/utils"
	"os"
)

var (
	cmdMediaComments = &Command{
		Callback: mediaComments,
		Usage:    "media_comments [MEDIA_ID]",
		Short:    "Get the list of comments for this MEDIA_ID.",
		Long:     `Get the list of comments for this MEDIA_ID.`,
	}
	cmdAddComment = &Command{
		Callback: addComment,
		Usage:    "add_comment [MEDIA_ID] [TEXT]",
		Short:    "Add comment for MEDIA_ID.",
		Long:     `Add comment for MEDIA_ID.`,
	}
	cmdDelComment = &Command{
		Callback: delComment,
		Usage:    "del_comment [MEDIA_ID] [COMMENT_ID]",
		Short:    "Delete comment COMMENT_ID from MEDIA_ID.",
		Long:     `Delete comment COMMENT_ID from MEDIA_ID.`,
	}
)

func mediaComments(r *Runner, c *Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", c.Usage)
		os.Exit(1)
	}
	mediaId := args[0]

	inst := r.Client.Instagram
	comments, err := inst.Comments.MediaComments(mediaId)
	utils.Check(err)

	// TODO: Supply next
	utils.CommentSlicePrinter(comments, nil, "-max-id")
	os.Exit(0)
}

func addComment(r *Runner, c *Command, args []string) {
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", c.Usage)
		os.Exit(1)
	}
	mediaId, text := args[0], args[1]

	inst := r.Client.Instagram
	err := inst.Comments.Add(mediaId, []string{text})
	utils.Check(err)

	os.Exit(0)
}

func delComment(r *Runner, c *Command, args []string) {
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", c.Usage)
		os.Exit(1)
	}
	mediaId, cid := args[0], args[1]

	inst := r.Client.Instagram
	err := inst.Comments.Delete(mediaId, cid)
	utils.Check(err)

	os.Exit(0)
}
