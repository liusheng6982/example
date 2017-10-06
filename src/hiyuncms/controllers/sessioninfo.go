package controllers


const(
	BACK_USER_SESSION = "hiyuncms.back.user"
)
type BackendUserSession struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`
}