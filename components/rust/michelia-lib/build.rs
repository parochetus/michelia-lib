use std::path::PathBuf;

fn main() {
    let out_dir = PathBuf::from("src/proto-rs/v0");
    tonic_build::configure()
        .out_dir(out_dir)
        .compile(&["../../../protos/version.proto"], &["../../../protos"])
        .unwrap();
}
