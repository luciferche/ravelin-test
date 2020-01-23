

  var errors = [];
  const CURRENT_SITE_URL = getCurrentSiteUrl();
  var originalHeight, originalWidth;
  var startTime;

  const eventService = new EventService();
  const EVENT_TYPES = {
    PASTE: 'copyAndPaste',
    RESIZE: 'screenResize',
    SUBMIT: 'timeTaken'
  }
  var inputEmail, inputCardNumber, inputCVV;


  function postUserData(event) {
    event.preventDefault();
    errors = validateForm();
    if(errors.length != 0) {
      showErrors(errors)
      return;
    }
    // prepareData(email, cardNumber, )
    let userForm = document.forms['userForm'];
    console.log('posted something', userForm);
    errors = [];
  }

  clearErrors = () => {
    var errorBlock = document.getElementById('errorBlock');
    errorBlock.style.display = 'none';

    document.getElementById("userForm").reset();

  }
  function prepareEventData(params) {
    return {...params, 
      sessionId: sessionId,
      websiteUrl: getCurrentSiteUrl()
    };
  }

  function sendEvent(event) {
    eventService.postEvent(event)
        .then(() => {
          alert('success')
        })
        .catcth(error => {
          console.error('error posting event - '+ event.event, error)
          alert('ERROR');
        })
  }
  function getCurrentSiteUrl() {
    return window.location.href;
  }
  function initProperties() {
    this.originalWidth = window.screen.width;
    this.originalHeight = window.screen.height;
    window.addEventListener("resize", eventHandlers.postOnResize);
    this.inputEmail = document.getElementById('inputEmail');
    this.inputCardNumber = document.getElementById('inputCardNumber');
    this.inputCVV = document.getElementById('inputCVV');
    this.inputEmail.addEventListener('paste', eventHandlers.postOnPaste);
    this.inputCardNumber.addEventListener('paste', eventHandlers.postOnPaste);
    this.inputCVV.addEventListener('paste', eventHandlers.postOnPaste);
    this.inputEmail.addEventListener('oninput', eventHandlers.onStartTyping);
    this.inputCardNumber.addEventListener('oninput', eventHandlers.onStartTyping);
    this.inputCVV.addEventListener('oninput', eventHandlers.onStartTyping);
    this.sessionId = getSessionId();
  }
  showErrors = (errors) => {
    var errorBlock = document.getElementById('errorBlock');
    var errorsElement = document.createElement('ul');
    for(let error in errors) {
      console.error('error with', error)
      errorsElement.append(document.createElement('li').innerHTML('Error - ' + error));
    }

  }

  var eventHandlers = {
    //call post when paste occurs 
    postOnPaste: (event) => {
      console.log('pasted', event.clipboardData)
      debugger;
      var eventParams = prepareEventData({ 
        'pasted' : true, 
        event: EVENT_TYPES.PASTE, 
        formId: event.target.id
      });
      sendEvent(eventParams);
    },
    //call post when paste occurs 
    postOnResize: (event) => {
      width = event.target.outerWidth;
      height = event.target.outerHeight;
      console.log('resized', event)

      window.removeEventListener("resize", eventHandlers.postOnResize);

    },

    postOnSubmit: (event) => {
      let email = inputEmail.value;
      let cardNumber = inputCardNumber.value;
      let cvv = inputCVV.value;

      var eventParams = prepareEventData({ 
        pasted : true,
        event: EVENT_TYPES.PASTE, 
        formId: 'nesto'
      });
      sendEvent(eventParams);
    },
    onStartTyping: (event) => {
      console.log('started typing -THIS', this)
      if(!startTime) {
        startTime = Date.now();
        inputEmail.removeEventListener('oninput', this.onStartTyping)
        inputCVV.removeEventListener('oninput', this.onStartTyping)
        inputCardNumber.removeEventListener('oninput', this.onStartTyping)
      }
    }
  };
  initProperties();

  /*
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
  */
    

