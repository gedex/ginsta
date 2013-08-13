package utils

import (
	"fmt"
	"github.com/gedex/go-instagram/instagram"
)

func UserPrinter(u *instagram.User) {
	fmt.Printf("%-16s : %v\n", "User ID", u.ID)
	fmt.Printf("%-16s : %v\n", "Username", u.Username)

	if u.FullName != "" {
		fmt.Printf("%-16s : %v\n", "Full name", u.FullName)
	}

	fmt.Printf("%-16s : %v\n", "Profile Pic URL", u.ProfilePicture)

	if u.Bio != "" {
		fmt.Printf("%-16s : %v\n", "Bio", u.Bio)
	}

	if u.Website != "" {
		fmt.Printf("%-16s : %v\n", "Website", u.Website)
	}

	if u.Counts != nil {
		fmt.Printf("%-16s : %v\n", "Total media", u.Counts.Media)
		fmt.Printf("%-16s : %v\n", "Total followings", u.Counts.Follows)
		fmt.Printf("%-16s : %v\n", "Total followers", u.Counts.FollowedBy)
	}
}

func UserSlicePrinter(u []instagram.User, rp *instagram.ResponsePagination, s string) {
	for i, _u := range u {
		UserPrinter(&_u)
		if i != len(u)-1 {
			fmt.Println()
		}
	}
	if rp != nil && rp.NextMaxID != "" {
		fmt.Printf("\nGet next page by supplying: %s %v\n", s, rp.NextMaxID)
	}
}

func MediaPrinter(m instagram.Media) {
	fmt.Printf("%-20s : %v\n", "Media ID", m.ID)
	fmt.Printf("%-20s : %v\n", "Media type", m.Type)
	fmt.Printf("%-20s : %v\n", "Filter", m.Filter)
	fmt.Printf("%-20s : %v\n", "Link", m.Link)
	if m.Caption != nil {
		fmt.Printf("%-20s : %v\n", "Caption", m.Caption.Text)
		fmt.Printf("%-20s : %v (%v)\n", "Caption by", m.Caption.From.Username, m.Caption.From.ID)
	}
	if m.Comments != nil {
		fmt.Printf("%-20s : %v\n", "Number of comments", m.Comments.Count)
	}
	if m.Likes != nil {
		fmt.Printf("%-20s : %v\n", "Number of likes", m.Likes.Count)
	}
	fmt.Printf("%-20s : %v\n", "Created at", m.CreatedTime)
	fmt.Printf("%-20s : %v (%v)\n", "Uploaded by", m.User.Username, m.User.ID)
}

func MediaSlicePrinter(m []instagram.Media, rp *instagram.ResponsePagination, s string) {
	for i, _m := range m {
		MediaPrinter(_m)
		if i != len(m)-1 {
			fmt.Println()
		}
	}
	if rp != nil && rp.NextMaxID != "" {
		fmt.Printf("\nGet next page by supplying: %v %v\n", s, rp.NextMaxID)
	}
}

func RelationshipPrinter(r *instagram.Relationship) {
	fmt.Printf("%-20s : %v\n", "Outgoing status", r.OutgoingStatus)
	fmt.Printf("%-20s : %v\n", "Incoming status", r.IncomingStatus)
}

func CommentPrinter(c instagram.Comment) {
	fmt.Printf("%-20s : %v\n", "Comment ID", c.ID)
	fmt.Printf("%-20s : %v\n", "Text", c.Text)
	fmt.Printf("%-20s : %v\n", "Created at", c.CreatedTime)
	fmt.Printf("%-20s : %v (%v)\n", "From", c.From.Username, c.From.ID)
}

func CommentSlicePrinter(c []instagram.Comment, rp *instagram.ResponsePagination, s string) {
	for i, _c := range c {
		CommentPrinter(_c)
		if i != len(c)-1 {
			fmt.Println()
		}
	}
	if rp != nil && rp.NextMaxID != "" {
		fmt.Printf("\nGet next page by supplying: %s %v\n", s, rp.NextMaxID)
	}
}

func TagPrinter(t *instagram.Tag) {
	fmt.Printf("%-20s : %v\n", "Name", t.Name)
	fmt.Printf("%-20s : %v\n", "Media count", t.MediaCount)
}

func TagSlicePrinter(t []instagram.Tag, rp *instagram.ResponsePagination, s string) {
	for i, _t := range t {
		TagPrinter(&_t)
		if i != len(t)-1 {
			fmt.Println()
		}
	}
	if rp != nil && rp.NextMaxID != "" {
		fmt.Printf("\nGet next page by supplying: %s %v\n", s, rp.NextMaxID)
	}
}
