use mbanking_project;
show tables from mbanking_project;


create table Users(
	id int primary key not null auto_increment,
	name varchar(50) not null,
	phone varchar(13) unique not null,
	password varchar(8) not null,
	saldo int default 0,
	created_at datetime default current_timestamp,
	update_at datetime default current_timestamp,
	deleted_at datetime
);

create table Transfer(
	id int primary key not null auto_increment,
	user_id_pengirim int,
	user_id_penerima int,
	value int not null,
	created_at datetime default current_timestamp,
	constraint fk_data_user foreign key (user_id_pengirim) references users(id) on update cascade on delete cascade,
	constraint fk_data_users foreign key (user_id_penerima) references users(id) on update cascade on delete cascade
);

create table TopUp(
	id int primary key not null auto_increment,
	user_id int,
	value int not null,
	created_at datetime default current_timestamp,
	constraint fk_data_user1 foreign key (user_id) references users(id) on update cascade on delete cascade
);

select * from Users;

