/*
  EventService function is used to encapsulate API sending  and session generation in one classs
  thus hidinng the sensitive variables and exposing only methods we want other users of the class to use
*/
function EventService() {
  const API_URL = 'http://localhost:8080/api';
  const POST_EVENT_PATH = '/events';
  var sessionId;

  var that = this;


  that.postPromised = (params) => {
    return new Promise(function(resolve, reject){
      var req = new XMLHttpRequest();
      req.open('POST', API_URL + POST_EVENT_PATH);

      //Send the proper header information along with the request
      req.setRequestHeader("Content-Type", "application/json;charset=UTF-8");

      req.onreadystatechange = function() { // Call a function when the state changes.
          if (this.readyState === XMLHttpRequest.DONE && this.status === 200) {
              // Request finished. Do processing here.
              resolve();
          }
      }
      req.send(JSON.stringify(params));
      // xhr.send(new Int8Array()); 
      // xhr.send(document);
      req.onload = function() {
        // This is called even on 404 etc
        // so check the status
        console.log(req.status + ' - returned from POST', req.readyState)
        if (req.readyState === XMLHttpRequest.DONE && req.status == 200) {
          // Resolve the promise with the response text
          resolve(req.response);
        }
        else {
          // Otherwise reject with the status text
          // which will hopefully be a meaningful error
          reject(Error(req.statusText));
        }
      };

      // Handle network errors
      req.onerror = function() {
        reject(Error("Network Error"));
      };

    });
  }
  that.getSessionId = () => {

    if(!sessionId) {
      sessionId = getRandomNumberOfLength(6)+ '-' + 
                  getRandomNumberOfLength(6) + '-' + 
                  getRandomNumberOfLength(9);
      

    }
    console.log('sessionID',sessionId);
    return sessionId;
  }
  
  that.postEvent = function (params) {
    return new Promise(function(resolve, reject) {
      axios.post(API_URL + POST_EVENT_PATH, {...params})
      .then(function (response) {
        console.log('response from POST ', response);
        resolve(response);
      })
      .catch(function (error) {
        console.log('error from POST ',error);
        reject(error);
      });
    });
    
  }
  
  //Init sessionid
  that.getSessionId();

  function getRandomNumberOfLength(n) {
    var result           = '';
    var characters       = '0123456789';
    var charactersLength = characters.length;
    for ( var i = 0; i < n; i++ ) {
       result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
  }
  return that;
}