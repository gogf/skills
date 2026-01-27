#!/bin/bash
# validate_skill.sh - 验证技能结构和格式
# 使用方法: bash scripts/validate_skill.sh <skill-dir>

set -e

SKILL_DIR=$1

# 获取脚本所在目录
SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)
# 获取 skills/ 目录 (脚本目录的父目录的父目录)
SKILLS_DIR=$(dirname "$SCRIPT_DIR")

# 如果是相对路径，相对于 skills/ 目录解析
if [[ "$SKILL_DIR" != /* ]]; then
  SKILL_DIR="$SKILLS_DIR/$SKILL_DIR"
fi

if [ ! -d "$SKILL_DIR" ]; then
  echo "❌ 错误: 目录不存在: $SKILL_DIR"
  exit 1
fi

echo "🔍 验证技能: $SKILL_DIR"
echo ""

# 检查 SKILL.md
if [ ! -f "$SKILL_DIR/SKILL.md" ]; then
  echo "❌ 错误: 缺少 SKILL.md"
  exit 1
fi
echo "✓ SKILL.md 存在"

# 检查 frontmatter
if ! grep -q "^---$" "$SKILL_DIR/SKILL.md"; then
  echo "❌ 错误: SKILL.md 缺少 YAML frontmatter (---)"
  exit 1
fi
echo "✓ YAML frontmatter 格式正确"

# 检查 name 字段
if ! grep -q "^name:" "$SKILL_DIR/SKILL.md"; then
  echo "❌ 错误: SKILL.md 缺少 name 字段"
  exit 1
fi
echo "✓ name 字段存在"

# 检查 description 字段
if ! grep -q "^description:" "$SKILL_DIR/SKILL.md"; then
  echo "❌ 错误: SKILL.md 缺少 description 字段"
  exit 1
fi
echo "✓ description 字段存在"

# 检查行数
LINES=$(wc -l < "$SKILL_DIR/SKILL.md" | awk '{print $1}')
if [ $LINES -gt 500 ]; then
  echo "⚠️  警告: SKILL.md 过长 ($LINES 行)，建议 < 500 行"
  echo "   请考虑将详细内容拆分到 references/ 目录"
fi
echo "✓ SKILL.md 行数: $LINES"

# 检查是否包含 README.md (按照规范不应该包含)
if [ -f "$SKILL_DIR/README.md" ]; then
  echo "⚠️  警告: 包含 README.md (按照 skill 规范不应该包含)"
  echo "   打包前请移除此文件"
fi

# 检查是否有 scripts 目录
if [ -d "$SKILL_DIR/scripts" ]; then
  echo "✓ scripts/ 目录存在"
else
  echo "⚠️  提示: 建议创建 scripts/ 目录"
fi

# 检查是否有 references 目录
if [ -d "$SKILL_DIR/references" ]; then
  echo "✓ references/ 目录存在"
else
  echo "⚠️  提示: 建议创建 references/ 目录"
fi

# 检查 scripts 是否有可执行文件
if [ -d "$SKILL_DIR/scripts" ]; then
  EXEC_COUNT=$(find "$SKILL_DIR/scripts" -type f -executable | wc -l)
  if [ $EXEC_COUNT -gt 0 ]; then
    echo "✓ scripts/ 中有 $EXEC_COUNT 个可执行文件"
  fi
fi

echo ""
echo "✅ 验证通过"
