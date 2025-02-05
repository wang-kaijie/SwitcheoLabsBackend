# Switcheo Labs Backend

## _Set Up:_

### Pre-requisite:
1. This project is build and run with **Ignite CLI @v28.7.0**, **Cosmos SDK @v0.50.11**


### Start the blockchain

1. cd to `crude` folder
2. Run `ignite chain serve` in terminal


## _Resource:_ 
This project is built using ignite and GO and the blockchain keeps track of `users` details
### User

- User data structure: `id`, `name`, `email`, `username`, `password`, `address`

### Interface functionalities (test cases)

- Create a User : `cruded tx crude create-user "Alice" "alice@example.com" 25 "Female" "123 Blockchain St." --from alice --chain-id crude --yes`
- Get All User: `cruded query crude allUsers --chain-id crude`
- Get User by ID: `cruded query crude specificUser 0 --chain-id crude`
- Get User by ID (error): `cruded query crude specificUser 9999 --chain-id crude`
- Get Users by (basic flitering)
- Get Users by (basic flitering)
- Update User Detail: `cruded tx crude update-user 1 "Alice Updated" "alice_new@example.com" 26 "Female" "789 Blockchain Blvd." --from alice --chain-id crude --yes`
- Delete a User: `cruded tx crude delete-user 0 --from alice --chain-id crude --yes`
- Delete a User(error): `cruded tx crude delete-user 9999 --from alice --chain-id crude --yes`
