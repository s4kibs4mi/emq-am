## EMQ-AM
EMQ-AM let you control Authentication & ACL of EMQ's MQTT Broker using simplified rest API
.

* #### User Registration

###### [ POST {{host}}/users ]
```json
{
	"user_name": "user1",
	"email": "user1@domain.com",
	"password": "123456789"
}
```
###### Response
```
{
    "code": 200,
    "details": "User successfully created",
    "data": {
        "id": "5a1316b92315e6823a298390",
        "user_name": "s4kibs4mi12",
        "email": "root12@sakib.ninja",
        "type": "default",
        "created_at": "2017-11-20T23:54:01.967+06:00",
        "updated_at": "2017-11-20T23:54:01.967+06:00"
    }
}
```
Note : First user will be registered as Admin and then are as default.

* #### User Login

###### [ POST {{host}}/session ]
```json
{
	"user_name": "s4kibs4mi",
	"password": "123456789"
}
```
###### Response
```
{
    "code": 200,
    "data": {
        "user_id": "5a11edeb2315e67eaef32f9b",
        "access_token": "502bfb33-94a9-40c0-9ebd-a99bdc6ae308",
        "refresh_token": "c0906a75-4df1-488c-ab80-fc392e14471c",
        "created_at": "2017-11-20T02:48:25.977248+06:00",
        "expire_at": "2017-11-21T02:48:25.977248+06:00"
    }
}
```

* #### Append Publish Topic

###### [ POST {{host}}/publish ]
- Header
    - user_id: "123123123123"
    - access_token: "adsfjnajslkJLKDJASLDKN"
```json
{
	"user_id": "5a15b8212315e690dfce5889",
	"topic": "hello3"
}
```
###### Response
```json
{
    "code": 200,
    "details": "Publish topic updated",
    "data": {
        "id": "5a15b8212315e690dfce5889",
        "user_name": "s4kibs4mi1",
        "email": "root1@sakib.ninja",
        "publish_topics": [
            "hello3"
        ],
        "subscribe_topics": [
            "hello3"
        ],
        "type": "default",
        "status": "unbanned",
        "created_at": "2017-11-22T23:47:13.85+06:00",
        "updated_at": "2017-11-22T23:47:13.85+06:00"
    }
}
```

* #### Remove Publish Topic

###### [ DELETE {{host}}/publish ]
- Header
    - user_id: "123123123123"
    - access_token: "adsfjnajslkJLKDJASLDKN"
```json
{
	"user_id": "5a15b8212315e690dfce5889",
	"topic": "hello3"
}
```
###### Response
```json
{
    "code": 200,
    "details": "Publish topic removed",
    "data": {
        "id": "5a15b8212315e690dfce5889",
        "user_name": "s4kibs4mi1",
        "email": "root1@sakib.ninja",
        "subscribe_topics": [
            "hello3"
        ],
        "type": "default",
        "status": "unbanned",
        "created_at": "2017-11-22T23:47:13.85+06:00",
        "updated_at": "2017-11-22T23:47:13.85+06:00"
    }
}
```

* #### Append Subscribe Topic

###### [ POST {{host}}/subscribe ]
- Header
    - user_id: "123123123123"
    - access_token: "adsfjnajslkJLKDJASLDKN"
```json
{
	"user_id": "5a15b8212315e690dfce5889",
	"topic": "hello3"
}
```
###### Response
```json
{
    "code": 200,
    "details": "Subscribe topic updated",
    "data": {
        "id": "5a15b8212315e690dfce5889",
        "user_name": "s4kibs4mi1",
        "email": "root1@sakib.ninja",
        "publish_topics": [
            "hello3"
        ],
        "subscribe_topics": [
            "hello3"
        ],
        "type": "default",
        "status": "unbanned",
        "created_at": "2017-11-22T23:47:13.85+06:00",
        "updated_at": "2017-11-22T23:47:13.85+06:00"
    }
}
```

* #### Remove Subscribe Topic

###### [ DELETE {{host}}/subscribe ]
- Header
    - user_id: "123123123123"
    - access_token: "adsfjnajslkJLKDJASLDKN"
```json
{
	"user_id": "5a15b8212315e690dfce5889",
	"topic": "hello3"
}
```
###### Response
```json
{
    "code": 200,
    "details": "Subscribe topic removed",
    "data": {
        "id": "5a15b8212315e690dfce5889",
        "user_name": "s4kibs4mi1",
        "email": "root1@sakib.ninja",
        "subscribe_topics": [
            "hello3"
        ],
        "type": "default",
        "status": "unbanned",
        "created_at": "2017-11-22T23:47:13.85+06:00",
        "updated_at": "2017-11-22T23:47:13.85+06:00"
    }
}
```

* #### EMQ Authentication

###### [ POST {{host}}/auth ]
```text
username=${User_Id}&password=${Access_Token}
```
###### Response
- Http Status : 200 if authenticated

* #### EMQ Authorization

###### [ POST {{host}}/acl ]
```text
username=${User_Id}&topic=hello&access=2
```
###### Response
- Http Status : 200 if authenticated


### Configuration
Change configuration in etc/config.json

```json
{
  "app": {
    "address": ":8090"
  },
  "databases": {
    "mongodb": {
      "uri": "mongodb://localhost:27017",
      "name": "emqam",
      "auth_collection": "emqauth",
      "acl_collection": "emqacl",
      "session_collection": "sessions"
    }
  },
  "security": {
    "registration_enabled": true,
    "key": "hello",
    "secret": "12345"
  },
  "pagination": {
    "per_page": 20
  }
}
```

### Docker Deployment
```
docker pull sakibsami/emq-am:latest
docker run -p 8090:8090 sakibsami/emq-am:latest
```
