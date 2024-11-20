package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateMd(t *testing.T) {
	// 创建临时输出目录
	outputDir := "tmp/output"
	os.MkdirAll(outputDir, 0755)
	defer os.RemoveAll(outputDir)

	// 设置配置
	config["output_dir"] = outputDir
	config["output_file"] = "output.md"

	// 设置问题列表
	problems = []Problem{
		{"p1", "tree", "done", "need", "mid", "2024-11-20"},
		{"p2", "array", "done", "need", "mid", "2024-11-20"},
		{"p3", "array", "doing", "need", "mid", "2024-11-20"},
		{"p4", "array", "doing", "need", "mid", "2024-11-20"},
		{"p5", "array", "doing", "basic", "mid", "2024-11-20"},
		{"p7", "array", "doing", "basic", "easy", "2024-11-20"},
		{"p6", "array", "doing", "basic", "easy", "2024-11-20"},
	}

	sortProblem(problems)

	// 调用 generateMd 函数
	err := generateMd()
	if err != nil {
		t.Errorf("generateMd failed: %v", err)
	}

	// 读取生成的 Markdown 文件
	outputFilePath := filepath.Join(outputDir, "output.md")
	data, err := os.ReadFile(outputFilePath)
	if err != nil {
		t.Errorf("Failed to read output file: %v", err)
	}

	expectedContent := `# 进展

|catagory            |status              |recommend           |difficulty          |name                |lastupdate          |
|--------------------|--------------------|--------------------|--------------------|--------------------|--------------------|
|array               |doing               |basic               |easy                |p6                  |2024-11-20          |
|array               |doing               |basic               |easy                |p7                  |2024-11-20          |
|array               |doing               |basic               |mid                 |p5                  |2024-11-20          |
|array               |doing               |need                |mid                 |p3                  |2024-11-20          |
|array               |doing               |need                |mid                 |p4                  |2024-11-20          |
|array               |done                |need                |mid                 |p2                  |2024-11-20          |
|tree                |done                |need                |mid                 |p1                  |2024-11-20          |
`
	if string(data) != expectedContent {
		t.Errorf("Generated content mismatch: expected %q, got %q", expectedContent, string(data))
	}
}
