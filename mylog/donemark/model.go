package main

import (
	"slices"
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

	if len(name) > 30 {
		name = name[:30]
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

func oldsortProblem(problems []Problem) {
	pCmp := func(a, b Problem) int {
		c1 := NewCatagory(a.catagory)
		c2 := NewCatagory(b.catagory)

		if c1 != c2 {

			if c1 == CATAGORY_UNKONWN || c2 == CATAGORY_UNKONWN {
				// 按照名称排序，需要手动维护，比如 1.sort 2.tree 3.list
				return strings.Compare(a.catagory, b.catagory)
			} else {
				return int(c1 - c2)
			}
		}

		s1 := NewStatus(a.status)
		s2 := NewStatus(b.status)

		if s1 != s2 {
			return int(s1 - s2)
		}

		r1 := NewRecommend(a.recommend)
		r2 := NewRecommend(b.recommend)
		if r1 != r2 {
			return int(r1 - r2)
		}

		d1 := NewDifficulty(a.difficulty)
		d2 := NewDifficulty(b.difficulty)

		if d1 != d2 {
			return int(d1 - d2)
		}

		return strings.Compare(a.name, b.name)

	}

	slices.SortFunc(problems, pCmp)

}

// 定义状态枚举
type Status int

const (
	STATUS_PLAN Status = iota
	STATUS_DOING
	STATUS_DONE
	STATUS_UNKONWN
)

var statusMap = map[Status]string{
	STATUS_PLAN:    "plan",
	STATUS_DOING:   "doing",
	STATUS_DONE:    "done",
	STATUS_UNKONWN: "unknown",
}

func (s Status) String() string {
	return statusMap[s]
}

func NewStatus(s string) Status {
	for k, v := range statusMap {
		if v == s {
			return k
		}
	}
	return STATUS_UNKONWN
}

// 定义 catagory 枚举
type Catagory int

const (
	CATAGORY_ARRAY Catagory = iota
	CATAGORY_STRING
	CATAGORY_TREE
	CATAGORY_UNKONWN
)

var catagoryMap = map[Catagory]string{
	CATAGORY_ARRAY:   "array",
	CATAGORY_STRING:  "string",
	CATAGORY_TREE:    "tree",
	CATAGORY_UNKONWN: "unkonwn",
}

func (c Catagory) String() string {
	return catagoryMap[c]
}

func NewCatagory(s string) Catagory {
	for k, v := range catagoryMap {
		if v == s {
			return k
		}
	}
	return CATAGORY_UNKONWN
}

// 定义 difficulty 枚举
type Difficulty int

const (
	DIFFICULTY_EASY Difficulty = iota
	DIFFICULTY_MID
	DIFFICULTY_HARD
	DIFFICULTY_UNKONWN
)

var difficultyMap = map[Difficulty]string{
	DIFFICULTY_EASY:    "easy",
	DIFFICULTY_MID:     "mid",
	DIFFICULTY_HARD:    "hard",
	DIFFICULTY_UNKONWN: "unkonwn",
}

func (d Difficulty) String() string {
	return difficultyMap[d]
}

func NewDifficulty(s string) Difficulty {
	for k, v := range difficultyMap {
		if v == s {
			return k
		}
	}
	return DIFFICULTY_UNKONWN
}

// 定义 Recommend  掌握程度
type Recommend int

const (
	RECOMMEND_BASIC     Recommend = iota // 基础内容，简单必要
	RECOMMEND_NEED                       // 必要但是全面，
	RECOMMEND_SKILL                      // 需要熟练掌握
	RECOMMEND_CHALLENGE                  // 挑战
	RECOMMEND_UNKONWN
)

var recommendMap = map[Recommend]string{
	RECOMMEND_BASIC:     "basic",
	RECOMMEND_NEED:      "need",
	RECOMMEND_SKILL:     "skill",
	RECOMMEND_CHALLENGE: "challenge",
	RECOMMEND_UNKONWN:   "unkonwn",
}

func (r Recommend) String() string {
	return recommendMap[r]
}

func NewRecommend(s string) Recommend {
	for k, v := range recommendMap {
		if v == s {
			return k
		}
	}
	return RECOMMEND_UNKONWN
}
