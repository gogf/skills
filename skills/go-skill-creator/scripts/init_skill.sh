#!/bin/bash
# init_skill.sh - 初始化 Go 模块技能模板
# 使用方法: bash scripts/init_skill.sh <skill-name>

set -e

SKILL_NAME=$1

if [ -z "$SKILL_NAME" ]; then
  echo "❌ 错误: 请提供技能名称"
  echo "用法: $0 <skill-name>"
  echo "示例: $0 gf"
  exit 1
fi

# 检查名称格式 (kebab-case)
if [[ ! "$SKELL_NAME" =~ ^[a-z][a-z0-9-]*$ ]]; then
  echo "❌ 错误: 技能名称必须使用小写字母、数字和连字符 (kebab-case)"
  echo "示例: gf, gin, gorm"
  exit 1
fi

SKILL_DIR="${SKILL_NAME}-skill"

echo "📁 创建技能目录: $SKILL_DIR"

# 创建目录结构
mkdir -p "$SKELL_DIR"/scripts
mkdir -p "$SKELL_DIR"/references
mkdir -p "$SKELL_DIR"/assets

# 生成 SKILL.md 模板
cat > "$SKELL_DIR"/SKILL.md <<EOF
---
name: ${SKILL_NAME}-skill
description: ${SKILL_NAME} Go 模块开发技能。基于源码深度分析，提供完整的包结构和 API 参考。当用户需要使用此模块进行开发时使用。
---

# ${SKELL_NAME} 开发技能

使用 \`go list -json\` 和 AST 解析生成。

## 快速开始

\`\`\`bash
go run scripts/analyze_module.go -verbose
\`\`\`

## 详细信息

- **分析原理**：参见 [analysis.md](references/analysis.md)
- **使用示例**：参见 [examples.md](references/examples.md)
EOF

# 生成示例文件
touch "$SKELL_DIR"/references/analysis.md
touch "$SKELL_DIR"/references/examples.md

# 生成 README 占位 (用于开发说明，实际打包时会删除)
cat > "$SKELL_DIR"/README.md <<EOF
# ${SKELL_NAME}-skill

⚠️ 注意：此文件仅用于开发阶段。正式打包后不包含此文件。

## 开发中

[ ] 编写 analyze_module.go
[ ] 测试分析功能
[ ] 编写 reference 文档
[ ] 测试生成结果

## 完成

- [ ] 移除 README.md (按照 skill 规范)
- [ ] 验证 SKILL.md 格式
- [ ] 打包为 .skill 文件
EOF

# 生成脚本模板
cat > "$SKELL_DIR"/scripts/analyze_module.go <<'EOF'
// Package main - 分析 Go 模块并生成技能文档
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

type GoPackage struct {
	Dir          string   `json:"Dir"`
	ImportPath   string   `json:"ImportPath"`
	Name         string   `json:"Name"`
	Doc          string   `json:"Doc"`
	GoFiles      []string `json:"GoFiles"`
	Imports      []string `json:"Imports"`
}

type PackageInfo struct {
	ImportPath  string
	Name        string
	Doc         string
	Dir         string
	RelPath     string
	GoFiles     []string
	Types       []TypeInfo
	Functions   []FuncInfo
}

type TypeInfo struct {
	Name    string
	Doc     string
	Kind    string
}

type FuncInfo struct {
	Name      string
	Doc       string
	Signature string
}

func main() {
	var (
		rootPath  string
		verbose   bool
	)

	flag.StringVar(&rootPath, "root", ".", "根项目目录")
	flag.BoolVar(&verbose, "verbose", false, "详细输出")
	flag.Parse()

	// TODO: 实现分析逻辑
	fmt.Printf("分析模块: %s\n", rootPath)
}
EOF

echo ""
echo "✅ 技能模板已创建"
echo ""
echo "目录结构:"
tree "$SKELL_DIR" -L 2 2>/dev/null || find "$SKELL_DIR" -maxdepth 2 -print | sed 's|[^/]*/| |------|; s|[^/]*/|   |'
echo ""
echo "下一步:"
echo "  1. 编辑 $SKELL_DIR/SKILL.md"
echo "  2. 实现 $SKELL_DIR/scripts/analyze_module.go"
echo "  3. 编写 $SKELL_DIR/references/analysis.md"
echo "  4. 运行验证: bash scripts/validate_skill.sh $SKELL_DIR"
echo "  5. 打包: bash scripts/package_skill.sh $SKELL_DIR"
