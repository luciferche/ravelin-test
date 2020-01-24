Ravelin Code Test 
===============================================
Repository for the test project for Ravelin
Project consists of frontend built with JS and a backend server written in Go

Remarks - 
===============================================
## Backend
### I tried to provide enough information through comments within the files but this is a brief overview:
App that consists of server object that serves HTTP requests and processes the data according to the specs. 
Data struct is used for storing all the information about the user session, it has it's own methods for manipulating it's data. 
Main function just instantiates the servers and assigns handlers for different routes.
eventsApi function does all the work once `POST /api/events/` is received.
It first parses params into Event structure.
Whenever event is received, if it is parsed properly into an Event struct(model for params from request) I store it in
a map with sessionIds  as keys. Every time new event comes with the same sessionId the session in the map 
gets updated with newest values based on the `eventType` passed in the params. Once the struct is updated with FormCompletionTime I assume that the form has been submitted and that it is finished so I print the struct and delete the session from the map(even though it wasn't explicitly stated in the requirements, it could give us false data if we assume that session ids are unique)
In Data structure I use reflection to read out properties of the object and format it out properly with field names.

## Frontend - 
Plain JS/HTML/CSS used. Axios library used for http requests. Client sessionid generates on page load and gets returned and added to every request fired. I encapsulated eventService functions in an object in order to hide constants, vars and helpers. The app responds to Window resize, Paste in a field and submit button. There is form validation inside HTML but also one before sending the requests.

-
## What's not working: 

Basically main client app functionality isn't working. `OPTIONS` request is fired under the hood and 
it doesn't pass CORS check policy. My first problem was writing a json response, and Content-Type should not 
be specified in response headers on OPTIONS request. But even after fixing that issue(by returning 
Content-Type: json only on actual POST/GET requests) - requests from the client kept failing with report that 
Content-Type header is present. 
However, whenever I tried with Postman everything worked smoothly and I wasn't getting specified header with OPTION request.

I initially developed the backend first using **Postman** to test the api and that's what I focused on prior to switching to FE. 
On the frontend I think that everything should be working but it isn't unfortunately.
I tried using standard ```XMLHttpRequest``` object instead of ```axios``` library, thinking there might be the problem under the hood with the OPTIONS request but it wasn't the case.

Hashing function isn't finished because I lost too much time on trying to solve problems with ```OPTIONS``` request.
So I am sorry to say I am submitting incomplete solution. 

Ravelin Code Test - Original Requirements
===============================================

## Summary
We need an HTTP server that will accept any POST request (JSON) from multiple clients' websites. Each request forms part of a struct (for that particular visitor) that will be printed to the terminal when the struct is fully complete. 

For the JS part of the test please feel free to use any libraries that may help you **but please only use the Go standard library for the backend**. Remember to keep things simple.

## Frontend (JS)
Insert JavaScript into the index.html (supplied) that captures and posts data every time one of the below events happens; this means you will be posting multiple times per visitor. Assume only one resize occurs.

  - if the screen resizes, the before and after dimensions
  - copy & paste (for each field)
  - time taken from the 1st character typed to clicking the submit button

### Example JSON Requests
```javascript
{
  "eventType": "copyAndPaste",
  "websiteUrl": "https://ravelin.com",
  "sessionId": "123123-123123-123123123",
  "pasted": true,
  "formId": "inputCardNumber"
}

{
  "eventType": "timeTaken",
  "websiteUrl": "https://ravelin.com",
  "sessionId": "123123-123123-123123123",
  "time": 72, // seconds
}

...

## Backend (Go)

### Part 1

The Backend should:

1. Create a Server
2. Accept POST requests in JSON format similar to those specified above
3. Map the JSON requests to relevant sections of the data struct (specified below)
4. Print the struct for each stage of its construction
5. Also print the struct when it is complete (i.e. when the form submit button has been clicked)

### Part 2

6. Write a simple hashing function (your implementation - either of
   your own design or a known algorithm), that given a string will
   calculate a hash of that string.  We are not looking for you to
   wrap a standard function, but to provide the implementation itself.
7. Use that function to calculate the hash of the `WebSiteurl` field
   and print the hash, and print out the hash once calculated.

### Go Struct
```go
type Data struct {
	WebsiteUrl         string
	SessionId          string
	ResizeFrom         Dimension
	ResizeTo           Dimension
	CopyAndPaste       map[string]bool // map[fieldId]true
	FormCompletionTime int // Seconds
}

type Dimension struct {
	Width  string
	Height string
}


