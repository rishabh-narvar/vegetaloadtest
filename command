jq -ncM 'while(true; .+10010) | {method: "POST", url: "https://ws.narvar.qa/api/v1/event/", header: {"Content-Type": ["application/json"], "Authorization":["Basic YWRkYmNjNGZhM2Y3NGRiNDk1MDA3Y2QwZmQ5YTU1ZTk6ZjA1YzEyZGNjNTgzNDQ2OWEzYmQ1Y2IzZjI1MGUxZTM="]}, body: {"guid":.,"name":"forgot_password","timestamp":"2019-04-08T02:51:38Z","customer":{"firstName":"cs","lastName":"cs","customerId":"12345","email":"chethan.sindhie@narvar.com","channels":["email","sms"],"phoneNumber":"9663399077"},"params":{"forgetpass_otp":"875233","forgetpass_otp_flag":"N","change_passwordurl":"http://example.com"}} | @base64 }' | \
  vegeta attack -rate=3/s -lazy -format=json -duration=1s | \
  tee results.bin | \
  vegeta report