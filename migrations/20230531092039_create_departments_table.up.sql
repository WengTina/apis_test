CREATE TABLE  departments(
    departments_id uuid default uuid_generate_v4()  NOT NULL,
    departments_code TEXT default departments_code() not null,
    department_name TEXT NOT NULL,
    created_by uuid null ,
    created_at timestamp default now() not null
);
create unique index departments_id
    on departments (departments_id);

alter table departments
    add constraint departments_pk
        primary key (departments_id);

create index departments_id_index
    on departments USING hash(departments_id);

create index departments_code_index
    on departments USING hash(departments_code);

create index departments_department_name_index
    on departments (department_name);


create index departments_created_at_index
    on departments (created_at desc );
