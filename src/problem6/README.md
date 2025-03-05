# Task

This is a system design question. Describe in detail (~500-1000 words) the specifications on how you would design a transaction broadcaster service. You may additionally attach a drawings/diagrams/illustrations if you wish.

- You should focus on designing the software abstractions and architectural flow required to fulfill the service requirements, instead of choosing cloud services or software packages.
- Your submission will be graded on correctness, scalability, and robustness.

1. The broadcast service exposes an internal api for broadcast requests that will be used by other services.
   It returns HTTP `200`, or HTTP `4xx`-`5xx` .

   ````jsx
   POST /transaction/broadcast

       {"message_type": "add_weight(address _addr, uint256 _weight)", "data": "0xd71363280000000000000000000000005eb715d601c2f27f83cb554b6b36e047822fb70a00000000000000000000000000000000000000000000000000000000000000fa"}
       ```
   ````

2. Using the post request parameters, the broadcaster service signs the `data` and outputs a `signed transaction`. Next, it broadcasts the `signed transaction` to an EVM-compatible blockchain network.

   1. A broadcasted transaction might fail and if it fails, it should be retried automatically.
   2. To broadcast a signed transaction, you make a RPC request to a blockchain node.
      1. 1% of the time, it does not respond earlier than 30 seconds.
      2. 95% of the time it responds with a success code within 20-30 seconds.
      3. The rest of the time it returns a failure code.
   3. There should also be a page that shows the list of transactions that passed or failed.

Additional Requirements

1. If `POST /transaction/broadcast` returns HTTP `200 OK`, it is assumed that the transaction will eventually be broadcasted successfully. If the broadcaster service restarts unexpectedly, it should still fulfill them.
2. An admin is able to, at any point in time, retry a failed broadcast.

## Design

### 1. System Components

#### A. Persistence Storage

- **Role:** Stores transaction metadata to enable retry mechanisms for failed transactions.
- **Schema:**
  - `transaction_id`: Unique identifier.
  - `message_type`: Type of transaction.
  - `data`: Transaction data.
  - `signed_data`: Signed transaction data.
  - `status`: Current status (`PENDING`, `SUCCESS`).
  - `retry_count`: Number of retries attempted.
  - `created_at`: Timestamp of transaction creation.
  - `updated_at`: Timestamp of last status update attempt.
  - `error_message`: Error message on failure.

#### B. Transaction API Handler Worker

- **Role:** Exposes HTTP POST endpoint, receives transaction requests, signs them and persists them in the database.

#### C. Transaction Broadcaster Master

- **Role:** Schedule transactions for broadcasting to available workers, update transaction status in the database, and handle manual retries from the admin dashboard.

#### D. Transaction Broadcaster Worker

- **Role:** Broadcasts signed transactions to the blockchain node and sends the status back to the master.

#### E. Admin Dashboard Service

- **Role:** Provides a user interface for transaction monitoring and manual intervention.

### 2. Architectural Flow

#### Step 1: Receiving Transaction Requests

- Client posts a request at `/transaction/broadcast`:

  ```json
  POST /transaction/broadcast
  {
    "message_type": "add_weight(address _addr, uint256 _weight)",
    "data": "0xd71363...00000fa"
  }
  ```

- Upon receiving the request, the **Transaction API Handler Worker**:
  - Validates the request format.
  - Signs the transaction data.
  - Stores the signed transaction in **Persistence Storage** with status `PENDING`.
  - Returns HTTP `200 OK` immediately upon successful persistence, or appropriate error codes `4xx`-`5xx` for invalid requests or internal errors.

#### Step 2: Automatic Broadcasting Transactions

- **Transaction Broadcaster Master**:
  - Polls for `PENDING` transactions from **Persistence Storage** periodically.
  - Schedules transactions for broadcasting based on their retry count and last attempt timestamp according to an exponential backoff strategy. For example, retry delays: 5s → 10s → 20s → 40s → 80s → ...
  - Assigns transactions to available workers for broadcasting.
  - Updates transaction status in the persistent databse based on worker responses:
    - `SUCCESS` if transaction successfully broadcasted.
    - Keep status `PENDING`, add error message and increment retry count if the transaction failed, then schedule for retry later.
- **Transaction Broadcaster Worker**:
  - Broadcasts signed transactions assigned by the master via RPC to an EVM-compatible blockchain node.
  - Sends the status back to the master, success or failure/timeout with error message (failure code). Since 95% of the time, the response is within 20-30 seconds, and 1% of the time it does not respond earlier than 30 seconds, the worker should set a maximum timeout for the RPC request to be slightly above 30 seconds (maybe 60 seconds).

#### Step 3: Manual Broadcasting Transactions & Admin Dashboard

- **Admin Dashboard Service**:

  - Display a paginated list of all transactions with information such as current status, retry count, time created, and last updated.
  - Provide basic searching and filtering options for the transactions based on status, retry count, date range, etc., to allow admins to efficiently monitor transaction flows.
  - Allow manual retry of `PENDING` transactions immediately by sending the request to the broadcaster master.

- **Transaction Broadcaster Master**:
  - Upon manual retry request, immediately schedule the transaction for broadcasting.

### 3. Robustness & Scalability

- **Robustness**:
  - If the broadcaster service restarts unexpectedly, the **Transaction Broadcaster Master** will recover the state of `PENDING` transactions from the persistent storage and reschedule them for broadcasting, ensuring no data loss.
  - Automatic retries with exponential backoff strategy ensure that failed transactions are eventually broadcasted successfully.

- **Scalability**:
  - The system can be scaled horizontally by adding more **Transaction Broadcaster Workers** and **Transaction API Handler Workers** to handle increased transaction volume.
  - A load balancer can be used to distribute incoming transaction requests across multiple instances of the **Transaction API Handler Worker**.
  - A single **Transaction Broadcaster Master** can manage multiple **Transaction Broadcaster Workers** to handle broadcasting of transactions concurrently.

### 4. Notes

- Currently, the system keeps retrying failed transactions indefinitely whenever it sends a HTTP `200 OK` back until it succeeds. It might be better to set a maximum number of retries to avoid spamming the blockchain network with repeated failed transactions, then the transaction should be marked as `FAILED` and require manual intervention, although this might affect the requirement that if `POST /transaction/broadcast` returns HTTP `200 OK`, it is assumed that the transaction will eventually be broadcasted successfully. This can be further discussed.
