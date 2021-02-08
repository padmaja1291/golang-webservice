# golang-webservice
A post method using golang used for basic user authentication

**API details**  
POST http://localhost:8081/login  
username: example@domain.com  
password: ********  
token: 1234 


Service accepts a POST endpoint with a JSON payload with username, password and token  
The backend should successfully validate the following credentials:  
* Username = c137@onecause.com
* Password = #th@nH@rm#y#r!$100%D0p#
* One time token = the 2 digit hour and 2 digit minute at time of submission  
<br/>
The backend should invalidate any other credentials or variations from the above

