-- create user da with password '1234';
drop table if exists moto_images;
drop table if exists moto_videos;
drop table if exists motorcycle_parameters;
drop table if exists motorcycles;
drop table if exists moto_category_parameters;
drop table if exists moto_parameter_values;
drop table if exists moto_parameters;
drop table if exists moto_models;
drop table if exists moto_brands;
drop table if exists moto_categories;
drop type if exists price_type_enum;

drop table if exists user_likes;
drop table if exists temp_users;
drop table if exists configurations;
drop table if exists images;
drop table if exists videos;
drop table if exists vehicles;
drop table if exists profiles;
drop table if exists regions;
drop table if exists cities;
drop table if exists generation_modifications;
drop table if exists colors;
drop table if exists services;
drop table if exists service_types;
drop table if exists fuel_types;
drop table if exists drivetrains;
drop table if exists engines;
drop table if exists transmissions;
drop table if exists generations;
drop table if exists body_types;
drop table if exists models;
drop table if exists brands;
drop table if exists users;
drop table if exists admins;
drop table if exists ownership_types;

create table temp_users (
    "id" serial primary key,
    "email" varchar(100),
    "username" varchar(100) not null,
    "role_id" int not null default 1, -- 1 user, 2 dealer, 3 logistic, 4 broker
    "password" varchar(100) not null,
    "phone" varchar(100),
    "status" int not null default 1, -- 1 active, 2 inactive
    "registered_by" varchar(20) not null,
    "updated_at" timestamp default now(),
    "created_at" timestamp default now(),
    unique("email"),
    unique("phone")
);

create table users (
    "id" serial primary key,
    "email" varchar(100),
    "username" varchar(100) not null,
    "role_id" int not null default 1, -- 1 user, 2 dealer, 3 logistic, 4 broker
    "password" varchar(100) not null,
    "phone" varchar(100),
    "status" int not null default 1, -- 1 active, 2 inactive
    "updated_at" timestamp default now(),
    "created_at" timestamp default now(),
    unique("email"),
    unique("phone")
);

create table user_parameters (
    "id" serial primary key, 
    "name" varchar(50) not null,
    "updated_at" timestamp not null default now(),
    "created_at" timestamp not null default now()
);



create table user_parameter_values (
    "id" serial primary key, 
    "name" varchar(50) not null,
    "parameter_id" int not null,
    "updated_at" timestamp not null default now(),
    "created_at" timestamp not null default now(),
    constraint fk_user_parameter_values_parameter_id
        foreign key ("parameter_id")
            references user_parameters("id")
                on delete cascade
                on update cascade,
    unique("name", "parameter_id")
);

create table user_type_parameters (
    "id" serial primary key, 
    "user_type" int not null,
    "parameter_id" int not null,
    "updated_at" timestamp not null default now(),
    "created_at" timestamp not null default now(),
    constraint fk_user_type_parameters_parameter_id
        foreign key ("parameter_id")
            references user_parameters("id")
                on delete cascade
                on update cascade,
    unique("name", "parameter_id")
)

create table messages (
    "id" serial primary key,
    "sender_id" int not null,
    "receiver_id" int not null,
    "message" varchar(500) not null, --  it is an id if type "item".
    "type" int not null default 1, -- 1-text, 2-item
    "created_at" timestamp default now(),
    constraint messages_sender_id_fk
        foreign key (sender_id)
            references users(id)
                on delete cascade
                on update cascade,
    constraint messages_receiver_id_fk
        foreign key (receiver_id)
            references users(id)
                on delete cascade
                on update cascade
);

create table cities (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now()
);

create table profiles (
    "id" serial primary key, 
    "user_id" int not null,
    "city_id" int,
    "driving_experience" int,
    "notification" boolean default false,
    "username" varchar(100) not null,
    "registered_by" varchar(20) not null,
    "google" varchar(200),
    "avatar" varchar(200),
    "birthday" date,
    "about_me" varchar(300),
    "last_active_date" timestamp default now(),
    "created_at" timestamp default now(),
    constraint profiles_user_id_fk 
        foreign key (user_id) 
            references users(id) 
                on delete cascade 
                on update cascade,
    constraint profiles_city_id_fk 
        foreign key (city_id) 
            references cities(id) 
                on delete cascade 
                on update cascade,
    unique (user_id)
);


create table admins (
    "id" serial primary key,
    "username" varchar(255) not null,
    "email" varchar(255) not null,
    "password" varchar(255) not null,
    "last_active_date" timestamp default now(),
    "created_at" timestamp default now()
);


create table brands (
    "id" serial primary key,
    "name" varchar(255) not null,
    "logo" varchar(255) not null,
    "model_count" int not null default 0,
    "popular" boolean default false,
    "updated_at" timestamp default now(),
    unique("name")
);


create table models (
    "id" serial primary key,
    "name" varchar(255) not null,
    "brand_id" int not null,
    "popular" boolean default false,
    "updated_at" timestamp default now(),
    constraint models_brand_id_fk 
        foreign key (brand_id) 
            references brands(id)
                on delete cascade
                on update cascade
);



create table body_types (
    "id" serial primary key,
    "name" varchar(255) not null,
    "image" character varying(255) not null,
    "created_at" timestamp default now()
);

create table transmissions (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now(),
    unique("name")
);

create table engines (
    "id" serial primary key,
    "value" varchar(255) not null,
    "created_at" timestamp default now(),
    unique("value")
);

create table drivetrains (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now(),
    unique("name")
);

create table fuel_types (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now(),
    unique("name")
);

create table regions (
    "id" serial primary key,
    "name" varchar(255) not null,
    "city_id" int not null,
    "created_at" timestamp default now(),
    constraint regions_city_id_fk
        foreign key (city_id)
            references cities(id)
                on delete cascade
                on update cascade
);


create table service_types (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now()
);

create table services (
    "id" serial primary key,
    "name" varchar(255) not null,
    "service_type_id" int not null,
    "created_at" timestamp default now(),
    constraint services_service_type_id_fk
        foreign key (service_type_id)
            references service_types(id)
                on delete cascade
                on update cascade
);



create table generations (
    "id" serial primary key,
    "name" varchar(255) not null,
    "model_id" int not null,
    "start_year" int not null,
    "end_year" int not null,
    "wheel" boolean not null default true,
    "image" varchar(255) not null,
    "created_at" timestamp default now(),
    constraint generations_model_id_fk
        foreign key (model_id)
            references models(id)
                on delete cascade
);

create table configurations (
    "id" serial primary key,
    "body_type_id" int not null,
    "generation_id" int not null,
    constraint configurations_generation_id_fk
        foreign key (generation_id)
            references generations(id)
                on delete cascade
                on update cascade,
    constraint configurations_body_type_id_fk
        foreign key (body_type_id)
            references body_types(id)
                on delete cascade
                on update cascade
);

create table generation_modifications (
    "id" serial primary key,
    "generation_id" int not null,
    "body_type_id" int not null,
    "engine_id" int not null,
    "fuel_type_id" int not null, 
    "drivetrain_id" int not null,
    "transmission_id" int not null, 
    unique(generation_id, body_type_id, engine_id, fuel_type_id, drivetrain_id, transmission_id),
    constraint generation_modifications_generation_id_fk
        foreign key (generation_id)
            references generations(id)
                on delete cascade
                on update cascade,
    constraint generation_modifications_engine_id_fk
        foreign key (engine_id)
            references engines(id)
                on delete cascade
                on update cascade,
    constraint generation_modifications_fuel_type_id_fk
        foreign key (fuel_type_id)
            references fuel_types(id)
                on delete cascade
                on update cascade,
    constraint generation_modifications_drivetrain_id_fk
        foreign key (drivetrain_id)
            references drivetrains(id)
                on delete cascade
                on update cascade,
    constraint generation_modifications_transmission_id_fk
        foreign key (transmission_id)
            references transmissions(id)
                on delete cascade
                on update cascade,
    constraint generation_modifications_body_type_id_fk
        foreign key (body_type_id)
            references body_types(id)
                on delete cascade
                on update cascade
);

create table ownership_types (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now()
);


create table colors (
    "id" serial primary key,
    "name" varchar(255) not null,
    "image" varchar(255) not null,
    "created_at" timestamp default now()
);


create table vehicles (
    "id" serial primary key,
    "user_id" int not null,
    "modification_id" int not null,
    "brand_id" int,
    "region_id" int,
    "city_id" int default 1,
    "model_id" int,
    "ownership_type_id" int not null default 1,
    "owners" int not null default 1,
    "view_count" int not null default 0,
    "year" int not null,
    "popular" int not null default 0,
    "description" text,
    "credit" boolean not null default false,
    "wheel" boolean not null default true, -- true left, false right
    "crash" boolean not null default false,
    "odometer" int not null default 0,
    "vin_code" varchar(255) not null,
    "phone_numbers" varchar(255)[] not null,
    "price" int not null,
    "new" boolean not null default false,
    "color_id" int not null,
    "trade_in" int not null default 1, -- 1. No exchange 2. Equal value 3. More expensive 4. Cheaper 5. Not a car
    "status" int not null default 3, -- 1-pending, 2-not sale (my cars), 3-on sale,
    "updated_at" timestamp default now(),
    "created_at" timestamp default now(),
    constraint vehicles_color_id_fk
        foreign key (color_id)
            references colors(id)
                on delete set null
                on update cascade,
    constraint vehicles_ownership_type_id_fk
        foreign key (ownership_type_id)
            references ownership_types(id)
                on delete cascade
                on update cascade,
    constraint vehicles_user_id_fk
        foreign key (user_id)
            references users(id)
                on delete cascade
                on update cascade,
    constraint vehicles_brand_id_fk
        foreign key (brand_id)
            references brands(id)
                on delete cascade
                on update cascade,
    constraint vehicles_model_id_fk
        foreign key (model_id)
            references models(id)
                on delete cascade
                on update cascade,
    constraint vehicles_modification_id_fk
        foreign key (modification_id)
            references generation_modifications(id)
                on delete cascade
                on update cascade,
    constraint vehicles_region_id_fk
        foreign key (region_id)
            references regions(id)
                on delete cascade
                on update cascade,
    constraint vehicles_city_id_fk
        foreign key (city_id)
            references cities(id)
                on delete cascade
                on update cascade
);


(1, 310, 3232, 1, 1, 1, 1, 2020, 50000, 'VIN1234567890', false, 1, false, ARRAY['123456789'], 25000, false, 1, 1, 3),
(1, 310, 3232, 1, 1, 1, 1, 2020, 30000, 'VIN1234567891', false, 1, true, ARRAY['123456790'], 28000, false, 2, 1, 3),
(2, 310, 3232, 1, 1, 2, 1, 2019, 70000, 'VIN1234567892', false, 2, false, ARRAY['123456791'], 22000, false, 3, 2, 3),
(1, 310, 3232, 1, 1, 1, 1, 2020, 15000, 'VIN1234567893', false, 1, true, ARRAY['123456792'], 32000, false, 1, 1, 3),
(2, 310, 3232, 1, 1, 3, 1, 2018, 85000, 'VIN1234567894', true, 2, false, ARRAY['123456793'], 18000, false, 2, 3, 3),
(1, 310, 3233, 1, 1, 1, 1, 2020, 25000, 'VIN1234567895', false, 1, true, ARRAY['123456794'], 35000, false, 3, 1, 3),
(2, 310, 3233, 1, 1, 2, 1, 2020, 40000, 'VIN1234567896', false, 1, false, ARRAY['123456795'], 30000, false, 1, 2, 3),
(1, 310, 3233, 1, 1, 1, 1, 2020, 12000, 'VIN1234567897', false, 1, true, ARRAY['123456796'], 38000, false, 2, 1, 3),
(2, 310, 3233, 1, 1, 3, 1, 2019, 60000, 'VIN1234567898', false, 2, false, ARRAY['123456797'], 27000, false, 3, 4, 3),
(1, 310, 3233, 1, 1, 1, 1, 2020, 8000, 'VIN1234567899', false, 1, true, ARRAY['123456798'], 42000, true, 1, 1, 3),
(2, 310, 3234, 1, 1, 2, 1, 2020, 35000, 'VIN1234567900', false, 1, false, ARRAY['123456799'], 28000, false, 2, 2, 3),
(1, 310, 3234, 1, 1, 1, 1, 2020, 22000, 'VIN1234567901', false, 1, true, ARRAY['123456800'], 33000, false, 3, 1, 3),
(2, 310, 3234, 1, 1, 3, 1, 2019, 75000, 'VIN1234567902', true, 3, false, ARRAY['123456801'], 24000, false, 1, 3, 3),
(1, 310, 3234, 1, 1, 1, 1, 2020, 18000, 'VIN1234567903', false, 1, true, ARRAY['123456802'], 36000, false, 2, 1, 3),
(2, 310, 3234, 1, 1, 2, 1, 2023, 5000, 'VIN1234567904', false, 1, false, ARRAY['123456803'], 45000, true, 3, 5, 3),
(1, 310, 3232, 1, 1, 1, 1, 2020, 95000, 'VIN1234567905', false, 3, false, ARRAY['123456804'], 16000, false, 1, 2, 3),
(2, 310, 3232, 1, 1, 3, 1, 2021, 28000, 'VIN1234567906', false, 1, true, ARRAY['123456805'], 31000, false, 2, 1, 3),
(1, 310, 3232, 1, 1, 1, 1, 2020, 45000, 'VIN1234567907', false, 2, false, ARRAY['123456806'], 26000, false, 3, 4, 3),
(2, 310, 3233, 1, 1, 2, 1, 2022, 16000, 'VIN1234567908', false, 1, true, ARRAY['123456807'], 37000, false, 1, 1, 3),
(1, 310, 3233, 1, 1, 1, 1, 2020, 65000, 'VIN1234567909', true, 2, false, ARRAY['123456808'], 25000, false, 2, 3, 3),
(2, 310, 3233, 1, 1, 3, 1, 2023, 3000, 'VIN1234567910', false, 1, true, ARRAY['123456809'], 48000, true, 3, 1, 3),
(1, 310, 3234, 1, 1, 1, 1, 2020, 52000, 'VIN1234567911', false, 2, false, ARRAY['123456810'], 29000, false, 1, 2, 3),
(2, 310, 3234, 1, 1, 2, 1, 2021, 33000, 'VIN1234567912', false, 1, true, ARRAY['123456811'], 34000, false, 2, 1, 3),
(1, 310, 3234, 1, 1, 1, 1, 2020, 88000, 'VIN1234567913', true, 3, false, ARRAY['123456812'], 20000, false, 3, 3, 3),
(2, 310, 3232, 1, 1, 3, 1, 2022, 14000, 'VIN1234567914', false, 1, true, ARRAY['123456813'], 39000, false, 1, 1, 3),
(1, 310, 3232, 1, 1, 1, 1, 2020, 72000, 'VIN1234567915', false, 2, false, ARRAY['123456814'], 23000, false, 2, 4, 3),
(2, 310, 3233, 1, 1, 2, 1, 2023, 6000, 'VIN1234567916', false, 1, true, ARRAY['123456815'], 46000, true, 3, 1, 3),
(1, 310, 3233, 1, 1, 1, 1, 2020, 38000, 'VIN1234567917', false, 1, false, ARRAY['123456816'], 32000, false, 1, 2, 3),
(2, 310, 3234, 1, 1, 3, 1, 2021, 26000, 'VIN1234567918', false, 1, true, ARRAY['123456817'], 35000, false, 2, 1, 3),
(1, 310, 3234, 1, 1, 1, 1, 2020, 19000, 'VIN1234567919', false, 1, false, ARRAY['123456818'], 41000, false, 3, 5, 3);

create table images (
    "vehicle_id" int not null,
    "image" varchar(255) not null,
    "created_at" timestamp not null default now(),
    constraint images_vehicle_id_fk
        foreign key (vehicle_id)
            references vehicles(id)
                on delete cascade
                on update cascade
);

CREATE TABLE user_likes (
    user_id INT NOT NULL,
    vehicle_id INT NOT NULL,
    PRIMARY KEY (user_id, vehicle_id),
    constraint user_likes_vehicle_id_fk
        foreign key (vehicle_id)
            references vehicles(id)
                on delete cascade
                on update cascade,
    constraint user_likes_user_id_fk
        foreign key (user_id)
            references users(id)
                on delete cascade
                on update cascade
);

create table videos (
    "vehicle_id" int not null,
    "video" varchar(255) not null,
    "created_at" timestamp not null default now(),
    constraint videos_vehicle_id_fk
        foreign key (vehicle_id)
            references vehicles(id)
                on delete cascade
                on update cascade
);


create table moto_categories (
    "id" serial primary key,
    "name" varchar(100) not null,
    "created_at" timestamp not null default now()
);

create table moto_brands (
    "id" serial primary key,
    "name" varchar(100) not null,
    "image" varchar(255) not null,
    "moto_category_id" integer not null,
    "created_at" timestamp not null default now(),
    constraint moto_brands_moto_category_id_fk
        foreign key (moto_category_id)
            references moto_categories(id)
                on delete cascade
                on update cascade
);


create table moto_models (
    "id" serial primary key,
    "name" varchar(100) not null,
    "moto_brand_id" integer not null,
    "created_at" timestamp not null default now(),
    constraint moto_brand_models_moto_brand_id_fk
        foreign key (moto_brand_id)
            references moto_brands(id)
                on delete cascade
                on update cascade
);

create table moto_parameters (
    "id" serial primary key,
    "moto_category_id" int not null,
    "name" varchar(100) not null,
    "created_at" timestamp default now(),
    constraint moto_parameters_moto_category_id_fk
        foreign key (moto_category_id)
            references moto_categories(id)
                on delete cascade
                on update cascade,
    unique("name", "moto_category_id")
);

create table moto_parameter_values (
    "id" serial primary key,
    "moto_parameter_id" int not null,
    "name" varchar(100) not null,
    "created_at" timestamp default now(),
    constraint moto_parameter_values_moto_parameter_id_fk
        foreign key (moto_parameter_id)
            references moto_parameters(id)
                on delete cascade
                on update cascade
);

create table moto_category_parameters (
    "moto_category_id" int not null,
    "moto_parameter_id" int not null,
    "created_at" timestamp not null default now(),
    constraint moto_category_parameters_moto_category_id_fk
        foreign key (moto_category_id)
            references moto_categories(id)
                on delete cascade
                on update cascade,
    constraint moto_category_parameters_moto_parameter_id_fk
        foreign key (moto_parameter_id)
            references moto_parameters(id)
                on delete cascade
                on update cascade
);

CREATE TYPE price_type_enum AS ENUM ('USD', 'AED', 'RUB', 'EUR');

create table motorcycles (
    "id" serial primary key,
    "user_id" int not null,
    "moto_category_id" int not null,
    "moto_brand_id" int not null,
    "moto_model_id" int not null,
    "fuel_type_id" int not null,
    "city_id" int not null,
    "color_id" int not null,
    "engine" int, -- cm3
    "power" int, -- hp
    "year" int not null,
    "number_of_cycles" int,
    "odometer" int not null default 0,
    "crash" boolean,
    "not_cleared" boolean,
    "owners" int,
    "date_of_purchase" varchar (50),
    "warranty_date" varchar(50),
    "ptc" boolean,
    "vin_code" varchar(50) not null,
    "certificate" varchar(50),
    "description" text,
    "can_look_coordinate" varchar(50),
    "phone_number" varchar(50) not null,
    "refuse_dealers_calls" boolean,
    "only_chat" boolean,
    "protect_spam" boolean,
    "verified_buyers" boolean,
    "contact_person" varchar(50), -- email or user_id
    "email" varchar(50),
    "price" int not null,
    "price_type" price_type_enum not null default 'USD',
    "status" int not null default 3, -- 1-pending, 2-not sale (my cars), 3-on sale,
    "updated_at" timestamp not null default now(),
    "created_at" timestamp not null default now(),
    constraint motorcycles_user_id_fk
        foreign key (user_id)
            references users(id)
                on delete cascade
                on update cascade,
    constraint motorcycles_category_id_fk
        foreign key (moto_category_id)
            references moto_categories(id)
                on delete cascade
                on update cascade,
    constraint motorcycles_brand_id_fk
        foreign key (moto_brand_id)
            references moto_brands(id)
                on delete cascade
                on update cascade,
    constraint motorcycles_model_id_fk
        foreign key (moto_model_id)
            references moto_models(id)
                on delete cascade
                on update cascade,
    constraint motorcycles_fuel_type_id_fk
        foreign key (fuel_type_id)
            references fuel_types(id)
                on delete cascade
                on update cascade,
    constraint motorcycles_color_id_fk
        foreign key (color_id)
            references colors(id)
                on delete cascade
                on update cascade,
    constraint motorcycles_city_id_fk
        foreign key (city_id)
            references cities(id)
                on delete cascade
                on update cascade
);


create table motorcycle_parameters (
    "id" serial primary key,
    "motorcycle_id" int not null,
    "moto_parameter_id" int not null,
    "moto_parameter_value_id" int not null,
    "created_at" timestamp default now(),
    constraint motorcycle_parameters_motorcycle_id_fk
        foreign key (motorcycle_id)
            references motorcycles(id)
                on delete cascade
                on update cascade,
    constraint motorcycle_parameters_moto_parameter_id_fk
        foreign key (moto_parameter_id)
            references moto_parameters(id)
                on delete cascade
                on update cascade,
    constraint motorcycle_parameters_moto_parameter_value_id_fk
        foreign key (moto_parameter_value_id)
            references moto_parameter_values(id)
                on delete cascade
                on update cascade,
    unique("motorcycle_id", "moto_parameter_id")
);

create table moto_images (
    "id" serial primary key,
    "moto_id" int not null,
    "image" varchar(255) not null,
    "created_at" timestamp not null default now(),
    constraint moto_images_moto_id_fk
        foreign key (moto_id)
            references motorcycles(id)
                on delete cascade
                on update cascade
);


create table moto_videos (
    "id" serial primary key,
    "moto_id" int not null,
    "video" varchar(255) not null,
    "created_at" timestamp not null default now(),
    constraint moto_videos_moto_id_fk
        foreign key (moto_id)
            references motorcycles(id)
                on delete cascade
                on update cascade
);

drop table if exists comtran_videos;
drop table if exists comtran_images;
drop table if exists comtran_parameters;
drop table if exists com_category_parameters;
drop table if exists com_parameter_values;
drop table if exists comtrans;
drop table if exists com_models;
drop table if exists com_brands;
drop table if exists com_parameters;
drop table if exists com_categories;

create table com_categories (
    "id" serial primary key,
    "name" varchar(100) not null,
    "created_at" timestamp not null default now()
);

create table com_brands (
    "id" serial primary key,
    "name" varchar(100) not null,
    "image" varchar(255) not null,
    "comtran_category_id" integer not null,
    "created_at" timestamp not null default now(),
    constraint com_brands_comtran_category_id_fk
        foreign key (comtran_category_id)
            references com_categories(id)
                on delete cascade
                on update cascade
);

create table com_models (
    "id" serial primary key,
    "name" varchar(100) not null,
    "comtran_brand_id" integer not null,
    "created_at" timestamp not null default now(),
    constraint com_brand_models_comtran_brand_id_fk
        foreign key (comtran_brand_id)
            references com_brands(id)
                on delete cascade
                on update cascade
);

create table com_parameters (
    "id" serial primary key,
    "comtran_category_id" int not null,
    "name" varchar(100) not null,
    "created_at" timestamp default now(),
    constraint com_parameters_comtran_category_id_fk
        foreign key (comtran_category_id)
            references com_categories(id)
                on delete cascade
                on update cascade,
    unique("name", "comtran_category_id")
);


create table com_parameter_values (
    "id" serial primary key,
    "comtran_parameter_id" int not null,
    "name" varchar(100) not null,
    "created_at" timestamp default now(),
    constraint com_parameter_values_comtran_parameter_id_fk
        foreign key (comtran_parameter_id)
            references com_parameters(id)
                on delete cascade
                on update cascade
);

create table com_category_parameters (
    "comtran_category_id" int not null,
    "comtran_parameter_id" int not null,
    "created_at" timestamp not null default now(),
    constraint com_category_parameters_comtran_category_id_fk
        foreign key (comtran_category_id)
            references com_categories(id)
                on delete cascade
                on update cascade,
    constraint com_category_parameters_comtran_parameter_id_fk
        foreign key (comtran_parameter_id)
            references com_parameters(id)
                on delete cascade
                on update cascade
);

create table comtrans (
    "id" serial primary key,
    "user_id" int not null,
    "comtran_category_id" int not null,
    "comtran_brand_id" int not null,
    "comtran_model_id" int not null,
    "fuel_type_id" int not null,
    "city_id" int not null,
    "color_id" int not null,
    "engine" int, -- cm3
    "power" int, -- hp
    "year" int not null,
    "number_of_cycles" int,
    "odometer" int not null default 0,
    "crash" boolean,
    "not_cleared" boolean,
    "owners" int,
    "date_of_purchase" varchar (50),
    "warranty_date" varchar(50),
    "ptc" boolean,
    "vin_code" varchar(50) not null,
    "certificate" varchar(50),
    "description" text,
    "can_look_coordinate" varchar(50),
    "phone_number" varchar(50) not null,
    "refuse_dealers_calls" boolean,
    "only_chat" boolean,
    "protect_spam" boolean,
    "verified_buyers" boolean,
    "contact_person" varchar(50), -- email or user_id
    "email" varchar(50),
    "price" int not null,
    "price_type" price_type_enum not null default 'USD',
    "status" int not null default 3, -- 1-pending, 2-not sale (my cars), 3-on sale,
    "updated_at" timestamp not null default now(),
    "created_at" timestamp not null default now(),
    constraint comtrans_user_id_fk
        foreign key (user_id)
            references users(id)
                on delete cascade
                on update cascade,
    constraint comtrans_category_id_fk
        foreign key (comtran_category_id)
            references com_categories(id)
                on delete cascade
                on update cascade,
    constraint comtrans_brand_id_fk
        foreign key (comtran_brand_id)
            references com_brands(id)
                on delete cascade
                on update cascade,
    constraint comtrans_model_id_fk
        foreign key (comtran_model_id)
            references com_models(id)
                on delete cascade
                on update cascade,
    constraint comtrans_fuel_type_id_fk
        foreign key (fuel_type_id)
            references fuel_types(id)
                on delete cascade
                on update cascade,
    constraint comtrans_color_id_fk
        foreign key (color_id)
            references colors(id)
                on delete cascade
                on update cascade,
    constraint comtrans_city_id_fk
        foreign key (city_id)
            references cities(id)
                on delete cascade
                on update cascade
);


create table comtran_parameters (
    "id" serial primary key,
    "comtran_id" int not null,
    "comtran_parameter_id" int not null,
    "comtran_parameter_value_id" int not null,
    "created_at" timestamp default now(),
    constraint comtran_parameters_comtran_id_fk
        foreign key (comtran_id)
            references comtrans(id)
                on delete cascade
                on update cascade,
    constraint comtran_parameters_comtran_parameter_id_fk
        foreign key (comtran_parameter_id)
            references com_parameters(id)
                on delete cascade
                on update cascade,
    constraint comtran_parameters_comtran_parameter_value_id_fk
        foreign key (comtran_parameter_value_id)
            references com_parameter_values(id)
                on delete cascade
                on update cascade,
    unique("comtran_id", "comtran_parameter_id")
);


create table comtran_images (
    "id" serial primary key,
    "comtran_id" int not null,
    "image" varchar(255) not null,
    "created_at" timestamp not null default now(),
    constraint comtran_images_comtran_id_fk
        foreign key (comtran_id)
            references comtrans(id)
                on delete cascade
                on update cascade
);


create table comtran_videos (
    "id" serial primary key,
    "comtran_id" int not null,
    "video" varchar(255) not null,
    "created_at" timestamp not null default now(),
    constraint comtran_videos_comtran_id_fk
        foreign key (comtran_id)
            references comtrans(id)
                on delete cascade
                on update cascade
);
