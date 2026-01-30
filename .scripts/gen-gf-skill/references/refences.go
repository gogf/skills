package references

import (
	`fmt`

	`github.com/gogf/gf/v2/container/garray`
	`github.com/gogf/gf/v2/encoding/gyaml`
	`github.com/gogf/gf/v2/os/gfile`
	`github.com/gogf/gf/v2/text/gregex`
	`github.com/gogf/gf/v2/text/gstr`
	`github.com/gogf/gf/v2/util/gconv`
)

type Reference struct {
	NewAbsPath   string // New absolute path to save the reference file.
	RelativePath string // Relative path from the original document folder.
	Description  string // Description extracted from front matter.
	Content      string // Main content of the reference.
}

type Manager struct {
	gfDocsPath     string
	skillName      string
	skillsDirPath  string
	referencesPath string
}

const (
	notSkillReference      = `notSkillReference`
	minLengthOfMainContent = 200
	frontMatterPattern     = `^---\n([\s\S]+?)\n---`
)

func New(gfDocsPath, skillName, skillsDirPath string) *Manager {
	return &Manager{
		gfDocsPath:     gfDocsPath,
		skillName:      skillName,
		skillsDirPath:  skillsDirPath,
		referencesPath: gfile.Join(skillsDirPath, skillName, "references"),
	}
}

// SaveReferences saves the generated references to their new absolute paths.
func (m *Manager) SaveReferences(references *garray.TArray[Reference]) {
	// Clear existing references directory.
	err := gfile.RemoveAll(m.referencesPath)
	if err != nil {
		panic(err)
	}
	// Save each reference to its new path.
	for _, ref := range references.Slice() {
		err = gfile.PutContents(ref.NewAbsPath, ref.Content)
		if err != nil {
			panic(err)
		}
	}
}

// GenReferencesTableContent generates the markdown table content for the references.
func (m *Manager) GenReferencesTableContent(references *garray.TArray[Reference]) string {
	tableContent := "| 参考资料 | 资料介绍 |\n| --- | --- |\n"
	for _, ref := range references.Slice() {
		referenceName := gstr.TrimLeft(ref.RelativePath, "references/")
		referenceName = gstr.TrimRight(referenceName, gfile.Ext(referenceName))
		referencePath := gstr.Replace(ref.RelativePath, " ", "%20")
		tableContent += fmt.Sprintf(
			"| [%s](%s) | %s |\n",
			referenceName, referencePath, ref.Description,
		)
	}
	return tableContent
}

// GenReferences generates references from markdown files in the specified document folder.
func (m *Manager) GenReferences(docFolder string) *garray.TArray[Reference] {
	references := garray.NewTArray[Reference]()
	docsPath := m.gfDocsPath + docFolder + "/"
	files, err := gfile.ScanDirFile(docsPath, "*", true)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		extName := gfile.ExtName(file)
		if !gstr.InArray([]string{"md", "MD"}, extName) {
			continue
		}
		var (
			relativePath = gstr.Replace(file, docsPath, "")
			newFilePath  = gfile.Join(m.referencesPath, relativePath)
			content      = gfile.GetContents(file)
			frontMatter  = m.getFrontMatter(content)
			mainContent  = m.getMainContent(content)
		)
		// Skip if marked as notSkillReference in front matter.
		if gconv.Bool(frontMatter[notSkillReference]) {
			continue
		}
		// Skip if main content is too short.
		if len(mainContent) < minLengthOfMainContent {
			continue
		}
		// Skip if main content has too few lines.
		if gstr.Count(mainContent, "\n") <= 10 {
			continue
		}
		references.Append(Reference{
			NewAbsPath:   newFilePath,
			RelativePath: fmt.Sprintf(`references/%s`, relativePath),
			Description:  gconv.String(frontMatter["description"]),
			Content:      mainContent,
		})
	}
	return references
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

// getMainContent extracts the main content from the markdown file,
func (m *Manager) getMainContent(content string) string {
	mainContent, _ := gregex.ReplaceString(frontMatterPattern, "", content)
	// remove image references like ![alt text](image_url)
	mainContent, _ = gregex.ReplaceString(`!\[.*?\]\(.*?\)`, "", mainContent)

	// remove import DocCardList from '@theme/DocCardList';
	// remove <DocCardList />
	mainContent = gstr.Replace(mainContent, `import DocCardList from '@theme/DocCardList';`, "")
	mainContent = gstr.Replace(mainContent, `<DocCardList />`, "")

	// remove multiple consecutive new lines
	mainContent, _ = gregex.ReplaceString(`[\n\r]{3,}`, "\n\n", mainContent)
	mainContent = gstr.Trim(mainContent)
	return mainContent
}
