default:
    @echo "Usage: just <dayXX>"

FNAME:
    cat input/{{FNAME}}.txt | cargo run --release --bin {{FNAME}}
