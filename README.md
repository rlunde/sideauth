# sideauth
A simple REST microservice to use with any other service to handle common auth tasks

# disclaimer
DISCLAIMER: I'm not expecting anyone to look at this project, or use it, or contribute to it, for
quite a while. I'm using it as a learning exercise more than anything else at the moment, with the
goal of eventually knowing golang better (in this case). I will remove this disclaimer when
I think the project has reached MVP (minimum viable project) stage. That won't be for a while.

NOTE: I'm not abandoning this project, but I've decided to spend some time really focusing on
another project (goodwin), which I will then use to try to help make this project.

## motivation
I found myself trying to make several different sample services, that I wanted
to run myself as default webservices, as well as (hopefully) provide a base for
other people. The problem was I always got hung up on the first step of that,
which is just to let users create an account and log in (or use OAuth2) so they
could save their data across sessions.

It occurred to me that I could just set up a "side" service with a small set of
REST APIs that handled all the tedious chores related to that:
 - create an account (if it doesn't already exist)
 - change a password
 - email a link to reset a lost password
 - log in and create a session
 - log out and destroy a session
 - validate a session

That way all of the projects I'm working on in whatever language (python, go,
javascript) could just pass through those operations.

At least that's the theory. We'll see how it goes.

The Madness to my Method:

Anyone looking at this (and other of my repos) will see that I've got a lot of little tiny
commits, many of which are hardly more than documentation. I'm trying out something called
the Seinfeld Method, in which the important thing is just to make a daily effort, not
how much you get done in any one day.

http://lifehacker.com/5886128/how-seinfelds-productivity-secret-fixed-my-procrastination-problem

Combine that with Dr. B.J.Fogg's method for habit development, which says that tiny
incremental changes are not just OK, but they're a good idea, and you've got a good
explanation of what I'm doing.

http://www.foggmethod.com/
