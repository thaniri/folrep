# folrep

The repository release tracking service.

## About

folrep is short for follow repository.

The problem we are trying to solve is to send notifications when open source software that you use releases a new version.

The initial idea is to deliver a weekly report of what software you want to track, and if a new release is a different version with what you currently are using.

This is useful for things like:

* your own interest in keeping track of open source software
* configuration management repositories, where software might be deployed from a tarball to a server.

## The first high level design

![rough architecture sketch](/images/folrep-rough-sketch.jpg)

The idea behind this is a workflow working something like this.

A user creates an account on folrep, adds their Github API key, and inputs into the app which repositories they want to track. Optionally, they can input the version they are currently using.

A seperate emailing service will periodically query a database for all the repos that all of the users are tracking. Then on a per-user basis it generates a report of the latest versions of each repo they are tracking, and if any repos have gotten a new version that is different than the one they are currently using. Some options to reduce noise will be to do things like only reporting on releases that came out that week, or only reporting on releases where the version is different than the one the user is using.

At first this service will be a weekly report.
