set shell := ["bash", "-cu"]

default:
  @echo "Usage: just run <dayXX>"
  @echo "Example: just run day01"

run fname:
  cat "input/{{fname}}.txt" | uv run python -m advent_of_code_2025."{{fname}}"

tc:
  uv run ty check

fmt:
  uv run ruff format

pydoc name:
  uv run python -m pydoc {{name}}
