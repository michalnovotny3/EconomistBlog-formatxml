package formatxml

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func OpenJSON(path string) ([]byte, error) {
	out := []byte("")
	if path == "" {
		return nil, errors.New("No path found")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	out, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	} else {
		return out, nil
	}
}

func Finalize(input []byte) ([]byte, error) {
	if input == nil || len(input) == 0 {
		return nil, errors.New("Wrong input")
	}
	var out []byte
	var tms string = string(input)
	tms = strings.Replace(tms, "&lt;", "<", -1)
	tms = strings.Replace(tms, "&gt;", ">", -1)
	tms = strings.Replace(tms, "&#34;", "\"", -1)
	tms = strings.Replace(tms, "</p><p>", "</p><p>", -1)
	tms = strings.Replace(tms, "<body.content>", "<body.content>\n\t\t\t", 1)
	tms = strings.Replace(tms, "</body.content>", "\n\t\t\t</body.content>", 1)
	tms = strings.Replace(tms, "></title>", "/>", 1)
	tms = strings.Replace(tms, "></docdata>", "/>", 1)
	tms = strings.Replace(tms, "></pubdata>", "/>", 1)
	tms = strings.Replace(tms, "></body.end>", "/>", 1)
	tms = strings.Replace(tms, "></location>", "/>", 1)
	out = []byte(tms)
	return out, nil
}

func ParseToXML(b []byte) ([]byte, error) {
	if b == nil || len(b) < 1 {
		return nil, errors.New("Empty Input")
	}
	var obj Core
	var out []byte
	var res ResultXML
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	} else {
		res.SetDefaults()
		res.Head.PubData.PubDatePublication, err = obj.GetPubdate()
		if err != nil {
			return nil, err
		}
		res.Head.PubData.PubPosSection, err = obj.GetSection()
		if err != nil {
			return nil, err
		}
		for _, tmp_segment := range obj.Text {
			res.Body.BodyContent += tmp_segment.GetContent()
		}
		res.Body.BodyHead.Headline.Hl1 = obj.HeadLine
		res.Body.BodyHead.Headline.Hl2 = obj.Subheadline
		res.Body.BodyHead.Abstract.Hl2 = obj.HeadLine
		if obj.Description != "" {
			res.Body.BodyHead.Abstract.Paragraph = obj.Description
		}
	}
	out, err = xml.MarshalIndent(res, "", "	")
	if err != nil {
		return nil, err
	}
	out, err = Finalize(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Core) GetSection() (string, error) {
	if len(c.ArticleSection.Internal) < 1 {
		return "", errors.New("No article section found")
	}
	for _, tmp_Core := range c.ArticleSection.Internal {
		if tmp_Core.TegType == "taxonomy" {
			return tmp_Core.HeadLine, nil
		}
	}
	return "", errors.New("No taxonomy found")
}

func (c *Core) GetPubdate() (string, error) {
	if c.DatePublished == "" {
		return "", errors.New("No date of publish found")
	}
	tmp_Date := strings.Split(c.DatePublished, "T")
	date := strings.Split(tmp_Date[0], "-")
	return date[0] + date[1] + date[2], nil
}

func (s *Segment) GetContent() string {
	var out string
	if s.Type == "tag" {
		var atr string
		if len(s.Attribs) > 0 {
			for key, val := range s.Attribs {
				atr += " " + key + "=\"" + val + "\""
			}
		}
		out = "<" + s.Name + atr + ">"
	} else if s.Type == "text" {
		out = s.Data
		return out
	}
	if len(s.Children) > 0 {
		for _, tem_Segment := range s.Children {
			out += tem_Segment.GetContent()
		}
	}
	out += "</" + s.Name + ">"
	return out
}

func (r *ResultXML) SetDefaults() {
	r.Version = "-//IPTC-NAA//DTD NITF-XML 2.0s//EN"
	r.Changedate = "29 March 2000"
	r.Changetime = "1900"
	r.Head.PubData.PubType = "web"
	r.Head.PubData.PubEditName = "The Economist online"
	r.Head.PubData.PubIssue = "1"
}

type ResultXML struct {
	XMLName    xml.Name `xml:"nitf"`
	Version    string   `xml:"version,attr"`
	Changedate string   `xml:"change.date,attr"`
	Changetime string   `xml:"change.time,attr"`
	Head       struct {
		XMLName   xml.Name `xml:"head"`
		HeadTytle string   `xml:"title,allowempty"`
		DocData   string   `xml:"docdata"`
		PubData   struct {
			XMLName            xml.Name `xml:"pubdata"`
			PubType            string   `xml:"type,attr"`
			PubIssue           string   `xml:"issue,attr"`
			PubDatePublication string   `xml:"date.publication,attr"`
			PubPosSection      string   `xml:"position.section,attr"`
			PubEditName        string   `xml:"edition.name,attr"`
		}
	}
	Body struct {
		XMLName  xml.Name `xml:"body"`
		BodyHead struct {
			XMLName  xml.Name `xml:"body.head"`
			Headline struct {
				XMLName xml.Name `xml:"hedline"`
				Hl1     string   `xml:"hl1"`
				Hl2     string   `xml:"hl2"`
			}
			Dateline struct {
				XMLName  xml.Name `xml:"dateline"`
				Location string   `xml:"location"`
			}
			Abstract struct {
				XMLName   xml.Name `xml:"abstract"`
				Hl2       string   `xml:"hl2"`
				Paragraph string   `xml:"p"`
			}
		}
		BodyContent string `xml:"body.content"`
		Bodyend     string `xml:"body.end"`
	}
}

type Core struct {
	Id          string `json:"id"`
	TegID       string `json:"tegID"`
	TegType     string `json:"tegType"`
	HeadLine    string `json:"headline"`
	Subheadline string `json:"subheadline"`
	Byline      string `json:"byline"`
	Description string `json:"description"`
	Url         struct {
		Canonical string `json:"Cannonical"`
	}
	Cannel              *Core  `json:"channel"`
	DateCreated         string `json:"dateCreated"`
	DateModified        string `json:"dateModified"`
	DatePublished       string `json:"datePublished"`
	IsAccessibleForFree bool   `json:"isAccessibleForFree"`
	Print               struct {
		Section *Core `json:"Section"`
	}
	ArticleSection struct {
		Internal []*Core `json:"Internal"`
	}
	Publication []*Core   `json:"publication"`
	Image       Image     `json:"image"`
	Text        []Segment `json:"text"`
}

type Image struct {
	Main   *Core
	Inline string
	Cover  json.RawMessage `json:""`
}

type Segment struct {
	Type     string
	Name     string
	Attribs  map[string]string
	Children []Segment
	Data     string
}
