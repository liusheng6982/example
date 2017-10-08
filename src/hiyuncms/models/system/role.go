package system


type Role struct {
	Id int64  `xorm:"pk BIGINT autoincr"`
	RoleName string `xorm:"varchar(25) notnull unique"`
}