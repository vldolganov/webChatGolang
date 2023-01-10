``Need .env file``
for google auth:
```
CLIENT_ID="891267034828-70e6fu6er6v2im3de354v7h49rv2nm2h.apps.googleusercontent.com"

CLIENT_SECRET="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"

REDIRECT_URL="http://localhost:8000/auth/google-callback"
```
Routing 
```
signUp | localhost:8000/auth/sign-up | method: POST | accepts: login(not null), first_name, last_name, email(not null), 
password(not null), avatar | return: json object with id and other fields 

signIn | localhost:8000/auth/sign-in | method: POST | accepts: login, password | return: json object with id and 
other fields

logOut | localhost:8000/auth/logout | method: GET | no accepts | return cleared cookie

googleLogin | localhost:8000/auth/google-login | method: GET | no accepts | redirect to google-callback

googleCallback | localhost:8000/auth/google-callback | method: GET | no accepts | return jsonObj with email, first_name,
last_name and id

getUsersList | localhost:8000/user/list | only for auth user | method: GET | return full users list in JSON objects
from database

editUser | localhost:8000/user/list  | only for auth user | method:PUT | accepts : email or avatar or login 
or all of this fields | return updated fields 

getUserInfo | localhost:8000/user/ | only for auth user | method:GET | accepts auth user id | 
return only auth user info

sendMessage | localhost:8000/chat/ | only for auth user | method:POST | accepts:from user id(from cookie) and toUserId(consumer user)
| return JSON obj with msg id, owner id and consumer id

getMessages | localhost:8000/chat/ | only for auth user | method:GET | accepts:toUserId(consumer user)
| return JSON obj with all msg id, owner id and consumer id in selected chat
```
for start project type this command in terminal 
`go run main.go`

