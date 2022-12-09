+++
title = "UX in team development"
date = "2020-08-01T18:00:00Z"
author = "Sandro Lain"
authorTwitter = "sandro_lain" #do not include @
cover = "/cover/alvaro-reyes-qWwpHwip31M-unsplash.jpg"
tags = []
keywords = []
description = "Well-designed UX as the first analysis and debug tool"
showFullContent = false
readingTime = false
hideComments = false
color = "" #color from the theme settings
+++

*Photo by [Alvaro Reyes](https://unsplash.com/@alvarordesign?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) on [Unsplash](https://unsplash.com/?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)*


Often, working in a team, happens to get our hands on a code written by other developers,
present or past, and therefore on the functionality that the code implements.
As developers, both ethically and practically, we should strive to write clean, easily readable code for others.

A few days ago while I was working I asked myself some questions:

*Does only the code contribute to making the software effectively interpretable and validable by others?*
*As long as the software is running and we interact with it, what does it tell us if the operations are successful and are giving the expected results?*

The answer is that in addition to the logical flow of the code, you can also **see the application running** and it can communicate with us.
To cover this aspect, however, it is important that a **good interface** is first designed and developed, especially for the **user experience**.

A wrong or poor UX does not make it clear to end-users who are picking up the software (let alone the developers)if it is working properly.
On the contrary, a clear and complete communication allows to better verify the execution status,
and therefore if the software is giving the expected results or if there is something wrong.

I describe below some examples deriving from my experience.


## Error messages

Error messages that are too general or are the same for different cases may not make it clear whether the error depends on the user or the software.
From the developer's point of view, it is necessary to research the origin of the problem in several places.

The absence of error messages, on the other hand, does not make it clear whether the software is working or not,
and makes it even more difficult to understand the problem (if it is really there).


## Blank screens

Screens and blank areas should be avoided.

In the case of dynamic lists that do not have any element, or in the absence of information or resource not found,
an appropriate message and an alternative action should be shown that make it clear to the user, as to the developer,
that the loading of the content carried out by the software was executed as expected and gave a result, albeit empty.


## Actions that cannot be implemented

Buttons, tools or sections inaccessible without indication of the reason, could frustrate the user
and not let the developer understand if it is a lawful state or an error.

On the contrary, an explanation, and possibly advice on how to solve it, can help the user to take the right path,
and make the developer understand that the logics are carried out correctly.


## Actions accessible when you can't

The lack of a preventive control over the permissions or accessibility to some functions
can lead the user to perform actions to which he cannot actually have access,
with the consequent frustration of seeing his expectations denied.

From the developer's point of view, there is no verification of whether the software is properly dressing permissions
and access controls.


## Lack of execution or completion indicators

The execution of processes that require long times without indicating the progress or completion
does not make it clear whether the software is performing the required actions,
and therefore whether it is necessary to wait or can continue with the risk of invalidating or corrupting the processing.

From his point of view, the developer does not know if the software is correctly managing the long-term action.


## To conclude

Therefore a well-designed UX can function as the first analysis and debug tool, allowing us to identify
where any problems originate from, so as to intervene in a more targeted way.
Complete feedback and expressiveness from the software equal the verbosity of the log messages,
while the error messages verify the correct handling of the exceptions.
