+++
title = "A simple fetch() interceptor with VanillaJS"
date = "2020-07-11T18:00:00Z"
author = "Sandro Lain"
authorTwitter = "sandro_lain" #do not include @
cover = "/cover/kevin-horvat-Pyjp2zmxuLk-unsplash.jpg"
tags = []
keywords = []
description = "Control over the requests and their responses"
showFullContent = false
readingTime = false
hideComments = false
color = "" #color from the theme settings
+++

*Photo by [Kevin Horvat](https://unsplash.com/@hidd3n?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) on [Unsplash](https://unsplash.com/s/photos/monitors?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)*

The introduction of the **Fetch API** brought a modern and comprehensive method of making HTTP calls to browsers.

In the development of more complex web applications, however, there is a need to have **more control** over the requests and their responses.

Some frameworks and libraries provide **alternative or similar methods to fetch()**,
and with it some tools to **intercept** the requests before they are made or the responses before they are digested, allowing their **manipulation** to normalize them according to business needs (e.g. add authentication HTTP headers), or to perform **predefined actions** in some circumstances (e.g. toast notifications in case of error responses).

## A *Vanilla.js* interceptor

Here below I report the **ES6 class** code of an interceptor, which proves to be very simple and tries to exploit the already rich native **fetch API** without distorting them,
therefore it does **not force to adopt proprietary formats** when creating the request or reading the response.

```javascript
class Interceptor {

  // The arrays that will contain the interception functions are prepared at the instantiation,
  // for requests and responses
  constructor () {
    this.requests = [];
    this.responses = [];
  }

  // Method for defining a function that intercepts and, if necessary, modifies HTTP requests
  interceptRequest (fn) {
    this.requests.push(fn);
  }

  // Method for defining a function that intercepts and, if necessary, modifies HTTP responses
  interceptResponse (fn) {
    this.responses.push(fn);
  }

  // Method that returns a function with the same signature as the native fetch()
  // (therefore you can rely on the standard),
  // but which manages the previously defined interceptors
  getFetcher () {
    return async (input, init) => {

      // The request interceptor manages the native "Request" type,
      // so if an instance of Request has not already been passed
      // the function expects to create it
      let request = (input instanceof Request) ? input : new Request(input, init);

      // The request instance is processed through all the request interceptor functions
      for(const fn of this.requests) {
        // The interceptor response must be an instance of Request
        // or a Promise that resolves with an instance of Request
        const res = fn(request);
        request = (res instanceof Promise) ? await res : res;
      }

      // The request is made via the native fetch() function, passing the processed "Request" instance
      let response = await window.fetch(request);

      // The response instance is processed through all the response interceptor functions
      for(const fn of this.responses) {
        // The interceptor response must be an instance of Response
        // or a Promise that resolves with an instance of Response
        const res = fn(response);
        response = (res instanceof Promise) ? await res : res;
      }

      // The result of the function is the processed "Response" instance
      return response;
    };
  }
}

```

## An example of use

Here is an example of use, very simple in this case.
By taking advantage of the **ES modules** and dividing the code over *different files*,
you can develop a modern and simple method to make fetch() calls.

```javascript

// Create the instance of interceptor management
const interc = new Interceptor();

// Set an interceptor function for the request
interc.interceptRequest((request) => {
  // Create a new Request with adding Authorization header,
  // for example, as value, an authentication token stored in the web-app
  return new Request(request, {
    headers: {
      "Authorization": `Bearer ${storedAuthToken}`
    }
  })
});

// Set an interceptor function for the response
interc.interceptResponse(async (response) => {
  // Clone the answer to be able to process it in the interceptor
  // (the response buffer can be used only once, in doing so I leave it available for subsequent processing)
  const responseClone = response.clone();
  if(responseClone.status === 401) {
    window.location = "https://.../login-page";
  }
  // Return the original response
  return response;
});

// Get a function that uses fetch() with interceptors
const interceptedFetch = interc.getFetcher();

// ...

// Where necessary I use the function obtained to make requests.
// The authentication header will be added before making the request
// Before returning the response, the status will be analyzed and predefined operations executed
interceptedFetch(`https://.../my-endpoint`, {
  method: "POST",
  body: JSON.stringify({foo: "bar"})
});

```

[View Live Example](/002-example.html)


In my **HTNA-tools** repository on github you can find the original **TypeScript** version:
[Click here to go to the exact line](https://github.com/sandrolain/HTNA-tools/blob/8dd09d233e3856ce78199ac2dff608f21555a507/src/netw.ts#L50).

