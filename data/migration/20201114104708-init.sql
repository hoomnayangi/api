
-- +migrate Up
create table users (
	id serial unique primary key,
	created_at timestamp(6) with time zone default now(),
	updated_at timestamp(6) with time zone,
	deleted_at timestamp(6) with time zone,
	name varchar(200),
	email varchar(200),
	avatar text
);

create table recipes (
	id serial unique primary key,
	created_at timestamp(6) with time zone default now(),
	updated_at timestamp(6) with time zone,
	deleted_at timestamp(6) with time zone,
	name varchar(200) not null,
	description text,
	picture text,
	steps jsonb
);

create table ingredients (
	id serial unique primary key,
	created_at timestamp(6) with time zone default now(),
	updated_at timestamp(6) with time zone,
	deleted_at timestamp(6) with time zone,
	name varchar(200) not null,
	description text,
	picture text
);

create table recipes_ingredients (
	recipe_id integer references recipes (id) on delete cascade not null,
	ingredient_id integer references ingredients (id) on delete cascade not null,
	amount varchar(255) not null,
	unit varchar(20),
	primary key(recipe_id, ingredient_id)
);

create table vendors (
	id serial unique primary key,
	created_at timestamp(6) with time zone default now(),
	updated_at timestamp(6) with time zone,
	deleted_at timestamp(6) with time zone,
	name varchar(200) not null,
	lat float,
	long float,
	address text
);

create table vendors_ingredients (
	vendor_id integer references vendors (id) on delete cascade not null,
	ingredient_id integer references ingredients (id) on delete cascade not null,
	amount varchar(255) not null,
	unit varchar(20),
	primary key(vendor_id, ingredient_id)
);

create table orders (
	id serial unique primary key,
	created_at timestamp(6) with time zone default now(),
	updated_at timestamp(6) with time zone,
	deleted_at timestamp(6) with time zone,
	user_id integer references users (id) on delete cascade not null,
	lat float,
	long float,
	description text
);

create table orders_recipes (
	order_id integer references orders (id) on delete cascade not null,
	recipe_id integer references recipes (id) on delete cascade not null,
	amount integer not null,
	primary key(order_id, recipe_id)
);

create table kitchens (
	id serial unique primary key,
	created_at timestamp(6) with time zone default now(),
	updated_at timestamp(6) with time zone,
	deleted_at timestamp(6) with time zone,
	lat float,
	long float,
	description text
);

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION distance(lat1 FLOAT, lon1 FLOAT, lat2 FLOAT, lon2 FLOAT) RETURNS FLOAT AS $$
DECLARE                                                   
    x float = 69.1 * (lat2 - lat1);                           
    y float = 69.1 * (lon2 - lon1) * cos(lat1 / 57.3);        
BEGIN                                                     
    RETURN sqrt(x * x + y * y);                               
END  
$$ LANGUAGE plpgsql;
-- +migrate StatementEnd

-- +migrate Down
drop table kitchens;
drop table orders_recipes;
drop table orders;
drop table vendors_ingredients;
drop table vendors;
drop table recipes_ingredients;
drop table ingredients;
drop table recipes;
drop table users;
drop function distance;