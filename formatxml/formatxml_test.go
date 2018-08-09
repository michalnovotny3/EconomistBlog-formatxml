package formatxml

import (
	"testing"
)

var trymsg string = `{
    "id": "/content/la1sask15cqjl06ecq9jb7c26237033m",
    "tegID": "la1sask15cqjl06ecq9jb7c26237033m",
    "tegType": "article",
    "headline": "Why corporate America loves Donald Trump",
    "subheadline": "The Trump trade",
    "byline": "",
    "description": "American executives are betting that the president is good for business. Not in the long run",
    "url": {
        "canonical": "https://stage.economist.com/leaders/2018/05/24/why-corporate-america-loves-donald-trump"
    },
    "channel": {
        "id": "/content/j53t6hsedat4l7rkbb1le98u73262sh5",
        "tegID": "j53t6hsedat4l7rkbb1le98u73262sh5",
        "tegType": "page",
        "headline": "About us",
        "subheadline": "",
        "byline": "",
        "description": "",
        "url": {
            "canonical": "https://stage.economist.com/help/about-us"
        }
    },
    "dateCreated": "2018-05-24T08:33:21Z",
    "dateModified": "2018-07-25T09:22:23Z",
    "datePublished": "2018-05-24T04:00:00Z",
    "isAccessibleForFree": false,
    "print": {
        "section": {
            "id": "/content/0eqdc3kshoq96naup34ruco7pktc6n51",
            "tegID": "0eqdc3kshoq96naup34ruco7pktc6n51",
            "tegType": "taxonomy",
            "headline": "Leaders",
            "subheadline": "",
            "byline": "",
            "description": "",
            "url": {
                "canonical": "https://stage.economist.com/leaders/"
            }
        }
    },
    "articleSection": {
        "internal": [
            {
                "id": "/content/c1b8fm6psufl0h57e27eup9hb8nmv0t3",
                "tegID": "c1b8fm6psufl0h57e27eup9hb8nmv0t3",
                "tegType": "taxonomy",
                "headline": "Leaders",
                "subheadline": "",
                "byline": "",
                "description": "",
                "url": {
                    "canonical": "https://stage.economist.com/sections/leaders"
                }
            }
        ]
    },
    "publication": [
        {
            "id": "/content/g8tboctaspgrdfqb9nrou09b87fm3a7d",
            "tegID": "g8tboctaspgrdfqb9nrou09b87fm3a7d",
            "tegType": "issue",
            "headline": "Why corporate America loves Donald Trump",
            "subheadline": "",
            "byline": "",
            "description": "",
            "url": {
                "canonical": "https://stage.economist.com/printedition/2018-05-26"
            },
            "image": {
                "main": null
            }
        }
    ],
    "image": {
        "main": {
            "id": "/content/lug9kr3e57pn3bua046pgn6608qs2lgm",
            "tegID": "lug9kr3e57pn3bua046pgn6608qs2lgm",
            "tegType": "image",
            "headline": "",
            "subheadline": "",
            "byline": "",
            "description": "",
            "url": {
                "canonical": "https://stage.economist.com/sites/default/files/images/print-edition/20180526_LDD001_0.jpg"
            }
        },
        "inline": null,
        "cover": null
    },
    "text": [
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "data": "MOST American  experts warn of a looming constitutional crisis.",
                    "type": "text"
                }
            ]
        },
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "data": "Amid the tumult soar.",
                    "type": "text"
                }
            ]
        },
        {
            "type": "tag",
            "name": "figure",
            "attribs": {
                "role": "presentation",
                "itemtype": "https://schema.org/WPAdBlock"
            },
            "children": []
        },
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "data": "The financial, America Inc is being short-sighted and sloppy.",
                    "type": "text"
                }
            ]
        },
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "type": "tag",
                    "name": "strong",
                    "attribs": {},
                    "children": [
                        {
                            "data": "The view from the C-suite",
                            "type": "text"
                        }
                    ]
                }
            ]
        },
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "data": "Since winning Congress  cutting headline -tax profits (it accounts for a tenth of the fiscal deficit).",
                    "type": "text"
                }
            ]
        },
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "data": "Deregulation is in full swing. This week saw a relaxation of banking rules (see ",
                    "type": "text"
                },
                {
                    "type": "tag",
                    "name": "a",
                    "attribs": {
                        "href": "https://www.economist.com/news/finance-and-economics/21743112-americas-dodd-frank-act-gets-tweak-not-rewrite-rare-bipartisan-moment-allows"
                    },
                    "children": [
                        {
                            "data": "article",
                            "type": "text"
                        }
                    ]
                },
                {
                    "data": "). The leaders of earnings.",
                    "type": "text"
                }
            ]
        },
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "data": "The trouble is that.",
                    "type": "text"
                }
            ]
        },
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "data": "As the contours offfs on their calls with investors. Over time, a mesh of distortions will build up.",
                    "type": "text"
                }
            ]
        },
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "data": "Because trade ady introduced new regimes this year for financial instruments and data.",
                    "type": "text"
                }
            ]
        },
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "data": "The expense ecome self-sufficient (see ",
                    "type": "text"
                },
                {
                    "type": "tag",
                    "name": "a",
                    "attribs": {
                        "href": "https://www.economist.com/news/china/21743127-chinese-officials-still-have-plenty-reasons-worry-threatened-trade-war-between-china-and"
                    },
                    "children": [
                        {
                            "data": "article",
                            "type": "text"
                        }
                    ]
                },
                {
                    "data": "), or Mr Trump is mocked ent up.",
                    "type": "text"
                }
            ]
        },
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "type": "tag",
                    "name": "strong",
                    "attribs": {},
                    "children": [
                        {
                            "data": "The new laws of the jungle",
                            "type": "text"
                        }
                    ]
                }
            ]
        },
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "data": "Another too. a baable business environment that results will raise the cost of capital.",
                    "type": "text"
                }
            ]
        },
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "data": "As Afiring workers and slashing costs has become toxic.",
                    "type": "text"
                }
            ]
        },
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "data": "Republicans sense.",
                    "type": "text"
                }
            ]
        }
    ]
}`

var example2 string = `{
    "id": "/content/la1sask15cqjl06ecq9jb7c26237033m",
    "tegID": "la1sask15cqjl06ecq9jb7c26237033m",
    "tegType": "article",
    "headline": "Why corporate America loves Donald Trump",
    "subheadline": "The Trump trade",
    "byline": "",
    "description": "American executives are betting that the president is good for business. Not in the long run",
    "url": {
        "canonical": "https://stage.economist.com/leaders/2018/05/24/why-corporate-america-loves-donald-trump"
    },
    "channel": {
        "id": "/content/j53t6hsedat4l7rkbb1le98u73262sh5",
        "tegID": "j53t6hsedat4l7rkbb1le98u73262sh5",
        "tegType": "page",
        "headline": "About us",
        "subheadline": "",
        "byline": "",
        "description": "",
        "url": {
            "canonical": "https://stage.economist.com/help/about-us"
        }
    },
    "print": {
        "section": {
            "id": "/content/0eqdc3kshoq96naup34ruco7pktc6n51",
            "tegID": "0eqdc3kshoq96naup34ruco7pktc6n51",
            "tegType": "taxonomy",
            "headline": "Leaders",
            "subheadline": "",
            "byline": "",
            "description": "",
            "url": {
                "canonical": "https://stage.economist.com/leaders/"
            }
        }
    },
    "articleSection": {
        "internal": [
            {
                "id": "/content/c1b8fm6psufl0h57e27eup9hb8nmv0t3",
                "tegID": "c1b8fm6psufl0h57e27eup9hb8nmv0t3",
                "tegType": "taxonomy",
                "headline": "Leaders",
                "subheadline": "",
                "byline": "",
                "description": "",
                "url": {
                    "canonical": "https://stage.economist.com/sections/leaders"
                }
            }
        ]
    },
    "image": {
        "main": {
            "id": "/content/lug9kr3e57pn3bua046pgn6608qs2lgm",
            "tegID": "lug9kr3e57pn3bua046pgn6608qs2lgm",
            "tegType": "image",
            "headline": "",
            "subheadline": "",
            "byline": "",
            "description": "",
            "url": {
                "canonical": "https://stage.economist.com/sites/default/files/images/print-edition/20180526_LDD001_0.jpg"
            }
        },
        "inline": null,
        "cover": null
    },
    "text": [
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "data": "MOST American elites believe that the Trump presidency is hurting their country. Foreign-policy mandarins are terrified that security alliances are being wrecked. Fiscal experts warn that borrowing is spiralling out of control. Scientists deplore the rejection of climate change. And some legal experts warn of a looming constitutional crisis.",
                    "type": "text"
                }
            ]
        }
    ]
}`

var example3 string = `{
    "id": "/content/la1sask15cqjl06ecq9jb7c26237033m",
    "tegID": "la1sask15cqjl06ecq9jb7c26237033m",
    "tegType": "article",
    "headline": "Why corporate America loves Donald Trump",
    "subheadline": "The Trump trade",
    "byline": "",
    "description": "American executives are betting that the president is good for business. Not in the long run",
    "url": {
        "canonical": "https://stage.economist.com/leaders/2018/05/24/why-corporate-america-loves-donald-trump"
    },
    "channel": {
        "id": "/content/j53t6hsedat4l7rkbb1le98u73262sh5",
        "tegID": "j53t6hsedat4l7rkbb1le98u73262sh5",
        "tegType": "page",
        "headline": "About us",
        "subheadline": "",
        "byline": "",
        "description": "",
        "url": {
            "canonical": "https://stage.economist.com/help/about-us"
        }
    },
    "dateCreated": "2018-05-24T08:33:21Z",
    "dateModified": "2018-07-25T09:22:23Z",
    "datePublished": "2018-05-24T04:00:00Z",
    "isAccessibleForFree": false,
    "print": {
        "section": {
            "id": "/content/0eqdc3kshoq96naup34ruco7pktc6n51",
            "tegID": "0eqdc3kshoq96naup34ruco7pktc6n51",
            "tegType": "taxonomy",
            "headline": "Leaders",
            "subheadline": "",
            "byline": "",
            "description": "",
            "url": {
                "canonical": "https://stage.economist.com/leaders/"
            }
        }
    },
    "image": {
        "main": {
            "id": "/content/lug9kr3e57pn3bua046pgn6608qs2lgm",
            "tegID": "lug9kr3e57pn3bua046pgn6608qs2lgm",
            "tegType": "image",
            "headline": "",
            "subheadline": "",
            "byline": "",
            "description": "",
            "url": {
                "canonical": "https://stage.economist.com/sites/default/files/images/print-edition/20180526_LDD001_0.jpg"
            }
        },
        "inline": null,
        "cover": null
    },
    "text": [
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "data": "MOST American elites believe that the Trump presidency is hurting their country. Foreign-policy mandarins are terrified that security alliances are being wrecked. Fiscal experts warn that borrowing is spiralling out of control. Scientists deplore the rejection of climate change. And some legal experts warn of a looming constitutional crisis.",
                    "type": "text"
                }
            ]
        }
    ]
}`

var example4 string = `{
    "id": "/content/la1sask15cqjl06ecq9jb7c26237033m",
    "tegID": "la1sask15cqjl06ecq9jb7c26237033m",
    "tegType": "article",
    "headline": "Why corporate America loves Donald Trump",
    "subheadline": "The Trump trade",
    "byline": "",
    "description": "American executives are betting that the president is good for business. Not in the long run",
    "url": {
        "canonical": "https://stage.economist.com/leaders/2018/05/24/why-corporate-america-loves-donald-trump"
    },
    "channel": {
        "id": "/content/j53t6hsedat4l7rkbb1le98u73262sh5",
        "tegID": "j53t6hsedat4l7rkbb1le98u73262sh5",
        "tegType": "page",
        "headline": "About us",
        "subheadline": "",
        "byline": "",
        "description": "",
        "url": {
            "canonical": "https://stage.economist.com/help/about-us"
        }
    },
    "dateCreated": "2018-05-24T08:33:21Z",
    "dateModified": "2018-07-25T09:22:23Z",
    "datePublished": "2018-05-24T04:00:00Z",
    "isAccessibleForFree": false,
    "print": {
    },
    "articleSection": {
        "internal": [
            {
            }
        ]
    },
    "publication": [
        {
            "id": "/content/g8tboctaspgrdfqb9nrou09b87fm3a7d",
            "tegID": "g8tboctaspgrdfqb9nrou09b87fm3a7d",
            "tegType": "issue",
            "headline": "Why corporate America loves Donald Trump",
            "subheadline": "",
            "byline": "",
            "description": "",
            "url": {
                "canonical": "https://stage.economist.com/printedition/2018-05-26"
            },
            "image": {
                "main": null
            }
        }
    ],
    "image": {
        "main": {
            "id": "/content/lug9kr3e57pn3bua046pgn6608qs2lgm",
            "tegID": "lug9kr3e57pn3bua046pgn6608qs2lgm",
            "tegType": "image",
            "headline": "",
            "subheadline": "",
            "byline": "",
            "description": "",
            "url": {
                "canonical": "https://stage.economist.com/sites/default/files/images/print-edition/20180526_LDD001_0.jpg"
            }
        },
        "inline": null,
        "cover": null
    },
    "text": [
        {
            "type": "tag",
            "name": "p",
            "attribs": {},
            "children": [
                {
                    "data": "MOST American  experts warn of a looming constitutional crisis.",
                    "type": "text"
                }
            ]
        }
    ]
}`

func TestOpenJSON(t *testing.T) {
	res, err := OpenJSON("path")
	if err == nil || res != nil {
		t.Error("Wrong path Failed")
	}
	if res, err = OpenJSON(""); err == nil || res != nil {
		t.Error("Empty Path Failed")
	}
	if _, err = OpenJSON("example.json"); err == nil {
		t.Log("Pass")
	} else {
		t.Error("Unexpected Error")
	}
}

func TestFinalize(t *testing.T) {
	b := []byte("")
	res, err := Finalize(b)
	if err == nil || res != nil {
		t.Errorf("Empty byte slice Failed")
	}
	res, err = Finalize([]byte("faaefeadgtrghdt"))
	if err == nil {
		t.Log("Passed")
	}
}

func TestParseToXML(t *testing.T) {
	b := []byte("")
	res, err := ParseToXML(b)
	if err == nil || res != nil {
		t.Errorf("Empty byte slice Failed")
	} else {
		t.Log("Empty byte slice Pass")
	}
	if res, err = ParseToXML([]byte(example2)); err == nil {
		t.Error("missing date of publication fail")
	}
	if res, err = ParseToXML([]byte(example3)); err == nil {
		t.Error("Missing article section fail")
	}
	if res, err = ParseToXML([]byte(example4)); err == nil {
		t.Error("Missing taxonomy section fail")
	}
	res, err = ParseToXML([]byte(trymsg))
	if err != nil || len(res) < 1 {
		t.Error("Correct JSON failed")
	} else {
		t.Log("Passed")
	}
}
