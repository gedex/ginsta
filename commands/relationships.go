package commands

import (
	"fmt"
	"github.com/gedex/ginsta/utils"
	"github.com/gedex/go-instagram/instagram"
	"os"
)

var (
	cmdUserFollowings = &Command{
		Callback: userFollowings,
		Usage:    "user_followings [USER_ID]",
		Short:    "Get the list of users this user follows.",
		Long: `Get the list of users this user (with ID USER_ID if specified, or
authenticated user if USER_ID is not set) follows.
`,
	}
	cmdUserFollowers = &Command{
		Callback: userFollowers,
		Usage:    "user_followers [USER_ID]",
		Short:    "Get the list of users this user is followed by.",
		Long: `Get the list of users this user (with ID USER_ID if specified, or
authenticated user if USER_ID is not set) is followed by.
`,
	}
	cmdUserRequestedBy = &Command{
		Callback: userRequestedBy,
		Usage:    "user_requested_by",
		Short:    "Get the list of users who have requested this user's permission to follow",
		Long:     `Get the list of users who have requested this user's permission to follow.`,
	}
	cmdRelationshipWith = &Command{
		Callback: relationshipWith,
		Usage:    "relationship_with [USER_ID]",
		Short:    "Get information about a relationship to another user.",
		Long:     "Get information about a relationship to another user.",
	}
	cmdFollowUser = &Command{
		Callback: followUser,
		Usage:    "follow_user [USER_ID]",
		Short:    "Follow a user specified with USER_ID",
		Long:     "Follow a user specified with USER_ID",
	}
	cmdUnfollowUser = &Command{
		Callback: unfollowUser,
		Usage:    "unfollow_user [USER_ID]",
		Short:    "Unfollow a user specified with USER_ID",
		Long:     "Unfollow a user specified with USER_ID",
	}
	cmdBlockUser = &Command{
		Callback: blockUser,
		Usage:    "block_user [USER_ID]",
		Short:    "Block a user specified with USER_ID",
		Long:     "Block a user specified with USER_ID",
	}
	cmdUnblockUser = &Command{
		Callback: unblockUser,
		Usage:    "unblock_user [USER_ID]",
		Short:    "Unblock a user specified with USER_ID",
		Long:     "Unblock a user specified with USER_ID",
	}
	cmdApproveUser = &Command{
		Callback: approveUser,
		Usage:    "approve_user [USER_ID]",
		Short:    "Approve request from user specified with USER_ID",
		Long:     "Approve request from user specified with USER_ID",
	}
	cmdDenyUser = &Command{
		Callback: denyUser,
		Usage:    "deny_user [USER_ID]",
		Short:    "Deny request from user specified with USER_ID",
		Long:     "Deny request from user specified with USER_ID",
	}
)

func userFollowings(r *Runner, c *Command, args []string) {
	userId := "self"
	if len(args) == 1 {
		userId = args[0]
	}

	inst := r.Client.Instagram
	u, next, err := inst.Relationships.Follows(userId)
	utils.Check(err)

	utils.UserSlicePrinter(u, next, "-max-id")
	os.Exit(0)
}

func userFollowers(r *Runner, c *Command, args []string) {
	userId := "self"
	if len(args) == 1 {
		userId = args[0]
	}

	inst := r.Client.Instagram
	u, next, err := inst.Relationships.FollowedBy(userId)
	utils.Check(err)

	utils.UserSlicePrinter(u, next, "-max-id")
	os.Exit(0)
}

func userRequestedBy(r *Runner, c *Command, args []string) {
	inst := r.Client.Instagram
	u, next, err := inst.Relationships.RequestedBy()
	utils.Check(err)

	utils.UserSlicePrinter(u, next, "-max-id")
	os.Exit(0)
}

func relationshipWith(r *Runner, c *Command, args []string) {
	userRelationshipAction("relationship_with", r, c, args)
}

func followUser(r *Runner, c *Command, args []string) {
	userRelationshipAction("follow", r, c, args)
}

func unfollowUser(r *Runner, c *Command, args []string) {
	userRelationshipAction("unfollow", r, c, args)
}

func blockUser(r *Runner, c *Command, args []string) {
	userRelationshipAction("block", r, c, args)
}

func unblockUser(r *Runner, c *Command, args []string) {
	userRelationshipAction("unblock", r, c, args)
}

func approveUser(r *Runner, c *Command, args []string) {
	userRelationshipAction("approve", r, c, args)
}

func denyUser(r *Runner, c *Command, args []string) {
	userRelationshipAction("deny", r, c, args)
}

func userRelationshipAction(action string, r *Runner, c *Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", c.Usage)
		os.Exit(1)
	}
	userId := args[0]

	inst := r.Client.Instagram

	var err error
	var rel *instagram.Relationship
	switch action {
	case "relationship_with":
		rel, err = inst.Relationships.Relationship(userId)
	case "follow":
		rel, err = inst.Relationships.Follow(userId)
	case "unfollow":
		rel, err = inst.Relationships.Unfollow(userId)
	case "block":
		rel, err = inst.Relationships.Block(userId)
	case "unblock":
		rel, err = inst.Relationships.Unblock(userId)
	case "approve":
		rel, err = inst.Relationships.Approve(userId)
	case "deny":
		rel, err = inst.Relationships.Deny(userId)
	}
	utils.Check(err)

	utils.RelationshipPrinter(rel)
	os.Exit(0)
}
