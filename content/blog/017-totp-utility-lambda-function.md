+++
title = "TOTP generator Lambda Function"
date = "2022-12-26T15:14:52Z"
author = "Sandro Lain"
authorTwitter = "sandro_lain" #do not include @
cover = "/cover/towfiqu-barbhuiya-FnA5pAzqhMM-unsplash.jpg"
tags = []
keywords = []
description = "Big as fast update"
showFullContent = false
readingTime = false
hideComments = false
color = "" #color from the theme settings
+++

For the need for work and tests on some projects I created a Lambda Netlify function on this site that performs the function of TOTP generator.
Passing a TOTP configuration URI as a Query-String parameter, the function calculates and returns the relative temporary code.
Having to pass through the parameter the URI including secret, the use is designed only for testing and development purposes.

## Example

Having Uri Totp as follows:

```
otpauth://totp/identity:example@example.com?algorithm=SHA1&digits=6&issuer=identity&period=30&secret=RTNQONWEJPKSB27LXQA4TCBXE6DBDAR
```

and the function served to URL:

```
https://www.sandrolain.com/.netlify/functions/totpCode
```

the final url is:

```
https://www.sandrolain.com/.netlify/functions/totpCode?uri=otpauth%3A%2F%2Ftotp%2Fidentity%3Aexample%40example.com%3Falgorithm%3DSHA1%26digits%3D6%26issuer%3Didentity%26period%3D30%26secret%3DRTNQONWEJPKSB27LXQA4TCBXE6DBDAR5
```

Specifically, the use I would like to try is on a [Postman](https://www.postman.com/) workflow to test two-factor authentication.

