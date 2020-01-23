function EventService() {
  const API_URL = 'http://localhost:8080/api';
  const GET_SESSION_PATH = '/session-id';
  const POST_EVENT_PATH = '/event';
  const STORAGE_SESSION_KEY = 'sessionId';
  var sessionId;
  getSessionId = () => {

    let sessionId = localStorage.getItem(STORAGE_SESSION_KEY);
    if(!sessionId) {
      sessionId = getRandomNumberOfLength(6)+ '-' + 
                  getRandomNumberOfLength(6) + '-' + 
                  getRandomNumberOfLength(9);
      

    }
    console.log('sessionID',sessionId);
    return sessionId;
  }
  storeSession = () => {
    localStorage.setItem(STORAGE_SESSION_KEY, getSessionId());
  }
  clearStorage = () => {
    localStorage.removeItem(STORAGE_SESSION_KEY,null)
  }
  getRandomNumberOfLength = (n) => {
    var result           = '';
    var characters       = '0123456789';
    var charactersLength = characters.length;
    for ( var i = 0; i < n; i++ ) {
       result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
  }
  postEvent = (params) => {
    axios.post(API_URL + POST_EVENT_PATH, {...params})
    .then(function (response) {
      console.log('response from POST ', response);
    })
    .catch(function (error) {
      console.log('error from POST ',error);
    });
  }
  
  //Init sessionid
  storeSession();
	// set the public value TODO CHECK IF IT WORKS
	// Object.assign(this, {
	// 	postEvent(params) {
  //     axios.post(API_URL + POST_EVENT_PATH, {...params})
  //     .then(function (response) {
  //       console.log('response from POST ', response);
  //     })
  //     .catch(function (error) {
  //       console.log('error from POST ',error);
  //     });
  //   }
	// });
}
const API_URL = 'http://localhost:8000/api';
const GET_SESSION_PATH = '/session-id';
const POST_EVENT_PATH = '/event';

postEvent = (params) => {
  axios.post(API_URL + POST_EVENT_PATH, {...params})
  .then(function (response) {
    console.log('response from POST ', response);
  })
  .catch(function (error) {
    console.log('error from POST ',error);
  });
}
