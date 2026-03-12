## Use Claude Code，auto generate commit message and push to remote branch
.PHONY: up
up:
	@if git diff --quiet HEAD && git diff --cached --quiet && [ -z "$$(git ls-files --others --exclude-standard)" ]; then \
		echo "nothing changes"; \
		exit 0; \
	fi
	@git add -A
	@echo "analysing commit message using claude..."
	@MSG=$$(git diff --cached --stat && echo "---" && git diff --cached | head -2000 | \
		claude -p "analyze git diff and generate a concise english commit message (one line, under 72 chars, lowercase, no quotes)" \
		--model haiku 2>/dev/null) && \
	COMMIT_MSG=$$(echo "$$MSG" | tail -1) && \
	echo "Commit: $$COMMIT_MSG" && \
	git commit -m "$$COMMIT_MSG" && \
	git push origin main