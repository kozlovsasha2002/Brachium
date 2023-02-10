create table tasks (
    task_id bigserial,
    description varchar not null,
    maximum_execution_value int not null,
    actual_execution_value int not null,
    percentage_of_completion smallint not null,
    primary key (task_id)
);

create table periods (
    period_id bigserial,
    task_id bigint not null,
    overall_percent_complete smallint,
    start_date date not null,
    end_date date,
    amount_of_days integer,
    primary key (period_id),
    foreign key (period_id)
        references tasks (task_id)
        on delete cascade
);

create table users (
    user_id bigserial,
    nickname varchar not null,
    email varchar not null,
    period_id bigint,
    password varchar not null,
    primary key (user_id),
    foreign key (period_id)
        references periods (period_id)
        on delete cascade 
);