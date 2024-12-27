package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Problem struct {
	name       string
	alias      string
	catagory   string
	status     string
	recommend  string // 1~3
	difficulty string
	lastupdate string //2024-11-20
	path       string
}

func newProblem(name string, path string, desc string) *Problem {

	if len(name) > 200 {
		name = name[:200]
	}
	result := Problem{name: name, path: path}

	if desc == "" {
		return &result
	}

	desc = strings.TrimSpace(desc)
	kvs := strings.Split(desc, ";")

	for _, item := range kvs {
		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			continue // 跳过格式不正确的项
		}
		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])

		switch key {
		case "alias":
			result.alias = value
		case "status":
			result.status = value
		case "catagory":
			result.catagory = value
		case "recommend":
			result.recommend = value
		case "difficulty":
			result.difficulty = value
		case "lastupdate":
			result.lastupdate = value
		}
	}

	return &result
}
func sortProblem(problems []Problem) {
	pCmp := func(a, b Problem) int {

		if a.catagory != b.catagory {
			return strings.Compare(a.catagory, b.catagory)
		}

		if a.status != b.status {
			return strings.Compare(a.status, b.status)
		}

		if a.recommend != b.recommend {
			return strings.Compare(a.recommend, b.recommend)
		}

		if a.difficulty != b.difficulty {
			return strings.Compare(a.difficulty, b.difficulty)
		}

		return strings.Compare(a.name, b.name)
	}

	slices.SortFunc(problems, pCmp)

}

func (p *Problem) ColorRender() {
	// 标记  status
	if strings.HasPrefix(p.status, "01") {
		p.recommend = colorRender(p.recommend)
		p.difficulty = colorRender(p.difficulty)
	}
	p.status = colorRender(p.status)

	if len(p.catagory) > 1 {
		num, ok := strconv.Atoi(p.catagory[:2])
		if ok != nil {
			return
		}
		if num%2 == 0 {
			p.catagory = fmt.Sprintf("<span style=\"color:#FEFFA7;\">%s</span>", p.catagory)
		} else {
			p.catagory = fmt.Sprintf("<span style=\"color:#D4F6FF;\">%s</span>", p.catagory)
		}

	}

}

func colorRender(input string) string {
	switch {
	case strings.HasPrefix(input, "01"):

		input = fmt.Sprintf("<span style=\"color:#00FF9C;\">%s</span>", input)
	case strings.HasPrefix(input, "02"):

		input = fmt.Sprintf("<span style=\"color:#FEEE91;\">%s</span>", input)
	case strings.HasPrefix(input, "03"):

		input = fmt.Sprintf("<span style=\"color:#FF9D3D;\">%s</span>", input)
	case strings.HasPrefix(input, "04"):

		input = fmt.Sprintf("<span style=\"color:#FF7D29;\">%s</span>", input)
	case strings.HasPrefix(input, "05"):

		input = fmt.Sprintf("<span style=\"color:#FA4032;\">%s</span>", input)
	}
	return input
}
