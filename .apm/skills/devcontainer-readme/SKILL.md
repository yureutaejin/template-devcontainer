---
name: devcontainer-readme
description: "Document Dev Container contents in this README.md. Use when adding or updating included components, development tools, VS Code extensions, version links, release pages, registries, Markdown tables, or collapsible README sections."
argument-hint: "Update the Dev Container README documentation"
---

# Dev Container README

## When to Use

Use this skill when documenting the development environment defined in `.devcontainer/Dockerfile` or `.devcontainer/devcontainer.json`.

## Procedure

1. Read the relevant Dev Container configuration before editing `README.md`.
2. Document only tools, versions, features, mounts, and extensions explicitly configured in the repository.
3. Add top-level runtime components to the `## Included Components` table.
4. Keep development packages in a separate `### Development Tools` section. Put its table inside a `<details>` element when a collapsible list is requested, but keep the H3 outside the element.
5. Keep VS Code extensions in a separate `### VS Code Extensions` section.
6. Link tools to the most specific official source available:
   - Use an official release page for version-pinned tools.
   - Use a package registry page for distribution packages.
   - Use an image registry page for container images and Dev Container features.
   - Do not use direct asset-download URLs when an official release page exists.
7. Use local relative links for repository resources, such as a bind-mounted `.codex` directory.
8. Validate the edited README for Markdown diagnostics.

## Template

Use [the README documentation template](./assets/readme-template.md) as the starting structure.
