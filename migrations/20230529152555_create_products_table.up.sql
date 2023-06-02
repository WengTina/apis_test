CREATE TABLE  products(
    products_id uuid default uuid_generate_v4()  NOT NULL,
    product TEXT NOT NULL,
    product_use TEXT NOT NULL,
    unit TEXT NOT NULL,
    price INTEGER NOT NULL,
    created_by uuid null ,
    created_at timestamp default now() not null
);
create unique index products_id
    on products (products_id);

alter table products
    add constraint products_pk
        primary key (products_id);

create index products_id_index
    on products USING hash(products_id);

create index products_product_index
    on products (product);

create index products_product_use_index
    on products (product_use);

create index products_unit_index
    on products (unit);

create index products_price_index
    on products (price);

create index products_created_at_index
    on products (created_at desc );
