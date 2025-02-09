package crawlerLab

import (
	"fmt"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/gocolly/colly/v2"
)

func MyColly() {
	// 建立一個新的 Collector
	collector := colly.NewCollector()

	// 設定訪問 HTML 標籤的回呼函式
	collector.OnHTML(".quote", func(e *colly.HTMLElement) {
		quote := e.ChildText(".text")    // 抓取名言
		author := e.ChildText(".author") // 抓取作者
		fmt.Printf("名言: %s - %s\n", quote, author)
	})

	// 訪問網頁
	err := collector.Visit("https://quotes.toscrape.com/")
	if err != nil {
		fmt.Println("爬取失敗:", err)
	}
}

func MyRod() {
	// 啟動 Chromium 並忽略 WebGL 檢測 (某些網站會用來偵測機器人)
	url := launcher.New().Headless(true).MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	// 開啟 HLTV 頁面
	page := browser.MustPage("https://www.hltv.org/matches?predefinedFilter=top_tier")

	// 等待 JavaScript 加載完成
	page.WaitLoad()

	// 選取所有 upcomingMatch 區塊
	days := page.MustElements(".upcomingMatchesSection")

	fmt.Println("爬取比賽資訊中...")

	for _, day := range days {
		day_text := day.MustElement(".matchDayHeadline").MustText()
		matches := day.MustElements(".upcomingMatch")

		for _, match := range matches {
			// 取得比賽時間
			time_text := match.MustElement(".matchTime").MustText()

			team_1 := match.MustElement(".team1")
			team_1_name := team_1.MustElement(".matchTeamName").MustText()
			// team_1_logo := team_1.MustElement(".matchTeamLogoContainer").MustElement(".matchTeamLogo").MustAttribute("src")
			// fmt.Println(*team_1_logo) // MustAttribute("src") 回傳 *string，所以要 *logoSrc 才能取得真正的 string

			team_2 := match.MustElement(".team2")
			team_2_name := team_2.MustElement(".matchTeamName").MustText()
			// team_2_logo := team_2.MustElement(".matchTeamLogoContainer").MustElement(".matchTeamLogo").MustAttribute("src")
			// fmt.Println(*team_2_logo) // MustAttribute("src") 回傳 *string，所以要 *logoSrc 才能取得真正的 string

			// 取得比賽名稱
			event := match.MustElement(".matchEvent")
			event_name := event.MustElement(".matchEventName").MustText()
			event_logo := event.MustElement(".matchEventLogoContainer").MustElement(".matchEventLogo").MustAttribute("src")
			event_logo_string := *event_logo
			event_logo_url := strings.ReplaceAll(event_logo_string, "\\u0026", "&")
			fmt.Println(event_logo_url)

			// 印出比賽資訊
			fmt.Printf("時間: %s - %s - %s - %s vs %s \n", day_text, time_text, event_name, team_1_name, team_2_name)
		}
	}
}
