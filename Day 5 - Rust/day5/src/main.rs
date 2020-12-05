use std::fs;
const SEATS: i32 = 127;
const COLUMNS: i32 = 8;

fn main() {
  let contents = fs::read_to_string("input.txt").unwrap();
  let ids = parse_ids(contents);
  println!("Part 1: {}", part1(ids.clone()));
  println!("Part 2: {}", part2(ids.clone()));
}


fn parse_ids(contents: String) -> Vec<i32> {
    let mut ids = Vec::new();

    for line in contents.split("\n") {
        let mut seats: Vec<i32> = (1..SEATS+1).collect();
        let mut columns: Vec<i32> = (0..COLUMNS).collect();

        for char in line.chars() {
           let _ = match char {
                'F' => seats.drain(seats.len()/2..seats.len()),
                'B' => seats.drain(0..seats.len()/2),
                'R' => columns.drain(0..columns.len()/2),
                'L' => columns.drain(columns.len()/2..columns.len()),
                _   => columns.drain(0..0),  // Just to satisfy exhaustive matching
            };
        }
        ids.push((seats[0] * 8 + columns[0]) as i32);
    }

    return ids;
}


fn part1(contents: Vec<i32>) -> i32 {
    let mut highest = 0;
    for id in contents {
        if id > highest {
            highest = id.to_owned();
        }
    }
    return highest;
}


fn part2(mut contents: Vec<i32>) -> i32 {
    let mut seat: i32 = 0;
    contents.sort();

    for i in 1..contents.len() {
        if contents[i] != contents[i-1]+1 {
            seat = contents[i]-1;
        }
    }

    return seat;
}