package main

import (
	`fmt`

	`main/examples`
	`main/references`

	`github.com/gogf/gf/v2/container/garray`
	`github.com/gogf/gf/v2/os/gfile`
	`github.com/gogf/gf/v2/text/gstr`
)

const (
	skillName      = "goframe-v2"
	gfDocsPath     = `/Users/john/Workspace/github/gogf/gf-site/docs/`
	gfExamplesPath = `/Users/john/Workspace/github/gogf/examples/`
	skillsDirPath  = `/Users/john/Workspace/github/gogf/skills/.claude/skills`
	skillTplPath   = `SKILL.tpl.md`
	skillVersion   = "v0.0.1"
)

type SkillTplData struct {
	SkillName       string
	SkillReferences string
	SkillExamples   string
}

var (
	docFolders = []string{"docs", "quick"}
)

func main() {
	var (
		refsManager     = references.New(gfDocsPath, skillName, skillsDirPath)
		examplesManager = examples.New(gfExamplesPath, skillName, skillsDirPath)
	)
	// Generate references from document folders.
	var (
		refsArray            = garray.NewTArray[references.Reference]()
		examplesArray        = examplesManager.GenExamples()
		examplesTableContent = examplesManager.GenExampleTableContent(examplesArray)
	)
	for _, docFolder := range docFolders {
		array := refsManager.GenReferences(docFolder)
		refsArray.Append(array.Slice()...)
	}
	// Save references to new paths.
	refsManager.SaveReferences(refsArray)
	examplesManager.SaveExamples(examplesArray)

	updateSkillFileContent(SkillTplData{
		SkillName:       skillName,
		SkillReferences: refsManager.GenReferencesTableContent(refsArray),
		SkillExamples:   examplesTableContent,
	})

	fmt.Println("Done!")
}

// updateSkillFileContent updates or generates the SKILL.md file for the skill.
func updateSkillFileContent(data SkillTplData) {
	var (
		skillTplContent = gfile.GetContents(skillTplPath)
		skillFilePath   = gfile.Join(skillsDirPath, skillName, "SKILL.md")
	)
	skillContent := gstr.ReplaceByMap(skillTplContent, map[string]string{
		"{{SKILL_NAME}}":    data.SkillName,
		"{{SKILL_VERSION}}": skillVersion,
	})
	err := gfile.PutContents(skillFilePath, skillContent)
	if err != nil {
		panic(err)
	}

	createReferencesIndexFile(data.SkillReferences)

	createExamplesIndexFile(data.SkillExamples)
}

// createReferencesIndexFile creates or updates the README.md file in the references directory
// with the generated references table content.
func createReferencesIndexFile(referencesIndexContent string) {
	indexFilePath := gfile.Join(skillsDirPath, skillName, "references", "README.MD")
	err := gfile.PutContents(indexFilePath, referencesIndexContent)
	if err != nil {
		panic(err)
	}
}

// createExamplesIndexFile creates or updates the README.md file in the examples directory
// with the generated examples table content.
func createExamplesIndexFile(examplesIndexContent string) {
	indexFilePath := gfile.Join(skillsDirPath, skillName, "examples", "README.MD")
	err := gfile.PutContents(indexFilePath, examplesIndexContent)
	if err != nil {
		panic(err)
	}
}
