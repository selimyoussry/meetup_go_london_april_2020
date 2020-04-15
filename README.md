# Hacker News Golang articles extractor

This repository contains the code written for a presentation at the Golang London Meetup on April 15th, 2020. It is purely for illustrative purposes.

This is a small app that:

- periodically scraps the Hacker News homepage to look for articles about the Go language
- stores these articles in a database

We are trying to propose some default structure to start a new app by:

- clearly separating the core business logic, from the external services (like Hacker News), and from internal services (like the database we store articles in)
- adopting a test-driven approach as much as possible
- allow for rapid prototyping
