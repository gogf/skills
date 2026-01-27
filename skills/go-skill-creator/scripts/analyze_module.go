// Package main provides a tool to analyze Go modules and generate skill documentation.
// Usage: go run analyze_module.go [options] [module_path]
//
// This tool uses `go list -json` to analyze package structure and generate
// comprehensive skill documentation for AI assistants.
//
// Features:
// - Automatically finds all go.mod files in the project (supports multi-module projects)
// - Uses `go list -json` for accurate package information
// - Parses Go AST for type and function extraction
// - No depth limit by default (scans all subdirectories)
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

// GoPackage represents the JSON output from `go list -json`
type GoPackage struct {
	Dir          string   `json:"Dir"`
	ImportPath   string   `json:"ImportPath"`
	Name         string   `json:"Name"`
	Doc          string   `json:"Doc"`
	Root         string   `json:"Root"`
	GoFiles      []string `json:"GoFiles"`
	Imports      []string `json:"Imports"`
	Deps         []string `json:"Deps"`
	TestGoFiles  []string `json:"TestGoFiles"`
	XTestGoFiles []string `json:"XTestGoFiles"`
	Module       *struct {
		Path      string `json:"Path"`
		Main      bool   `json:"Main"`
		Dir       string `json:"Dir"`
		GoMod     string `json:"GoMod"`
		GoVersion string `json:"GoVersion"`
	} `json:"Module"`
}

// PackageInfo represents analyzed package information
type PackageInfo struct {
	ImportPath  string
	Name        string
	Doc         string
	Dir         string
	RelPath     string // Relative path within module
	ModulePath  string // Which module this package belongs to
	GoFiles     []string
	Imports     []string
	InternalDep []string // Internal module dependencies
	ExternalDep []string // External dependencies
	Types       []TypeInfo
	Functions   []FuncInfo
	IsInternal  bool // Is this an internal package
}

// TypeInfo represents a type definition
type TypeInfo struct {
	Name    string
	Doc     string
	Kind    string // struct, interface, alias
	Methods []FuncInfo
}

// FuncInfo represents a function or method
type FuncInfo struct {
	Name      string
	Doc       string
	Signature string
	Receiver  string // For methods
}

// ModuleInfo represents a Go module
type ModuleInfo struct {
	Path    string // Module path (e.g., github.com/gogf/gf/v2)
	Dir     string // Directory containing go.mod
	GoMod   string // Full path to go.mod
	RelPath string // Relative path from root project
}

// ModuleAnalyzer analyzes Go modules
type ModuleAnalyzer struct {
	RootDir      string                  // Root project directory
	Modules      []ModuleInfo            // All found modules
	MainModule   string                  // Main module path
	MainModuleVersion string              // Version of main module (e.g., v2, latest)
	Packages     map[string]*PackageInfo // All packages from all modules
	Categories   map[string][]string     // Category -> package paths
	SkillName    string
	SkillDir     string
	IncludeTests bool
	TempDir      string                  // Temporary directory for go mod operations
}

func main() {
	var (
		rootPath  string
		skillName string
		outputDir string
		moduleRef string // Module reference (e.g., github.com/gogf/gf/v2@latest)
		internal  bool
		verbose   bool
	)

	flag.StringVar(&rootPath, "root", ".", "Root project directory to analyze")
	flag.StringVar(&skillName, "name", "", "Skill name (default: extracted from main module name)")
	flag.StringVar(&outputDir, "output", "", "Output directory for skill (default: .codebuddy/skills/<name>-dev)")
	flag.StringVar(&moduleRef, "module", "", "Module reference with version (e.g., github.com/gogf/gf/v2@latest)")
	flag.BoolVar(&internal, "internal", false, "Include internal packages")
	flag.BoolVar(&verbose, "verbose", false, "Verbose output")

	flag.Parse()

	if flag.NArg() > 0 {
		rootPath = flag.Arg(0)
	}

	analyzer := &ModuleAnalyzer{
		Packages:   make(map[string]*PackageInfo),
		Categories: make(map[string][]string),
	}

	if err := analyzer.Analyze(rootPath, moduleRef, internal, verbose); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Clean up temporary directory if it was created
	defer func() {
		if analyzer.TempDir != "" && analyzer.TempDir != analyzer.RootDir {
			if err := os.RemoveAll(analyzer.TempDir); err != nil && verbose {
				fmt.Fprintf(os.Stderr, "Warning: failed to clean temp dir %s: %v\n", analyzer.TempDir, err)
			}
		}
	}()

	// Set skill name
	if skillName != "" {
		analyzer.SkillName = skillName
	} else {
		// Extract name from main module path
		parts := strings.Split(analyzer.MainModule, "/")
		analyzer.SkillName = parts[len(parts)-1]
		// Remove version suffix like v2, v3
		if strings.HasPrefix(analyzer.SkillName, "v") && len(analyzer.SkillName) <= 3 {
			if len(parts) > 1 {
				analyzer.SkillName = parts[len(parts)-2]
			}
		}
	}

	// Set output directory
	if outputDir != "" {
		analyzer.SkillDir = outputDir
	} else {
		analyzer.SkillDir = filepath.Join(analyzer.RootDir, ".codebuddy", "skills", analyzer.SkillName+"-skill")
	}

	// Generate skill
	if err := analyzer.GenerateSkill(verbose); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating skill: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n✅ Skill generated: %s\n", analyzer.SkillDir)
	fmt.Printf("   Modules found: %d\n", len(analyzer.Modules))
	fmt.Printf("   Packages analyzed: %d\n", len(analyzer.Packages))
	fmt.Printf("   Categories: %d\n", len(analyzer.Categories))
	if analyzer.MainModuleVersion != "" {
		fmt.Printf("   Main module version: %s\n", analyzer.MainModuleVersion)
	}
	
	fmt.Printf("\n📝 Required: Optimize the description\n")
	fmt.Printf("   The current description is auto-generated. You MUST optimize it according to the skill-creator standard.\n")
	fmt.Printf("   \n")
	fmt.Printf("   Description should include:\n")
	fmt.Printf("   1. Module functionality summary with common aliases/nicknames\n")
	fmt.Printf("   2. Complete Go import path (as unique identifier)\n")
	fmt.Printf("   3. Version (if applicable)\n")
	fmt.Printf("   4. 3-4 specific use cases/scenarios for when to use this Skill\n")
	fmt.Printf("   \n")
	fmt.Printf("   Steps:\n")
	fmt.Printf("   1. Read the generated SKILL.md frontmatter description\n")
	fmt.Printf("   2. Ask AI to optimize it per the standards (see go-skill-creator/SKILL.md for format)\n")
	fmt.Printf("   3. Replace the description in SKILL.md frontmatter with the optimized version\n")
	fmt.Printf("   \n")
	fmt.Printf("   Reference: .codebuddy/skills/go-skill-creator/SKILL.md (Description 优化流程)\n")
}

// Analyze analyzes all Go modules in the project
func (a *ModuleAnalyzer) Analyze(rootPath string, moduleRef string, includeInternal bool, verbose bool) error {
	// Resolve root directory
	absPath, err := filepath.Abs(rootPath)
	if err != nil {
		return fmt.Errorf("failed to resolve path: %w", err)
	}
	a.RootDir = absPath

	if verbose {
		fmt.Printf("Scanning for go.mod files in: %s\n", a.RootDir)
	}

	// If moduleRef is provided, use it to download and analyze the module
	if moduleRef != "" {
		if err := a.analyzeModuleRef(moduleRef, includeInternal, verbose); err != nil {
			return err
		}
		return nil
	}

	// Otherwise, find and analyze local modules
	// Find all go.mod files
	if err := a.findAllModules(verbose); err != nil {
		return fmt.Errorf("failed to find modules: %w", err)
	}

	if len(a.Modules) == 0 {
		return fmt.Errorf("no go.mod found in %s", absPath)
	}

	if verbose {
		fmt.Printf("Found %d module(s)\n", len(a.Modules))
		for _, mod := range a.Modules {
			fmt.Printf("  - %s (at %s)\n", mod.Path, mod.RelPath)
		}
	}

	// Set main module (the one at root or the first one)
	for _, mod := range a.Modules {
		if mod.RelPath == "." || mod.RelPath == "" {
			a.MainModule = mod.Path
			break
		}
	}
	if a.MainModule == "" && len(a.Modules) > 0 {
		a.MainModule = a.Modules[0].Path
	}

	// Analyze each module
	for _, mod := range a.Modules {
		if err := a.analyzeModule(mod, includeInternal, verbose); err != nil {
			if verbose {
				fmt.Printf("Warning: failed to analyze module %s: %v\n", mod.Path, err)
			}
			continue
		}
	}

	return nil
}

// analyzeModuleRef analyzes a Go module specified by reference (e.g., github.com/gogf/gf/v2@latest)
func (a *ModuleAnalyzer) analyzeModuleRef(moduleRef string, includeInternal bool, verbose bool) error {
	// Parse module reference
	modulePath := moduleRef
	version := "latest"
	
	if idx := strings.LastIndex(moduleRef, "@"); idx != -1 {
		modulePath = moduleRef[:idx]
		version = moduleRef[idx+1:]
	}
	
	if verbose {
		fmt.Printf("Analyzing module: %s (version: %s)\n", modulePath, version)
	}
	
	// Create temporary directory for go mod operations
	tempDir, err := os.MkdirTemp("", "go-skill-analyzer-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	a.TempDir = tempDir
	
	if verbose {
		fmt.Printf("Using temporary directory: %s\n", tempDir)
	}
	
	// Create a temporary go.mod file
	goModContent := fmt.Sprintf("module temp\n\ngo 1.21\n\nrequire %s %s\n", modulePath, version)
	goModPath := filepath.Join(tempDir, "go.mod")
	if err := os.WriteFile(goModPath, []byte(goModContent), 0644); err != nil {
		return fmt.Errorf("failed to create temporary go.mod: %w", err)
	}
	
	// Run go mod download to fetch the module
	cmd := exec.Command("go", "mod", "download")
	cmd.Dir = tempDir
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to download module: %w\n%s", err, string(output))
	}
	
	if verbose {
		fmt.Printf("Module downloaded successfully\n")
	}
	
	// Get module information
	cmd = exec.Command("go", "list", "-m", "-json", moduleRef)
	cmd.Dir = tempDir
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get module info: %w", err)
	}
	
	var modInfo struct {
		Path      string `json:"Path"`
		Version   string `json:"Version"`
		Dir       string `json:"Dir"`
		GoVersion string `json:"GoVersion"`
	}
	if err := json.Unmarshal(output, &modInfo); err != nil {
		return fmt.Errorf("failed to parse module info: %w", err)
	}
	
	a.MainModule = modInfo.Path
	a.MainModuleVersion = modInfo.Version
	
	if verbose {
		fmt.Printf("Module path: %s, version: %s, dir: %s\n", modInfo.Path, modInfo.Version, modInfo.Dir)
	}
	
	// Create ModuleInfo for the downloaded module
	a.Modules = []ModuleInfo{
		{
			Path:    modInfo.Path,
			Dir:     modInfo.Dir,
			GoMod:   filepath.Join(modInfo.Dir, "go.mod"),
			RelPath: ".",
		},
	}
	
	// Analyze the module
	if err := a.analyzeModule(a.Modules[0], includeInternal, verbose); err != nil {
		return fmt.Errorf("failed to analyze module: %w", err)
	}
	
	return nil
}

// findAllModules finds all go.mod files in the project
func (a *ModuleAnalyzer) findAllModules(verbose bool) error {
	return filepath.Walk(a.RootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip inaccessible paths
		}

		// Skip hidden directories and common non-source directories
		if info.IsDir() {
			name := info.Name()
			if strings.HasPrefix(name, ".") || name == "vendor" || name == "node_modules" || name == "testdata" {
				return filepath.SkipDir
			}
			return nil
		}

		// Check for go.mod
		if info.Name() == "go.mod" {
			modDir := filepath.Dir(path)
			relPath, _ := filepath.Rel(a.RootDir, modDir)
			if relPath == "" {
				relPath = "."
			}

			// Get module path from go.mod
			modPath, err := getModulePath(modDir)
			if err != nil {
				if verbose {
					fmt.Printf("Warning: failed to get module path from %s: %v\n", path, err)
				}
				return nil
			}

			a.Modules = append(a.Modules, ModuleInfo{
				Path:    modPath,
				Dir:     modDir,
				GoMod:   path,
				RelPath: relPath,
			})
		}

		return nil
	})
}

// getModulePath extracts module path from go.mod using go list
func getModulePath(dir string) (string, error) {
	cmd := exec.Command("go", "list", "-m", "-json")
	cmd.Dir = dir
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	var modInfo struct {
		Path string `json:"Path"`
	}
	if err := json.Unmarshal(output, &modInfo); err != nil {
		return "", err
	}
	return modInfo.Path, nil
}

// analyzeModule analyzes a single Go module
func (a *ModuleAnalyzer) analyzeModule(mod ModuleInfo, includeInternal bool, verbose bool) error {
	if verbose {
		fmt.Printf("\nAnalyzing module: %s\n", mod.Path)
	}

	// List all packages in this module
	cmd := exec.Command("go", "list", "-json", "./...")
	cmd.Dir = mod.Dir
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to list packages: %w", err)
	}

	// Parse JSON stream
	decoder := json.NewDecoder(strings.NewReader(string(output)))
	for decoder.More() {
		var pkg GoPackage
		if err := decoder.Decode(&pkg); err != nil {
			continue
		}

		// Skip test packages
		if strings.HasSuffix(pkg.Name, "_test") {
			continue
		}

		// Skip main packages (usually cmd/)
		if pkg.Name == "main" {
			continue
		}

		// Calculate relative path within the module
		relPath := strings.TrimPrefix(pkg.ImportPath, mod.Path)
		relPath = strings.TrimPrefix(relPath, "/")

		// For sub-modules, include the module's relative path in category
		fullRelPath := relPath
		if mod.RelPath != "." && mod.RelPath != "" {
			if relPath == "" {
				fullRelPath = mod.RelPath
			} else {
				fullRelPath = mod.RelPath + "/" + relPath
			}
		}

		// Check internal
		isInternal := strings.Contains(fullRelPath, "internal/") || strings.HasPrefix(fullRelPath, "internal") ||
			strings.Contains(fullRelPath, "/internal")
		if isInternal && !includeInternal {
			continue
		}

		// Create package info
		pkgInfo := &PackageInfo{
			ImportPath: pkg.ImportPath,
			Name:       pkg.Name,
			Doc:        pkg.Doc,
			Dir:        pkg.Dir,
			RelPath:    fullRelPath,
			ModulePath: mod.Path,
			GoFiles:    pkg.GoFiles,
			Imports:    pkg.Imports,
			IsInternal: isInternal,
		}

		// Separate internal and external dependencies
		for _, imp := range pkg.Imports {
			// Check if it's from any of our modules
			isInternalDep := false
			for _, m := range a.Modules {
				if strings.HasPrefix(imp, m.Path) {
					isInternalDep = true
					break
				}
			}
			if isInternalDep {
				pkgInfo.InternalDep = append(pkgInfo.InternalDep, imp)
			} else if !isStdLib(imp) {
				pkgInfo.ExternalDep = append(pkgInfo.ExternalDep, imp)
			}
		}

		// Parse source files to extract types and functions
		if err := a.parsePackageSource(pkgInfo, verbose); err != nil {
			if verbose {
				fmt.Printf("Warning: failed to parse %s: %v\n", pkg.ImportPath, err)
			}
		}

		a.Packages[pkg.ImportPath] = pkgInfo

		// Categorize by top-level directory
		category := "root"
		if fullRelPath != "" {
			parts := strings.Split(fullRelPath, "/")
			category = parts[0]
		}
		a.Categories[category] = append(a.Categories[category], pkg.ImportPath)

		if verbose {
			fmt.Printf("  - %s (%d types, %d funcs)\n", fullRelPath, len(pkgInfo.Types), len(pkgInfo.Functions))
		}
	}

	return nil
}

// parsePackageSource parses Go source files to extract types and functions
func (a *ModuleAnalyzer) parsePackageSource(pkg *PackageInfo, verbose bool) error {
	fset := token.NewFileSet()

	for _, filename := range pkg.GoFiles {
		filePath := filepath.Join(pkg.Dir, filename)
		file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
		if err != nil {
			continue
		}

		for _, decl := range file.Decls {
			switch d := decl.(type) {
			case *ast.GenDecl:
				if d.Tok == token.TYPE {
					for _, spec := range d.Specs {
						typeSpec := spec.(*ast.TypeSpec)
						if !typeSpec.Name.IsExported() {
							continue
						}

						typeInfo := TypeInfo{
							Name: typeSpec.Name.Name,
						}

						// Get doc
						if d.Doc != nil {
							typeInfo.Doc = d.Doc.Text()
						}

						// Determine kind
						switch typeSpec.Type.(type) {
						case *ast.StructType:
							typeInfo.Kind = "struct"
						case *ast.InterfaceType:
							typeInfo.Kind = "interface"
						default:
							typeInfo.Kind = "type"
						}

						pkg.Types = append(pkg.Types, typeInfo)
					}
				}

			case *ast.FuncDecl:
				if !d.Name.IsExported() {
					continue
				}

				funcInfo := FuncInfo{
					Name: d.Name.Name,
				}

				// Get doc
				if d.Doc != nil {
					funcInfo.Doc = strings.TrimSpace(d.Doc.Text())
				}

				// Get receiver
				if d.Recv != nil && len(d.Recv.List) > 0 {
					recv := d.Recv.List[0]
					funcInfo.Receiver = exprToString(recv.Type)
				}

				// Build signature
				funcInfo.Signature = buildFuncSignature(d)

				if funcInfo.Receiver != "" {
					// Find and attach to type
					receiverName := strings.TrimPrefix(funcInfo.Receiver, "*")
					for i := range pkg.Types {
						if pkg.Types[i].Name == receiverName {
							pkg.Types[i].Methods = append(pkg.Types[i].Methods, funcInfo)
							break
						}
					}
				} else {
					pkg.Functions = append(pkg.Functions, funcInfo)
				}
			}
		}
	}

	return nil
}

// GenerateSkill generates the skill documentation
func (a *ModuleAnalyzer) GenerateSkill(verbose bool) error {
	// Create directories
	refsDir := filepath.Join(a.SkillDir, "references")
	if err := os.MkdirAll(refsDir, 0755); err != nil {
		return err
	}

	// Generate SKILL.md
	if err := a.generateSkillMD(verbose); err != nil {
		return err
	}

	// Generate reference files by category
	if err := a.generateReferences(verbose); err != nil {
		return err
	}

	return nil
}

// generateSkillMD generates the main SKILL.md file
func (a *ModuleAnalyzer) generateSkillMD(verbose bool) error {
	var sb strings.Builder

	// Frontmatter
	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("name: %s-skill\n", a.SkillName))
	
	// Build description with full module path and alias
	description := fmt.Sprintf("Go 模块 `%s` 的开发指南。", a.MainModule)
	if a.MainModuleVersion != "" {
		description += fmt.Sprintf("版本: %s。", a.MainModuleVersion)
	}
	description += "基于源码深度分析，提供完整的包结构和 API 参考。"
	sb.WriteString(fmt.Sprintf("description: %s\n", description))
	sb.WriteString("---\n\n")

	// Title
	sb.WriteString(fmt.Sprintf("# %s 开发技能\n\n", a.SkillName))

	// Module info
	sb.WriteString("## 模块信息\n\n")
	sb.WriteString(fmt.Sprintf("- **主模块**: `%s`\n", a.MainModule))
	if a.MainModuleVersion != "" {
		sb.WriteString(fmt.Sprintf("- **版本**: `%s`\n", a.MainModuleVersion))
	}
	sb.WriteString(fmt.Sprintf("- **子模块数量**: %d\n", len(a.Modules)))
	sb.WriteString(fmt.Sprintf("- **包数量**: %d\n", len(a.Packages)))
	sb.WriteString(fmt.Sprintf("- **分类数量**: %d\n\n", len(a.Categories)))

	// List all modules
	if len(a.Modules) > 1 {
		sb.WriteString("### 模块列表\n\n")
		sb.WriteString("| 模块路径 | 相对位置 |\n")
		sb.WriteString("|---------|----------|\n")
		for _, mod := range a.Modules {
			sb.WriteString(fmt.Sprintf("| `%s` | `%s` |\n", mod.Path, mod.RelPath))
		}
		sb.WriteString("\n")
	}

	// Package index by category
	sb.WriteString("## 包索引\n\n")

	// Sort categories
	categories := make([]string, 0, len(a.Categories))
	for cat := range a.Categories {
		categories = append(categories, cat)
	}
	sort.Strings(categories)

	for _, cat := range categories {
		pkgPaths := a.Categories[cat]
		sort.Strings(pkgPaths)

		catTitle := cat
		if cat == "root" {
			catTitle = "根包"
		}
		sb.WriteString(fmt.Sprintf("### %s\n\n", catTitle))
		sb.WriteString("| 包名 | 导入路径 | 说明 | Reference |\n")
		sb.WriteString("|------|---------|------|----------|\n")

		for _, pkgPath := range pkgPaths {
			pkg := a.Packages[pkgPath]
			doc := pkg.Doc
			if len(doc) > 60 {
				doc = doc[:57] + "..."
			}
			doc = strings.ReplaceAll(doc, "\n", " ")
			doc = strings.ReplaceAll(doc, "|", "\\|")

			// Determine reference file path
			refPath := getRefPath(pkg.RelPath)
			sb.WriteString(fmt.Sprintf("| %s | `%s` | %s | `references/%s` |\n", pkg.Name, pkg.ImportPath, doc, refPath))
		}
		sb.WriteString("\n")
	}

	// Usage guide
	sb.WriteString("## 使用指南\n\n")
	sb.WriteString("当用户询问某个包的用法时：\n\n")
	sb.WriteString("1. **查找对应的 Reference 文件**：根据上方索引定位对应的 `.md` 文件\n")
	sb.WriteString("2. **读取包详情**：使用 `read_file` 读取 reference 文件获取完整的类型和函数定义\n")
	sb.WriteString("3. **查看源码实现**：如需更详细信息，Reference 文件中包含源码路径，可直接读取源文件\n")
	sb.WriteString("4. **提供示例代码**：基于包的 API 提供准确的使用示例\n\n")
	sb.WriteString("### 查看源码\n\n")
	sb.WriteString("每个 Reference 文件都包含 `源码位置` 章节，记录了包的实际目录路径。\n")
	sb.WriteString("当需要查看具体实现细节时，可以直接使用 `read_file` 读取对应的 `.go` 源文件。\n\n")
	sb.WriteString("```bash\n")
	sb.WriteString("# 获取包的源码目录\n")
	sb.WriteString("go list -f '{{.Dir}}' <import_path>\n")
	sb.WriteString("```\n\n")

	// Quick reference
	sb.WriteString("## 快速参考\n\n")
	sb.WriteString("### 核心类型\n\n")

	// List top types from each category
	for _, cat := range categories {
		if cat == "internal" || cat == "cmd" {
			continue
		}
		pkgPaths := a.Categories[cat]
		for _, pkgPath := range pkgPaths {
			pkg := a.Packages[pkgPath]
			if len(pkg.Types) > 0 {
				sb.WriteString(fmt.Sprintf("**%s** (`%s`):\n", pkg.Name, pkg.ImportPath))
				count := 0
				for _, t := range pkg.Types {
					if count >= 5 {
						sb.WriteString(fmt.Sprintf("  - ... (%d more)\n", len(pkg.Types)-5))
						break
					}
					sb.WriteString(fmt.Sprintf("  - `%s` (%s)\n", t.Name, t.Kind))
					count++
				}
				sb.WriteString("\n")
			}
		}
	}

	// Write file
	skillPath := filepath.Join(a.SkillDir, "SKILL.md")
	if err := os.WriteFile(skillPath, []byte(sb.String()), 0644); err != nil {
		return err
	}

	if verbose {
		fmt.Printf("Generated: %s\n", skillPath)
	}
	return nil
}


// generateReferences generates reference files for each category
func (a *ModuleAnalyzer) generateReferences(verbose bool) error {
	refsDir := filepath.Join(a.SkillDir, "references")

	for cat, pkgPaths := range a.Categories {
		// Create category directory
		catDir := refsDir
		if cat != "root" {
			catDir = filepath.Join(refsDir, cat)
			if err := os.MkdirAll(catDir, 0755); err != nil {
				return err
			}
		}

		for _, pkgPath := range pkgPaths {
			pkg := a.Packages[pkgPath]

			// Determine filename based on relative path
			var refFilePath string
			if pkg.RelPath == "" {
				refFilePath = filepath.Join(refsDir, "root.md")
			} else {
				// For nested packages, create subdirectories
				parts := strings.Split(pkg.RelPath, "/")
				if len(parts) > 1 {
					// Create subdirectory structure
					subDir := filepath.Join(refsDir, filepath.Join(parts[:len(parts)-1]...))
					if err := os.MkdirAll(subDir, 0755); err != nil {
						return err
					}
					refFilePath = filepath.Join(subDir, parts[len(parts)-1]+".md")
				} else {
					refFilePath = filepath.Join(refsDir, parts[0]+".md")
				}
			}

			if err := a.generatePackageRef(pkg, refFilePath, verbose); err != nil {
				return err
			}
		}
	}

	return nil
}

// generatePackageRef generates a reference file for a single package
func (a *ModuleAnalyzer) generatePackageRef(pkg *PackageInfo, refPath string, verbose bool) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# %s\n\n", pkg.Name))
	sb.WriteString(fmt.Sprintf("> Package: `%s`\n\n", pkg.ImportPath))

	sb.WriteString("```go\n")
	sb.WriteString(fmt.Sprintf("import \"%s\"\n", pkg.ImportPath))
	sb.WriteString("```\n\n")

	// Package doc
	if pkg.Doc != "" {
		sb.WriteString("## 概述\n\n")
		sb.WriteString(pkg.Doc + "\n\n")
	}

	// Module info if from sub-module
	if pkg.ModulePath != a.MainModule {
		sb.WriteString("## 所属模块\n\n")
		sb.WriteString(fmt.Sprintf("此包属于子模块: `%s`\n\n", pkg.ModulePath))
	}

	// Source code location - for AI to find actual source files
	sb.WriteString("## 源码位置\n\n")
	sb.WriteString(fmt.Sprintf("**导入路径**: `%s`\n\n", pkg.ImportPath))
	sb.WriteString("如需查看完整源码实现，可以通过以下命令获取包的实际路径：\n\n")
	sb.WriteString("```bash\n")
	sb.WriteString(fmt.Sprintf("# 获取包的源码目录\n"))
	sb.WriteString(fmt.Sprintf("go list -f '{{.Dir}}' %s\n\n", pkg.ImportPath))
	sb.WriteString(fmt.Sprintf("# 或者查看 GOPATH/GOMODCACHE\n"))
	sb.WriteString(fmt.Sprintf("go env GOPATH GOMODCACHE\n"))
	sb.WriteString("```\n\n")

	// Source files (just filenames, use go list to get full path)
	sb.WriteString("## 源文件\n\n")
	for _, f := range pkg.GoFiles {
		sb.WriteString(fmt.Sprintf("- `%s`\n", f))
	}
	sb.WriteString("\n")

	// Types
	if len(pkg.Types) > 0 {
		sb.WriteString("## 类型定义\n\n")
		for _, t := range pkg.Types {
			sb.WriteString(fmt.Sprintf("### %s\n\n", t.Name))
			sb.WriteString(fmt.Sprintf("**类型**: %s\n\n", t.Kind))
			if t.Doc != "" {
				sb.WriteString(t.Doc + "\n\n")
			}

			// Methods
			if len(t.Methods) > 0 {
				sb.WriteString("**方法**:\n\n")
				for _, m := range t.Methods {
					sb.WriteString(fmt.Sprintf("- `%s`\n", m.Signature))
					if m.Doc != "" {
						// First line only
						docLine := strings.Split(m.Doc, "\n")[0]
						if len(docLine) > 80 {
							docLine = docLine[:77] + "..."
						}
						sb.WriteString(fmt.Sprintf("  %s\n", docLine))
					}
				}
				sb.WriteString("\n")
			}
		}
	}

	// Functions
	if len(pkg.Functions) > 0 {
		sb.WriteString("## 函数\n\n")
		for _, f := range pkg.Functions {
			sb.WriteString(fmt.Sprintf("### %s\n\n", f.Name))
			sb.WriteString(fmt.Sprintf("```go\n%s\n```\n\n", f.Signature))
			if f.Doc != "" {
				sb.WriteString(f.Doc + "\n\n")
			}
		}
	}

	// Dependencies
	if len(pkg.InternalDep) > 0 {
		sb.WriteString("## 内部依赖\n\n")
		for _, dep := range pkg.InternalDep {
			// Show relative path
			relDep := dep
			for _, mod := range a.Modules {
				if strings.HasPrefix(dep, mod.Path) {
					relDep = strings.TrimPrefix(dep, mod.Path+"/")
					break
				}
			}
			sb.WriteString(fmt.Sprintf("- `%s`\n", relDep))
		}
		sb.WriteString("\n")
	}

	if len(pkg.ExternalDep) > 0 {
		sb.WriteString("## 外部依赖\n\n")
		for _, dep := range pkg.ExternalDep {
			sb.WriteString(fmt.Sprintf("- `%s`\n", dep))
		}
		sb.WriteString("\n")
	}

	if err := os.WriteFile(refPath, []byte(sb.String()), 0644); err != nil {
		return err
	}

	if verbose {
		fmt.Printf("Generated: %s\n", refPath)
	}
	return nil
}

// Helper functions

func isStdLib(pkg string) bool {
	return !strings.Contains(pkg, ".")
}

func exprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + exprToString(t.X)
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name
	case *ast.ArrayType:
		return "[]" + exprToString(t.Elt)
	case *ast.MapType:
		return "map[" + exprToString(t.Key) + "]" + exprToString(t.Value)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.Ellipsis:
		return "..." + exprToString(t.Elt)
	case *ast.FuncType:
		return "func"
	case *ast.ChanType:
		return "chan " + exprToString(t.Value)
	default:
		return ""
	}
}

func buildFuncSignature(decl *ast.FuncDecl) string {
	var sb strings.Builder
	sb.WriteString("func ")

	if decl.Recv != nil && len(decl.Recv.List) > 0 {
		recv := decl.Recv.List[0]
		sb.WriteString("(")
		if len(recv.Names) > 0 {
			sb.WriteString(recv.Names[0].Name + " ")
		}
		sb.WriteString(exprToString(recv.Type))
		sb.WriteString(") ")
	}

	sb.WriteString(decl.Name.Name)
	sb.WriteString("(")

	// Parameters
	if decl.Type.Params != nil {
		params := make([]string, 0)
		for _, field := range decl.Type.Params.List {
			typeStr := exprToString(field.Type)
			if len(field.Names) > 0 {
				for _, name := range field.Names {
					params = append(params, name.Name+" "+typeStr)
				}
			} else {
				params = append(params, typeStr)
			}
		}
		sb.WriteString(strings.Join(params, ", "))
	}
	sb.WriteString(")")

	// Results
	if decl.Type.Results != nil && len(decl.Type.Results.List) > 0 {
		results := make([]string, 0)
		for _, field := range decl.Type.Results.List {
			typeStr := exprToString(field.Type)
			if len(field.Names) > 0 {
				for _, name := range field.Names {
					results = append(results, name.Name+" "+typeStr)
				}
			} else {
				results = append(results, typeStr)
			}
		}
		if len(results) == 1 {
			sb.WriteString(" " + results[0])
		} else {
			sb.WriteString(" (" + strings.Join(results, ", ") + ")")
		}
	}

	return sb.String()
}

func getSortedKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func getRefPath(relPath string) string {
	if relPath == "" {
		return "root.md"
	}
	return relPath + ".md"
}
