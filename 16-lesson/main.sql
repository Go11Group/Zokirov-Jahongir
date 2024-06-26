create table People(
    id int not null primary key,
    first_name varchar(255) ,
    last_name varchar(255) ,
    age int ,
    gender varchar(255) ,
    nation varchar(255) ,
    field varchar(255) ,
    parent_name varchar(255) ,
    city varchar(255) 
);
insert into People(id, first_name, last_name, age, gender,nation, field, parent_name, city) values 
(1, 'John', 'Doe', 30, 'Male', 'USA', 'Engineering', 'Jane Doe', 'New York'),
(2, 'Jane', 'Smith', 25, 'Female', 'UK', 'Medicine', 'John Smith', 'London'),
(3, 'Michael', 'Johnson', 40, 'Male', 'Canada', 'Finance', 'Alice Johnson', 'Toronto'),
(4, 'Emily', 'Brown', 35, 'Female', 'Australia', 'Education', 'Robert Brown', 'Sydney'),
(5, 'James', 'Davis', 28, 'Male', 'USA', 'Technology', 'Sarah Davis', 'San Francisco'),
(6, 'Olivia', 'Wilson', 32, 'Female', 'UK', 'Science', 'David Wilson', 'Manchester'),
(7, 'William', 'Moore', 45, 'Male', 'Canada', 'Healthcare', 'Linda Moore', 'Vancouver'),
(8, 'Sophia', 'Taylor', 27, 'Female', 'Australia', 'Law', 'Thomas Taylor', 'Melbourne'),
(9, 'Benjamin', 'Anderson', 33, 'Male', 'USA', 'Engineering', 'Patricia Anderson', 'Chicago'),
(10, 'Ava', 'Thomas', 29, 'Female', 'UK', 'Design', 'George Thomas', 'Birmingham'),
(11, 'Liam', 'Jackson', 31, 'Male', 'Canada', 'Marketing', 'Barbara Jackson', 'Calgary'),
(12, 'Isabella', 'White', 26, 'Female', 'Australia', 'Arts', 'Edward White', 'Perth'),
(13, 'Mason', 'Harris', 34, 'Male', 'USA', 'Technology', 'Jennifer Harris', 'Los Angeles'),
(14, 'Mia', 'Martin', 28, 'Female', 'UK', 'Science', 'Richard Martin', 'Liverpool'),
(15, 'Ethan', 'Thompson', 37, 'Male', 'Canada', 'Finance', 'Nancy Thompson', 'Ottawa'),
(16, 'Charlotte', 'Garcia', 30, 'Female', 'Australia', 'Education', 'Joseph Garcia', 'Brisbane'),
(17, 'Alexander', 'Martinez', 29, 'Male', 'USA', 'Engineering', 'Karen Martinez', 'Houston'),
(18, 'Amelia', 'Robinson', 25, 'Female', 'UK', 'Medicine', 'Charles Robinson', 'Edinburgh'),
(19, 'Lucas', 'Clark', 38, 'Male', 'Canada', 'Healthcare', 'Deborah Clark', 'Montreal'),
(20, 'Harper', 'Rodriguez', 27, 'Female', 'Australia', 'Law', 'Paul Rodriguez', 'Adelaide'),
(21, 'Henry', 'Lewis', 32, 'Male', 'USA', 'Finance', 'Michelle Lewis', 'Phoenix'),
(22, 'Ella', 'Lee', 24, 'Female', 'UK', 'Arts', 'Frank Lee', 'Sheffield'),
(23, 'Aiden', 'Walker', 35, 'Male', 'Canada', 'Marketing', 'Sandra Walker', 'Halifax'),
(24, 'Grace', 'Hall', 29, 'Female', 'Australia', 'Design', 'Raymond Hall', 'Gold Coast'),
(25, 'Sebastian', 'Allen', 31, 'Male', 'USA', 'Technology', 'Brenda Allen', 'San Diego'),
(26, 'Zoey', 'Young', 28, 'Female', 'UK', 'Science', 'Bruce Young', 'Glasgow'),
(27, 'Matthew', 'King', 36, 'Male', 'Canada', 'Engineering', 'Laura King', 'Winnipeg'),
(28, 'Aria', 'Wright', 26, 'Female', 'Australia', 'Medicine', 'Henry Wright', 'Darwin'),
(29, 'Jackson', 'Lopez', 33, 'Male', 'USA', 'Education', 'Megan Lopez', 'Dallas'),
(30, 'Scarlett', 'Hill', 25, 'Female', 'UK', 'Finance', 'Jerry Hill', 'Bristol'),
(31, 'Logan', 'Scott', 30, 'Male', 'Canada', 'Law', 'Maria Scott', 'Quebec City'),
(32, 'Chloe', 'Green', 27, 'Female', 'Australia', 'Arts', 'Steve Green', 'Hobart'),
(33, 'David', 'Adams', 34, 'Male', 'USA', 'Marketing', 'Kathy Adams', 'Denver'),
(34, 'Mila', 'Baker', 28, 'Female', 'UK', 'Design', 'Carl Baker', 'Leeds'),
(35, 'Joseph', 'Gonzalez', 29, 'Male', 'Canada', 'Technology', 'Diana Gonzalez', 'Saskatoon'),
(36, 'Lily', 'Nelson', 32, 'Female', 'Australia', 'Science', 'Gary Nelson', 'Canberra'),
(37, 'Samuel', 'Carter', 30, 'Male', 'USA', 'Healthcare', 'Betty Carter', 'Atlanta'),
(38, 'Layla', 'Mitchell', 26, 'Female', 'UK', 'Engineering', 'Eric Mitchell', 'Nottingham'),
(39, 'Owen', 'Perez', 31, 'Male', 'Canada', 'Medicine', 'Lisa Perez', 'Victoria'),
(40, 'Riley', 'Roberts', 29, 'Female', 'Australia', 'Education', 'Kevin Roberts', 'Townsville'),
(41, 'Daniel', 'Turner', 33, 'Male', 'USA', 'Finance', 'Martha Turner', 'Boston'),
(42, 'Nora', 'Phillips', 27, 'Female', 'UK', 'Law', 'Brian Phillips', 'Leicester'),
(43, 'Caleb', 'Campbell', 34, 'Male', 'Canada', 'Arts', 'Joan Campbell', 'Edmonton'),
(44, 'Victoria', 'Parker', 28, 'Female', 'Australia', 'Marketing', 'Alan Parker', 'Wollongong'),
(45, 'Luke', 'Evans', 30, 'Male', 'USA', 'Design', 'Rose Evans', 'Seattle'),
(46, 'Hannah', 'Edwards', 24, 'Female', 'UK', 'Technology', 'Roger Edwards', 'Coventry'),
(47, 'Jack', 'Collins', 32, 'Male', 'Canada', 'Science', 'Dorothy Collins', 'Hamilton'),
(48, 'Avery', 'Stewart', 26, 'Female', 'Australia', 'Engineering', 'Ralph Stewart', 'Geelong'),
(49, 'Gabriel', 'Sanchez', 29, 'Male', 'USA', 'Medicine', 'Ann Sanchez', 'Miami'),
(50, 'Eleanor', 'Morris', 27, 'Female', 'UK', 'Education', 'Fred Morris', 'Aberdeen');