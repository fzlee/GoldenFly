package other

import (
	"github.com/gorilla/feeds"
	"github.com/ikeikeikeike/go-sitemap-generator/v2/stm"
	"golden_fly/config"
	"golden_fly/page"
	"time"
)



func GenerateRSS(pages *[] page.Page) (string, error){
	conf := config.Get()

	now := time.Now()
	feed := &feeds.Feed{
		Title:       conf.SiteName,
		Link:        &feeds.Link{Href: conf.SiteURL},
		Description: conf.SiteDescription,
		Author:      &feeds.Author{Name: conf.SiteAuthor},
		Created:     now,
	}

	feed.Items = make([]*feeds.Item, len(*pages))

	for i:= range *pages {
		page := (*pages)[i]
		url := conf.SiteURL + "/articles/" + page.URL
		feed.Items[i] = &feeds.Item{
			Title:       page.Title,
			Link:        &feeds.Link{Href: url},
			Description: page.Title,
			Id:          url,
			Updated:     *page.UpdateTime,
			Created:     *page.CreateTime,
			Content:     page.HTML,
		}

		if page.NeedKey {
			feed.Items[i].Content = "文章已被加密，请输入密码访问"
			feed.Items[i].Description = feed.Items[i].Content
		}
	}

	return feed.ToRss()
}


func GenerateSitemap(pages *[]page.Page) *stm.Sitemap{
	conf := config.Get()

	sm := stm.NewSitemap(1)
	sm.SetDefaultHost(conf.SiteURL)
	sm.Create()
	sm.Add(stm.URL{{"loc", "/"}, {"changefreq", "daily"}})


	for i:= range *pages {
		page := (*pages)[i]
		sm.Add(stm.URL{
			{"loc", "/articles/" + page.URL},
			{"priority", 0.8},
		})
	}

	return sm
}
