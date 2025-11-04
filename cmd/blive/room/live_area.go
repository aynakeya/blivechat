package room

import (
	"errors"
	"github.com/charmbracelet/huh"
	"github.com/tidwall/gjson"
	"sort"
	"strconv"
)

func getAreas() (map[string]map[string]int64, error) {
	resp, err := client.R().Get("https://api.live.bilibili.com/xlive/app-blink/v1/preLive/GetAreaListForLive?show_pinyin=1&platform=pc")
	if err != nil {
		return nil, err
	}
	result := make(map[string]map[string]int64)
	data := gjson.ParseBytes(resp.Bytes())
	for _, key := range data.Get("data.area_v1_info").Array() {
		result[key.Get("name").String()] = make(map[string]int64)
		for _, sub := range key.Get("list").Array() {
			id, _ := strconv.ParseInt(sub.Get("id").String(), 10, 64)
			result[key.Get("name").String()][sub.Get("name").String()] = id
		}
	}
	return result, nil
}

func runAreaUI() (int64, error) {
	area, err := getAreas()
	if err != nil {
		return 0, err
	}
	if len(area) == 0 {
		return 0, errors.New("no areas available")
	}

	v1Names := make([]string, 0, len(area))
	for k := range area {
		v1Names = append(v1Names, k)
	}
	sort.Strings(v1Names)

	var v1Choice string
	var v2Choice string
	v1Choice = v1Names[0]

	v2ListOf := func(v1 string) []string {
		m := area[v1]
		names := make([]string, 0, len(m))
		for k := range m {
			names = append(names, k)
		}
		sort.Strings(names)
		return names
	}

	v2Names := v2ListOf(v1Choice)
	if len(v2Names) == 0 {
		return 0, errors.New("selected v1 has no sub areas")
	}
	v2Choice = v2Names[0]

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("选择一级分区").
				Value(&v1Choice).
				Options(func() []huh.Option[string] {
					opts := make([]huh.Option[string], 0, len(v1Names))
					for _, n := range v1Names {
						opts = append(opts, huh.NewOption(n, n))
					}
					return opts
				}()...).
				Height(4).WithTheme(huh.ThemeBase()),

			huh.NewSelect[string]().
				TitleFunc(func() string { return "选择二级分区" }, &v1Choice).
				Value(&v2Choice).
				OptionsFunc(func() []huh.Option[string] {
					v2Names = v2ListOf(v1Choice)
					if len(v2Names) == 0 {
						return []huh.Option[string]{}
					}
					v2Choice = v2Names[0]
					opts := make([]huh.Option[string], 0, len(v2Names))
					for _, n := range v2Names {
						opts = append(opts, huh.NewOption(n, n))
					}
					return opts
				}, &v1Choice).
				Height(4).WithTheme(huh.ThemeBase()),
		),
	)

	if err := form.Run(); err != nil {
		return 0, err
	}

	id, ok := area[v1Choice][v2Choice]
	if !ok {
		return 0, errors.New("invalid area selection")
	}
	return id, nil
}
