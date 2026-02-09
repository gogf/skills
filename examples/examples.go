package examples

import (
	`fmt`

	`github.com/gogf/gf/v2/container/garray`
	`github.com/gogf/gf/v2/encoding/gyaml`
	`github.com/gogf/gf/v2/os/gfile`
	`github.com/gogf/gf/v2/text/gregex`
	`github.com/gogf/gf/v2/text/gstr`
	`github.com/gogf/gf/v2/util/gconv`
)

type Example struct {
	NewAbsPath   string
	OldAbsPath   string
	RelativePath string
	Description  string
}

// Manager handles the generation and management of skill references.
type Manager struct {
	gfExamplesPath      string
	skillName           string
	skillsDirPath       string
	examplesInSkillPath string
}

const (
	notSkillReference      = `notSkillReference`
	minLengthOfMainContent = 200
	frontMatterPattern     = `^---\n([\s\S]+?)\n---`
)

// New creates a new Manager instance for handling skill references.
func New(gfExamplesPath, skillName, skillsDirPath string) *Manager {
	return &Manager{
		gfExamplesPath:      gfExamplesPath,
		skillName:           skillName,
		skillsDirPath:       skillsDirPath,
		examplesInSkillPath: gfile.Join(skillsDirPath, skillName, "examples"),
	}
}

func (m *Manager) SaveExamples(examples *garray.TArray[Example]) {
	// Clear existing references directory.
	err := gfile.RemoveAll(m.examplesInSkillPath)
	if err != nil {
		panic(err)
	}
	// Save each reference to its new path.
	for _, example := range examples.Slice() {
		if gfile.Exists(example.NewAbsPath) {
			continue
		}
		err = gfile.CopyDir(example.OldAbsPath, example.NewAbsPath)
		if err != nil {
			panic(err)
		}
	}
}

func (m *Manager) GenExampleTableContent(examples *garray.TArray[Example]) string {
	tableContent := "| 示例代码 | 示例介绍 |\n| --- | --- |\n"
	for _, ref := range examples.Slice() {
		referenceName := gstr.TrimLeft(ref.RelativePath, "examples/")
		referencePath := gstr.Replace(ref.RelativePath, " ", "%20")
		tableContent += fmt.Sprintf(
			"| [%s](%s) | %s |\n",
			referenceName, gfile.Join("examples", referencePath), ref.Description,
		)
	}
	return tableContent
}

func (m *Manager) GenExamples() *garray.TArray[Example] {
	examples := garray.NewTArray[Example]()
	files, err := gfile.ScanDirFile(m.gfExamplesPath, "*", true)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		extName := gfile.ExtName(file)
		if !gstr.InArray([]string{"md", "MD"}, extName) {
			continue
		}
		// Only include files with "ZH" in the filename.
		if !gstr.Contains(gfile.Basename(file), "ZH") {
			continue
		}

		if !m.isDirectExampleDirPath(gfile.Dir(file)) {
			continue
		}

		var (
			relativePath     = gfile.Dir(gstr.Replace(file, m.gfExamplesPath, ""))
			newSourceDirPath = gfile.Join(m.examplesInSkillPath, relativePath)
			readMeContent    = gfile.GetContents(file)
			frontMatter      = m.getFrontMatter(readMeContent)
		)
		// Skip if marked as notSkillReference in front matter.
		if gconv.Bool(frontMatter[notSkillReference]) {
			continue
		}
		examples.Append(Example{
			NewAbsPath:   newSourceDirPath,
			OldAbsPath:   gfile.Dir(file),
			RelativePath: relativePath,
			Description:  gconv.String(frontMatter["description"]),
		})
	}
	return examples
}

func (m *Manager) isDirectExampleDirPath(dirPath string) bool {
	var (
		hasGoFile bool
		hasGoMod  bool
	)
	files, err := gfile.ScanDirFile(dirPath, "*", false)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if gfile.ExtName(file) == "go" {
			hasGoFile = true
		}
		if gfile.Basename(file) == "go.mod" {
			hasGoMod = true
		}
	}
	return hasGoFile || hasGoMod
}

// getFrontMatter extracts the front matter from the markdown file content.
func (m *Manager) getFrontMatter(content string) map[string]any {
	content = gstr.TrimLeft(content)
	match, err := gregex.MatchString(frontMatterPattern, content)
	if err != nil {
		panic(err)
	}
	if len(match) < 2 {
		return map[string]any{}
	}
	result, err := gyaml.Decode([]byte(match[1]))
	if err != nil {
		panic(err)
	}
	return result
}
