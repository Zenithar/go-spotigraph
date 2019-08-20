package models

// PersonGroup defines user group contract
type PersonGroup interface {
	GetGroupType() string
	GetGroupID() string
}
