#!/bin/sh

set -eu

REPO="theoldzoom/gofm"
BINARY_NAME="gofm"
BIN_DIR="${GOFS_INSTALL_DIR:-$HOME/.local/bin}"

need_cmd() {
    if ! command -v "$1" >/dev/null 2>&1; then
        echo "Error: required command not found: $1" >&2
        exit 1
    fi
}

fetch() {
    url="$1"

    if command -v curl >/dev/null 2>&1; then
        curl -fsSL "$url"
        return
    fi

    if command -v wget >/dev/null 2>&1; then
        wget -qO- "$url"
        return
    fi

    echo "Error: install requires curl or wget" >&2
    exit 1
}

detect_os() {
    case "$(uname -s)" in
        Linux) printf '%s' "linux" ;;
        Darwin) printf '%s' "darwin" ;;
        *)
            echo "Error: unsupported operating system: $(uname -s)" >&2
            exit 1
            ;;
    esac
}

detect_arch() {
    case "$(uname -m)" in
        x86_64|amd64) printf '%s' "amd64" ;;
        aarch64|arm64) printf '%s' "arm64" ;;
        *)
            echo "Error: unsupported architecture: $(uname -m)" >&2
            exit 1
            ;;
    esac
}

need_cmd uname
need_cmd sed
need_cmd awk
need_cmd mktemp
need_cmd chmod
need_cmd mkdir
need_cmd mv

OS="$(detect_os)"
ARCH="$(detect_arch)"
ASSET_NAME="${BINARY_NAME}-${OS}-${ARCH}"
API_URL="https://api.github.com/repos/${REPO}/releases/latest"

echo "Fetching latest release metadata for ${REPO}..."
RELEASE_JSON="$(fetch "$API_URL")"

DOWNLOAD_URL="$(
    printf '%s\n' "$RELEASE_JSON" |
        sed -n 's/.*"browser_download_url":[[:space:]]*"\([^"]*\)".*/\1/p' |
        awk -v asset="$ASSET_NAME" 'index($0, "/" asset) { print; exit }'
)"

if [ -z "$DOWNLOAD_URL" ]; then
    echo "Error: could not find release asset '${ASSET_NAME}' in the latest release" >&2
    exit 1
fi

mkdir -p "$BIN_DIR"

TMP_FILE="$(mktemp)"
trap 'rm -f "$TMP_FILE"' EXIT INT TERM HUP

echo "Downloading ${ASSET_NAME}..."
fetch "$DOWNLOAD_URL" > "$TMP_FILE"

chmod +x "$TMP_FILE"
mv "$TMP_FILE" "$BIN_DIR/$BINARY_NAME"
trap - EXIT INT TERM HUP

echo "Installed ${BINARY_NAME} to ${BIN_DIR}/${BINARY_NAME}"

case ":$PATH:" in
    *":$BIN_DIR:"*)
        ;;
    *)
        echo "Warning: ${BIN_DIR} is not in your PATH" >&2
        echo "Add this to your shell profile:" >&2
        echo "  export PATH=\"${BIN_DIR}:\$PATH\"" >&2
        ;;
esac