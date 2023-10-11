package rankExample

import "strconv"

type Member interface {
	FormatInfo(memberID int) MemberInfo
}

type MemberInfo struct {
	MemberID int
	Name     string
	Image    string
	Uid      int
	Type     string
}

// user
type User struct {
	Type string
}

func NewUser(t string) Member {
	return &User{Type: t}
}

func (u *User) FormatInfo(memberID int) MemberInfo {
	// TODO: get user info ...
	return MemberInfo{
		MemberID: memberID,
		Name:     "username_" + strconv.Itoa(memberID),
		Image:    "image_" + strconv.Itoa(memberID),
		Uid:      memberID,
		Type:     u.Type,
	}
}

// room
type Room struct {
	Type string
}

func NewRoom(t string) Member {
	return &Room{Type: t}
}

func (r *Room) FormatInfo(memberID int) MemberInfo {
	// TODO: get room info ...
	return MemberInfo{
		MemberID: memberID,
		Name:     "roomName_" + strconv.Itoa(memberID),
		Image:    "image_" + strconv.Itoa(memberID),
		Uid:      memberID,
		Type:     r.Type,
	}
}
