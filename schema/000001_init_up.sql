create table users (
    user_id serial,
    nickname varchar not null,
    email varchar not null,
    password varchar not null,
    primary key (user_id)
);

create table check_lists (
    list_id serial,
    overall_percent_complete smallint,
    start_date varchar not null,
    end_date varchar,
    user_id int not null,
    primary key (list_id),
    foreign key (user_id) references users(user_id) on delete cascade
);

create table tasks (
    task_id serial,
    description varchar not null,
    maximum_execution_value int not null,
    actual_execution_value int not null,
    percentage_of_completion smallint not null,
    list_id int not null,
    primary key (task_id),
    foreign key (list_id) references check_lists(list_id) on delete cascade
);