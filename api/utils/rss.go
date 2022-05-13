package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"api/dao"
)

const format = "2006-01-02 15:04:05"

func Rss() error {
	posts, err := dao.GetAllPostsWithDetail()
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)

	for _, p := range posts {
		buf.WriteString(fmt.Sprintf(`<item><title>%s</title><tags>%s</tags><pubDate>%s</pubDate><updateDate>%s</updateDate><link>https://qianx1.xyz/posts/%s</link><guid>https://qianx1.xyz/posts/%s</guid></item>`, p.Title, p.Tags, time.Unix(p.CreateTime, 0).Format(format), time.Unix(p.UpdateTime, 0).Format(format), p.DisplayId, p.DisplayId))
	}

	rss := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?><rss version="2.0"><channel><title>千禧的博客</title><link>http://qianx1.xyz</link><description>qianxi's personal website</description><updateTime>%s</updateTime>%s</channel></rss>`, time.Now().Format(format), buf.String())

	// write file to disk
	err = ioutil.WriteFile("rss.xml", []byte(rss), 0o644)
	if err != nil {
		return err
	}

	return nil
}
