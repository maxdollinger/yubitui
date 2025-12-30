# Yubitui

Yubitui is a terminal user interface (TUI) for managing 2-factor-authentication (2FA) codes stored on a YubiKey. It provides a convenient way to view and manage your OTP credentials directly in the terminal, acting as a wrapper around the `ykman` CLI tool.

## Why?

This tool was created to provide a simple and efficient way to access YubiKey OTP codes within the terminal, avoiding the need to switch to a graphical application. It uses `ykman` as a dependency to interact with the YubiKey due to complexities with the native Go libraries, especially concerning account renaming.

## Installation

Before you begin, ensure you have `ykman` installed and accessible in your system's PATH.

1.  **Install `ykman`:**
    Follow the official instructions to install the [YubiKey Manager (ykman)](https://support.yubico.com/hc/en-us/articles/360013789900-YubiKey-Manager-CLI-ykman-User-Manual).

2.  **Install Yubitui:**
    ```sh
    go install github.com/mdollinger/yubitui@latest
    ```

## Usage

Simply run the application from your terminal:

```sh
yubitui
```

The interface will display your 2FA codes. You can navigate the list with the arrow keys and press `Enter` to copy a code to your clipboard.
