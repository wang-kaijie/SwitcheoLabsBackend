# Switcheo Labs Backend

## _Set Up:_

### Pre-requisite:
1. This project is build and run with **Ignite CLI @v28.7.0**, **Cosmos SDK @v0.50.11**



### Start the blockchain

1. cd to `crude` folder
2. Run `ignite chain serve` in terminal


### Start the interface

1. cd to `web` folder
2. Run `go run main.go` in terminal
3. cd to static folder and open `index.html` for the UI


## _Resource:_ 
This project is built using ignite and GO and the blockchain keeps track of `users` details
### User

- User data structure: `id`, `name`, `email`, `username`, `password`, `address`

### Interface functionalities

- Create a User
- Get All User
- Get User by ID 
- Get Users by (basic flitering)
- Get Users by (basic flitering)
- Update User Detail
- Update User Password Function
- Delete a User
