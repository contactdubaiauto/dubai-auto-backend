create user da with password '1234';
grant all privileges on database da to da;
grant all privileges on schema public to da;
grant all privileges on all tables in schema public to da;
grant all privileges on all sequences in schema public to da;
alter default privileges in schema public grant all on tables to da;
alter default privileges in schema public grant all on sequences to da;


drop table if exists images;
drop table if exists vehicles;
drop table if exists generation_modifications;
drop table if exists generation_body_types;
drop table if exists generation_transmissions;
drop table if exists generation_fuel_types;
drop table if exists generation_drivetrains;
drop table if exists colors;
drop table if exists services;
drop table if exists service_types;
drop table if exists regions;
drop table if exists cities;
drop table if exists fuel_types;
drop table if exists drivetrains;
drop table if exists engines;
drop table if exists transmissions;
drop table if exists body_types;
drop table if exists generations;
drop table if exists models;
drop table if exists brands;
drop table if exists users;
drop table if exists admins;
drop table if exists ownership_types;


create table users (
    "id" serial primary key,
    "username" varchar(100) not null,
    "email" varchar(100),
    "role_id" int not null default 1,
    "password" varchar(100) not null,
    "phone" varchar(100),
    "updated_at" timestamp default now(),
    "created_at" timestamp default now(),
    unique("email"),
    unique("phone")
);

insert into users (username, email, password, phone, created_at) 
    values ('user', 'user@gmail.com', '$2a$10$Cya9x0xSJSnRknBmJpW.Bu8ukZpVTqzwgrQgAYNPXdrX2HYGRk33W', '01234567890', now()); -- password: 12345678

create table profiles (
    "id" serial primary key, 
    "user_id" int not null,
    "notification" boolean default false,
    "last_active_date" timestamp default now(),
    "created_at" timestamp default now(),
    constraint profiles_user_id_fk 
    foreign key (user_id) references users(id) on delete cascade on update cascade
)

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
    "car_count" int not null default 0,
    "updated_at" timestamp default now()
);

insert into brands (name, logo, car_count) values ('Toyota', '/images/logo/toyota.png', 0);
insert into brands (name, logo, car_count) values ('Honda', '/images/logo/honda.png', 0);
insert into brands (name, logo, car_count) values ('Ford', '/images/logo/ford.png', 0);
insert into brands (name, logo, car_count) values ('Chevrolet', '/images/logo/chevrolet.png', 0);
insert into brands (name, logo, car_count) values ('Nissan', '/images/logo/nissan.png', 0);
insert into brands (name, logo, car_count) values ('Hyundai', '/images/logo/hyundai.png', 0);
insert into brands (name, logo, car_count) values ('Kia', '/images/logo/kia.png', 0);



create table models (
    "id" serial primary key,
    "name" varchar(255) not null,
    "brand_id" int not null,
    "updated_at" timestamp default now(),
    constraint models_brand_id_fk 
        foreign key (brand_id) 
            references brands(id)
                on delete cascade
                on update cascade
);

-- toyota
insert into models (name, brand_id) values ('Camry', 1);
insert into models (name, brand_id) values ('Corolla', 1);
insert into models (name, brand_id) values ('Rav4', 1);
insert into models (name, brand_id) values ('Land Cruiser', 1);

-- honda
insert into models (name, brand_id) values ('Civic', 2);
insert into models (name, brand_id) values ('Accord', 2);
insert into models (name, brand_id) values ('CR-V', 2);

-- ford
insert into models (name, brand_id) values ('F-150', 3);
insert into models (name, brand_id) values ('Mustang', 3);
insert into models (name, brand_id) values ('Explorer', 3);
insert into models (name, brand_id) values ('Bronco', 3);

-- chevrolet
insert into models (name, brand_id) values ('Camaro', 4);
insert into models (name, brand_id) values ('Corvette', 4);
insert into models (name, brand_id) values ('Equinox', 4);
insert into models (name, brand_id) values ('Silverado', 4);

-- nissan
insert into models (name, brand_id) values ('Altima', 5);
insert into models (name, brand_id) values ('Pathfinder', 5);
insert into models (name, brand_id) values ('Rogue', 5);
insert into models (name, brand_id) values ('Sentra', 5);

-- hyundai
insert into models (name, brand_id) values ('Elantra', 6);
insert into models (name, brand_id) values ('Sonata', 6);
insert into models (name, brand_id) values ('Kona', 6);

-- kia
insert into models (name, brand_id) values ('K5', 7);
insert into models (name, brand_id) values ('K7', 7);



create table body_types (
    "id" serial primary key,
    "name" varchar(255) not null,
    "image" character varying(255) not null,
    "created_at" timestamp default now()
);

insert into body_types (name, image) values ('Sedan','/images/body/sedan.png');
insert into body_types (name, image) values ('Hatchback', '/images/body/hatchback.png');
insert into body_types (name, image) values ('SUV', '/images/body/suv.png');
insert into body_types (name, image) values ('Crossover', '/images/body/crossover.png');
insert into body_types (name, image) values ('Coupe', '/images/body/coupe.png');
insert into body_types (name, image) values ('Convertible', '/images/body/convertible.png');
insert into body_types (name, image) values ('Wagon', '/images/body/wagon.png');
insert into body_types (name, image) values ('Pickup Truck', '/images/body/pickup_truck.png');
insert into body_types (name, image) values ('Van', '/images/body/van.png');
insert into body_types (name, image) values ('Minivan', '/images/body/minivan.png');
insert into body_types (name, image) values ('Roadster', '/images/body/roadster.png');
insert into body_types (name, image) values ('Sports Car', '/images/body/sports_car.png');
insert into body_types (name, image) values ('Off-Road', '/images/body/off_road.png');
insert into body_types (name, image) values ('Limousine', '/images/body/limousine.png'); 
insert into body_types (name, image) values ('Utility', '/images/body/utility.png');




create table transmissions (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now()
);

insert into transmissions (name) values ('Automatic');
insert into transmissions (name) values ('Manual');
insert into transmissions (name) values ('Semi-Automatic');


create table engines (
    "id" serial primary key,
    "value" varchar(255) not null,
    "created_at" timestamp default now()
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
    "created_at" timestamp default now()
);

insert into drivetrains (name) values ('Front-Wheel Drive');
insert into drivetrains (name) values ('Rear-Wheel Drive');
insert into drivetrains (name) values ('All-Wheel Drive');


create table fuel_types (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now()
);

insert into fuel_types (name) values ('Gasoline');
insert into fuel_types (name) values ('Diesel');
insert into fuel_types (name) values ('Electric');
insert into fuel_types (name) values ('Hybrid');

create table cities (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now()
);

insert into cities (name) values ('Dubai');
insert into cities (name) values ('Abu Dhabi');
insert into cities (name) values ('Sharjah');

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
    "image" varchar(255) not null,
    "end_year" int not null,
    "created_at" timestamp default now(),
    constraint generations_model_id_fk
        foreign key (model_id)
            references models(id)
                on delete cascade,
    unique(model_id, start_year, end_year)
);

insert into generations (name, model_id, start_year, end_year, image) values ('First Generation', 1, 2020, 2022, '/images/gens/1.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Second Generation', 1, 2023, 2025, '/images/gens/2.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('First Generation', 3, 2020, 2022, '/images/gens/3.jpg');

create table generation_body_types (
    "id" serial primary key,
    "generation_id" int not null,
    "body_type_id" int not null,
    "created_at" timestamp default now(),
    constraint generation_body_types_generation_id_fk
        foreign key (generation_id)
            references generations(id)
                on delete cascade
                on update cascade,
    constraint generation_body_types_body_type_id_fk
        foreign key (body_type_id)
            references body_types(id)
                on delete cascade
                on update cascade,
    unique(generation_id, body_type_id)
);

insert into generation_body_types (generation_id, body_type_id) values (1, 1); -- First Generation Camry - Sedan
insert into generation_body_types (generation_id, body_type_id) values (1, 2); -- First Generation Camry - Hatchback
insert into generation_body_types (generation_id, body_type_id) values (2, 1); -- Second Generation Camry - Sedan
insert into generation_body_types (generation_id, body_type_id) values (2, 2); -- Second Generation Camry - Hatchback
insert into generation_body_types (generation_id, body_type_id) values (3, 1); -- First Generation Rav4 - Sedan
insert into generation_body_types (generation_id, body_type_id) values (3, 2); -- First Generation Rav4 - Hatchback

create table generation_fuel_types (
    "id" serial primary key,
    "generation_id" int not null,
    "fuel_type_id" int not null,
    "created_at" timestamp default now(),
    constraint generation_fuel_types_generation_id_fk
        foreign key (generation_id)
            references generations(id)
                on delete cascade
                on update cascade,
    constraint generation_fuel_types_fuel_type_id_fk
        foreign key (fuel_type_id)
            references fuel_types(id)
                on delete cascade
                on update cascade,
    unique(generation_id, fuel_type_id)
);

insert into generation_fuel_types (generation_id, fuel_type_id) values (1, 1); -- First Generation Camry - Gasoline
insert into generation_fuel_types (generation_id, fuel_type_id) values (1, 2); -- First Generation Camry - Diesel
insert into generation_fuel_types (generation_id, fuel_type_id) values (2, 1); -- Second Generation Camry - Gasoline
insert into generation_fuel_types (generation_id, fuel_type_id) values (2, 2); -- Second Generation Camry - Diesel
insert into generation_fuel_types (generation_id, fuel_type_id) values (3, 1); -- First Generation Rav4 - Gasoline
insert into generation_fuel_types (generation_id, fuel_type_id) values (3, 2); -- First Generation Rav4 - Diesel

create table generation_drivetrains (
    "id" serial primary key,
    "generation_id" int not null,
    "drivetrain_id" int not null,
    "created_at" timestamp default now(),
    constraint generation_drivetrains_generation_id_fk
        foreign key (generation_id)
            references generations(id)
                on delete cascade
                on update cascade,
    constraint generation_drivetrains_drivetrain_id_fk
        foreign key (drivetrain_id)
            references drivetrains(id)
                on delete cascade
                on update cascade,
    unique(generation_id, drivetrain_id)
);

insert into generation_drivetrains (generation_id, drivetrain_id) values (1, 1); -- First Generation Camry - Front-Wheel Drive
insert into generation_drivetrains (generation_id, drivetrain_id) values (1, 2); -- First Generation Camry - Rear-Wheel Drive
insert into generation_drivetrains (generation_id, drivetrain_id) values (2, 1); -- Second Generation Camry - Front-Wheel Drive
insert into generation_drivetrains (generation_id, drivetrain_id) values (2, 2); -- Second Generation Camry - Rear-Wheel Drive
insert into generation_drivetrains (generation_id, drivetrain_id) values (3, 1); -- First Generation Rav4 - Front-Wheel Drive
insert into generation_drivetrains (generation_id, drivetrain_id) values (3, 2); -- First Generation Rav4 - Rear-Wheel Drive

create table generation_transmissions (
    "id" serial primary key,
    "generation_id" int not null,
    "transmission_id" int not null,
    "created_at" timestamp default now(),
    constraint generation_transmissions_generation_id_fk
        foreign key (generation_id)
            references generations(id)
                on delete cascade
                on update cascade,
    constraint generation_transmissions_transmission_id_fk
        foreign key (transmission_id)
            references transmissions(id)
                on delete cascade
                on update cascade,
    unique(generation_id, transmission_id)
);

insert into generation_transmissions (generation_id, transmission_id) values (1, 1); -- First Generation Camry - Automatic
insert into generation_transmissions (generation_id, transmission_id) values (1, 2); -- First Generation Camry - Manual
insert into generation_transmissions (generation_id, transmission_id) values (2, 1); -- Second Generation Camry - Automatic
insert into generation_transmissions (generation_id, transmission_id) values (2, 2); -- Second Generation Camry - Manual
insert into generation_transmissions (generation_id, transmission_id) values (3, 1); -- First Generation Rav4 - Automatic
insert into generation_transmissions (generation_id, transmission_id) values (3, 2); -- First Generation Rav4 - Manual

create table generation_modifications (
    "id" serial primary key,
    "generation_id" int not null,
    "body_type_id" int not null,
    "fuel_type_id" int not null, 
    "drivetrain_id" int not null,
    "transmission_id" int not null, 
    "title" character varying(100) not null,
    constraint generation_modifications_generation_id_fk
        foreign key (generation_id)
            references generations(id)
                on delete cascade
                on update cascade,
    constraint generation_modifications_body_type_id_fk
        foreign key (body_type_id)
            references body_types(id)
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
                on update cascade
);

insert into generation_modifications (generation_id, body_type_id, fuel_type_id, drivetrain_id, transmission_id, title) 
values 
    (1, 1, 1, 1, 1, '2.8 MT Gas (170 l.c.)'),    
    (1, 1, 1, 1, 1, '2.5 MT Gas (150 l.c.)'),
    (1, 1, 1, 1, 1, '3.0 MT Gas (180 l.c.)'),
    (1, 1, 1, 1, 1, '3.0 MT Gas (200 l.c.)'),
    (1, 1, 1, 1, 1, '2.0 MT Gas (140 l.c.)'),
    (1, 1, 1, 1, 1, '2.4 MT Gas (160 l.c.)'),

    (1, 1, 1, 1, 2, '2.5 MT Gas (150 l.c.)'),
    (1, 1, 1, 1, 2, '3.0 MT Gas (180 l.c.)'),
    (1, 1, 1, 1, 2, '3.0 MT Gas (200 l.c.)'),

    (1, 1, 1, 2, 1, '3.0 MT Gas (180 l.c.)'),
    (1, 1, 1, 2, 2, '3.0 MT Gas (200 l.c.)'),
    (1, 2, 1, 1, 1, '2.0 MT Gas (140 l.c.)'),
    (1, 2, 1, 1, 2, '2.4 MT Gas (160 l.c.)'),
    (1, 2, 1, 2, 1, '2.5 MT Gas (155 l.c.)'),
    (1, 2, 1, 2, 2, '2.7 MT Gas (165 l.c.)'),
    (2, 1, 2, 1, 1, '2.8 AT Diesel (180 l.c.)'),
    (2, 1, 2, 1, 2, '2.5 AT Diesel (160 l.c.)'),
    (2, 1, 2, 2, 1, '3.0 AT Diesel (200 l.c.)'),
    (2, 1, 2, 2, 2, '3.2 AT Diesel (220 l.c.)'),
    (2, 2, 2, 1, 1, '2.0 AT Diesel (130 l.c.)'),
    (2, 2, 2, 1, 2, '2.4 AT Diesel (150 l.c.)'),
    (2, 2, 2, 2, 1, '2.5 AT Diesel (170 l.c.)'),
    (2, 2, 2, 2, 2, '2.7 AT Diesel (175 l.c.)'),
    (3, 1, 3, 1, 1, '2.0 CVT Electric (120 l.c.)'),
    (3, 1, 3, 1, 2, '2.5 CVT Electric (140 l.c.)'),
    (3, 1, 3, 2, 1, '3.0 CVT Electric (160 l.c.)'),
    (3, 1, 3, 2, 2, '3.2 CVT Electric (180 l.c.)'),
    (3, 2, 3, 1, 1, '2.0 CVT Electric (110 l.c.)'),
    (3, 2, 3, 1, 2, '2.4 CVT Electric (130 l.c.)'),
    (3, 2, 3, 2, 1, '2.5 CVT Electric (150 l.c.)'),
    (3, 2, 3, 2, 2, '2.7 CVT Electric (170 l.c.)'),
    (1, 1, 4, 1, 1, '2.5 MT Hybrid (160 l.c.)'),
    (1, 1, 4, 1, 2, '2.8 MT Hybrid (180 l.c.)'),
    (1, 1, 4, 2, 1, '3.0 MT Hybrid (200 l.c.)'),
    (1, 1, 4, 2, 2, '3.2 MT Hybrid (220 l.c.)'),
    (2, 2, 4, 1, 1, '2.0 AT Hybrid (140 l.c.)'),
    (2, 2, 4, 1, 2, '2.4 AT Hybrid (160 l.c.)'),
    (2, 2, 4, 2, 1, '2.5 AT Hybrid (180 l.c.)');


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
    "hex_code" varchar(7) not null,
    "created_at" timestamp default now()
);

insert into colors (name, hex_code) values ('Black', '#000000');
insert into colors (name, hex_code) values ('White', '#FFFFFF');
insert into colors (name, hex_code) values ('Red', '#FF0000');
insert into colors (name, hex_code) values ('Blue', '#0000FF');
insert into colors (name, hex_code) values ('Green', '#008000');
insert into colors (name, hex_code) values ('Yellow', '#FFFF00');
insert into colors (name, hex_code) values ('Silver', '#C0C0C0');
insert into colors (name, hex_code) values ('Gray', '#808080');
insert into colors (name, hex_code) values ('Orange', '#FFA500');
insert into colors (name, hex_code) values ('Purple', '#800080');


create table vehicles (
    "id" serial primary key,
    "user_id" int,
    "brand_id" int,
    "region_id" int,
    "city_id" int,
    "model_id" int,
    "generation_id" int,
    "transmission_id" int,
    "engine_id" int,
    "drivetrain_id" int,
    "body_type_id" int,
    "fuel_type_id" int,
    "ownership_type_id" int not null default 1,
    "announcement_type" int not null default 0,
    "view_count" int not null default 0,
    "year" int not null,
    "exchange" boolean not null default false,
    "credit" boolean not null default false,
    "right_hand_drive" boolean not null default false,
    "modification_id" int,
    "odometer" int,
    "vin_code" varchar(255),
    "door_count" int,
    "phone_number" varchar(255) not null,
    "price" int not null,
    "new" boolean not null default false,
    "color_id" int,
    "interior_color_id" int,
    "mileage_km" int,
    "crash" boolean not null default false,
    "negotiable" boolean not null default false,
    "credit_price" int,
    "status" int not null default 1,
    "updated_at" timestamp default now(),
    "created_at" timestamp default now(),
    constraint vehicles_generation_id_fk
        foreign key (generation_id)
            references generations(id)
                on delete set null
                on update cascade,
    constraint vehicles_color_id_fk
        foreign key (color_id)
            references colors(id)
                on delete set null
                on update cascade,
    constraint vehicles_interior_color_id_fk
        foreign key (interior_color_id)
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
    constraint vehicles_transmission_id_fk
        foreign key (transmission_id)
            references transmissions(id)
                on delete cascade
                on update cascade,
    constraint vehicles_engine_id_fk
        foreign key (engine_id)
            references engines(id)
                on delete cascade
                on update cascade,
    constraint vehicles_drivetrain_id_fk
        foreign key (drivetrain_id)
            references drivetrains(id)
                on delete cascade
                on update cascade,
    constraint vehicles_body_type_id_fk
        foreign key (body_type_id)
            references body_types(id)
                on delete cascade
                on update cascade,
    constraint vehicles_fuel_type_id_fk
        foreign key (fuel_type_id)
            references fuel_types(id)
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
                on update cascade,
    constraint vehicles_modification_id_fk
        foreign key (modification_id)
            references generation_modifications(id)
                on delete cascade 
                on update cascade
);

insert into vehicles (
    user_id, brand_id, city_id, model_id, body_type_id, fuel_type_id, ownership_type_id, year, 
    exchange, credit, right_hand_drive, odometer, vin_code, door_count, phone_number, price, new,
    credit_price, status, crash, negotiable, drivetrain_id, mileage_km, interior_color_id, color_id
)  values (
    1, 1, 1, 1, 1, 1, 1, 2020, false, false, false, 100000, '1234567890', 4, '01234567890', 100000, true, 100000, 1, false, true, 2, 200000, 
    2, 1
);
insert into vehicles (
    user_id, brand_id, city_id, model_id, body_type_id, fuel_type_id, ownership_type_id, year, 
    exchange, credit, right_hand_drive, odometer, vin_code, door_count, phone_number, price, new,
    credit_price, status, crash, negotiable, drivetrain_id, mileage_km, interior_color_id, color_id
)  values (
    1, 2, 1, 2, 2, 2, 1, 2021, false, false, false, 50000, '0987654321', 4, '01234567890', 150000, true,
    150000, 1, false, true, 3, 100000,
    3, 2
);
insert into vehicles (
    user_id, brand_id, city_id, model_id, body_type_id, fuel_type_id, ownership_type_id, year,
    exchange, credit, right_hand_drive, odometer, vin_code, door_count, phone_number, price, new,
    credit_price, status, crash, negotiable, drivetrain_id, mileage_km, interior_color_id, color_id
)  values (
    1, 3, 1, 3, 3, 3, 1, 2022, false, false, false, 30000, '1122334455', 4, '01234567890', 200000, true,
    200000, 1, false, true, 1, 50000,
    5, 3
);


create table images (
    "vehicle_id" int not null,
    "image" varchar(255) not null,
    constraint images_vehicle_id_fk
        foreign key (vehicle_id)
            references vehicles(id)
                on delete cascade
                on update cascade
);

insert into images (vehicle_id, image) values (1, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (1, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (1, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (1, '/images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- steps for add vehicle
-- 1. vin code, 2. marka, 3. model, 4. year, 5. generations, 6. body types, 7. engine, 8. driveride, 
-- 9. transmission, 10. modification, 11. color, 12. images
