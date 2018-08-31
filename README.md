# goriila_bot
slack app when member join channel, send message to user 

# Note
this bot is only working http, it's not working https

# How to use
### 1. change config.yaml

```
protocol: http # http or https
port: :8080 # server port
endpoint: /slack/gorilla # event api request url 
sslcertificatefile: config/server.crt # ssl cert file *but it's not working...
sslcertificatekeyfile: config/server.key # ssl private key file *but it's not working...
urlverifytoken: Jhj5dZrVaK7ZwHHjRyZWjbDl # set your verification token 
authorizationtoken: Jhj5dZrVaK7ZwHHjRyZWjbDl # set your bot user's auth token
messagefile: config/message.txt # send message read from this txet
watchchannels: # send message when member join this channels
    - about-me
    - question
```

### 2. build bot
```
cd $GOPATH/src
git clone https://github.com/skanehira/gorilla_bot.git
cd gorilla_bot
go get ./...
go build
```

### 3. start bot
```
sudo bash start_bot.sh
```

### 4. if you want see log
```
tail -f gorilla_bot.YYYY-mm-dd.log
```
