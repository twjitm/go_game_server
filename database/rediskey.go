package database

import "fmt"

const UNFRIENDLIEST string = "h:user:info"
const (
	SESSION_POOL    string = "session:pool"
	SSO_SESSION_KEY string = "com.graduation.online.sso.entity:SsoPo:"
)

func GetSSoSessionKey(sessionId string) string {
	return fmt.Sprintf(SSO_SESSION_KEY, sessionId)
}
