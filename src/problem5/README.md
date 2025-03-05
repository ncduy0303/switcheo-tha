# Task

Create a practical blockchain with Cosmos SDK that provides a basic set of CRUD interface that allow a user to interact with the blockchain.

1. Interface functionalities:
   1. Create a resource.
   2. List resources with basic filters.
   3. Get details of a resource.
   4. Update resource details.
   5. Delete a resource.
2. Create a git branch with a _consensus-breaking_ change introduced to your blockchain, and document in a `README.md` file.
   1. Explain what does it mean by breaking consensus.
   2. Explain why your change would break the consensus.

## Context

In the dynamic landscape of blockchain technology, we encounter two primary architectural paradigms: the virtual machine blockchain and the application-specific blockchain. Each of these frameworks presents a unique approach to the development of decentralized applications. The virtual machine blockchain, exemplified by the Ethereum Virtual Machine (EVM), is distinguished by its use of languages like Solidity for smart contract deployment. On the other hand, the application-specific blockchain, as demonstrated by Cosmos SDK, creates decentralized applications through modules.

In this task, we invite you to delve into the Cosmos SDK approach. Your challenge is to demonstrate your understanding by creating a practical CRUD-based decentralized application using Cosmos SDK. This exercise will not only gauge your technical skills but also your ability to leverage Cosmos SDK's unique features for data manipulation and transaction execution.

## Implementation

Refer to the `crude` directory where I added the module `addressbook` to the chain. This decentralized address book provides a basic CRUD interface for managing contacts where each contact has a name, phone number, email, address as well as a unique id. Users can create, list, get, update, and delete contacts. To launch the blockchain, ensure Ignite CLI is installed and execute the following commands in the terminal (`config.yml` is the configuration file for the chain with sample genesis data):

```bash
cd crude
ignite chain serve -c config.yml
```

### Create a Contact

To create a contact, execute the following command:

```bash
cruded tx addressbook create-contact "John Smith" "555-555-5555" "johns@yahoo.com" "123 Main St" --from alice
```

### Update a Contact

Each contact can only be updated by the user who created it. To update a contact, execute the following command:

```bash
cruded tx addressbook update-contact 3 "John Smith" "444-444-4444" "johnsmith@yahoo.com" "123 Main St" --from alice
```

### Delete a Contact

Each contact can only be deleted by the user who created it. To delete a contact by id, execute the following command:

```bash
cruded tx addressbook delete-contact 3 --from alice
```

### Show a Contact

To get the details of a contact by id, execute the following command:

```bash
cruded query addressbook show-contact 0
```

### List Contacts

To list all contacts, execute the following command:

```bash
cruded query addressbook list-contact
```

To filter contacts by name, phone number, email, or address, execute the following command:

```bash
# cruded query addressbook list-contact-filter [name] [phone] [email] [address]
cruded query addressbook list-contact-filter "Johnson" "" "" ""
```

## Consensus-Breaking Change

A consensus breaking change is one that alters the state of the blockchain in a way that is incompatible with the previous state. This change can result in nodes disagreeing on the state of the blockchain, causing inconsistencies and potential issues with transaction validation and block creation. It requires all nodes to update to the new state to maintain consensus.

In the `problem5-consensus-breaking` branch, I added a new field `Affiliation` to the `Contact` structure in the `addressbook` module. This field is not compatible with the previous state format, which only contained `Id`, `Name`, `Phone`, `Email`, and `Address`. As a result, nodes running the previous version of the blockchain will not be able to correctly interpret or interact with contacts that have the new structure. This change would break the consensus because the old chain would not recognize the new field.
