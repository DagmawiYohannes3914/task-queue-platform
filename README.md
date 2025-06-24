Task Queue Processing Platform (Golang)

🏗 Project Overview

A full-stack distributed asynchronous task processing system built entirely in Golang using:

✅ REST API (Chi router)

✅ PostgreSQL (GORM ORM)

✅ NATS (message queue)

✅ JWT-based authentication

✅ API Key support for external integrations

✅ Worker service for job processing

✅ Fully Dockerized and scalable

🎯 System Architecture
API server handles registration, authentication, job submission, and job status query.

NATS message queue used to decouple job processing.

Worker service processes jobs asynchronously from queue.

Secure multi-tenant system using JWT and API Keys.

🛠 Technologies
Stack
Language Golang
Database PostgreSQL
Queue NATS
Auth JWT + API Keys
Docker Full multi-service deployment
