# Updates

_8/16/2025_
- going to use net/http standard library to build backend instead of chi or gin frameworks
- might help solidify my REST API foundations like what exact strings are "curl"ed 
- because I'm learning Go for distributed systems and concurrency-ness, manually setting up a TCP server might help understand the underlying network libraries in Go

_10/12/2025_

Missing a bunch of updates since 8/16/2025 (commit history should exist) but adapted golang backend to CodeBox bootcamp requirements: added a bunch of AWS microservices including deployment.
- Github Actions CI/CD fully integrated with deployment on AWS EC2
- AWS RDS Database connected and functional
still lacks a proper frontend and requires curl to demo or test. may include AWS security vulnerabilities, please be nice and dont spam my endpoint if you end up finding it.
Note to self: needs cleanup of folders for better presentation and fixes to .env variable loading logic on EC2 
