CREATE TABLE persons (
	id bigserial not null,
	name VARCHAR(100) not null,
	surname VARCHAR(100) not null,
	patronymic VARCHAR(100),
	age int not null,
	gender varchar(6) not null,
	nationality VARCHAR(2) not null
);