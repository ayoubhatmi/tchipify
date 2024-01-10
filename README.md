# TP middleware 
### Group : BAHOU yassine & HATMI Ayoub
### 

## Project Structure

- **songs**: Go API for handling songs data
  url = http://localhost:8080/songs/
- **users**: Go API for managing user information
  url = http://localhost:8081/users/
- **ratings**: API for managing ratings
  url = https://ratings-mike.edu.forestier.re
- **flask_api**: Flask API for integrating data from the three other APIs
  url = http://localhost:8888/
- **front**: Vue.js frontend (work in progress)


## Getting Started

### Go APIs

#### Run the following commands in the terminal:
Tidy / download modules :
```
go mod tidy
```
Build & run :
```
go run cmd/main.go
```

---
Or build : 
```
go build -o middleware_collections cmd/main.go
```
Then run : 
```
./middleware_collections
```

### Flask API
#### Run the following commands in the terminal:

Download modules :
```
pip install -r requirements.txt
```
On Debian : `pip install -r requirements.txt --break-system-packages`  

Run (command line) from base directory (*flask_base*) :
```
PYTHONPATH=$PYTHONPATH:$(pwd) python3 src/app.py
```

If a warning about your PATH appears :  
```
export PATH=$PATH:$HOME/.local/bin
```

#### Documentation

Documentation is visible in **/api/docs** when running the app.

