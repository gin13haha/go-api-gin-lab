# Student API with GIN

## How to Run
```bash
go run main.go
```
OUTPUT
```bash
Listening and serving HTTP on :8080
```
## How to Test
### 1. Test Get ALL
```bash
GET /students
```

### 2. GET by Student ID
```bash
GET /students/66090001
```

### 3. Test POST
```bash
POST /students
```
```json
{
  "id": "66090003",
  "name": "Alice",
  "major": "Computer Science",
  "gpa": 3.7
}
```
**Validation TEST**
```json
{
  "id": "",
  "name": "",
  "major": "CS",
  "gpa": 5.0
}
```
**Expect output**
```json
{
  "error": "invalid input"
}
```

### 4. Test GET by ID
```bash
GET /students/66090003
```
Expect output
```json
{
  "id": "66090003",
  "name": "Alice",
  "major": "Computer Science",
  "gpa": 3.7
}
```
**Test Not Found**
```bash
GET /students/9999
```
**Expect output**
```json
{
  "error": "student not found"
}
```

### 5. Test PUT
```bash
PUT /students/66090003
```
Body → raw → JSON

```json
{
  "name": "Alice Updated",
  "major": "IT",
  "gpa": 3.9
}
```
Expect result :
```200 OK```

### 6. Test DELETE
```bash
DELETE /students/66090003
```
Expect result
```
204 No Content
No body
```

```bash
GET /students/66090003
```
Expect result :
```404```