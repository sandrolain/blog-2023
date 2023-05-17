+++
title = "Micro Frontend|Service Analogy"
date = "2023-05-17T12:21:49Z"
author = "Sandro Lain"
authorTwitter = "sandro_lain" #do not include @
cover = "/cover/_3b124ff5-4dce-4873-b52c-0a9110c55fd4.jpeg"
tags = []
keywords = []
description = "An analogy between a microservices architecture and a microfrontend one"
showFullContent = false
readingTime = false
hideComments = false
color = "" #color from the theme settings
+++

In the last few days I have had some opportunities to reflect and compare the frontend architecture of a SPA developed using custom elements with the elements of a microservices backend architecture.  
Here are some points of comparison.

## Microservices

Custom element can be seen as microservices (micro frontend) if they are developed with an **isolated logic**.  
However, the structure of the web pages being tree-like and made up of numerous elements, one could conceive that different instances of the same microservice operate for the use of different sub-domains of the same application.

## Networking

**CSS selectors** can be considered the equivalent of a **DNS system**.  
CSS selectors allows to refer to an element abstractly from their real memory location, similar to how domain names relate to a **network address**.

The obtained **DOM nodes** are equivalent to the **addresses/port** for the reference of the service.  
Also due to the tree nature, the sub-elements can be understood both as services and as the single doors of a parent element depending on the level of dependency they have between them.

## Direct communication

Custom element can expose **methods** that act as synchronous APIs to the service of other elements, similarly to an **RPC** communication.

## Event driven communication

DOM event listeners allow for asynchronous and decoupled communication between different elements.  
Container elements, and in particular `window` and `document`, can act as pub/sub communication brokers.  
Event types are comparable to **channels/topics**.  
Individual elements can emit events of a certain type on the global broker and others subscribe to them.

## Environment Variables

Custom element **attributes** can be thought of as the analog of **environment variables** or other type of configuration source for the service.

## Tests

Frontend unit tests can often be seen as **integration tests** with browser-provided services (webapis), and as with backend integration tests what you can do is use mock interfaces to the service/external API, or implement real communication by running tests **on a browser** environment equivalent to running **test-containers** for external services.
