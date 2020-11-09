struct Sample {
    x: i32,
}

impl Sample {
    fn new(x: i32) -> Sample {
        Sample{x: x}
    }

    fn inc(&self) -> i32 {
        self.x + 1
    }

    fn instance_inc(&mut self) {
        self.x = self.x + 1
    }
}


fn main() {
    println!("Hello, world!");

    let mut s = Sample::new(555);

    s.inc();
    s.inc();
    s.inc();
    let inced_num = s.inc();

    println!("Sample's x is {}.", inced_num); // 556

    s.instance_inc();
    s.instance_inc();
    s.instance_inc();
    println!("Sample's x is {}.", s.x); // 558
}
