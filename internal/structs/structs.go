package structs

type UserInfo struct {
	Username     string
	UserPassHash string
	UserUUID     string
}

type MessageInfo struct {
	MsgUserName  string `json:"msgusername"`
	MsgText      string `json:"msgtext"`
	MsgTimestamp string `json:"msgtimestamp"`
}

// type ConnInfo struct {
// 	// ws       *websocket.Conn
// 	Token    string
// 	ExpireAt int64
// }
