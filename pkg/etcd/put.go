package etcd

import (
	"encoding/json"
)

type IAppLink struct {
	TitleText       string      `json:"titleText"`
	URL             string      `json:"url"`
	IconURL         string      `json:"iconUrl"`
	LinkType        string      `json:"linkType"` // e.g. internal, external, iframe
	IconURLMap      *IconURLMap `json:"iconUrlMap,omitempty"`
	DescriptionText string      `json:"descriptionText,omitempty"`
	ButtonText      string      `json:"buttonText,omitempty"`
}

type IconURLMap struct {
	Disabled *string `json:"disabled,omitempty"`
	Active   *string `json:"active,omitempty"`
	Hover    *string `json:"hover,omitempty"`
}

func (i *IAppLink) new(titleText, linkUrl, iconUrl, linkType string) {
	i.TitleText = titleText
	i.URL = linkUrl
	i.IconURL = iconUrl
	i.LinkType = linkType
}

// IAppFeature---------------------------

type IAppFeature struct {
	IconURL   string `json:"iconUrl"`
	TitleText string `json:"titleText"`
}

func (i *IAppFeature) new(titleText, iconUrl string) {
	i.TitleText = titleText
	i.IconURL = iconUrl
}

// Put---------------------

func (c *EtcdCli) putIAppLink(i IAppLink) {
	// guard.Logger.Info("Put IAppLink")
	// fmt.Printf("Put IAppLink: %+v\n", i)
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	_, err = c.PutMetadata("iAppLink", string(b))
	if err != nil {
		panic(err)
	}
}

func (c *EtcdCli) putIAppFeature(i IAppFeature) {
	// guard.Logger.Info("Put IAppFeature")
	// fmt.Printf("Put IAppFeature: %+v\n", i)
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	_, err = c.PutMetadata("iAppFeature", string(b))
	if err != nil {
		panic(err)
	}
}
