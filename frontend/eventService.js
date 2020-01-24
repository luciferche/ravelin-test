/*
  EventService function is used to encapsulate API sending  and session generation in one classs
  thus hidinng the sensitive variables and exposing only methods we want other users of the class to use
*/
function EventService() {
  const API_URL = 'http://localhost:8080/api';
  const POST_EVENT_PATH = '/events';
  var sessionId;

  var that = this;


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