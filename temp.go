// front-end, html form, take user details, send json
// send json via hhtp post to heroku server
// parse to user details object at backend
// encrypt password at front-end and store it in encrypted form

type Task struct {
    Name string
    Count []string
}

type TaskMap struct {
    Tasks map[string]*Task
    Count int
    Session int
}

type UserDetails struct {
    uid int
    uname string
    pass string
    fname string
    lname string
    email string
    phone string
}

type User struct {
    details *UserDetails
    t *TaskMap
}

type UserDB struct {
    user map[string]*User
}