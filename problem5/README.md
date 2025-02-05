# Switcheo Labs Backend

## _Set Up:_

### Pre-requisite:
1. This project is build and run with **Ignite CLI @v28.7.0**, **Cosmos SDK @v0.50.11**


### Start the blockchain

1. cd to `crude` folder
2. Run `ignite chain serve` in terminal
3. For more details: `ignite chain serve --verbose`


## _Consensus Breaking Change_

### Change Made to Exsiting BlockChain

*   **Change**: Addition of `mobileNo` field of type `uint64` to the `User` data structure in user.proto
*   **Reason for Change**: To include the mobile number of users as one of the user details in the blockchain message.



### Questions

**_1. Explain what does it mean by breaking consensus_**

Breaking consensus means making a change to the blockchain protocol or data structure that:
- Causes nodes running the old version to disagree with nodes running the new version
- Invalidates previously accepted transactions or blocks for nodes that have not upgraded
- Requires all nodes to upgrade to stay in sync with the network


**_2. Explain why your change would break the consensus._**
    
a. Old Nodes Won’t Recognize the New Field. When they receive a transaction or block containing the new field (mobileNo), they cannot parse the data and will reject the block. Stored Data Format Changes

b. If you store User data in key-value format, the stored bytes will no longer match what old nodes expect. Old nodes might fail to decode new users, leading to failed queries and transactions.

c. Blockchains use cryptographic hashes to verify transactions. Since mobileNo changes the data structure, any transaction or block signature created under the old rules will not match the new format. Nodes running different versions will disagree on block validity.


