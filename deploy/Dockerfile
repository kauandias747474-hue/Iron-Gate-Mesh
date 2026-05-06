# Estágio 1: Build do Rust (Core)
FROM rust:1.75 as rust-builder
WORKDIR /app
COPY . .
RUN cargo build --release

# Estágio 2: Imagem Final
FROM debian:bookworm-slim
WORKDIR /root/
COPY --from=rust-builder /app/target/release/irongate-core .
# Copia os dashboards e scripts necessários
COPY irongate-visuals/dashboards ./dashboards
CMD ["./irongate-core"]
