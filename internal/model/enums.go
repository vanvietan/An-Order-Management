package model

type Role string
type Status string
type Operation string

/*
Enums for Role in User table
Status in Order table
Operation in Audit_trail table
*/
const (
	RoleUser  Role = "USER"
	RoleAdmin Role = "ADMIN"
)

const (
	StatusApproved        Status = "APPROVED"
	StatusApprovalPending Status = "APPROVAL_PENDING"
	StatusShipping        Status = "SHIPPING"
	StatusShipped         Status = "SHIPPED"
)

const (
	OperationModified Operation = "MODIFIED"
	OperationDeleted  Operation = "DELETED"
)
