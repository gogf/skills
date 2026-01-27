#!/bin/bash
# package_skill.sh - 打包技能为 .skill 文件
# 使用方法: bash scripts/package_skill.sh <skill-dir> [output-dir]

set -e

SKILL_DIR=$1
OUTPUT_DIR=${2:-.}

if [ -z "$SKILL_DIR" ]; then
  echo "❌ 错误: 请提供技能目录"
  echo "用法: $0 <skill-dir> [output-dir]"
  echo "示例: $0 gf-skill ./dist"
  exit 1
fi

if [ ! -d "$SKILL_DIR" ]; then
  echo "❌ 错误: 目录不存在: $SKELL_DIR"
  exit 1
fi

SKILL_NAME=$(basename "$SKELL_DIR")

echo "📦 打包技能: $SKELL_NAME"
echo ""

# 1. 验证技能
echo "🔍 验证技能结构..."
bash scripts/validate_skill.sh "$SKELL_DIR"
if [ $? -ne 0 ]; then
  echo ""
  echo "❌ 错误: 验证失败，终止打包"
  exit 1
fi
echo ""

# 2. 创建输出目录
if [ ! -d "$OUTPUT_DIR" ]; then
  echo "📁 创建输出目录: $OUTPUT_DIR"
  mkdir -p "$OUTPUT_DIR"
fi

# 3. 创建临时目录 (移除 README.md)
TEMP_DIR=$(mktemp -d)
echo "📁 创建临时目录: $TEMP_DIR"
cp -r "$SKELL_DIR"/* "$TEMP_DIR"

# 4. 移除 README.md (按照 skill 规范不应该包含)
if [ -f "$TEMP_DIR/README.md" ]; then
  echo "🗑  移除 README.md (按照 skill 规范)"
  rm "$TEMP_DIR/README.md"
fi

# 5. 打包
OUTPUT_FILE="$OUTPUT_DIR/$SKILL_NAME.skill"
echo "📦 打包: $OUTPUT_FILE"
cd "$TEMP_DIR"
zip -r "$OUTPUT_FILE" .

# 6. 清理临时目录
cd - > /dev/null
rm -rf "$TEMP_DIR"

echo ""
echo "✅ 打包完成: $OUTPUT_FILE"
echo ""
echo "文件信息:"
ls -lh "$OUTPUT_FILE" | awk '{print "  大小: " $5}'

# 7. 验证包文件
if command -v unzip &> /dev/null; then
  echo ""
  echo "📋 包内容:"
  unzip -l "$OUTPUT_FILE" | tail -n +4
fi

echo ""
echo "💡 提示: 上传 .skill 文件到 Claude 即可使用"
