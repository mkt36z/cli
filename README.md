# mkt36z

AI-powered marketing intelligence from your terminal. Replace your entire marketing team with a coordinated system of AI agents that think like the world's top marketers.

## Install

```sh
curl -sSL https://install.mkt36z.com | sh
```

This works on **macOS** and **Linux** (amd64 and arm64). The script auto-detects your OS and architecture, downloads the correct binary from GitHub Releases, verifies the SHA256 checksum, and installs to `/usr/local/bin`.

### Other methods

```sh
# Homebrew (macOS/Linux)
brew install mkt36z/tap/mkt36z

# Go (requires Go 1.24+)
go install github.com/mkt36z/cli/cmd/mkt36z@latest

# Docker
docker run --rm ghcr.io/mkt36z/cli mkt36z version

# Windows (Scoop)
scoop bucket add mkt36z https://github.com/mkt36z/scoop-bucket
scoop install mkt36z
```

### Environment variables

| Variable | Description |
|----------|-------------|
| `MKT36Z_VERSION` | Install a specific version instead of latest |
| `MKT36Z_DIR` | Custom install directory (default: `/usr/local/bin`) |

```sh
# Example: install v0.2.0 to ~/bin
MKT36Z_VERSION=0.2.0 MKT36Z_DIR=~/bin curl -sSL https://install.mkt36z.com | sh
```

## Quick start

```sh
# 1. Authenticate
mkt36z auth login

# 2. Set up your brand context (product, audience, voice)
mkt36z init

# 3. Generate marketing copy
mkt36z generate headline "AI CRM for startups"
mkt36z generate email "Product launch announcement"
mkt36z generate landing-page "SaaS analytics tool"

# 4. Run deep marketing playbooks
mkt36z playbook list
mkt36z playbook run product-launch "AI CRM for startups"

# 5. Analyze your positioning
mkt36z analyze positioning

# 6. Check your setup
mkt36z doctor
```

## What you can do

### Generate copy

```sh
mkt36z generate headline "your product"        # Headlines
mkt36z generate email "topic"                   # Email sequences
mkt36z generate landing-page "product"          # Landing pages
mkt36z generate ad "campaign theme"             # Ad copy
```

### Run playbooks

44 deep marketing playbooks encoding frameworks from Schwartz, Hormozi, Ogilvy, and 20+ marketing masters.

```sh
mkt36z playbook list                            # Browse all playbooks
mkt36z playbook show brand-foundation           # Read the playbook
mkt36z playbook run brand-foundation "my SaaS"  # Execute with AI agents
```

### Analyze and plan

```sh
mkt36z analyze positioning                      # Positioning diagnosis
mkt36z analyze competitors                      # Competitive analysis
mkt36z plan create "Launch AI CRM"              # 90-day marketing plan
mkt36z campaign create "Q1 launch"              # Multi-agent campaign
```

### Templates and workflows

```sh
mkt36z template list                            # 19 ready-to-use templates
mkt36z template run cold-email "SaaS outreach"  # Execute a template
mkt36z workflow list                            # 12 step-by-step workflows
mkt36z workflow run content-calendar "Q1 blog"  # Run a workflow
```

## Architecture

mkt36z is a **thin CLI client**. All intelligence lives server-side:

```
Your terminal                    Cloud
  mkt36z CLI  ──── HTTPS ────>  mkt36z API (Cloudflare Workers)
  (this repo)                     │
                                  ├── 16 specialized AI agents
                                  ├── 44 marketing playbooks
                                  ├── Multi-agent orchestration
                                  └── Governance & quality gates
```

The CLI fetches content from the API and caches it locally (`~/.cache/mkt36z/`) for 24 hours. After the first fetch, list and show commands are instant.

## Configuration

```sh
# Global config
~/.mkt36z/config.yaml           # API URL, preferences
~/.mkt36z/auth.json             # Authentication token
~/.mkt36z/context/              # Your brand context files

# Cache
~/.cache/mkt36z/                # Cached playbooks, templates, history

# Project-local (optional)
.mkt36z.yaml                    # Per-project config
.mkt36z/playbooks/              # Custom playbooks for this project
.mkt36z/templates/              # Custom templates for this project
```

## Documentation

- [Full documentation](https://docs.mkt36z.com)
- [API reference](https://docs.mkt36z.com/api)
- [Releases & changelog](https://github.com/mkt36z/cli/releases)

## License

MIT
