package rankExample

import "fmt"

type memberCreator func(typ string) Member

type MemberRank struct {
	creators map[string]memberCreator
}

func NewMemberRank() *MemberRank {
	return &MemberRank{
		creators: map[string]memberCreator{
			"user": NewUser,
			"room": NewRoom,
		},
	}
}

func (m *MemberRank) CreateMember(typ string) (Member, error) {
	memberCreator, ok := m.creators[typ]
	if !ok {
		return nil, fmt.Errorf("member typ: %s is not supported yet", typ)
	}

	return memberCreator(typ), nil
}
