package core

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type CustomPages struct {
	Root  xml.Name     `xml:"pages"`
	Pages []CustomPage `xml:"page"`
}

type CustomPage struct {
	Name       string      `json:"name"`
	Title      string      `xml:"title" json:"title"`
	Icon       string      `xml:"icon" json:"icon"`
	Tooltip    string      `xml:"tooltip" json:"tooltip"`
	Body       HtmlContent `xml:"body"`
	BodyString string      `json:"body"`
}

type HtmlContent struct {
	Value string `xml:",innerxml"`
}

func ReadCustomPagesXmlFile() []CustomPage {
	exe, err := os.Executable()
	if err != nil {
		DebugError(err)
		return nil
	}
	dir := path.Dir(exe)

	data, err := ioutil.ReadFile(path.Join(dir, "pages.xml"))
	if err != nil {
		DebugError(err)
		return nil
	}

	pages := &CustomPages{}

	err = xml.Unmarshal(data, &pages)
	if err != nil {
		DebugError(err)
		return nil
	}

	md5Hash := md5.New()
	for i := 0; i < len(pages.Pages); i++ {
		pages.Pages[i].Title = strings.TrimSpace(pages.Pages[i].Title)
		pages.Pages[i].Icon = strings.TrimSpace(pages.Pages[i].Icon)
		pages.Pages[i].Tooltip = strings.TrimSpace(pages.Pages[i].Tooltip)
		pages.Pages[i].BodyString = pages.Pages[i].Body.Value

		if pages.Pages[i].Icon == "" {
			pages.Pages[i].Icon = "fa-external-link-alt"
		}

		if pages.Pages[i].Title != "" {
			md5Hash.Write([]byte(pages.Pages[i].Title))
			pages.Pages[i].Name = hex.EncodeToString(md5Hash.Sum(nil))
			md5Hash.Reset()
		}
	}

	return pages.Pages
}
