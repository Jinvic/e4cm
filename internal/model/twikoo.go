package model

import (
	cryptoUtil "e4cm/internal/util/crypto"
	"strings"
	"time"
)

type TwikooComment struct {
	ID       string `json:"id"`
	UID      string `json:"uid"`
	Nick     string `json:"nick"`
	Mail     string `json:"mail"`
	MailMd5  string `json:"mailMd5"`
	Link     string `json:"link"`
	UA       string `json:"ua"`
	IP       string `json:"ip"`
	Master   bool   `json:"master"`
	URL      string `json:"url"`
	Href     string `json:"href"`
	Comment  string `json:"comment"`
	Created  int64  `json:"created"`
	Updated  int64  `json:"updated"`
	IPRegion string `json:"ipRegion"`
}

func (t *TwikooComment) GetOldID() string {
	oldID := strings.TrimPrefix(t.URL, "/echo/")
	return oldID
}

func (t *TwikooComment) ToComment(newID string) Comment {
	return Comment{
		EchoID:    newID,
		Nickname:  t.Nick,
		Email:     t.Mail,
		Website:   t.Link,
		Content:   t.Comment,
		Status:    StatusPending,
		Hot:       false,
		IPHash:    cryptoUtil.HashClientIP(t.IP),
		UserAgent: t.UA,
		Source:    SourceGuest,
		CreatedAt: time.Unix(t.Created/1000, 0),
		UpdatedAt: time.Unix(t.Updated/1000, 0),
	}
}
