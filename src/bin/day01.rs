use anyhow::{anyhow, Result};


enum Rotation {
    L(u16),
    R(u16),
}

fn parse_lines() -> Result<Vec<Rotation>> {
    let stdin = std::io::stdin();
    stdin.lines().map(|line| {

        let l = line?;

        let (dir, num) = l.split_at(1);

        match dir {
            "L" => Ok(Rotation::L(num.parse()?)),
            "R" => Ok(Rotation::R(num.parse()?)),
            _ => Err(anyhow!("Unknown direction"))
        }


    }).collect()
}





fn main() -> Result<()> {

    let rotations = parse_lines()?;

    println!("{:?}", rotations)

    Ok(())

}
