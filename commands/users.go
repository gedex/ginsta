package commands

import (
	"fmt"
	"github.com/gedex/ginsta/utils"
	"github.com/gedex/go-instagram/instagram"
	"os"
)

var (
	cmdUserInfo = &Command{
		Callback: userInfo,
		Usage:    "user_info [USER_ID]",
		Short:    "Get basic information about a user",
		Long:     `Get basic information about a user`,
	}
	cmdUserFeed = &Command{
		Callback: userFeed,
		Usage:    "user_feed [-c COUNT] [-min_id MIN_ID] [-max_id MAX_ID]",
		Short:    "Get authenticated user's feed.",
		Long:     `Get authenticated user's feed. May return a mix of both image and video type.`,
	}
	cmdUserRecentMedia = &Command{
		Callback: userRecentMedia,
		Usage:    "user_recent_media [-u USER_ID] [-c COUNT] [-min_ts MIN_TIMESTAMP] [-max_ts MAX_TIMESTAMP] [-min_id MIN_ID] [-max_id MAX_ID",
		Short:    "Get the most recent media published by a user.",
		Long: `Get the most recent media published by a user. If -u is not specified
then USER_ID is implicitly set to 'self' whichh currently authenticated user.`,
	}
	cmdUserLikedMedia = &Command{
		Callback: userLikedMedia,
		Usage:    "user_liked_media [-c COUNT] [-max_id MAX_LIKED_ID]",
		Short:    "Get the authenticated user's list of media they've liked.",
		Long: `Get the authenticated user's list of media they've liked. The list is ordered
by the order in which the user liked the media."`,
	}
	cmdSearchUser = &Command{
		Callback: searchUser,
		Usage:    "search_user [-c COUNT] [QUERY]",
		Short:    "Search for a user by name",
		Long:     `Search for a user by name. To limit the results, -c can be used.`,
	}
	flagUserCount                uint64
	flagUserMinID, flagUserMaxID string
	flagUserMinTS, flagUserMaxTS int64
	flagUserID                   string
)

func init() {
	cmdUserFeed.Flag.Uint64Var(&flagUserCount, "c", 0, "COUNT")
	cmdUserFeed.Flag.StringVar(&flagUserMinID, "min_id", "", "MIN_ID")
	cmdUserFeed.Flag.StringVar(&flagUserMaxID, "max_id", "", "MAX_ID")

	cmdUserRecentMedia.Flag.StringVar(&flagUserID, "u", "", "USER_ID")
	cmdUserRecentMedia.Flag.Uint64Var(&flagUserCount, "c", 0, "COUNT")
	cmdUserRecentMedia.Flag.Int64Var(&flagUserMinTS, "min-ts", 0, "MIN_TS")
	cmdUserRecentMedia.Flag.Int64Var(&flagUserMaxTS, "max-ts", 0, "MAX_TS")
	cmdUserRecentMedia.Flag.StringVar(&flagUserMinID, "min_id", "", "MIN_ID")
	cmdUserRecentMedia.Flag.StringVar(&flagUserMaxID, "max_id", "", "MAX_ID")

	cmdUserLikedMedia.Flag.Uint64Var(&flagUserCount, "c", 0, "COUNT")
	cmdUserLikedMedia.Flag.StringVar(&flagUserMaxID, "max_id", "", "MAX_LIKE_ID")

	cmdSearchUser.Flag.Uint64Var(&flagUserCount, "c", 0, "COUNT")
}

func userInfo(r *Runner, c *Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", c.Usage)
		os.Exit(1)
	}
	userId := args[0]

	inst := r.Client.Instagram
	u, err := inst.Users.Get(userId)
	utils.Check(err)

	utils.UserPrinter(u)
	os.Exit(0)
}

func userFeed(r *Runner, c *Command, args []string) {
	inst := r.Client.Instagram
	params := new(instagram.Parameters)

	if flagUserCount > 0 {
		params.Count = flagUserCount
	}
	if flagUserMinID != "" {
		params.MinID = flagUserMinID
	}
	if flagUserMaxID != "" {
		params.MaxID = flagUserMaxID
	}

	media, next, err := inst.Users.MediaFeed(params)
	utils.Check(err)

	utils.MediaSlicePrinter(media, next, "-max_id")
	os.Exit(0)
}

func userRecentMedia(r *Runner, c *Command, args []string) {
	inst := r.Client.Instagram
	params := new(instagram.Parameters)

	userId := "self"
	if flagUserID != "" {
		userId = flagUserID
	}

	if flagUserCount > 0 {
		params.Count = flagUserCount
	}
	if flagUserMinID != "" {
		params.MinID = flagUserMinID
	}
	if flagUserMaxID != "" {
		params.MaxID = flagUserMaxID
	}
	if flagUserMinTS != 0 {
		params.MinTimestamp = flagUserMinTS
	}
	if flagUserMaxTS != 0 {
		params.MaxTimestamp = flagUserMaxTS
	}

	media, next, err := inst.Users.RecentMedia(userId, params)
	utils.Check(err)

	utils.MediaSlicePrinter(media, next, "-max_id")
	os.Exit(0)
}

func userLikedMedia(r *Runner, c *Command, args []string) {
	inst := r.Client.Instagram
	params := new(instagram.Parameters)

	if flagUserCount != 0 {
		params.Count = flagUserCount
	}
	if flagUserMaxID != "" {
		params.MaxID = flagUserMaxID
	}

	media, next, err := inst.Users.LikedMedia(params)
	utils.Check(err)

	utils.MediaSlicePrinter(media, next, "-max_id")
	os.Exit(0)
}

func searchUser(r *Runner, c *Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", c.Usage)
		os.Exit(1)
	}
	q := args[0]

	inst := r.Client.Instagram
	params := new(instagram.Parameters)

	if flagUserCount > 0 {
		params.Count = flagUserCount
	}

	u, next, err := inst.Users.Search(q, params)
	utils.Check(err)

	utils.UserSlicePrinter(u, next, "-max_id")
	os.Exit(0)
}
