create table tasks (
    task_id bigserial,
    description varchar not null,
    maximum_execution_value int not null,
    actual_execution_value int not null,
    percentage_of_completion smallint not null,
    primary key (task_id)
);