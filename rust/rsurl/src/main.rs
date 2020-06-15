use std::net::TcpStream;

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

        // ref: https://doc.rust-lang.org/std/net/struct.TcpStream.html#examples-1
        if let Ok(stream) = TcpStream::connect("127.0.0.1:8888") {
            // 別ターミナルで、goexec 'http.ListenAndServe(":8888", nil)'　を実行したあとで以下が出力される
            println!("Connected to the server!");
        } else {
            println!("Couldn't connect to server...");
        }
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

