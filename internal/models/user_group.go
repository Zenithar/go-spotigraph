package models

// UserGroup defines user group contract
type UserGroup interface {
	GetGroupType() string
	GetGroupID() string
}
