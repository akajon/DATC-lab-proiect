package dangers

/*danger_id int IDENTITY(1, 1) primary key,
category varchar(60) not null,
name varchar(60) not null,
description varchar(255) not null,
grade int not null*/

type CreateDangerRequest struct {
	Category    string
	Name        string
	Description string
	Grade       int
}

type DeleteDangerRequest struct {
	Id int
}
