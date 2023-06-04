CREATE TABLE  employees(
    employees_id uuid default uuid_generate_v4()  NOT NULL,
    employees_name TEXT NOT NULL,
    roles_id uuid NOT NULL,
    departments_id uuid NOT NULL,
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

create index employees_roles_id_index
    on employees USING hash(roles_id);

create index employees_departments_id_index
    on employees  USING hash (departments_id);


create index employees_created_at_index
    on employees (created_at desc );
