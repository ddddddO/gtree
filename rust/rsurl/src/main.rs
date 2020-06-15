struct Cli {
    protocol: String,
    target: String
}

fn new_cli(protocol: String, target: String) -> Cli {
    Cli {
        protocol,
        target
    }
}

impl Cli {
    fn info(&self) {
        println!("protocol: {}, target: {}", self.protocol, self.target);
    }

    fn send(&self, num: i32) {
        println!("deadline time: {}s", num);
    }
}

fn main() {
    let cli = new_cli(
        std::env::args().nth(1).expect("no protocol given"),
        std::env::args().nth(2).expect("no target given")
    );

    cli.info();
    cli.send(5);
}

