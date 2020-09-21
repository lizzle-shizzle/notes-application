# Notes application

A simple note taking application where you can create a note, which gets saved to the database, and you can view all your created notes.

### Run:
1. Run cd frontend && sudo npm i from project directory
2. Run cd ../
3. Run chmod u+x build-images.sh run.sh cleanup.sh ./backend/build.sh ./frontend/build.sh
4. Run sudo ./build-images.sh - binary is already included for Go backend so Go doesn't need to be installed
5. Run sudo ./run.sh
6. Access application at www.notes-application.co.za:3001
7. Run sudo ./cleanup.sh and enter 'y'

### Endpoints
- POST localhost:8080/notes
	- Request body:
```
	{
		"text": "My first note"
	}
```
- GET localhost:8080/notes
    - Response:
```
	[
		{
			"id": 1,
			"text": "My first note",
			"created_timestamp": "2020-01-01T12:00:00"
		}
	]
```
	
