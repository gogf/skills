## Use Claude Code，auto generate commit message and push to remote branch
.PHONY: up
up:
	@if git diff --quiet HEAD && git diff --cached --quiet && [ -z "$$(git ls-files --others --exclude-standard)" ]; then \
		echo "No changes to commit"; \
		exit 0; \
	fi
	@git add -A
	@echo "Analyzing changes and generating commit message via AI..."
	@MSG=$$(git diff --cached --stat && echo "---" && git diff --cached | head -2000 | \
		claude -p "Analyze the git diff above and generate a concise commit message (single line, max 72 chars, lowercase, no quotes). Output only the commit message itself, nothing else." \
		--model haiku 2>/dev/null) && \
	COMMIT_MSG=$$(echo "$$MSG" | tail -1) && \
	echo "Commit: $$COMMIT_MSG" && \
	git commit -m "$$COMMIT_MSG" && \
	git push origin main