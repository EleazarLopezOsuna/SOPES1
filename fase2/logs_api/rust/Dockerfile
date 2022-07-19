FROM rust:1.60-slim-buster as build

RUN USER=root cargo new --bin mserver
WORKDIR /mserver

COPY ./Cargo.lock ./Cargo.lock
COPY ./Cargo.toml ./Cargo.toml

RUN cargo build --release 
RUN rm src/*.rs

COPY ./src ./src

RUN rm ./target/release/deps/mserver*
RUN cargo build --release

FROM debian:buster-slim
#RUN apt-get update && apt-get install -y build-essential musl-tools && \
 #rustup target add x86-64-unknown-linux-musl && rm -rf /var/lib/apt/lists/*
COPY --from=build /mserver/target/release/mserver /usr/src/mserver
CMD ["/usr/src/mserver"]
EXPOSE 8080