namespace go demo

//--------------------request & response--------------
struct College {
    1: required string name(go.tag = 'json:"name"'),
    2: string address(go.tag = 'json:"address"'),
}

struct Student {
    1: required i32 id(api.body='id'),
    2: required string name(api.body='name'),
    3: required College college(api.body='college'),
    4: optional list<string> email(api.body='email'),
}

struct Teacher {
    1: required i32 id(api.body='id'),
    2: required string name(api.body='name'),
    3: required College college(api.body='college'),
    4: optional list<string> email(api.body='email'),
}

struct RegisterResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

struct QueryReq {
    1: required i32 id(api.query='id')
}

//----------------------service-------------------
service StudentService {
    RegisterResp StuRegister(1: Student student)(api.post = '/student/add-student-info')
    Student StuQuery(1: QueryReq req)(api.get = '/student/query')
}

service TeacherService {
    RegisterResp TeacherRegister(1: Teacher teacher)(api.post = '/teacher/add-teacher-info')
    Teacher TeacherQuery(1: QueryReq req)(api.get = '/teacher/query')
}