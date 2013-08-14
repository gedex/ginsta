// Copyright 2013 The go-instagram AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package commands

import (
	"flag"
	"fmt"
	"github.com/gedex/ginsta/clients"
	"strings"
)

// Command represents command after ginsta
type Command struct {
	// Callback function that get executed for this command
	Callback func(r *Runner, cmd *Command, args []string)

	// Flag available for this command
	Flag flag.FlagSet

	// Command's usage that prefixed with command's name
	// and followed by flags and args
	Usage string

	// Short description of this command that shown when calling
	// global help
	Short string

	// Long description of this command that shown when calling
	// help on this command
	Long string
}

var (
	UsersCommands = []*Command{
		cmdUserInfo,
		cmdUserFeed,
		cmdUserRecentMedia,
		cmdUserLikedMedia,
		cmdSearchUser,
	}
	RelationshipsCommands = []*Command{
		cmdUserFollowings,
		cmdUserFollowers,
		cmdUserRequestedBy,
		cmdRelationshipWith,
		cmdFollowUser,
		cmdUnfollowUser,
		cmdBlockUser,
		cmdUnblockUser,
		cmdApproveUser,
		cmdDenyUser,
	}
	MediaCommands = []*Command{
		cmdMediaInfo,
		cmdSearchMedia,
		cmdPopularMedia,
	}
	CommentsCommands = []*Command{
		cmdMediaComments,
		cmdAddComment,
		cmdDelComment,
	}
	LikesCommands = []*Command{
		cmdMediaLikes,
		cmdAddLike,
		cmdDelLike,
	}
	TagsCommands = []*Command{
		cmdTagInfo,
		cmdRecentMediaByTag,
		cmdSearchTag,
	}
	LocationsCommands = []*Command{
		cmdLocationInfo,
		cmdRecentMediaByLocation,
		cmdSearchLocation,
	}
	BasicCommands = []*Command{
		cmdHelp,
		cmdVersion,
		cmdConfig,
		cmdTokenGet,
		cmdGeocoding,
		cmdReverseGeocoding,
	}
)

func (c *Command) PrintUsage() {
	if c.Runnable() {
		fmt.Printf("Usage: %s %s\n\n", clients.Name, c.Usage)
	}
	fmt.Println(strings.Trim(c.Long, "\n"))
}

func (c *Command) Name() string {
	name := c.Usage
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) Runnable() bool {
	return c.Callback != nil
}

func (c *Command) List() bool {
	return c.Short != ""
}

func All() []*Command {
	commands := make([]*Command, 0)
	commands = append(commands, UsersCommands...)
	commands = append(commands, RelationshipsCommands...)
	commands = append(commands, MediaCommands...)
	commands = append(commands, CommentsCommands...)
	commands = append(commands, LikesCommands...)
	commands = append(commands, TagsCommands...)
	commands = append(commands, LocationsCommands...)
	commands = append(commands, BasicCommands...)

	return commands
}
