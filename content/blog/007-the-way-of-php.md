+++
title = "The way to PHP"
date = "2021-05-29T11:51:03Z"
author = "Sandro Lain"
authorTwitter = "sandro_lain" #do not include @
cover = "/cover/ben-gAe1pHGc6ms-unsplash.jpg"
tags = []
keywords = []
description = "Journey into an old world"
showFullContent = false
readingTime = false
hideComments = false
color = "" #color from the theme settings
+++

*Photo by [Ben](https://unsplash.com/@benofthenorth?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) on [Unsplash](https://unsplash.com/s/photos/php?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)*

This week I have been experimenting a bit with PHP some new approaches I learned with [Node.js](https://nodejs.org/) and [GOlang](https://golang.org/).
Trying to follow a "[frameworkless](https://www.frameworklessmovement.org/)" philosophy I tried to expand the functionality of the chosen router library to add the use of middleware and manage the request and response information using representative classes, such as [express](https://expressjs.com/) or [fastify](https://www.fastify.io/).

At the moment I have created a couple of simple middleware.
The first intercepts and allows you to manage requests to static files, specifying a specific directory where to locate them.
With the second I tried to create automated management of a user session through the use of [JWT](https://jwt.io/) on cookies; this approach should replace the use of PHP sessions that rely on disk files in the server and session cookies, but I am also thinking of using authorization headers to expand the functionality of the middleware to API calls from SPA as well.

Being quite rusty with PHP I had to add a log manager to log critical and non-critical errors.

Obviously I had to choose new libraries.
After some research I chose [Monolog](https://seldaek.github.io/monolog/) for the simplicity and integration of [FirePHP](https://firephp.org/), which allows me to see the logs collected in the browser console.

The list of libraries has evolved as follows:

- JWT: [lcobucci/jwt](https://github.com/lcobucci/jwt) ***(New Entry)***
- Logging: [Monolog](https://seldaek.github.io/monolog/) ***(New Entry)***
- Env Vars: [PHPdotenv](https://github.com/vlucas/phpdotenv)
- Routing: [FastRoute](https://github.com/nikic/FastRoute)
- MySQL Client: [MeekroDB](https://meekro.com/)
- Template: [Twig](https://twig.symfony.com/)
- Unit Test: [PHPunit](https://phpunit.de/)

I am considering whether to start a small concrete project that I have been thinking about for some time; I think it may be the best way to explore more concrete aspects of development.
At the moment I have to say that PHP with version 8 has really modernized, but never as much as TypeScript 4.

***"My Little Framework"*** evolves.
