CREATE TABLE  employees(
    employees_id uuid default uuid_generate_v4()  NOT NULL,
    employees_name TEXT NOT NULL,
    roles uuid NOT NULL,
    department TEXT NOT NULL,
    created_by uuid NULL,
    created_at timestamp default now() NOT NULL
);
create unique index employees_id
    on employees (employees_id);

alter table employees
    add constraint employees_pk
        primary key (employees_id);

create index employees_id_index
    on employees USING hash(employees_id);

create index employees_employees_name_index
    on employees (employees_name);

create index roles_index
    on employees USING hash(roles);

create index employees_department_index
    on employees (department);


create index employees_created_at_index
    on employees (created_at desc );
