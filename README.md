# Template Dev Container

## Included Components

| Component | Included |
| --- | --- |
| Base image | [Ubuntu 24.04 (Noble Numbat)](https://hub.docker.com/_/ubuntu) |
| Go | [1.26.5](https://go.dev/dl/) |
| Python | [3.13](https://docs.astral.sh/uv/guides/install-python/) |
| Rust | [1.97.1](https://releases.rs/docs/1.97.1/) |
| uv | [0.12.5](https://github.com/astral-sh/uv/releases/tag/0.12.5) |
| Docker CLI | [Docker outside of Docker Dev Container Feature](https://ghcr.io/devcontainers/features/docker-outside-of-docker) |

The container runs as the `vscode` user. It bind-mounts the local [`.codex`](.codex) directory at `/home/vscode/.codex`, mounts the workspace at `/workspaces/<workspace-name>`, uses the host network, and maps `host.docker.internal` to the host gateway.

### Development Tools

<details>
<summary>Show tools</summary>

| Tool | Package Registry |
| --- | --- |
| [Agent Package Manager](https://github.com/microsoft/apm/releases) | `apm` |
| [bash-completion](https://packages.ubuntu.com/noble/bash-completion) | `bash-completion` |
| [Build Essential](https://packages.ubuntu.com/noble/build-essential) | `build-essential` |
| [CA Certificates](https://packages.ubuntu.com/noble/ca-certificates) | `ca-certificates` |
| [Clang LLDB](https://packages.ubuntu.com/noble/lldb) | `lldb` |
| [cURL](https://packages.ubuntu.com/noble/curl) | `curl` |
| [Git](https://packages.ubuntu.com/noble/git) | `git` |
| [OpenSSL development libraries](https://packages.ubuntu.com/noble/libssl-dev) | `libssl-dev` |
| [pkg-config](https://packages.ubuntu.com/noble/pkg-config) | `pkg-config` |
| [SSH client](https://packages.ubuntu.com/noble/ssh) | `ssh` |
| [Starship](https://github.com/starship/starship/releases/tag/v1.26.0) | `starship` 1.26.0 |
| [sudo](https://packages.ubuntu.com/noble/sudo) | `sudo` |
| [Vim](https://packages.ubuntu.com/noble/vim) | `vim` |

</details>

### VS Code Extensions

| Extension | Purpose |
| --- | --- |
| [bierner.markdown-mermaid](https://marketplace.visualstudio.com/items?itemName=bierner.markdown-mermaid) | Render Mermaid diagrams in Markdown previews. |
| [charliermarsh.ruff](https://marketplace.visualstudio.com/items?itemName=charliermarsh.ruff) | Python linting and formatting with Ruff. |
| [eamodio.gitlens](https://marketplace.visualstudio.com/items?itemName=eamodio.gitlens) | Git history, blame, and repository insights. |
| [golang.Go](https://marketplace.visualstudio.com/items?itemName=golang.Go) | Go language support. |
| [GitHub.vscode-pull-request-github](https://marketplace.visualstudio.com/items?itemName=GitHub.vscode-pull-request-github) | GitHub pull request and issue integration. |
| [mhutchie.git-graph](https://marketplace.visualstudio.com/items?itemName=mhutchie.git-graph) | Git commit graph visualization. |
| [openai.chatgpt](https://marketplace.visualstudio.com/items?itemName=openai.chatgpt) | OpenAI ChatGPT integration. |
| [rust-lang.rust-analyzer](https://marketplace.visualstudio.com/items?itemName=rust-lang.rust-analyzer) | Rust language analysis and editor support. |
| [tamasfe.even-better-toml](https://marketplace.visualstudio.com/items?itemName=tamasfe.even-better-toml) | TOML language support. |
| [vadimcn.vscode-lldb](https://marketplace.visualstudio.com/items?itemName=vadimcn.vscode-lldb) | Native debugging with LLDB. |
| [yzhang.markdown-all-in-one](https://marketplace.visualstudio.com/items?itemName=yzhang.markdown-all-in-one) | Markdown editing, shortcuts, and formatting. |