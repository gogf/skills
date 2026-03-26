## Use Claude Code，auto generate commit message and push to remote branch
.PHONY: up
up:
	@set -e; \
	if git diff --quiet HEAD && git diff --cached --quiet && [ -z "$$(git ls-files --others --exclude-standard)" ]; then \
		if git diff --quiet HEAD origin/$$(git branch --show-current) 2>/dev/null; then \
			echo "No changes to commit and nothing to push"; \
		else \
			echo "No local changes, pushing unpushed commits..."; \
			git push origin $$(git branch --show-current); \
		fi; \
		exit 0; \
	fi; \
	git add -A; \
	echo "Analyzing changes and generating commit message via AI (model: $(or $(m),haiku))..."; \
	MSG=$$(git diff --cached --stat && echo "---" && git diff --cached | head -2000 | \
		claude -p "Analyze the git diff above and generate a concise commit message (single line, max 72 chars, lowercase, no quotes). Output only the commit message itself, nothing else." \
		--model $(or $(m),haiku)) || { echo "Error: Claude command failed"; exit 1; }; \
	COMMIT_MSG=$$(echo "$$MSG" | tail -1); \
	if [ -z "$$COMMIT_MSG" ]; then \
		echo "Error: Failed to generate commit message"; \
		exit 1; \
	fi; \
	echo "Commit: $$COMMIT_MSG"; \
	git commit -m "$$COMMIT_MSG" && \
	git push origin $$(git branch --show-current)

## Generate skill content by running main.go
.PHONY: skills
skills:
	@echo "Generating skill content..."
	@go run main.go
	@echo "Skill content generated successfully!"
