package main

import (
	`main/references`

	`github.com/gogf/gf/v2/container/garray`
	`github.com/gogf/gf/v2/os/gfile`
	`github.com/gogf/gf/v2/text/gstr`
)

const (
	skillName      = "gf-master"
	gfDocsPath     = `/Users/john/Workspace/github/gogf/gf-site/docs/`
	gfExamplesPath = `/Users/john/Workspace/github/gogf/examples`
	skillsDirPath  = `/Users/john/Workspace/github/gogf/skills/.github/skills`
	skillTplPath   = `SKILL.tpl.md`
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
		refsManager = references.New(gfDocsPath, skillName, skillsDirPath)
	)
	// Generate references from document folders.
	refsArray := garray.NewTArray[references.Reference]()
	for _, docFolder := range docFolders {
		array := refsManager.GenReferences(docFolder)
		refsArray.Append(array.Slice()...)
	}
	// Save references to new paths.
	refsManager.SaveReferences(refsArray)
	updateSkillFileContent(SkillTplData{
		SkillName:       skillName,
		SkillReferences: refsManager.GenReferencesTableContent(refsArray),
		SkillExamples:   "",
	})
}

// updateSkillFileContent updates or generates the SKILL.md file for the skill.
func updateSkillFileContent(data SkillTplData) {
	var (
		skillTplContent = gfile.GetContents(skillTplPath)
		skillFilePath   = gfile.Join(skillsDirPath, skillName, "SKILL.md")
	)
	skillContent := gstr.ReplaceByMap(skillTplContent, map[string]string{
		"{{SKILL_NAME}}":       data.SkillName,
		"{{SKILL_REFERENCES}}": data.SkillReferences,
		"{{SKILL_EXAMPLES}}":   data.SkillExamples,
	})
	err := gfile.PutContents(skillFilePath, skillContent)
	if err != nil {
		panic(err)
	}
}
