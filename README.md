# userRegistration
Sample golang rest api for user registration

Routes 
=======

**GET /user** - get all users

Sample Output
```
[{
EmailID: "prk@gmail.com",
Password: "test",
Name: "prk",
City: "SF"
},
{
EmailID: "prk1@gmail.com",
Password: "test",
Name: "prk",
City: "sf"
}]
```
**GET /user?emailid={string}** - get specific user by id

Sample Output
```
{
EmailID: "prk@gmail.com",
Password: "test",
Name: "prk",
City: "SF"
}
```

**POST /user** - create new user

##### Headers
	Content-type - application/json
##### Body 
    emailid {string}
    password {string}
    city {string}
    name {string}
