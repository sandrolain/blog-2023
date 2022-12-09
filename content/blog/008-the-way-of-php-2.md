+++
title = "The way to PHP #2"
date = "2021-06-05T13:42:00Z"
author = "Sandro Lain"
authorTwitter = "sandro_lain" #do not include @
cover = "/cover/ben-gAe1pHGc6ms-unsplash.jpg"
tags = []
keywords = []
description = "Journey into an old world, 2nd part"
showFullContent = false
readingTime = false
hideComments = false
color = "" #color from the theme settings
+++

*Photo by [Ben](https://unsplash.com/@benofthenorth?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) on [Unsplash](https://unsplash.com/s/photos/php?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)*

After the latest additions to create a solid foundation for route management and user navigation session I went into more specifics with the functional objectives of the PHP project.

The application will provide for the management of user accounts, therefore registration and login.
To simplify the work, I chose to use the [Auth0](https://auth0.com/) services, so I had to integrate the service library, in a few lines I have management of user authentication without having to integrate the entire registration process.
In doing so I ran into the fact that Auth0 uses an older major version of the JWT library I had chosen and the two implementations are incompatible, so I had to rewrite the JWT middleware integration using the old version of the library, losing some minor functionality.

I also did a trial of using [MongoDB](https://www.mongodb.com/) as database with [MongoDB Atlas](https://www.mongodb.com/cloud/atlas), but I retraced my steps and I think I will use MySQL (classic LAMP).

The list of libraries has evolved as follows:

- User Auth: [Auth0 PHP](https://github.com/auth0/auth0-PHP) ***(New Entry)***
- JWT: [Auth0 PHP JWT](https://github.com/auth0/php-jwt) ***(New Entry, replace "lcobucci/jwt")***
- Filesystem: [Symfony Filesystem](https://github.com/symfony/filesystem) ***(New Entry)***
- MIME Types: [Mimey](https://github.com/ralouphie/mimey) ***(New Entry)***
- QRCode: [Simple QrCode](https://github.com/SimpleSoftwareIO/simple-qrcode) ***(New Entry)***
- Slugify: [cocur/slugify](https://github.com/cocur/slugify) ***(New Entry)***
- MongoDB Client: [Xenus](https://abellion.github.io/xenus/#/) ***(New Entry)***
- Logging: [Monolog](https://seldaek.github.io/monolog/)
- Env Vars: [PHPdotenv](https://github.com/vlucas/phpdotenv)
- Routing: [FastRoute](https://github.com/nikic/FastRoute)
- MySQL Client: [MeekroDB](https://meekro.com/)
- Template: [Twig](https://twig.symfony.com/)
- Unit Test: [PHPunit](https://phpunit.de/)

From here on I will try to implement the application interface.

***"My Little Framework"*** is growing.

