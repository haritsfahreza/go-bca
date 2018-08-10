package bca

//Error represent BCA error response messsage
type Error struct {
	ErrorCode    string
	ErrorMessage ErrorLang
}

//ErrorLang represent BCA error response message language
type ErrorLang struct {
	Indonesian string
	English    string
}
