### GoLang_Authentication_Practise

Pre-Requisite: <br>
- Visual Studio Code  <br>
- Install Go from https://go.dev/<br>
- Postman app from https://www.postman.com/downloads/ <br>

Check Go Configuaration using VS Terminal by running below commands <br>
- go --version <br>
- go env <br>

### Code Practises <br>
1) Marshal Sample -- Converting Objects to json <br>
2) Unmarshal Sample -- Converting json to objects <br>
3) Encode and Decode Sample Program <br>
4) Storingpassword Sample Program using bcrypt <br>
    #### for more password encryption golang examples <br>
    https://golang.hotexamples.com/examples/golang.org.x.crypto.scrypt/-/Key/golang-key-function-examples.html <br>
5) HMAC Sample - Hash Message Authentication Code (HMAC) as defined in U.S. Federal Information Processing Standards Publication 198. An HMAC is a cryptographic hash that uses a key to sign a message. The receiver verifies the hash by recomputing it using the same key. <br>
6) HMAC Sample 2- with Client Response and Saved Key in Cookie<br>
7) JWT - 
    #### JWT INTRO - https://github.com/golang-jwt/jwt <br>
    Hint: <br>
    type UserClaims struct { 
	jwt.StandardClaims 
	SessionID int64 
    } 
    after typing above code run go mod tidy and restart your VS to get the jwt intellisense. <br>

    to know more about JWT in GoLang visit https://learn.vonage.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr/ <br>

    for encode and decode jwt tokens vsiit https://jwt.io/ <br>

    for uuid in goLang download package using go get github.com/gofrs/uuid
8) base64 encode-decode <br>
9) Crypto Sample <br>
10) Crypto Sample - Abstract encrypt writer function <br>

### Output Screenshots:
1) Marshal Sample <br>
<img src="Screenshots/Marshal_Sample.png" /><br>
2) UnMarshaling Sample <br>
<img src="Screenshots/unMarshal_Sample.png" /><br>
3) Encode and Decode Sample Program <br>
<img src="Screenshots/encode_decode_sample.png" /><br>
Encode postman Request:<br>
<img src="Screenshots/postman_encode_sample.png" /><br>
Decode postman Request:<br>
<img src="Screenshots/postman_decode_sample.png" /><br>
4) Storingpassword Sample Program using bcrypt <br>
<img src="Screenshots/Storingpassword_sample.png"/><br>
5) HMAC Sample <br>
<img src="Screenshots/HMAC_Sample.png"><br>
6) HMAC Sample 2<br>
<img src="Screenshots/HMAC_Sample2.png"><br>
7) JWT Sample <br>
<img src="Screenshots/JWT_Sample.png"><br>
<img src="Screenshots/JWT_Sample_encode_decode.png"><br>
8) base64 encode-decode <br>
<img src="Screenshots/base64_encode_decode.png"><br>
9) Crypto Sample <br>
<img src="Screenshots/Crypto_sample.png"><br>
10) Crypto Sample - Abstract encrypt writer function <br>
<img src="Screenshots/Crypto_sample2.png"><br>