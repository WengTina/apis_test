CREATE TABLE  purchases(
    purchases_id uuid default uuid_generate_v4()  NOT NULL,
    applicant_name TEXT  NOT NULL,
    product TEXT NOT NULL,
    quantity INTEGER NOT NULL,
    purchases_reason TEXT NOT NULL,
    purchases_date DATE NOT NULL,
    demand_date DATE NOT NULL,
    remark TEXT NOT NULL,
    created_by uuid null ,
    created_at timestamp default now() not null
);
create unique index purchases_id_uindex
    on purchases (purchases_id);

alter table purchases
    add constraint purchases_pk
        primary key (purchases_id);

create index purchases_id_index
    on purchases USING hash(purchases_id);



create index purchases_applicant_name_index
    on purchases (applicant_name);

create index purchases_product_index
    on purchases (product);

create index purchases_quantity_index
    on purchases (quantity);

create index purchases_purchases_reason_index
    on purchases (purchases_reason);

create index purchases_purchases_date_index
    on purchases (purchases_date);

create index purchases_demand_date_index
    on purchases (demand_date);



create index purchases_remark_index
    on purchases (remark);

create index purchases_created_at_index
    on purchases (created_at desc );
