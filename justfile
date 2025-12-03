set shell := ["bash", "-cu"]


default:
  @echo "Usage: just run <dayXX>"
  @echo "Example: just run day01"

run fname:
  @if [ -f "{{fname}}.go" ]; then \
    cat "input/{{fname}}.txt" | go run "{{fname}}.go"; \
  else \
    cat "input/{{fname}}.txt" | uv run python -m "{{fname}}"; \
  fi

tc:
  uv run ty check

fmt:
  uv run ruff format
  go fmt .

pydoc name:
  uv run python -m pydoc {{name}}

godoc name:
  go doc {{name}} | less

godoc-serve:
  go doc -http
