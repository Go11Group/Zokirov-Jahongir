create table student(
    id uuid primary key default gen_random_uuid() not null,
    name varchar,
    age int
    );

create table course(
    id uuid primary key default gen_random_uuid() not null,
    name varchar
    );


create table student_course(
    id uuid primary key default gen_random_uuid() not null,
    student_id uuid references student(id),
    course_id uuid references course(id)
    );

create table grade(
    id uuid primary key default gen_random_uuid() not null,
    student_course_id uuid references student_course(id),
    grade int
    );


insert into student(name, age) 
values
    ('Muhammadzokir', 19),
    ('Azam', 20),
    ('muhamadaziz', 11),
    ('Polvon', 32),
    ('aziz', 21),
    ('Jahongir', 22),
    ('Bekzod', 20),
    ('sardor', 23),insert into coursstudent_coursee(name)
values
    ('Goland'),
    ('Fluter'),fmt.Println(db)
    ('Javaskript'),
    ('Rust');
    ('Sheker', 22);

insert into course(name)
values
    ('Golang'),
    ('Flutter'),
    ('Javascript'),
    ('Rust');

insert into student_course(student_id, course_id)
values
    ('fb8267ad-0c11-4bcd-83e0-baceaa41c8a6', '35136ec3-1336-4a10-9d9c-a67642d0452b'),
    ('b2ab2ecd-9e65-4721-ba93-e5e0529d5594', '479f07f8-4546-4769-8f75-1ed20badff6d'),
    ('d992345a-d4c5-4ed6-8d8d-b2a097ac3633', 'bca0aa50-ede8-4e84-ac8b-11b6d9ea46f8'),
    ('bd3dfef9-9496-4908-b817-a6d2457de8d7', 'bca0aa50-ede8-4e84-ac8b-11b6d9ea46f8'),
    ('bd3dfef9-9496-4908-b817-a6d2457de8d7', 'cf1f688a-cd44-486d-8ddd-55e720b0b379'),
    ('d6a645ba-5bf1-4cf0-bb32-f788550c66aa', 'beb93c7e-ea2f-4c49-992e-d02df4d6869c'),
    ('5d79f20b-fead-4663-ba7c-826e7ac02088', 'cf1f688a-cd44-486d-8ddd-55e720b0b379'),
    ('153f5982-f811-492c-84a7-6835b4a38f12', '479f07f8-4546-4769-8f75-1ed20badff6d'),
    ('4ae6f8da-5a29-4e1d-870d-20a6bdd35b03', '35136ec3-1336-4a10-9d9c-a67642d0452b'),
    ('897340cc-c32c-4fa0-b757-6e0a6b41524b', 'bca0aa50-ede8-4e84-ac8b-11b6d9ea46f8');

insert into grade(student_course_id, grade)
values
    ('94cc9a66-2d96-4ebe-b069-8ff423ca5e3a', 5),
    ('c2fd92a1-0b01-4607-97bb-602282f2d902', 5),
    ('c2fd92a1-0b01-4607-97bb-602282f2d902', 5),
    ('c2fd92a1-0b01-4607-97bb-602282f2d902', 5),
    ('c2fd92a1-0b01-4607-97bb-602282f2d902', 5),
    ('1a098273-df9d-4dfa-8df3-50e697d8f072', 4),
    ('1a098273-df9d-4dfa-8df3-50e697d8f072', 5),
    ('7f6207df-b841-4981-91b1-dfd94c41a389', 3),
    ('7f6207df-b841-4981-91b1-dfd94c41a389', 4),
    ('c9d0c662-c9f8-4f8d-ace4-e1e189971672', 4),
    ('c9d0c662-c9f8-4f8d-ace4-e1e189971672', 5),
    ('72a23e39-8f15-401b-a5fd-615e4513b0b6', 4),
    ('72a23e39-8f15-401b-a5fd-615e4513b0b6', 3),
    ('72a23e39-8f15-401b-a5fd-615e4513b0b6', 5),
    ('3d8ff7e6-fa38-4057-b977-f27cd3639035', 5),
    ('3d8ff7e6-fa38-4057-b977-f27cd3639035', 4),
    ('3547f7f2-d208-4750-a507-b8a77a964062', 2),
    ('3547f7f2-d208-4750-a507-b8a77a964062', 4),
    ('606f5d67-8459-4be8-b267-1efd90c863bb', 3),
    ('606f5d67-8459-4be8-b267-1efd90c863bb', 5),
    ('94cc9a66-2d96-4ebe-b069-8ff423ca5e3a', 4),
    ('94cc9a66-2d96-4ebe-b069-8ff423ca5e3a', 3),
    ('37ef63df-aaa3-4e3d-881e-2a3bc01c7152', 4),
    ('37ef63df-aaa3-4e3d-881e-2a3bc01c7152', 5);

select c.name, array_agg(s.name) as students
from course as c
right join student_course as sc
on sc.course_id = c.id
right join student as s
on sc.student_id = s.id
group by c.name;


select c.name, max(round(avg(g.grade)::numeric, 2)) as avarage_grade
    from student as s
    right join student_course as sc
    on s.id = sc.student_id
    right join course as c
    on c.id = sc.id
    right join grade as g
    on sc.id = g.student_course_id
    group by c.id;


with students_avg_grade as (
    select 
        s.id as student_id, sc.course_id as course_id, round(avg(g.grade)::numeric, 2) as avarage_grade
    from student as s
    right join 
        student_course as sc
    on 
        s.id = sc.student_id
    right join 
        grade as g
    on 
        sc.id = g.student_course_id
    group by 
        sc.course_id, s.id
),
max_scores as (
    select course_id, max(avarage_grade) as max_grade
    from students_avg_grade
    group by course_id
)

select c.name, m.max_grade, array_agg(s.name)
from course as c
join max_scores as m
on m.course_id = c.id
join students_avg_grade as sag
on sag.course_id = c.id and m.max_grade = sag.avarage_grade
join student as s
on sag.student_id = s.id 
group by c.name, m.max_grade;

select c.name, round(avg(g.grade)::numeric, 2) as avarage_grade
from course as c
right join student_course as sc
on sc.course_id = c.id
right join grade as g
on sc.id = g.student_course_id
group by c.name
order by avarage_grade desc;

with minageofgroups as(
    select
        c.id as course_id,
        min(s.age) as min_age
    from course as c
    right join student_course as sc
    on sc.course_id = c.id
    right join student as s
    on sc.student_id = s.id
    group by c.id
)
select c.name, array_agg(s.name)
from course as c
right join student_course as sc
on sc.course_id = c.id
right join minageofgroups as m
on m.course_id = c.id
join student as s
on sc.student_id = s.id and m.min_age = s.age
group by  c.name;

select c.name, round(avg(g.grade)::numeric, 2) as avarage_grade
from course as c
right join student_course as sc
on sc.course_id = c.id
right join grade as g
on sc.id = g.student_course_id
group by c.name
order by avarage_grade desc
limit 1;