set shell := ["bash", "-cu"]


default:
  @echo "Usage: just run <dayXX>"
  @echo "Example: just run day01"

run fname:
  @set -euo pipefail
  @if [ -f "{{fname}}/main.go" ]; then \
    cat "{{fname}}/input.txt" | go run "./{{fname}}"; \
  elif [ -f "{{fname}}/main.py" ]; then \
    cat "{{fname}}/input.txt" | uv run python "./{{fname}}/main.py"; \
  else \
    echo "No main.go or main.py found in {{fname}}" >&2; \
    exit 1; \
  fi

tc:
  uv run ty check

fmt:
  uv run ruff format
  rg --files -g '*.go' -0 | xargs -0 gofmt -w

pydoc name:
  uv run python -m pydoc {{name}}

godoc name:
  go doc {{name}} | less

godoc-serve:
  go doc -http
