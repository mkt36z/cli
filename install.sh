#!/bin/sh
# mkt36z CLI installer
# Usage: curl -sSL https://install.mkt36z.com | sh
#
# Environment variables:
#   MKT36Z_VERSION  — specific version to install (default: latest)
#   MKT36Z_DIR      — install directory (default: /usr/local/bin)

set -e

REPO="mkt36z/cli"
BINARY="mkt36z"
DEFAULT_DIR="/usr/local/bin"

# Colors (disabled if not a terminal)
if [ -t 1 ]; then
    BOLD='\033[1m'
    GREEN='\033[32m'
    RED='\033[31m'
    BLUE='\033[34m'
    DIM='\033[2m'
    RESET='\033[0m'
else
    BOLD='' GREEN='' RED='' BLUE='' DIM='' RESET=''
fi

info()  { printf "${BLUE}==>${RESET} ${BOLD}%s${RESET}\n" "$1"; }
ok()    { printf "${GREEN}==>${RESET} ${BOLD}%s${RESET}\n" "$1"; }
err()   { printf "${RED}error:${RESET} %s\n" "$1" >&2; exit 1; }

# Detect OS and architecture
detect_platform() {
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)

    case "$OS" in
        linux)  OS="linux" ;;
        darwin) OS="darwin" ;;
        *)      err "Unsupported OS: $OS. Use Linux or macOS." ;;
    esac

    case "$ARCH" in
        x86_64|amd64)   ARCH="amd64" ;;
        aarch64|arm64)  ARCH="arm64" ;;
        *)              err "Unsupported architecture: $ARCH" ;;
    esac

    PLATFORM="${OS}_${ARCH}"
}

# Get latest version from GitHub releases
get_version() {
    if [ -n "$MKT36Z_VERSION" ]; then
        VERSION="$MKT36Z_VERSION"
    else
        VERSION=$(curl -sSL "https://api.github.com/repos/${REPO}/releases/latest" \
            | grep '"tag_name"' | head -1 | sed 's/.*"v\(.*\)".*/\1/')

        if [ -z "$VERSION" ]; then
            err "Could not determine latest version. Set MKT36Z_VERSION manually."
        fi
    fi
}

# Download and verify
download() {
    INSTALL_DIR="${MKT36Z_DIR:-$DEFAULT_DIR}"
    ARCHIVE="mkt36z_${VERSION}_${PLATFORM}.tar.gz"
    URL="https://github.com/${REPO}/releases/download/v${VERSION}/${ARCHIVE}"
    CHECKSUM_URL="https://github.com/${REPO}/releases/download/v${VERSION}/checksums.txt"

    TMPDIR=$(mktemp -d)
    trap 'rm -rf "$TMPDIR"' EXIT

    info "Downloading mkt36z v${VERSION} for ${PLATFORM}..."
    curl -sSL -o "${TMPDIR}/${ARCHIVE}" "$URL" || err "Download failed. Check version and platform."

    # SECURITY (VULN-12): Verify checksum — never silently skip.
    # Require a checksum tool to be present; abort if missing.
    info "Verifying checksum..."
    curl -sSL -o "${TMPDIR}/checksums.txt" "$CHECKSUM_URL" || err "Failed to download checksums."
    EXPECTED=$(grep "$ARCHIVE" "${TMPDIR}/checksums.txt" | awk '{print $1}')

    if [ -z "$EXPECTED" ]; then
        err "No checksum found for ${ARCHIVE} in checksums.txt. Aborting."
    fi

    if command -v sha256sum >/dev/null 2>&1; then
        ACTUAL=$(sha256sum "${TMPDIR}/${ARCHIVE}" | awk '{print $1}')
    elif command -v shasum >/dev/null 2>&1; then
        ACTUAL=$(shasum -a 256 "${TMPDIR}/${ARCHIVE}" | awk '{print $1}')
    else
        err "sha256sum or shasum is required for integrity verification. Install coreutils and retry."
    fi

    if [ "$ACTUAL" != "$EXPECTED" ]; then
        err "Checksum mismatch! Expected ${EXPECTED}, got ${ACTUAL}. Binary may be tampered with."
    fi

    ok "Checksum verified."

    # Extract
    info "Extracting..."
    tar -xzf "${TMPDIR}/${ARCHIVE}" -C "${TMPDIR}"

    # Install
    if [ -w "$INSTALL_DIR" ]; then
        mv "${TMPDIR}/${BINARY}" "${INSTALL_DIR}/${BINARY}"
    else
        info "Installing to ${INSTALL_DIR} (requires sudo)..."
        sudo mv "${TMPDIR}/${BINARY}" "${INSTALL_DIR}/${BINARY}"
    fi
    chmod +x "${INSTALL_DIR}/${BINARY}"
}

# Post-install
post_install() {
    ok "mkt36z v${VERSION} installed to ${INSTALL_DIR}/${BINARY}"
    echo ""
    printf "${DIM}  Get started:${RESET}\n"
    printf "${DIM}    mkt36z auth login${RESET}\n"
    printf "${DIM}    mkt36z generate headline \"your product\"${RESET}\n"
    printf "${DIM}    mkt36z --help${RESET}\n"
    echo ""

    # Check if in PATH
    if ! command -v mkt36z >/dev/null 2>&1; then
        printf "${RED}  Warning:${RESET} ${INSTALL_DIR} is not in your PATH.\n"
        printf "  Add it:  export PATH=\"${INSTALL_DIR}:\$PATH\"\n"
    fi
}

# Main
main() {
    detect_platform
    get_version
    download
    post_install
}

main
