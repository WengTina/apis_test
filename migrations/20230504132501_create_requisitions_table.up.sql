CREATE TABLE  requisitions(
    requisitions_id uuid default uuid_generate_v4()  NOT NULL,
    requisitions_code TEXT default requisitions_code() not null,
    applicantname VARCHAR(50)  NOT NULL,
    company VARCHAR(50)  NOT NULL,
    department VARCHAR(50)  NOT NULL,
    product VARCHAR(50) NOT NULL,
    quantity integer NOT NULL,
    price integer NOT NULL,
    created_at timestamp default now() not null
);
create unique index requisitions_id_uindex
    on requisitions (requisitions_id);

alter table requisitions
    add constraint requisitions_pk
        primary key (requisitions_id);

create index requisitions_id_index
    on requisitions USING hash(requisitions_id);

create index requisitions_code_index
    on requisitions USING hash(requisitions_code);

create index requisitions_company_index
    on requisitions (company);

create index requisitions_department_index
    on requisitions (department);

create index requisitions_product_index
    on requisitions (product);

create index requisitions_applicantname_index
    on requisitions using gin
        (applicantname collate pg_catalog."default" gin_trgm_ops)
    tablespace pg_default;

create index requisitions_quantity_index
    on requisitions (quantity);

create index requisitions_price_index
    on requisitions (price);


create index requisitions_created_at_index
    on requisitions (created_at desc );
