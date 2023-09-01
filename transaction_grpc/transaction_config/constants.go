package transactionconstants

const ConnectionString = "mongodb+srv://jp:jp@cluster0.4ipjx6l.mongodb.net/?retryWrites=true&w=majority"
const Port = ":6010"




//json web tokens are created in different ways.They will take an object encrypted with long pharse time of expiration and salt(algorithm).
//For every request we will send the token and that server is going to intercept the token and decrypts it..If it is able to decryt then we consider it as a vaild token.
//Otherwise invaild token
