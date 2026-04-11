FROM rust:1.82 AS builder
 
WORKDIR /app
 
COPY Cargo.toml Cargo.lock* ./
 
RUN mkdir src && echo "fn main() {}" > src/main.rs
RUN cargo build --release
RUN rm -f target/release/deps/hello_world*

COPY src ./src
RUN cargo build --release
 
FROM debian:bookworm-slim
 
COPY --from=builder /app/target/release/hello-world /usr/local/bin/hello-world
 
CMD ["hello-world"]