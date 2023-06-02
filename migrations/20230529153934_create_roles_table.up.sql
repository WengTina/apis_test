CREATE TABLE  roles(
    roles_id uuid default uuid_generate_v4()  NOT NULL,
    role_name TEXT NOT NULL,
    created_by uuid null ,
    created_at timestamp default now() not null
);
create unique index roles_id
    on roles (roles_id);

alter table roles
    add constraint roles_pk
        primary key (roles_id);

create index roles_id_index
    on roles USING hash(roles_id);

create index roles_role_name_index
    on roles (role_name);


create index roles_created_at_index
    on roles (created_at desc );
