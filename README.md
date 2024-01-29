# Dummy Backend Wallet

[WIP]

This is for experimental backend wallet API service with [Shamir's Secret Sharing](https://en.wikipedia.org/wiki/Shamir%27s_secret_sharing) and Email or Social Login without operating users to generate Seed Phrase anymore.

But operate with User's Private Key onto server side itself is dangerous, So try with [aka. MPC Key Management](https://ja.wikipedia.org/wiki/%E7%A7%98%E5%8C%BF%E3%83%9E%E3%83%AB%E3%83%81%E3%83%91%E3%83%BC%E3%83%86%E3%82%A3%E8%A8%88%E7%AE%97) not exactly just by own server ways

## Environment

- MacOS Intel Core 13.2.1
- [Golang 1.21.x](https://go.dev/dl/)

## Prerequisite

- Hashicorp Vault

  - Installatiion

    ```shell
    brew tap hashicorp/tap

    brew install hashicorp/tap/vault

    ### To update to the latest, run
    brew upgrade hashicorp/tap/vault
    ```

  - Starting the server
    ```shell
    vault server -dev
    ```
