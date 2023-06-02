CREATE TABLE  purchases_products(
    purchases_products_id uuid default uuid_generate_v4()  NOT NULL,
    purchases_id uuid NOT NULL, 
    product TEXT NOT NULL,
    subtotal INTEGER NOT NULL,
    total INTEGER NOT NULL,
    created_by uuid null ,
    created_at timestamp default now() not null
);
create unique index purchases_products_id
    on purchases_products (purchases_products_id);

alter table purchases_products
    add constraint purchases_products_pk
        primary key (purchases_products_id);

create index purchases_products_id_index
    on purchases_products USING hash(purchases_products_id);

create index purchases_purchases_id_index
    on purchases_products USING hash(purchases_id);

create index purchases_products_product_index
    on purchases_products (product);

create index purchases_subtotal_index
    on purchases_products (subtotal);

create index purchases_products_total_index
    on purchases_products (total);



create index purchases_products_created_at_index
    on purchases_products (created_at desc );
