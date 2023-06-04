CREATE TABLE  purchases(
    purchases_id uuid default uuid_generate_v4()  NOT NULL,
    applicant_id uuid  NOT NULL,
    purchases_product_id uuid NOT NULL,
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



create index purchases_applicant_id_index
    on purchases USING hash (applicant_id);

create index purchases_purchases_product_id_index
    on purchases USING hash (purchases_product_id);

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
