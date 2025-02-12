# Switcheo Labs Backend

## _Set Up:_

### Pre-requisite:
1. This project is build and run with **Ignite CLI @v28.7.0**, **Cosmos SDK @v0.50.11**


### Start the blockchain

1. cd to `crude` folder
2. Run `ignite chain serve` in terminal
3. For more details: `ignite chain serve --verbose`


## _Resource:_ 
This project is built using ignite and GO and the blockchain keeps track of `users` details
### User

- User data structure: `id`, `name`, `email`, `age`, `gender`, `address`

### Interface functionalities (test cases, execution in sequence)

- **Create a User** : 
    - `cruded tx crude create-user "Alice" "alice@example.com" 25 "Female" "123 Blockchain St." --from alice --chain-id crude --yes`
    - `cruded tx crude create-user "Bob" "bob@example.com" 25 "Male" "456 Blockchain St." --from alice --chain-id crude --yes`
    - `cruded tx crude create-user "Charles" "charles@example.com" 30 "Male" "789 Blockchain St." --from alice --chain-id crude --yes`
- **Get All User**: 
    - `cruded query crude all-users --chain-id crude`
- **Get User by ID**: 
    - Alice: `cruded query crude specific-user 0 --chain-id crude`
    - Bob: `cruded query crude specific-user 1 --chain-id crude`
    - Charles: `cruded query crude specific-user 2 --chain-id crude`
- **Get User by ID (error):** 
    - `cruded query crude specific-user 9999 --chain-id crude`
- **Get Users by age (basic flitering for age(uint))**: 
    - Alice and Bob: `cruded query crude users-by-age 25 --chain-id crude`
- **Get Users by gender (basic flitering for "Female" or "Male"):** 
    - Alice: `cruded query crude users-by-gender "Female" --chain-id crude`
- **Update User Detail:** 
    - `cruded tx crude update-user 0 "Alice Updated" "alice_new@example.com" 26 "Female" "789 Blockchain Blvd." --from alice --chain-id crude --yes`
    - check: `cruded query crude specific-user 0 --chain-id crude`
- **Delete a User:** 
    - `cruded tx crude delete-user 0 --from alice --chain-id crude --yes`
    - check: `cruded query crude specific-user 0 --chain-id crude`






