###ENV FILE

GIT DATABASE

HOST="localhost"

USER="dolganoffadmin"

PASSWORD="dolganoffadmin"

DBNAME="web_chat"

PORT="5432"

JWT

REFRESH_SECRET="ugabuga"

ACCESS_SECRET="goroutine"

GOOGLE_AUTH

CLIENT_ID="891267034828-70e6fu6er6v2im3de354v7h49rv2nm2h.apps.googleusercontent.com"

CLIENT_SECRET="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"

REDIRECT_URL="http://localhost:8000/auth/google-callback"

USERINFO_EMAIL="https://www.googleapis.com/auth/userinfo.email"

USERINFO_PROFILE="https://www.googleapis.com/auth/userinfo.profile"

GOOGLE_APIS="https://www.googleapis.com/oauth2/v2/userinfo?access_token="

###TODO

Подключить google cloud platform, подключить бд

###ROUTING

SignUp -> /auth/sign-up

SignIn -> /auth/sign-in

LogOut -> /auth/log-out

GoogleLogin -> /auth/google-login

GoogleCallback -> /auth/google-callback

GetUserList -> /user/list

EditUser -> /user/edit

GetUserInfo -> /user/

SendMessage -> /chat/

