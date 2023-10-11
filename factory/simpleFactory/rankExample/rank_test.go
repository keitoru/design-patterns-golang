package rankExample

import "testing"

func Test_rank(t *testing.T) {
	mr := NewMemberRank()

	m1, err := mr.CreateMember("user")
	if err != nil {
		t.Error(err)
		return
	}
	info1 := m1.FormatInfo(111)
	t.Log("info1:", info1)

	m2, err := mr.CreateMember("room")
	if err != nil {
		t.Error(err)
		return
	}
	info2 := m2.FormatInfo(222)
	t.Log("info2:", info2)
}
