# Rysk V12 CLI

A command-line interface (CLI) for interacting with the Rysk v12 protocol via WebSockets.

This CLI allows you to connect to a WebSocket server, send signed messages for actions like approving token spending, initiating transfers, and sending quotes. It utilizes Unix sockets for inter-process communication, enabling you to send commands to a running WebSocket connection from other processes.

## Prerequisites

- **Go (Golang) installed:** This project is written in Go and requires a Go development environment to build.
- **Ethereum Node Access:** For the `approve` command, you'll need access to an Ethereum node (e.g., via Infura, Alchemy, or a local node) corresponding to the specified `rpc_url`.

## Installation

1.  **Clone the repository (if the code is in one):**

    ```bash
    git clone <repository_url>
    cd <repository_directory>
    ```

2.  **Build the CLI:**
    ```bash
    go build -o ryskV12
    ```
    This will create an executable file named `ryskV12` in the current directory.

## Usage

The `ryskV12` CLI provides the following commands:

### `approve`

Approves spending of the default strike asset for a given account.

```bash
./ryskV12 approve --chain_id <chain_id> --rpc_url <rpc_url> --amount <amount> --private_key <private_key>
```

Flags

- `--chain_id` (**required**): The ID of the blockchain.
- `--rpc_url` (**required**): The URL of the Ethereum RPC endpoint.
- `--amount` (**required**): The amount of the asset to approve for spending.
- `--private_key` (**required**): The private key of the Ethereum account performing the approval.

---

### `balances`

Retrieves USDC balances for the specified account.

```bash
./ryskV12 balances --channel_id <channel_id> --account <0xabc>
```

Flags

- `--account` (**required**): The address to query data for.
- `--channel_id` (**required**): Unique ID for the connection and named pipe (/tmp/<channel_id>).

---

### `connect`

Establishes a WebSocket connection and runs in daemon mode with a named pipe.

```bash
./ryskV12 connect --channel_id <channel_id> --url <websocket_url>
```

Flags

- `--channel_id` (**required**): Unique ID for the connection and named pipe (/tmp/<channel_id>).
- `--url` (**required**): WebSocket URL to connect to.

Endpoints:

- `wss://<base_url>/rfqs/<asset_address>` listen for rfqs for the specified asset
- `wss://<base_url>/maker` endpoint to send quotes and transfer requests

---

### `positions`

Retrieves positions (oToken details) for the specified account
```bash
./ryskV12 positions --channel_id <channel_id> --account <0xabc>
```

Flags

- `--account` (**required**): The address to query data for.
- `--channel_id` (**required**): Unique ID for the connection and named pipe (/tmp/<channel_id>).

Endpoints:

- `wss://<base_url>/rfqs/<asset_address>` listen for rfqs for the specified asset
- `wss://<base_url>/maker` endpoint to send quotes and transfer requests

---

### `quote`

Sends a signed quote for options trading through the WebSocket.

```bash
./ryskV12 quote --channel_id <channel_id> --rfq_id <rfq_id> --chain_id <chain_id> --expiry <expiry_timestamp> --is_put --is_taker_buy --maker <maker_address> --nonce <nonce> --price <price> --quantity <quantity> --strike <strike> --valid_until <valid_until_timestamp> --private_key <private_key>
```

Flags

- `--channel_id` (**required**): The unique ID of the WebSocket connection.
- `--rfq_id` (**required**): The unique ID of the rfq you are quoting for.
- `--chain_id` (**required**): The ID of the blockchain.
- `--expiry` (**required**): Option expiry timestamp.
- `--is_put`: present for put, not for call.
- `--is_taker_buy`: present if maker sells, not if maker buys.
- `--maker` (**required**): Address of the quote maker.
- `--nonce` (**required**): Unique nonce for the quote.
- `--price` (**required**): Option price.
- `--quantity` (**required**): Option quantity.
- `--strike` (**required**): Option strike price.
- `--valid_until` (**required**): Quote validity timestamp.
- `--private_key` (**required**): Private key for signing.

---

### `transfer`

Requests a transfer (deposit or withdrawal) through the WebSocket.

```bash
./ryskV12 transfer --channel_id <channel_id> --chain_id <chain_id> --asset <asset_address> --amount <amount> --is_deposit --nonce <nonce> --private_key <private_key>
```

Flags

- `--channel_id` (**required**): The unique ID of the WebSocket connection (matches connect --channel_id).
- `--chain_id` (**required**): The ID of the blockchain for the transfer.
- `--asset` (**required**): The address of the asset being transferred.
- `--amount` (**required**): The amount to transfer.
- `--is_deposit`: present if deposit, not for withdrawal.
- `--nonce` (**required**): A unique nonce for signing.
- `--private_key` (**required**): The private key for signing.
