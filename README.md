# TP middleware 

## Project Structure

- **songs**: Go API for handling songs data
- **users**: Go API for managing user information
- **ratings**: API for managing ratings
- **flask_api**: Flask API for integrating data from the three other APIs
- **front**: Vue.js frontend


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

#### Documentation

Documentation is visible in **api** directory.



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

