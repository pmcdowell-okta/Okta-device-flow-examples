## Sample Code Snippets for Device Authentication with Okta

Many of these examples can be found here: https://developer.okta.com/docs/guides/device-authorization-grant/main/

I also include a Postman Collection and a small code snippet so it might save folks a few minutes to setup. The Postman Collection is included in the file name: **"**`Device authentication Flow Samples.postman_collection.json`**"**

## Performing a request from a Client

**Sample Request**

    curl --location --request POST 'https://cbdevice.okta.com/oauth2/aus1l4r7i7vJE9zRF3l7/v1/device/authorize' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --header 'Cookie: JSESSIONID=CE0C1C29F46323C43F38A1D7F2AEAF9A' \
    --data-urlencode 'client_id=0oa1l4rl4c3IQluOA3l7' \
    --data-urlencode 'scope=openid profile offline_access'

**Sample Response**

    {"device_code":"da90717b-c2a6-400b-8bc4-b432e90ce7d6","user_code":"QLWKNBRW","verification_uri":"https://cbdevice.okta.com/activate","verification_uri_complete":"https://cbdevice.okta.com/activate?user_code=QLWKNBRW","expires_in":600,"interval":5}

**.. Then point a Browser to the URL called *verification_uri_complete***

If you have opened the Browser and performed the Authentication you will receive an Access_token and Refresh_token when you request it from the Token Endpoint.

    curl --location --request POST 'https://cbdevice.okta.com/oauth2/aus1l4r7i7vJE9zRF3l7/v1/token' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --header 'Cookie: JSESSIONID=01CCCCC5F92725C4AF8667695190E949' \
    --data-urlencode 'grant_type=urn:ietf:params:oauth:grant-type:device_code' \
    --data-urlencode 'client_id=0oa1l4rl4c3IQluOA3l7' \
    --data-urlencode 'device_code=880318a7-31b4-48c8-a6fb-a6f099227738'

**Sample response from Token Endpoint.**

    {
    "token_type": "Bearer",
    "expires_in": 3600,
    "access_token": "eyJraWQiOiI0TmhTaUpjSVpfWEJHSjVyRUpZeW9qR2lhOTdrYzlFZkpRVDdidUVlQ0YwIiwiYWxnIjoiUlMyNTYifQ.eyJ2ZXIiOjEsImp0aSI6Ik......




