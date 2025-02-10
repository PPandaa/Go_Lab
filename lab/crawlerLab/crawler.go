package crawlerLab

import (
	"fmt"
	"regexp"
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

func UCLRod() {
	// 啟動 Chromium 並忽略 WebGL 檢測 (某些網站會用來偵測機器人)
	url := launcher.New().Headless(true).MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	// 開啟頁面
	page := browser.MustPage("https://www.fotmob.com/leagues/42/matches/champions-league")

	// 等待 JavaScript 加載完成
	page.WaitLoad()

	// 選取區塊
	match_day_element := page.MustElement(".css-1hrxww6-LeagueMatchesSectionCSS")
	day_text := match_day_element.MustElement(".css-1b2yz7i-HeaderCSS").MustText()

	matchs_element := match_day_element.MustElements(".css-s4hjf6-MatchWrapper")

	for _, match_element := range matchs_element {
		event_name := "UCL"
		event_logo_string := "https://images.fotmob.com/image_resources/logo/leaguelogo/dark/42.png"

		time_text := match_element.MustElement(".css-hytar5-TimeCSS").MustText()
		re := regexp.MustCompile(`\d{1}:\d{2}`)
		time_text = re.FindString(time_text)

		team_1_element := match_element.MustElement(".css-1v13k0i-TeamBlockCSS")
		team_1_name_text := team_1_element.MustElement(".css-1wtw6ba-TeamName").MustText()
		team_1_logo := team_1_element.MustElement(".Image").MustAttribute("src")
		team_1_logo_string := *team_1_logo
		// fmt.Println(team_1_name_text, team_1_logo_string)

		team_2_element := match_element.MustElement(".css-so6otw-TeamBlockCSS")
		team_2_name_text := team_2_element.MustElement(".css-1wtw6ba-TeamName").MustText()
		team_2_logo := team_2_element.MustElement(".Image").MustAttribute("src")
		team_2_logo_string := *team_2_logo
		// fmt.Println(team_2_name_text, team_2_logo_string)

		// 印出比賽資訊
		fmt.Printf("Game - %s - %s - %s - %s vs %s \n", day_text, time_text, event_name, team_1_name_text, team_2_name_text)
		fmt.Println("Logo:", event_logo_string, team_1_logo_string, team_2_logo_string)

		// // 定義日期與時間格式
		// dateLayout := "Monday, January 2"
		// timeLayout := "15:04"

		// // 解析日期
		// parsedDate, _ := time.Parse(dateLayout, day_text)
		// // if err != nil {
		// // 	fmt.Println("Error parsing date:", err)
		// // 	return
		// // }

		// // 假設當前年份
		// currentYear := time.Now().Year()

		// // 解析時間
		// parsedTime, _ := time.Parse(timeLayout, time_text)
		// // if err != nil {
		// // 	fmt.Println("Error parsing time:", err)
		// // 	return
		// // }

		// // 合併日期與時間
		// finalTime := time.Date(currentYear, parsedDate.Month(), parsedDate.Day(),
		// 	parsedTime.Hour(), parsedTime.Minute(), 0, 0, time.Local)
		// fmt.Println(finalTime)
	}
}

func CSRod() {
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
			// 取得比賽名稱
			event := match.MustElement(".matchEvent")
			event_name := event.MustElement(".matchEventName").MustText()
			event_logo := event.MustElement(".matchEventLogoContainer").MustElement(".matchEventLogo").MustAttribute("src")
			event_logo_string := *event_logo
			event_logo_url := strings.ReplaceAll(event_logo_string, "\\u0026", "&")
			fmt.Println(event_logo_url)

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

			// 印出比賽資訊
			fmt.Printf("Game - %s - %s - %s - %s vs %s \n", day_text, time_text, event_name, team_1_name, team_2_name)
		}
	}
}
