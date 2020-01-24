

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


  var eventHandlers = {
    //call post when paste occurs 
    postOnPaste: (event) => {
      console.log('pasted', event.clipboardData);
      var eventParams = prepareEventData({ 
        'pasted' : true, 
        event: EVENT_TYPES.PASTE, 
        formId: event.target.id
      });
      sendEvent(eventParams)
        .then(() => {
          console.log('gotov paste')
          showToast('event sent');
        });
    },
    //call post when paste occurs 
    postOnResize: (event) => {
      width = event.target.outerWidth;
      height = event.target.outerHeight;
      console.log('resized', event);

      var eventParams = prepareEventData({ 
        resizeFrom: { 
          width: originalWidth, height: originalHeight
        },
        resizeTo: {
          width: width, height: height
        },
        event: EVENT_TYPES.PASTE
      });
      sendEvent(eventParams)
        .then(() => {
          console.log('event sent');
          showToast('sent event');
          window.removeEventListener("resize", eventHandlers.postOnResize);
        });

    },

    postOnSubmit: (event) => {
      //stop control taking us away from the page
      event.preventDefault();
      errors = validateForm();
      if(errors.length !== 0) {
        let i=0; s='';
        for(;i<errors.length;i++) {
          s += `<p>${i} - ${errros[i]} \n<p>`
        }
        showToast(s)
        return;
      }
      //form data isn't being sent at this moment
      console.log('time it took to submit form ', (Date.now() - this.startTime) / 1000)
      var eventParams = prepareEventData({ 
        event: EVENT_TYPES.PASTE, 
        time: (Date.now() - this.startTime) / 1000
      });
      this.sendEvent(eventParams)
        .then(() =>  console.log('goooTOVO'))
        .then(initSession());                           //reset session after submit
      
    },
    onStartTyping: (event) => {
      console.log('started typing -THIS', this)
      if(!startTime) {
        startTime = Date.now();
        inputEmail.removeEventListener('keyup', this.eventHandlers.onStartTyping)
        inputCVV.removeEventListener('keyup', this.eventHandlers.onStartTyping)
        inputCardNumber.removeEventListener('keyup', eventHandlers.onStartTyping)
      }
    }
  };

  /* append session id annd website url it params object */
  function prepareEventData(params) {
    return {...params, 
      sessionId: sessionId,
      websiteUrl: getCurrentSiteUrl()
    };
  }

  /* helper for caallinng eventService post method */
  function sendEvent(event) {
    return new Promise(function(resolve, reject){
      eventService.postPromised(event)
        .then(resolve)
        .catch(reject)
    });
  }


  //
  function initSession() {
    this.originalWidth = window.screen.width;
    this.originalHeight = window.screen.height;
    this.sessionId = eventService.getSessionId();
    console.log('session started -- ', this.sessionId)
    initListeners();
  }

  function initListeners() {
    window.addEventListener("resize", eventHandlers.postOnResize);
    this.inputEmail = document.getElementById('inputEmail');
    this.inputCardNumber = document.getElementById('inputCardNumber');
    this.inputCVV = document.getElementById('inputCVV');
    this.inputEmail.addEventListener('paste', eventHandlers.postOnPaste);
    this.inputCardNumber.addEventListener('paste', eventHandlers.postOnPaste);
    this.inputCVV.addEventListener('paste', eventHandlers.postOnPaste);
    this.inputEmail.addEventListener('keyup', eventHandlers.onStartTyping);
    this.inputCardNumber.addEventListener('keyup', eventHandlers.onStartTyping);
    this.inputCVV.addEventListener('keyup', eventHandlers.onStartTyping);
    this.snackbar = document.getElementById("snackbar");
  }

  /*  
    even though we have already simple form validation in the html with required
    just another check for data so it hasn't been messed with
    simple validation for empty string and agreed terms
  */
 function validateForm(event) {
  let isValid = true;
  if(!document.getElementById('cbTerms'.value)) {
    this.errors['terms'] = 'You have to accepts terms of service';
    isValid = false
  }
  if(!this.inputEmail.value) {
    this.errors['email'] = 'Value not provided';
    isValid = false
  }
  if(!this.inputCVV.value) {
    this.errors['CVV'] = 'Value not provided';
    isValid = false
  }
  if(!this.inputCardNumber.value) {
    this.errors['cardNumber'] = 'Value not provided';
    isValid = false;
  }
  return isValid
}
function showToast(message) {
    // Get the snackbar DIV
    snackbar.className = "show";
    snackbar.innerHTML = message
    // Add the "show" class to DIV
    snackbar.className = "show";
  
    // After 3 seconds, remove the show class from DIV
    setTimeout(function(){
      snackbar.className = snackbar.className.replace("show", "");
    }, 3000);
}
function getCurrentSiteUrl() {
  return window.location.href;
}


  initSession();

