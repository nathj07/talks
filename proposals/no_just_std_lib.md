# Elevator Pitch (300 characters)

Go has an excellent community. However, I see and hear a lot of "just use the standard library" as advice given to newcomers. We should acknowledge the and promote the community & libraries we find useful. This talk will build community by promoting the more nuanced response of "try this library..."

# Abstract
Go is a great language. Go has an excellent community. Those two make a powerful combination. However, I see and hear a lot of "just use the standard library" as advice given to newcomers. We need to stop this, we need to be more nuanced. It indicates how powerful the language is but it minimises how great the community is. This talk will cover a number of very useful libraries across a variety of domains that we, at Turnitin, have used and found very effective. The result will be that we, as a community, can start to give better, more welcoming advice to the newcomers and to the old hands alike.

# Introduction
By show of hands who has ever received the advice "just use the standard library" when asking how to do something in Go?
Again, by show of hands who has ever given the advice "just use the standard library" when asked how to do something in Go?
Finally, by show of hands who only ever uses the standard library when working in Go?

Interesting, many of us have received, many of us have given the advice, yet not so many of us, in honesty probably none of us actually follow that advice completely.

The fact that we can even consider giving that advice demonstrates how awesome the standard library is. However, this is terrible advice.

The reason this is bad advice is because the main strength of Go is not the language, it is the community. I've worked in many languages over many years and not encountered one with a community like this.

Yet that advice, whether we mean to or not, puts a barrier up between those in the community and those looking to join.

What I am going to do in this talk is demonstrate, from experience, some of the better options we could provide when asked "how to do you do x in Go?"

The aim is to get us all thinking more about the libraries we do use and how we can share the knowledge to mae the community more welcoming to newcomers.


# Conclusion
Based on the code examples shown we can see that, while our standard library is incredible, there are many excellent libraries available to us.

We should stop telling one another, stop telling newcomers, to re-invent the wheel.

So next time you encounter the question "how do I do x in Go?" and you find yourself thinking, typing, or saying "Just use the standard library" stop. Take a pause and evaluate the library you use for that, then give your advice.

In this way we, as a community, will be more welcoming, more encouraging and will receive the benefits of wider adoption, wider input, and more varied programming experience. In short our community will become even better.

# Outline

* Introduction outlining the issue at hand
    * what is the advice
    * how come we can even give it
    * why is it bad
    * examples of where it is bad
* The main areas where I've seen this are:
    * web development, especially in routing
    * logging
    * testing
* Other ares cover command line args, and output, there are more but we will cover the main 3
* routing
    * show how routes look using just the standard library
    * show how they look using open source packages
* logging
    * demonstrate logging using just the standard library
    * show logging with an open source library
    * highlight the benefits
    * anecdotes of how this can save you time
* testing
    * show a simple test using only the standard library
    * show a more complex test using the standard library
    * show how we have used testify to simplify testing
    * benefits of clearer code, clearer output and shorter time to write
* What advice to give
    * review the libraries you use, recommend them
    * contribute to the libraries you use
    * write posts or tweets about the libraries you use
    * we can still recommend the standard library, let's not swing too far the other way, as the Brits say don't throw the baby out with the bath water
* Conclusion
    * let's make our awesome community more welcoming
    * we will benefit as a community from encouraging people to join in
    * lowering the barrier to entry will bring more folks in and we will benefit from their experience

# Description

Go is a great language. Go has an excellent community. Those two make a powerful combination. However, I see and hear a lot of "just use the standard library" as advice given to newcomers.

You've seen, heard, or even engaged in the conversation like:
"What framework/package do I use to build a web app?"
"Oh, you just use the standard library."

We need to stop this, we need to be more nuanced.

The fact we can give this advice, and it will work indicates how powerful the language is. However, it minimises how great the community is. This is detrimental to building the community, it can put people off, it can create a false impression of the community and the general state of the language and ecosystem.

This talk will cover a number of very useful libraries across a variety of domains that we, at Turnitin, have used and found very effective. The advice, and examples are non-commercial and it comes out of experience building system tools, micro-services, monoliths, and everything in between.

The result will be that we, as a community, can start to give better, more welcoming advice to the newcomers and to the old hands alike. Let's be more nuanced and respond  with "you could try library ..."


# Why it's a fit (dotGo Specific)
Having been to dotGo and interacted with the audience before I think that there will be something for everyone here. At Turnitin we build system tools, micro-services, and monoliths. As such the useful libraries we've found cover domains that will be common to many if not all in the audience. If audience members are new to the language, new to software or highly experienced in both I'm sure there will be something to think about and learn from this talk, be it a library specifically or better community engagement.

# Notes
I've been using Go since 2013, where I started with just the standard library. I've built micro-services, monoliths, and command tool in Go shipping to a variety of environments. I've used a lot of libraries and as an active member of my local Go community I have moved from saying "just use the standard library" to saying "you could try the following library".

I think this experience, a light hearted approach to the subject, and deep love for the Go community make me a good person to talk on the subject.