set shell := ["bash", "-cu"]

default:
    @echo "Usage: just run <dayXX>"
    @echo "Example: just run day01"

run fname:
    cat "input/{{fname}}.txt" | cargo run --release --bin "{{fname}}"
