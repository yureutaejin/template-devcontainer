FROM --platform=$BUILDPLATFORM ubuntu:noble@sha256:561618e2c15bf2397621dd04f96926663a3b5616c189cf7e38db7e82f5c538ea

ARG TARGETOS
ARG TARGETARCH
ARG USERNAME=vscode
ARG USER_UID=1001
ARG USER_GID=${USER_UID}
ARG GO_VERSION=1.26.5
ARG RUSTUP_VERSION=1.28.2
ARG RUST_VERSION=1.97.1
ARG PYTHON_VERSION=3.13

ENV DEBIAN_FRONTEND=noninteractive
# astral-sh/uv
ENV UV_PYTHON_INSTALL_DIR=/usr/local/python
ENV UV_PYTHON_BIN_DIR=/usr/local/bin
# rust
ENV RUSTUP_HOME=/usr/local/rustup
ENV CARGO_HOME=/usr/local/cargo
ENV PATH=${CARGO_HOME}/bin:${PATH}

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    bash-completion \
    build-essential \
    ca-certificates \
    curl \
    git \
    libssl-dev \
    lldb \
    pkg-config \
    vim \
    sudo \
    ssh \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /etc/ssh /etc/sudoers.d

RUN echo 'github.com ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOMqqnkVzrm0SdG6UOoqKLsabgH5C9okWi0dh2l9GKJl' \
    >> /etc/ssh/ssh_known_hosts

RUN groupadd --gid ${USER_GID} ${USERNAME} \
    && useradd --uid ${USER_UID} --gid ${USER_GID} -m ${USERNAME} --shell /bin/bash \
    && echo "${USERNAME} ALL=(root) NOPASSWD:ALL" > /etc/sudoers.d/${USERNAME} \
    && chmod 0440 /etc/sudoers.d/${USERNAME}

# MARK: Install Rust
RUN <<EOF
    set -e
    curl -fsSL \
        -o /usr/local/bin/rustup-init \
        "https://static.rust-lang.org/rustup/archive/${RUSTUP_VERSION}/$(uname -m)-unknown-linux-gnu/rustup-init"
    chmod +x /usr/local/bin/rustup-init
    rustup-init -y --no-modify-path --default-toolchain none
    rm /usr/local/bin/rustup-init

    rustup toolchain install ${RUST_VERSION} --profile default \
        --component rustfmt --component rust-src --component clippy
    rustup default ${RUST_VERSION}
    chmod -R a+w ${RUSTUP_HOME} ${CARGO_HOME}
EOF

# MARK: Install Go
RUN <<EOF
    set -e
    curl -fsSL \
        https://go.dev/dl/go${GO_VERSION}.${TARGETOS}-${TARGETARCH}.tar.gz \
        | tar -C /usr/local -xz
    echo 'export PATH="/usr/local/go/bin:$PATH"' >> /etc/bash.bashrc
EOF

# MARK: Install uv & default python version
COPY --from=ghcr.io/astral-sh/uv:0.12.5 /uv /uvx /usr/local/bin/
RUN uv python install ${PYTHON_VERSION} --default

# MARK: Install user preferences
# List:
# - github.com/astral-sh/starship: A cross-shell prompt made with Rust
# - github.com/microsoft/apm: Agent Package Manager
RUN curl -fsSL \
    https://github.com/starship/starship/releases/download/v1.26.0/starship-$(uname -m)-unknown-linux-musl.tar.gz \
    | tar xz -C /usr/local/bin \
    && echo 'eval "$(/usr/local/bin/starship init bash)"' >> /etc/bash.bashrc

RUN <<EOF
  curl -fsSL https://github.com/microsoft/apm/releases/latest/download/apm-linux-${TARGETARCH}.tar.gz | tar -xz -C /usr/local/lib/
  mv /usr/local/lib/apm-${TARGETOS}-${TARGETARCH} /usr/local/lib/apm
  ln -sf /usr/local/lib/apm/apm /usr/local/bin/apm
EOF
