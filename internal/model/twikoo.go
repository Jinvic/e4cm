package model

type TwikooComment struct {
	ID string `json:"id"`
	UID string `json:"uid"`
	Nick string `json:"nick"`
	Mail string `json:"mail"`
	MailMd5 string `json:"mailMd5"`
	Link string `json:"link"`
	UA string `json:"ua"`
	IP string `json:"ip"`
	Master bool `json:"master"`
	URL string `json:"url"`
	Href string `json:"href"`
	Comment string `json:"comment"`
	Created int64 `json:"created"`
	Updated int64 `json:"updated"`
	IPRegion string `json:"ipRegion"`
}