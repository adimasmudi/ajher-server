# Ajher Team
- Masmudi
- Nicola Yani Alivant
- Okhi Sahrul Barkah
- Hesa Firdaus
From Google Developer Student Clubs Universitas Trunojoyo Madura

# Ajher App

Ajher is an app that can makes it easier for teachers to correct exam questions.

## General Information
- Programming Language : Go
- Database : Cloud Firestore
- API Documentation : Swagger
- Deployment : Cloud Run

```
We provide our firebase SDK publicly (not gitignore) and all environment variable inside .env.example so that the judges can try the app locally
```

## How to run the project

### clone this repository using `git clone`
### then you need to setup and install all required dependency
```
go mod init <module-name>
go mod download
go mod verify
go mod tidy
```
### run using air (like nodemon in nodejs)
```
air
```
### see the documentation in browser
- see the documentation in http://localhost:5000/docs/index.html

### enjoy🎉
