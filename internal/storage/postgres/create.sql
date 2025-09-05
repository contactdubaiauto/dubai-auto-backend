-- create user da with password '1234';
-- grant all privileges on database da to da;
-- grant all privileges on schema public to da;
-- grant all privileges on all tables in schema public to da;
-- grant all privileges on all sequences in schema public to da;
-- alter default privileges in schema public grant all on tables to da;
-- alter default privileges in schema public grant all on sequences to da;





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
    "registered_by" varchar(20) not null,
    "updated_at" timestamp default now(),
    "created_at" timestamp default now(),
    unique("email"),
    unique("phone")
);

create table users (
    "id" serial primary key,
    "email" varchar(100),
    "role_id" int not null default 1, -- 1 user, 2 dealer, 3 logistic, 4 broker
    "password" varchar(100) not null,
    "phone" varchar(100),
    "updated_at" timestamp default now(),
    "created_at" timestamp default now(),
    unique("email"),
    unique("phone")
);

create table messages (
    "id" serial primary key,
    "sender_id" int not null,
    "receiver_id" int not null,
    "message" text not null,
    "type" int not null default 1, -- 1-text, 2-image, 3-video, 4-audio, 5-file, 6-item
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

insert into users (email, password, phone, created_at) 
    values ('user@gmail.com', '$2a$10$Cya9x0xSJSnRknBmJpW.Bu8ukZpVTqzwgrQgAYNPXdrX2HYGRk33W', '01234567890', now()); -- password: 12345678

insert into users (email, password, phone, created_at) 
    values ('user2@gmail.com', '$2a$10$Cya9x0xSJSnRknBmJpW.Bu8ukZpVTqzwgrQgAYNPXdrX2HYGRk33W', '0111222222', now()); -- password: 12345678


create table cities (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now()
);

insert into cities (name) values ('Dubai');
insert into cities (name) values ('Abu Dhabi');
insert into cities (name) values ('Sharjah');


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

insert into profiles(
    user_id, username, driving_experience, notification, google, birthday, about_me, registered_by
)values
( 1, 'user1', 3, false, 'user1@gmail.com', '2025-04-14', 'im a f1 driver with 3 years of experiences', 'email'),
( 2, 'user2', 2, true, 'user2@gmail.com', '2025-04-14', 'im a truck driver with 2 years of experiences', 'email');

create table admins (
    "id" serial primary key,
    "username" varchar(255) not null,
    "email" varchar(255) not null,
    "password" varchar(255) not null,
    "last_active_date" timestamp default now(),
    "created_at" timestamp default now()
);

insert into admins (username, email, password, last_active_date, created_at) 
    values ('admin', 'admin@gmail.com', '$2a$10$wPb6//DXtLxpDjZgzVEMuOlqqHUtVmPQMbOhmBlkNlAkzzve..CZe', now(), now()); -- password: admin

create table brands (
    "id" serial primary key,
    "name" varchar(255) not null,
    "logo" varchar(255) not null,
    "model_count" int not null default 0,
    "popular" boolean default false,
    "updated_at" timestamp default now(),
    unique("name")
);

-- insert into brands (name, logo, car_count, popular) values ('Toyota', '/images/logo/toyota.png', 12, true);
-- insert into brands (name, logo, car_count, popular) values ('Honda', '/images/logo/honda.png', 8, false);

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

-- toyota
-- insert into models (name, brand_id, popular, car_count) values ('Camry', 1, true, 7);
-- insert into models (name, brand_id, popular, car_count) values ('Corolla', 1, true, 41);
-- insert into models (name, brand_id, popular, car_count) values ('Rav4', 1, false, 73);
-- insert into models (name, brand_id, popular, car_count) values ('Land Cruiser', 1, false, 1);

-- -- honda
-- insert into models (name, brand_id, popular, car_count) values ('Civic', 2, true, 34);
-- insert into models (name, brand_id, popular, car_count) values ('Accord', 2, false, 23);
-- insert into models (name, brand_id, popular, car_count) values ('CR-V', 2, false, 65);



create table body_types (
    "id" serial primary key,
    "name" varchar(255) not null,
    "image" character varying(255) not null,
    "created_at" timestamp default now()
);

insert into body_types (name, image) values ('Sedan','/images/body/sedan.png');
insert into body_types (name, image) values ('Hatchback', '/images/body/hatchback.png');
insert into body_types (name, image) values ('Liftback', '/images/body/liftback.png');
insert into body_types (name, image) values ('SUV', '/images/body/suv.png');
insert into body_types (name, image) values ('Crossover', '/images/body/crossover.png');
insert into body_types (name, image) values ('Coupe', '/images/body/coupe.png');
insert into body_types (name, image) values ('Convertible', '/images/body/convertible.png');
insert into body_types (name, image) values ('Wagon', '/images/body/wagon.png');
insert into body_types (name, image) values ('Pickup Truck', '/images/body/pickup.png');
insert into body_types (name, image) values ('Van', '/images/body/van.png');
insert into body_types (name, image) values ('Minivan', '/images/body/minivan.png');
insert into body_types (name, image) values ('Roadster', '/images/body/roadster.png');
insert into body_types (name, image) values ('Sports Car', '/images/body/sports_car.png');
insert into body_types (name, image) values ('Off-Road', '/images/body/off_road.png');
insert into body_types (name, image) values ('Limousine', '/images/body/limousine.png'); 
insert into body_types (name, image) values ('Utility', '/images/body/utility.png');
insert into body_types (name, image) values ('Universal', '/images/body/universal.png');
insert into body_types (name, image) values ('Cabriolet', '/images/body/cabri.webp');


create table transmissions (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now(),
    unique("name")
);

insert into transmissions (name) values ('Automatic');
insert into transmissions (name) values ('Manual');
insert into transmissions (name) values ('Semi-Automatic');


create table engines (
    "id" serial primary key,
    "value" varchar(255) not null,
    "created_at" timestamp default now(),
    unique("value")
);

insert into engines (value) values ('1.0L');
insert into engines (value) values ('1.5L');
insert into engines (value) values ('2.0L');
insert into engines (value) values ('2.5L');
insert into engines (value) values ('3.0L');
insert into engines (value) values ('4.0L');
insert into engines (value) values ('5.0L');
insert into engines (value) values ('6.0L');
insert into engines (value) values ('7.0L');


create table drivetrains (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now(),
    unique("name")
);

insert into drivetrains (name) values ('Front-Wheel Drive');
insert into drivetrains (name) values ('Rear-Wheel Drive');
insert into drivetrains (name) values ('All-Wheel Drive');


create table fuel_types (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now(),
    unique("name")
);

insert into fuel_types (name) values ('Gasoline');
insert into fuel_types (name) values ('Diesel');
insert into fuel_types (name) values ('Electric');
insert into fuel_types (name) values ('Hybrid');

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

-- dubai
insert into regions (name, city_id) values ('Dubai Marina', 1);
insert into regions (name, city_id) values ('Dubai Mall', 1);
insert into regions (name, city_id) values ('Dubai Creek', 1);
insert into regions (name, city_id) values ('Dubai Creek Harbour', 1);
insert into regions (name, city_id) values ('Dubai Creek Golf Club', 1);

-- abu dhabi
insert into regions (name, city_id) values ('Abu Dhabi Marina', 2);
insert into regions (name, city_id) values ('Abu Dhabi Mall', 2);
insert into regions (name, city_id) values ('Abu Dhabi Creek', 2);
insert into regions (name, city_id) values ('Abu Dhabi Creek Harbour', 2);
insert into regions (name, city_id) values ('Abu Dhabi Creek Golf Club', 2);

-- sharjah
insert into regions (name, city_id) values ('Sharjah Marina', 3);
insert into regions (name, city_id) values ('Sharjah Mall', 3);
insert into regions (name, city_id) values ('Sharjah Creek', 3);

create table service_types (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now()
);

insert into service_types (name) values ('Car Wash');
insert into service_types (name) values ('Car Detailing');
insert into service_types (name) values ('Car Repair');

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

-- car wash
insert into services (name, service_type_id) values ('Rocket Wash', 1);
insert into services (name, service_type_id) values ('Premium Wash', 1);
insert into services (name, service_type_id) values ('Express Wash', 1);
insert into services (name, service_type_id) values ('Self-Service Wash', 1);

-- car detailing
insert into services (name, service_type_id) values ('Full Detail', 2);
insert into services (name, service_type_id) values ('Basic Detail', 2);
insert into services (name, service_type_id) values ('Premium Detail', 2);
insert into services (name, service_type_id) values ('Express Detail', 2);
insert into services (name, service_type_id) values ('Self-Service Detail', 2);

-- car repair
insert into services (name, service_type_id) values ('Oil Change', 3);
insert into services (name, service_type_id) values ('Brake Repair', 3);
insert into services (name, service_type_id) values ('Tire Repair', 3);
insert into services (name, service_type_id) values ('Engine Repair', 3);
insert into services (name, service_type_id) values ('Transmission Repair', 3);
insert into services (name, service_type_id) values ('Suspension Repair', 3);
insert into services (name, service_type_id) values ('Electrical Repair', 3);



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

-- insert into generations (
--     name, model_id, start_year, end_year, image, wheel
-- ) values (
--     '1 generation', 1, 2010, 2020, '/images/gens/1.jpg', true
-- );

-- insert into generations (
--     name, model_id, start_year, end_year, image, wheel
-- ) values (
--     '2 generation', 1, 2005, 2025, '/images/gens/2.jpg', true
-- );


-- insert into generations (
--     name, model_id, start_year, end_year, image, wheel
-- ) values (
--     '2 generation', 1, 2005, 2025, '/images/gens/2.jpg', false
-- );

-- insert into generations (
--     name, model_id, start_year, end_year, image, wheel
-- ) values (
--     '3 generation', 2, 2020, 2025, '/images/gens/3.jpg', true
-- );

-- insert into generations (
--     name, model_id, start_year, end_year, image, wheel
-- ) values (
--     '3 generation', 2, 2020, 2025, '/images/gens/3.jpg', false
-- );

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

-- insert into generation_modifications (
--     generation_id, engine_id, fuel_type_id, drivetrain_id, transmission_id
-- ) 
-- values 
--     (1, 1, 1, 1, 1),
--     (1, 3, 2, 2, 2),
--     (2, 1, 1, 1, 1),
--     (2, 3, 2, 2, 2),
--     (1, 2, 2, 2, 2);


create table ownership_types (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now()
);

insert into ownership_types (name) values ('Dealership');
insert into ownership_types (name) values ('Private Owner');

create table colors (
    "id" serial primary key,
    "name" varchar(255) not null,
    "image" varchar(255) not null,
    "created_at" timestamp default now()
);

insert into colors (name, image) values ('White', '/images/colors/white.jpg');
insert into colors (name, image) values ('Red', '/images/colors/red.jpg');
insert into colors (name, image) values ('Blue', '/images/colors/blue.jpg');
insert into colors (name, image) values ('Green', '/images/colors/green.jpg');
insert into colors (name, image) values ('Yellow', '/images/colors/yellow.jpg');
insert into colors (name, image) values ('Orange', '/images/colors/orange.jpg');
insert into colors (name, image) values ('Purple', '/images/colors/purple.jpg');


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

-- Insert example vehicle data
insert into vehicles (user_id, brand_id, model_id, modification_id, region_id, city_id, ownership_type_id, year, odometer, vin_code, crash, owners, credit, phone_numbers, price, new, color_id, trade_in, status) values 
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

insert into moto_categories (name) values ('Moto');
insert into moto_categories (name) values ('Skuter');
insert into moto_categories (name) values ('Motovezdehody');
insert into moto_categories (name) values ('Snegohody');

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

insert into moto_brands (name, moto_category_id, image) values ('Honda', 1, '/images/logo/honda.png');
insert into moto_brands (name, moto_category_id, image) values ('Yamaha', 1, '/images/logo/yamaha.png');
insert into moto_brands (name, moto_category_id, image) values ('Suzuki', 2, '/images/logo/suzuki.png');
insert into moto_brands (name, moto_category_id, image) values ('Kawasaki', 2, '/images/logo/kawasaki.png');
insert into moto_brands (name, moto_category_id, image) values ('BMW', 3, '/images/logo/bmw.png');
insert into moto_brands (name, moto_category_id, image) values ('Ducati', 3, '/images/logo/ducati.png');
insert into moto_brands (name, moto_category_id, image) values ('BMW', 4, '/images/logo/bmw.png');
insert into moto_brands (name, moto_category_id, image) values ('Ducati', 4, '/images/logo/ducati.png');

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

insert into moto_models (name, moto_brand_id) values ('CBR1000RR', 1);
insert into moto_models (name, moto_brand_id) values ('CBR650R', 1);

insert into moto_models (name, moto_brand_id) values ('CBR1000RR', 2);
insert into moto_models (name, moto_brand_id) values ('CBR650R', 2);

insert into moto_models (name, moto_brand_id) values ('CBR1000RR', 3);
insert into moto_models (name, moto_brand_id) values ('CBR650R', 3);

insert into moto_models (name, moto_brand_id) values ('CBR1000RR', 4);
insert into moto_models (name, moto_brand_id) values ('CBR650R', 4);

insert into moto_models (name, moto_brand_id) values ('CBR1000RR', 5);
insert into moto_models (name, moto_brand_id) values ('CBR650R', 5);

insert into moto_models (name, moto_brand_id) values ('CBR1000RR', 6);
insert into moto_models (name, moto_brand_id) values ('CBR650R', 6);


insert into moto_models (name, moto_brand_id) values ('CBR1000RR', 7);
insert into moto_models (name, moto_brand_id) values ('CBR650R', 7);


insert into moto_models (name, moto_brand_id) values ('CBR1000RR', 8);
insert into moto_models (name, moto_brand_id) values ('CBR650R', 8);



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

insert into moto_parameters (name, moto_category_id) values ('Type', 1);
insert into moto_parameters (name, moto_category_id) values ('Drivetrain', 1);
insert into moto_parameters (name, moto_category_id) values ('Transmission', 1);
insert into moto_parameters (name, moto_category_id) values ('Cylinder Count', 1);
insert into moto_parameters (name, moto_category_id) values ('Cylinder Arrangement', 1);
insert into moto_parameters (name, moto_category_id) values ('Equipment', 1);

insert into moto_parameters (name, moto_category_id) values ('Type', 2);
insert into moto_parameters (name, moto_category_id) values ('Drivetrain', 2);
insert into moto_parameters (name, moto_category_id) values ('Transmission', 2);
insert into moto_parameters (name, moto_category_id) values ('Cylinder Count', 2);
insert into moto_parameters (name, moto_category_id) values ('Cylinder Arrangement', 2);
insert into moto_parameters (name, moto_category_id) values ('Equipment', 2);

insert into moto_parameters (name, moto_category_id) values ('Type', 3);
insert into moto_parameters (name, moto_category_id) values ('Drivetrain', 3);
insert into moto_parameters (name, moto_category_id) values ('Transmission', 3);
insert into moto_parameters (name, moto_category_id) values ('Cylinder Count', 3);
insert into moto_parameters (name, moto_category_id) values ('Cylinder Arrangement', 3);
insert into moto_parameters (name, moto_category_id) values ('Equipment', 3);

insert into moto_parameters (name, moto_category_id) values ('Type', 4);
insert into moto_parameters (name, moto_category_id) values ('Drivetrain', 4);
insert into moto_parameters (name, moto_category_id) values ('Transmission', 4);
insert into moto_parameters (name, moto_category_id) values ('Cylinder Count', 4);
insert into moto_parameters (name, moto_category_id) values ('Cylinder Arrangement', 4);
insert into moto_parameters (name, moto_category_id) values ('Equipment', 4);


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

insert into moto_parameter_values (moto_parameter_id, name) values (1, 'Sport');
insert into moto_parameter_values (moto_parameter_id, name) values (1, 'Touring');
insert into moto_parameter_values (moto_parameter_id, name) values (1, 'Cruiser');
insert into moto_parameter_values (moto_parameter_id, name) values (1, 'Sport Touring');
insert into moto_parameter_values (moto_parameter_id, name) values (1, 'Sport Tourer');

insert into moto_parameter_values (moto_parameter_id, name) values (2, 'Front');
insert into moto_parameter_values (moto_parameter_id, name) values (2, 'Rear');
insert into moto_parameter_values (moto_parameter_id, name) values (2, 'All');

insert into moto_parameter_values (moto_parameter_id, name) values (3, 'Automatic');
insert into moto_parameter_values (moto_parameter_id, name) values (3, 'Manual');
insert into moto_parameter_values (moto_parameter_id, name) values (3, 'Semi-Automatic');

insert into moto_parameter_values (moto_parameter_id, name) values (4, '4');
insert into moto_parameter_values (moto_parameter_id, name) values (4, '6');
insert into moto_parameter_values (moto_parameter_id, name) values (4, '8');

insert into moto_parameter_values (moto_parameter_id, name) values (5, 'V-Twin');
insert into moto_parameter_values (moto_parameter_id, name) values (5, 'In-Twin');

insert into moto_parameter_values (moto_parameter_id, name) values (6, 'Single');
insert into moto_parameter_values (moto_parameter_id, name) values (6, 'Twin');
insert into moto_parameter_values (moto_parameter_id, name) values (6, 'Twin-Cooled');


insert into moto_parameter_values (moto_parameter_id, name) values (7, 'Sport');
insert into moto_parameter_values (moto_parameter_id, name) values (7, 'Touring');
insert into moto_parameter_values (moto_parameter_id, name) values (7, 'Enduro');
insert into moto_parameter_values (moto_parameter_id, name) values (7, 'Cross');
insert into moto_parameter_values (moto_parameter_id, name) values (7, 'Enduro Cross');

insert into moto_parameter_values (moto_parameter_id, name) values (8, 'Sport');
insert into moto_parameter_values (moto_parameter_id, name) values (8, 'Touring');
insert into moto_parameter_values (moto_parameter_id, name) values (8, 'Enduro');
insert into moto_parameter_values (moto_parameter_id, name) values (8, 'Cross');
insert into moto_parameter_values (moto_parameter_id, name) values (8, 'Enduro Cross');

insert into moto_parameter_values (moto_parameter_id, name) values (9, 'Sport');
insert into moto_parameter_values (moto_parameter_id, name) values (9, 'Touring');
insert into moto_parameter_values (moto_parameter_id, name) values (9, 'Enduro');
insert into moto_parameter_values (moto_parameter_id, name) values (9, 'Cross');
insert into moto_parameter_values (moto_parameter_id, name) values (9, 'Enduro Cross');

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

insert into moto_category_parameters (moto_category_id, moto_parameter_id) values (1, 1);
insert into moto_category_parameters (moto_category_id, moto_parameter_id) values (1, 2);
insert into moto_category_parameters (moto_category_id, moto_parameter_id) values (1, 3);
insert into moto_category_parameters (moto_category_id, moto_parameter_id) values (1, 5);

insert into moto_category_parameters (moto_category_id, moto_parameter_id) values (2, 3);
insert into moto_category_parameters (moto_category_id, moto_parameter_id) values (2, 4);
insert into moto_category_parameters (moto_category_id, moto_parameter_id) values (2, 5);
insert into moto_category_parameters (moto_category_id, moto_parameter_id) values (2, 6);

insert into moto_category_parameters (moto_category_id, moto_parameter_id) values (3, 1);
insert into moto_category_parameters (moto_category_id, moto_parameter_id) values (3, 6);

insert into moto_category_parameters (moto_category_id, moto_parameter_id) values (4, 1);
insert into moto_category_parameters (moto_category_id, moto_parameter_id) values (4, 4);
insert into moto_category_parameters (moto_category_id, moto_parameter_id) values (4, 5);



-- Create enum type for price currency
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

insert into motorcycles (
    user_id, moto_category_id, moto_brand_id, moto_model_id, fuel_type_id, 
    city_id, color_id, engine, power, year, number_of_cycles, odometer, 
    crash, not_cleared, owners, date_of_purchase, warranty_date, ptc, vin_code, 
    certificate, description, can_look_coordinate, phone_number, refuse_dealers_calls, 
    only_chat, protect_spam, verified_buyers, contact_person, email, price, price_type
) values (
    1, 1, 1, 1, 1,
    1, 1, 1, 1, 2020, 1, 10000,
    true, false, 1, '2020-01-01', '2020-01-01', false, '1234567890',
    true, 'description', '1234567890', '1234567890', true,
    true, true, true, 'John Doe', 'john.doe@example.com', 10000, 'USD'
);

insert into motorcycles (
    user_id, moto_category_id, moto_brand_id, moto_model_id, fuel_type_id, 
    city_id, color_id, engine, power, year, number_of_cycles, odometer, 
    crash, not_cleared, owners, date_of_purchase, warranty_date, ptc, vin_code, 
    certificate, description, can_look_coordinate, phone_number, refuse_dealers_calls, 
    only_chat, protect_spam, verified_buyers, contact_person, email, price, price_type
) values (
    1, 1, 1, 1, 1,
    1, 1, 1, 1, 2020, 1, 10000,
    true, false, 1, '2020-01-01', '2020-01-01', false, '1234567890',
    true, 'description', '1234567890', '1234567890', true,
    true, true, true, 'John Doe', 'john.doe@example.com', 10000, 'USD'
);

insert into motorcycles (
    user_id, moto_category_id, moto_brand_id, moto_model_id, fuel_type_id, 
    city_id, color_id, engine, power, year, number_of_cycles, odometer, 
    crash, not_cleared, owners, date_of_purchase, warranty_date, ptc, vin_code, 
    certificate, description, can_look_coordinate, phone_number, refuse_dealers_calls, 
    only_chat, protect_spam, verified_buyers, contact_person, email, price, price_type
) values (
    1, 1, 1, 1, 1,
    1, 1, 1, 1, 2020, 1, 10000,
    true, false, 1, '2020-01-01', '2020-01-01', false, '1234567890',
    true, 'description', '1234567890', '1234567890', true,
    true, true, true, 'John Doe', 'john.doe@example.com', 10000, 'USD'
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

insert into motorcycle_parameters (motorcycle_id, moto_parameter_id, moto_parameter_value_id) values (1, 1, 1);
insert into motorcycle_parameters (motorcycle_id, moto_parameter_id, moto_parameter_value_id) values (1, 2, 1);
insert into motorcycle_parameters (motorcycle_id, moto_parameter_id, moto_parameter_value_id) values (1, 3, 1);
insert into motorcycle_parameters (motorcycle_id, moto_parameter_id, moto_parameter_value_id) values (1, 4, 1);
insert into motorcycle_parameters (motorcycle_id, moto_parameter_id, moto_parameter_value_id) values (1, 5, 1);

insert into motorcycle_parameters (motorcycle_id, moto_parameter_id, moto_parameter_value_id) values (2, 1, 1);
insert into motorcycle_parameters (motorcycle_id, moto_parameter_id, moto_parameter_value_id) values (2, 2, 1);
insert into motorcycle_parameters (motorcycle_id, moto_parameter_id, moto_parameter_value_id) values (2, 3, 1);
insert into motorcycle_parameters (motorcycle_id, moto_parameter_id, moto_parameter_value_id) values (2, 4, 1);
insert into motorcycle_parameters (motorcycle_id, moto_parameter_id, moto_parameter_value_id) values (2, 5, 1);

insert into motorcycle_parameters (motorcycle_id, moto_parameter_id, moto_parameter_value_id) values (3, 1, 1);
insert into motorcycle_parameters (motorcycle_id, moto_parameter_id, moto_parameter_value_id) values (3, 2, 1);
insert into motorcycle_parameters (motorcycle_id, moto_parameter_id, moto_parameter_value_id) values (3, 3, 1);
insert into motorcycle_parameters (motorcycle_id, moto_parameter_id, moto_parameter_value_id) values (3, 4, 1);
insert into motorcycle_parameters (motorcycle_id, moto_parameter_id, moto_parameter_value_id) values (3, 5, 1);


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

insert into moto_images (moto_id, image) values (1, 'https://via.placeholder.com/150');
insert into moto_images (moto_id, image) values (1, 'https://via.placeholder.com/150');
insert into moto_images (moto_id, image) values (1, 'https://via.placeholder.com/150');
insert into moto_images (moto_id, image) values (1, 'https://via.placeholder.com/150');

insert into moto_images (moto_id, image) values (2, 'https://via.placeholder.com/150');
insert into moto_images (moto_id, image) values (2, 'https://via.placeholder.com/150');

insert into moto_images (moto_id, image) values (3, 'https://via.placeholder.com/150');
insert into moto_images (moto_id, image) values (3, 'https://via.placeholder.com/150');


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

insert into moto_videos (moto_id, video) values (1, 'https://www.youtube.com/watch?v=dQw4w9WgXcQ');
insert into moto_videos (moto_id, video) values (1, 'https://www.youtube.com/watch?v=dQw4w9WgXcQ');

insert into moto_videos (moto_id, video) values (2, 'https://www.youtube.com/watch?v=dQw4w9WgXcQ');
insert into moto_videos (moto_id, video) values (2, 'https://www.youtube.com/watch?v=dQw4w9WgXcQ');

insert into moto_videos (moto_id, video) values (3, 'https://www.youtube.com/watch?v=dQw4w9WgXcQ');
insert into moto_videos (moto_id, video) values (3, 'https://www.youtube.com/watch?v=dQw4w9WgXcQ');


--  comtrans

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

insert into com_categories (name) values ('Truck');
insert into com_categories (name) values ('Bus');
insert into com_categories (name) values ('Trailer');
insert into com_categories (name) values ('Other');

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

insert into com_brands (name, image, comtran_category_id) values ('Volvo', 'https://via.placeholder.com/150', 1);
insert into com_brands (name, image, comtran_category_id) values ('Mercedes-Benz', 'https://via.placeholder.com/150', 1);
insert into com_brands (name, image, comtran_category_id) values ('MAN', 'https://via.placeholder.com/150', 1);
insert into com_brands (name, image, comtran_category_id) values ('Scania', 'https://via.placeholder.com/150', 1);
insert into com_brands (name, image, comtran_category_id) values ('Iveco', 'https://via.placeholder.com/150', 1);

insert into com_brands (name, image, comtran_category_id) values ('Volvo', 'https://via.placeholder.com/150', 2);
insert into com_brands (name, image, comtran_category_id) values ('Mercedes-Benz', 'https://via.placeholder.com/150', 2);
insert into com_brands (name, image, comtran_category_id) values ('MAN', 'https://via.placeholder.com/150', 2);
insert into com_brands (name, image, comtran_category_id) values ('Scania', 'https://via.placeholder.com/150', 2);
insert into com_brands (name, image, comtran_category_id) values ('Iveco', 'https://via.placeholder.com/150', 2);

insert into com_brands (name, image, comtran_category_id) values ('Volvo', 'https://via.placeholder.com/150', 3);
insert into com_brands (name, image, comtran_category_id) values ('Mercedes-Benz', 'https://via.placeholder.com/150', 3);
insert into com_brands (name, image, comtran_category_id) values ('MAN', 'https://via.placeholder.com/150', 3);
insert into com_brands (name, image, comtran_category_id) values ('Scania', 'https://via.placeholder.com/150', 3);
insert into com_brands (name, image, comtran_category_id) values ('Iveco', 'https://via.placeholder.com/150', 3);

insert into com_brands (name, image, comtran_category_id) values ('Volvo', 'https://via.placeholder.com/150', 4);
insert into com_brands (name, image, comtran_category_id) values ('Mercedes-Benz', 'https://via.placeholder.com/150', 4);
insert into com_brands (name, image, comtran_category_id) values ('MAN', 'https://via.placeholder.com/150', 4);
insert into com_brands (name, image, comtran_category_id) values ('Scania', 'https://via.placeholder.com/150', 4);
insert into com_brands (name, image, comtran_category_id) values ('Iveco', 'https://via.placeholder.com/150', 4);


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

insert into com_models (name, comtran_brand_id) values ('FH', 1);
insert into com_models (name, comtran_brand_id) values ('FM', 1);

insert into com_models (name, comtran_brand_id) values ('FH', 2);
insert into com_models (name, comtran_brand_id) values ('FM', 2);

insert into com_models (name, comtran_brand_id) values ('FH', 3);
insert into com_models (name, comtran_brand_id) values ('FM', 3);

insert into com_models (name, comtran_brand_id) values ('FH', 4);
insert into com_models (name, comtran_brand_id) values ('FM', 4);


insert into com_models (name, comtran_brand_id) values ('FH', 5);
insert into com_models (name, comtran_brand_id) values ('FM', 5);


insert into com_models (name, comtran_brand_id) values ('FH', 6);
insert into com_models (name, comtran_brand_id) values ('FM', 6);

insert into com_models (name, comtran_brand_id) values ('FH', 7);
insert into com_models (name, comtran_brand_id) values ('FM', 7);

insert into com_models (name, comtran_brand_id) values ('FH', 8);
insert into com_models (name, comtran_brand_id) values ('FM', 8);


insert into com_models (name, comtran_brand_id) values ('FH', 9);
insert into com_models (name, comtran_brand_id) values ('FM', 9);

insert into com_models (name, comtran_brand_id) values ('FH', 10);
insert into com_models (name, comtran_brand_id) values ('FM', 10);

insert into com_models (name, comtran_brand_id) values ('FH', 11);
insert into com_models (name, comtran_brand_id) values ('FM', 11);

insert into com_models (name, comtran_brand_id) values ('FH', 12);
insert into com_models (name, comtran_brand_id) values ('FM', 12);

insert into com_models (name, comtran_brand_id) values ('FH', 13);
insert into com_models (name, comtran_brand_id) values ('FM', 13);

insert into com_models (name, comtran_brand_id) values ('FH', 14);
insert into com_models (name, comtran_brand_id) values ('FM', 14);

insert into com_models (name, comtran_brand_id) values ('FH', 15);

insert into com_models (name, comtran_brand_id) values ('FM', 15);
insert into com_models (name, comtran_brand_id) values ('FH', 16);
insert into com_models (name, comtran_brand_id) values ('FM', 16);

insert into com_models (name, comtran_brand_id) values ('FH', 17);
insert into com_models (name, comtran_brand_id) values ('FM', 17);

insert into com_models (name, comtran_brand_id) values ('FH', 18);
insert into com_models (name, comtran_brand_id) values ('FM', 18);

insert into com_models (name, comtran_brand_id) values ('FH', 19);
insert into com_models (name, comtran_brand_id) values ('FM', 19);

insert into com_models (name, comtran_brand_id) values ('FH', 20);
insert into com_models (name, comtran_brand_id) values ('FM', 20);


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

insert into com_parameters (name, comtran_category_id) values ('Engine', 1);

insert into com_parameters (name, comtran_category_id) values ('Engine', 2);
insert into com_parameters (name, comtran_category_id) values ('Power', 2);
insert into com_parameters (name, comtran_category_id) values ('Number of Cycles', 2);
insert into com_parameters (name, comtran_category_id) values ('Odometer', 2);

insert into com_parameters (name, comtran_category_id) values ('Engine', 3);
insert into com_parameters (name, comtran_category_id) values ('Crash', 3);

insert into com_parameters (name, comtran_category_id) values ('Engine', 4);
insert into com_parameters (name, comtran_category_id) values ('Number of Cycles', 4);
insert into com_parameters (name, comtran_category_id) values ('Crash', 4);


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

insert into com_parameter_values (comtran_parameter_id, name) values (1, '1000');
insert into com_parameter_values (comtran_parameter_id, name) values (2, '1000');
insert into com_parameter_values (comtran_parameter_id, name) values (3, '1000');
insert into com_parameter_values (comtran_parameter_id, name) values (3, '800');
insert into com_parameter_values (comtran_parameter_id, name) values (4, '1000');
insert into com_parameter_values (comtran_parameter_id, name) values (5, '1000');
insert into com_parameter_values (comtran_parameter_id, name) values (6, '1000');
insert into com_parameter_values (comtran_parameter_id, name) values (7, '1000');
insert into com_parameter_values (comtran_parameter_id, name) values (7, '800');
insert into com_parameter_values (comtran_parameter_id, name) values (8, '1000');
insert into com_parameter_values (comtran_parameter_id, name) values (8, '200');
insert into com_parameter_values (comtran_parameter_id, name) values (9, '1000');
insert into com_parameter_values (comtran_parameter_id, name) values (9, '800');
insert into com_parameter_values (comtran_parameter_id, name) values (10, '1000');
insert into com_parameter_values (comtran_parameter_id, name) values (10, '800');

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

insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (1, 1);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (1, 2);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (1, 3);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (1, 4);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (1, 5);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (1, 6);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (1, 7);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (1, 8);

insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (2, 1);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (2, 2);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (2, 3);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (2, 4);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (2, 5);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (2, 6);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (2, 7);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (2, 8);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (2, 9);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (2, 10);

insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (3, 1);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (3, 2);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (3, 3);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (3, 4);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (3, 5);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (3, 6);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (3, 7);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (3, 8);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (3, 9);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (3, 10);

insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (4, 1);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (4, 2);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (4, 3);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (4, 4);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (4, 5);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (4, 6);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (4, 7);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (4, 8);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (4, 9);
insert into com_category_parameters (comtran_category_id, comtran_parameter_id) values (4, 10);


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

insert into comtrans (user_id, comtran_category_id, comtran_brand_id, comtran_model_id, fuel_type_id, city_id, color_id, engine, power, year, number_of_cycles, odometer, crash, not_cleared, owners, date_of_purchase, warranty_date, ptc, vin_code, certificate, description, can_look_coordinate, phone_number, refuse_dealers_calls, only_chat, protect_spam, verified_buyers, contact_person, email, price, price_type, status) values (1, 1, 1, 1, 1, 1, 1, 1000, 1000, 2020, 100000, 0, false, false, 1, '2020-01-01', '2020-01-01', true, '1234567890', '1234567890', 'description', '1234567890', '1234567890', true, true, true, true, 'contact_person', 'email', 10000, 'USD', 1);
insert into comtrans (user_id, comtran_category_id, comtran_brand_id, comtran_model_id, fuel_type_id, city_id, color_id, engine, power, year, number_of_cycles, odometer, crash, not_cleared, owners, date_of_purchase, warranty_date, ptc, vin_code, certificate, description, can_look_coordinate, phone_number, refuse_dealers_calls, only_chat, protect_spam, verified_buyers, contact_person, email, price, price_type, status) values (1, 2, 1, 1, 1, 1, 1, 1000, 1000, 2020, 100000, 0, false, false, 1, '2020-01-01', '2020-01-01', true, '1234567890', '1234567890', 'description', '1234567890', '1234567890', true, true, true, true, 'contact_person', 'email', 10000, 'USD', 1);
insert into comtrans (user_id, comtran_category_id, comtran_brand_id, comtran_model_id, fuel_type_id, city_id, color_id, engine, power, year, number_of_cycles, odometer, crash, not_cleared, owners, date_of_purchase, warranty_date, ptc, vin_code, certificate, description, can_look_coordinate, phone_number, refuse_dealers_calls, only_chat, protect_spam, verified_buyers, contact_person, email, price, price_type, status) values (1, 3, 1, 1, 1, 1, 1, 1000, 1000, 2020, 100000, 0, false, false, 1, '2020-01-01', '2020-01-01', true, '1234567890', '1234567890', 'description', '1234567890', '1234567890', true, true, true, true, 'contact_person', 'email', 10000, 'USD', 1);
insert into comtrans (user_id, comtran_category_id, comtran_brand_id, comtran_model_id, fuel_type_id, city_id, color_id, engine, power, year, number_of_cycles, odometer, crash, not_cleared, owners, date_of_purchase, warranty_date, ptc, vin_code, certificate, description, can_look_coordinate, phone_number, refuse_dealers_calls, only_chat, protect_spam, verified_buyers, contact_person, email, price, price_type, status) values (1, 4, 1, 1, 1, 1, 1, 1000, 1000, 2020, 100000, 0, false, false, 1, '2020-01-01', '2020-01-01', true, '1234567890', '1234567890', 'description', '1234567890', '1234567890', true, true, true, true, 'contact_person', 'email', 10000, 'USD', 1);


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

insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (1, 1, 1);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (1, 2, 2);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (1, 3, 3);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (1, 4, 4);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (1, 5, 5);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (1, 6, 6);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (1, 7, 7);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (1, 8, 8);

insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (2, 1, 1);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (2, 2, 2);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (2, 3, 3);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (2, 4, 4);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (2, 5, 5);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (2, 6, 6);

insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (3, 1, 1);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (3, 2, 2);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (3, 3, 3);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (3, 4, 4);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (3, 5, 5);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (3, 6, 6);

insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (4, 1, 1);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (4, 2, 2);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (4, 3, 3);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (4, 4, 4);
insert into comtran_parameters (comtran_id, comtran_parameter_id, comtran_parameter_value_id) values (4, 5, 5);


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

insert into comtran_images (comtran_id, image) values (1, 'https://via.placeholder.com/150');
insert into comtran_images (comtran_id, image) values (2, 'https://via.placeholder.com/150');
insert into comtran_images (comtran_id, image) values (3, 'https://via.placeholder.com/150');
insert into comtran_images (comtran_id, image) values (4, 'https://via.placeholder.com/150');


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

insert into comtran_videos (comtran_id, video) values (1, 'https://via.placeholder.com/150');
insert into comtran_videos (comtran_id, video) values (2, 'https://via.placeholder.com/150');
insert into comtran_videos (comtran_id, video) values (3, 'https://via.placeholder.com/150');
insert into comtran_videos (comtran_id, video) values (4, 'https://via.placeholder.com/150');

insert into comtran_videos (comtran_id, video) values (1, 'https://via.placeholder.com/150');
insert into comtran_videos (comtran_id, video) values (2, 'https://via.placeholder.com/150');
insert into comtran_videos (comtran_id, video) values (3, 'https://via.placeholder.com/150');
insert into comtran_videos (comtran_id, video) values (4, 'https://via.placeholder.com/150');

