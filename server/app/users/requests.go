package users

/*user_id int IDENTITY(1, 1) primary key,
first_name varchar(50) not null,
last_name varchar(50) not null,
email varchar(60) not null unique,
passw varchar(255) not null,
tax_reduction int,
rol varchar(50) not null,
deletion_date date*/

type CreateUserRequest struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UpdateDeleteDateRequest struct {
	Id int
}
