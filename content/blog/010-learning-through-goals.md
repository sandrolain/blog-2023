+++
title = "Learning through goals"
date = "2021-07-25T13:11:01Z"
author = "Sandro Lain"
authorTwitter = "sandro_lain" #do not include @
cover = "/cover/engin-akyurt-bPiuY2ZSlvU-unsplash.jpg"
tags = []
keywords = []
description = "More sophisticated, more skills"
showFullContent = false
readingTime = false
hideComments = false
color = "" #color from the theme settings
+++

*Photo by [engin akyurt](https://unsplash.com/@enginakyurt?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) on [Unsplash](https://unsplash.com/s/photos/target?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)*

I have always felt that having to face the various job challenges as a developer would lead me to add new pieces to the puzzle of my skills.

For this reason, in my little adventure with PHP (which from an experiment is becoming a product) I decided to add a bit of sophistication and modernity through the combination of microservices written in [GOlang](https://golang.org/) (even if at the moment it is only a monolith-service :D ).

In my current job I already find myself having to put my hand to the GO code, but not having much experience yet, I decided to hone my skills with small challenges that I have already been able to solve with other languages.
In the specific case I am trying to move the more complex processing parts from PHP to GO, so far with excellent results.

To connect the PHP part with the GO one at the moment I have chosen to use [Redis](https://redis.io/) as a Pub/Sub broker, and I do not exclude that when everything is stable enough I will also learn [Docker](https://www.docker.com/), of which at the moment I make a more passive use.

The result I have obtained so far is a parallelization of the most complex operations by delegating them to the service in GO, which in most cases consist of creating resources on disk (images, pdf, etc.) and sending emails, operations that if performed from PHP itself would result in a slowdown in returning the response to the user.

At the moment it is going well, let's see how it continues. ;)
