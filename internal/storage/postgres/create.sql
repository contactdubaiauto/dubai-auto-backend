create user da with password '1234';
grant all privileges on database da to da;
grant all privileges on schema public to da;
grant all privileges on all tables in schema public to da;
grant all privileges on all sequences in schema public to da;
alter default privileges in schema public grant all on tables to da;
alter default privileges in schema public grant all on sequences to da;


drop table if exists images;
drop table if exists vehicles;
drop table if exists profiles;
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

insert into users (username, email, password, phone, created_at) 
    values ('user2', 'user2@gmail.com', '$2a$10$Cya9x0xSJSnRknBmJpW.Bu8ukZpVTqzwgrQgAYNPXdrX2HYGRk33W', '0111222222', now()); -- password: 12345678

create table profiles (
    "id" serial primary key, 
    "user_id" int not null,
    "notification" boolean default false,
    "last_active_date" timestamp default now(),
    "created_at" timestamp default now(),
    constraint profiles_user_id_fk 
    foreign key (user_id) references users(id) on delete cascade on update cascade
);

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
    "popular" boolean default false,
    "updated_at" timestamp default now()
);

insert into brands (name, logo, car_count, popular) values ('Toyota', '/images/logo/toyota.png', 12, true);
insert into brands (name, logo, car_count, popular) values ('Honda', '/images/logo/honda.png', 8, false);
insert into brands (name, logo, car_count, popular) values ('Ford', '/images/logo/ford.png', 38, false);
insert into brands (name, logo, car_count, popular) values ('Chevrolet', '/images/logo/chevrolet.png', 5, false);
insert into brands (name, logo, car_count, popular) values ('Nissan', '/images/logo/nissan.png', 22, true);
insert into brands (name, logo, car_count, popular) values ('Hyundai', '/images/logo/hyundai.png', 74, false);
insert into brands (name, logo, car_count, popular) values ('Kia', '/images/logo/kia.png', 33, true);



create table models (
    "id" serial primary key,
    "name" varchar(255) not null,
    "brand_id" int not null,
    "popular" boolean default false,
    "car_count" int not null default 0,
    "updated_at" timestamp default now(),
    constraint models_brand_id_fk 
        foreign key (brand_id) 
            references brands(id)
                on delete cascade
                on update cascade
);

-- toyota
insert into models (name, brand_id, popular, car_count) values ('Camry', 1, true, 7);
insert into models (name, brand_id, popular, car_count) values ('Corolla', 1, true, 41);
insert into models (name, brand_id, popular, car_count) values ('Rav4', 1, false, 73);
insert into models (name, brand_id, popular, car_count) values ('Land Cruiser', 1, false, 1);

-- honda
insert into models (name, brand_id, popular, car_count) values ('Civic', 2, true, 34);
insert into models (name, brand_id, popular, car_count) values ('Accord', 2, false, 23);
insert into models (name, brand_id, popular, car_count) values ('CR-V', 2, false, 65);

-- ford
insert into models (name, brand_id, popular, car_count) values ('F-150', 3, true, 56);
insert into models (name, brand_id, popular, car_count) values ('Mustang', 3, true, 29);
insert into models (name, brand_id, popular, car_count) values ('Explorer', 3, false, 664);
insert into models (name, brand_id, popular, car_count) values ('Bronco', 3, false, 34);

-- chevrolet
insert into models (name, brand_id, popular, car_count) values ('Camaro', 4, true, 23);
insert into models (name, brand_id, popular, car_count) values ('Corvette', 4, true, 645);
insert into models (name, brand_id, popular, car_count) values ('Equinox', 4, false, 23);
insert into models (name, brand_id, popular, car_count) values ('Silverado', 4, false, 55);

-- nissan
insert into models (name, brand_id, popular, car_count) values ('Altima', 5, true, 23);
insert into models (name, brand_id, popular, car_count) values ('Pathfinder', 5, true, 56);
insert into models (name, brand_id, popular, car_count) values ('Rogue', 5, false, 22);
insert into models (name, brand_id, popular, car_count) values ('Sentra', 5, false, 53);

-- hyundai
insert into models (name, brand_id, popular, car_count) values ('Elantra', 6, true, 54);
insert into models (name, brand_id, popular, car_count) values ('Sonata', 6, false, 12);
insert into models (name, brand_id, popular, car_count) values ('Kona', 6, false, 45);

-- kia
insert into models (name, brand_id, popular, car_count) values ('K5', 7, true, 50);
insert into models (name, brand_id, popular, car_count) values ('K7', 7, false, 98);



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

insert into generations (name, model_id, start_year, end_year, image) values ('First Generation (model 1)', 1, 2020, 2022, '/images/gens/1.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Second Generation (model 1)', 1, 2023, 2025, '/images/gens/2.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Third Generation (model 1)', 1, 2012, 2022, '/images/gens/3.jpg');

insert into generations (name, model_id, start_year, end_year, image) values ('First Generation (model 2)', 2, 2010, 2022, '/images/gens/1.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Second Generation (model 2)', 2, 2007, 2025, '/images/gens/2.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Third Generation (model 2)', 2, 2001, 2022, '/images/gens/3.jpg');

insert into generations (name, model_id, start_year, end_year, image) values ('First Generation (model 3)', 3, 2000, 2022, '/images/gens/1.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Second Generation (model 3)', 3, 2004, 2025, '/images/gens/2.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Third Generation (model 3)', 3, 2003, 2022, '/images/gens/3.jpg');

insert into generations (name, model_id, start_year, end_year, image) values ('First Generation (model 4)', 4, 2000, 2022, '/images/gens/1.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Second Generation (model 4)', 4, 2004, 2025, '/images/gens/2.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Third Generation (model 4)', 4, 2003, 2022, '/images/gens/3.jpg');

insert into generations (name, model_id, start_year, end_year, image) values ('First Generation (model 5)', 5, 2000, 2022, '/images/gens/1.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Second Generation (model 5)', 5, 2004, 2025, '/images/gens/2.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Third Generation (model 5)', 5, 2003, 2022, '/images/gens/3.jpg');

insert into generations (name, model_id, start_year, end_year, image) values ('First Generation (model 6)', 6, 2000, 2022, '/images/gens/1.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Second Generation (model 6)', 6, 2004, 2025, '/images/gens/2.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Third Generation (model 6)', 6, 2003, 2022, '/images/gens/3.jpg');

insert into generations (name, model_id, start_year, end_year, image) values ('First Generation (model 7)', 7, 2000, 2022, '/images/gens/1.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Second Generation (model 7)', 7, 2004, 2025, '/images/gens/2.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Third Generation (model 7)', 7, 2003, 2022, '/images/gens/3.jpg');

insert into generations (name, model_id, start_year, end_year, image) values ('First Generation (model 8)', 8, 2000, 2022, '/images/gens/1.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Second Generation (model 8)', 8, 2004, 2025, '/images/gens/2.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Third Generation (model 8)', 8, 2003, 2022, '/images/gens/3.jpg');

insert into generations (name, model_id, start_year, end_year, image) values ('First Generation (model 9)', 9, 2000, 2022, '/images/gens/1.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Second Generation (model 9)', 9, 2004, 2025, '/images/gens/2.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Third Generation (model 9)', 9, 2003, 2022, '/images/gens/3.jpg');

insert into generations (name, model_id, start_year, end_year, image) values ('First Generation (model 10)', 10, 2000, 2022, '/images/gens/1.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Second Generation (model 10)', 10, 2004, 2025, '/images/gens/2.jpg');
insert into generations (name, model_id, start_year, end_year, image) values ('Third Generation (model 10)', 10, 2003, 2022, '/images/gens/3.jpg');

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

insert into generation_body_types (generation_id, body_type_id) values (1, 1);
insert into generation_body_types (generation_id, body_type_id) values (1, 2);
insert into generation_body_types (generation_id, body_type_id) values (2, 3);
insert into generation_body_types (generation_id, body_type_id) values (2, 4);
insert into generation_body_types (generation_id, body_type_id) values (3, 5);
insert into generation_body_types (generation_id, body_type_id) values (3, 6);

insert into generation_body_types (generation_id, body_type_id) values (4, 1);
insert into generation_body_types (generation_id, body_type_id) values (4, 2);
insert into generation_body_types (generation_id, body_type_id) values (5, 3);
insert into generation_body_types (generation_id, body_type_id) values (5, 4);
insert into generation_body_types (generation_id, body_type_id) values (6, 5);
insert into generation_body_types (generation_id, body_type_id) values (6, 6);

insert into generation_body_types (generation_id, body_type_id) values (7, 1);
insert into generation_body_types (generation_id, body_type_id) values (7, 2);
insert into generation_body_types (generation_id, body_type_id) values (8, 3);
insert into generation_body_types (generation_id, body_type_id) values (8, 4);
insert into generation_body_types (generation_id, body_type_id) values (9, 5);
insert into generation_body_types (generation_id, body_type_id) values (9, 6);

insert into generation_body_types (generation_id, body_type_id) values (10, 1);
insert into generation_body_types (generation_id, body_type_id) values (10, 2);
insert into generation_body_types (generation_id, body_type_id) values (11, 3);
insert into generation_body_types (generation_id, body_type_id) values (11, 4);
insert into generation_body_types (generation_id, body_type_id) values (12, 5);
insert into generation_body_types (generation_id, body_type_id) values (12, 6);

insert into generation_body_types (generation_id, body_type_id) values (13, 1);
insert into generation_body_types (generation_id, body_type_id) values (13, 2);
insert into generation_body_types (generation_id, body_type_id) values (14, 3);
insert into generation_body_types (generation_id, body_type_id) values (14, 4);
insert into generation_body_types (generation_id, body_type_id) values (15, 5);
insert into generation_body_types (generation_id, body_type_id) values (15, 6);

insert into generation_body_types (generation_id, body_type_id) values (16, 1);
insert into generation_body_types (generation_id, body_type_id) values (16, 2);
insert into generation_body_types (generation_id, body_type_id) values (17, 3);
insert into generation_body_types (generation_id, body_type_id) values (17, 4);
insert into generation_body_types (generation_id, body_type_id) values (18, 5);
insert into generation_body_types (generation_id, body_type_id) values (18, 6);

insert into generation_body_types (generation_id, body_type_id) values (19, 1);
insert into generation_body_types (generation_id, body_type_id) values (19, 2);
insert into generation_body_types (generation_id, body_type_id) values (20, 3);
insert into generation_body_types (generation_id, body_type_id) values (20, 4);
insert into generation_body_types (generation_id, body_type_id) values (21, 5);
insert into generation_body_types (generation_id, body_type_id) values (21, 6);

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

insert into generation_fuel_types (generation_id, fuel_type_id) values (1, 1);
insert into generation_fuel_types (generation_id, fuel_type_id) values (1, 2);
insert into generation_fuel_types (generation_id, fuel_type_id) values (2, 3);
insert into generation_fuel_types (generation_id, fuel_type_id) values (2, 4);

insert into generation_fuel_types (generation_id, fuel_type_id) values (3, 1);
insert into generation_fuel_types (generation_id, fuel_type_id) values (3, 2);
insert into generation_fuel_types (generation_id, fuel_type_id) values (4, 3);
insert into generation_fuel_types (generation_id, fuel_type_id) values (4, 4);

insert into generation_fuel_types (generation_id, fuel_type_id) values (5, 1);
insert into generation_fuel_types (generation_id, fuel_type_id) values (5, 2);
insert into generation_fuel_types (generation_id, fuel_type_id) values (6, 3);
insert into generation_fuel_types (generation_id, fuel_type_id) values (6, 4);

insert into generation_fuel_types (generation_id, fuel_type_id) values (7, 1);
insert into generation_fuel_types (generation_id, fuel_type_id) values (7, 2);
insert into generation_fuel_types (generation_id, fuel_type_id) values (8, 3);
insert into generation_fuel_types (generation_id, fuel_type_id) values (8, 4);

insert into generation_fuel_types (generation_id, fuel_type_id) values (9, 1);
insert into generation_fuel_types (generation_id, fuel_type_id) values (9, 2);
insert into generation_fuel_types (generation_id, fuel_type_id) values (10, 3);
insert into generation_fuel_types (generation_id, fuel_type_id) values (10, 4);

insert into generation_fuel_types (generation_id, fuel_type_id) values (11, 1);
insert into generation_fuel_types (generation_id, fuel_type_id) values (11, 2);
insert into generation_fuel_types (generation_id, fuel_type_id) values (12, 3);
insert into generation_fuel_types (generation_id, fuel_type_id) values (12, 4);

insert into generation_fuel_types (generation_id, fuel_type_id) values (13, 1);
insert into generation_fuel_types (generation_id, fuel_type_id) values (13, 2);
insert into generation_fuel_types (generation_id, fuel_type_id) values (14, 3);
insert into generation_fuel_types (generation_id, fuel_type_id) values (14, 4);


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

insert into generation_drivetrains (generation_id, drivetrain_id) values (1, 1);
insert into generation_drivetrains (generation_id, drivetrain_id) values (1, 2);
insert into generation_drivetrains (generation_id, drivetrain_id) values (2, 3);
insert into generation_drivetrains (generation_id, drivetrain_id) values (2, 1);
insert into generation_drivetrains (generation_id, drivetrain_id) values (3, 2);
insert into generation_drivetrains (generation_id, drivetrain_id) values (3, 3);

insert into generation_drivetrains (generation_id, drivetrain_id) values (4, 1);
insert into generation_drivetrains (generation_id, drivetrain_id) values (4, 2);
insert into generation_drivetrains (generation_id, drivetrain_id) values (5, 3);
insert into generation_drivetrains (generation_id, drivetrain_id) values (5, 1);
insert into generation_drivetrains (generation_id, drivetrain_id) values (6, 2);
insert into generation_drivetrains (generation_id, drivetrain_id) values (6, 3);

insert into generation_drivetrains (generation_id, drivetrain_id) values (7, 1);
insert into generation_drivetrains (generation_id, drivetrain_id) values (7, 2);
insert into generation_drivetrains (generation_id, drivetrain_id) values (8, 3);
insert into generation_drivetrains (generation_id, drivetrain_id) values (8, 1);
insert into generation_drivetrains (generation_id, drivetrain_id) values (9, 2);
insert into generation_drivetrains (generation_id, drivetrain_id) values (9, 3);

insert into generation_drivetrains (generation_id, drivetrain_id) values (10, 1);
insert into generation_drivetrains (generation_id, drivetrain_id) values (10, 2);
insert into generation_drivetrains (generation_id, drivetrain_id) values (11, 3);
insert into generation_drivetrains (generation_id, drivetrain_id) values (11, 1);
insert into generation_drivetrains (generation_id, drivetrain_id) values (12, 2);
insert into generation_drivetrains (generation_id, drivetrain_id) values (12, 3);

insert into generation_drivetrains (generation_id, drivetrain_id) values (13, 1);
insert into generation_drivetrains (generation_id, drivetrain_id) values (13, 2);
insert into generation_drivetrains (generation_id, drivetrain_id) values (14, 3);
insert into generation_drivetrains (generation_id, drivetrain_id) values (14, 1);
insert into generation_drivetrains (generation_id, drivetrain_id) values (15, 2);
insert into generation_drivetrains (generation_id, drivetrain_id) values (15, 3);

insert into generation_drivetrains (generation_id, drivetrain_id) values (16, 1);
insert into generation_drivetrains (generation_id, drivetrain_id) values (16, 2);
insert into generation_drivetrains (generation_id, drivetrain_id) values (17, 3);
insert into generation_drivetrains (generation_id, drivetrain_id) values (17, 1);
insert into generation_drivetrains (generation_id, drivetrain_id) values (18, 2);
insert into generation_drivetrains (generation_id, drivetrain_id) values (18, 3);

insert into generation_drivetrains (generation_id, drivetrain_id) values (19, 1);
insert into generation_drivetrains (generation_id, drivetrain_id) values (19, 2);
insert into generation_drivetrains (generation_id, drivetrain_id) values (20, 3);
insert into generation_drivetrains (generation_id, drivetrain_id) values (20, 1);
insert into generation_drivetrains (generation_id, drivetrain_id) values (21, 2);
insert into generation_drivetrains (generation_id, drivetrain_id) values (21, 3);

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


insert into generation_transmissions (generation_id, transmission_id) values (1, 1);
insert into generation_transmissions (generation_id, transmission_id) values (1, 2);
insert into generation_transmissions (generation_id, transmission_id) values (2, 3);
insert into generation_transmissions (generation_id, transmission_id) values (2, 1);
insert into generation_transmissions (generation_id, transmission_id) values (3, 2);
insert into generation_transmissions (generation_id, transmission_id) values (3, 3);

insert into generation_transmissions (generation_id, transmission_id) values (4, 1);
insert into generation_transmissions (generation_id, transmission_id) values (4, 2);
insert into generation_transmissions (generation_id, transmission_id) values (5, 3);
insert into generation_transmissions (generation_id, transmission_id) values (5, 1);
insert into generation_transmissions (generation_id, transmission_id) values (6, 2);
insert into generation_transmissions (generation_id, transmission_id) values (6, 3);

insert into generation_transmissions (generation_id, transmission_id) values (7, 1);
insert into generation_transmissions (generation_id, transmission_id) values (7, 2);
insert into generation_transmissions (generation_id, transmission_id) values (8, 3);
insert into generation_transmissions (generation_id, transmission_id) values (8, 1);
insert into generation_transmissions (generation_id, transmission_id) values (9, 2);
insert into generation_transmissions (generation_id, transmission_id) values (9, 3);

insert into generation_transmissions (generation_id, transmission_id) values (10, 1);
insert into generation_transmissions (generation_id, transmission_id) values (10, 2);
insert into generation_transmissions (generation_id, transmission_id) values (11, 3);
insert into generation_transmissions (generation_id, transmission_id) values (11, 1);
insert into generation_transmissions (generation_id, transmission_id) values (12, 2);
insert into generation_transmissions (generation_id, transmission_id) values (12, 3);

insert into generation_transmissions (generation_id, transmission_id) values (13, 1);
insert into generation_transmissions (generation_id, transmission_id) values (13, 2);
insert into generation_transmissions (generation_id, transmission_id) values (14, 3);
insert into generation_transmissions (generation_id, transmission_id) values (14, 1);
insert into generation_transmissions (generation_id, transmission_id) values (15, 2);
insert into generation_transmissions (generation_id, transmission_id) values (15, 3);

insert into generation_transmissions (generation_id, transmission_id) values (16, 1);
insert into generation_transmissions (generation_id, transmission_id) values (16, 2);
insert into generation_transmissions (generation_id, transmission_id) values (17, 3);
insert into generation_transmissions (generation_id, transmission_id) values (17, 1);
insert into generation_transmissions (generation_id, transmission_id) values (18, 2);
insert into generation_transmissions (generation_id, transmission_id) values (18, 3);

insert into generation_transmissions (generation_id, transmission_id) values (19, 1);
insert into generation_transmissions (generation_id, transmission_id) values (19, 2);
insert into generation_transmissions (generation_id, transmission_id) values (20, 3);
insert into generation_transmissions (generation_id, transmission_id) values (20, 1);
insert into generation_transmissions (generation_id, transmission_id) values (21, 2);
insert into generation_transmissions (generation_id, transmission_id) values (21, 3);


create table generation_modifications (
    "id" serial primary key,
    "generation_id" int not null,
    "body_type_id" int not null,
    "fuel_type_id" int not null, 
    "drivetrain_id" int not null,
    "transmission_id" int not null, 
    "name" character varying(100) not null,
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

insert into generation_modifications (generation_id, body_type_id, fuel_type_id, drivetrain_id, transmission_id, name) 
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
    "city_id" int default 1,
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
    "status" int not null default 2, -- 1-pending, 2-not sale (my cars), 3-on sale,
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
    user_id, brand_id, city_id, model_id, body_type_id, fuel_type_id, ownership_type_id, year, exchange, credit, right_hand_drive, odometer, vin_code, door_count, phone_number, price, new,
    credit_price, status, crash, negotiable, drivetrain_id, mileage_km, 
    interior_color_id, color_id, view_count 
)  values 
(
    1, 1, 1, 1, 1, 1, 1, 2020, false, false, false, 74839782, '238748927', 4, '23487827397', 39999, true, 
    388292, 3, false, true, 2, 16611888, 
    2, 1, 392
),
(
    1, 2, 1, 2, 2, 2, 1, 2019, false, false, false, 273854, '26873900987', 4, '23748798273942', 7828973, true,
    237848, 3, false, true, 3, 2772719,
    3, 2, 234
),
(
    2, 2, 1, 2, 2, 2, 1, 2010, false, false, false, 8373, '98987987987', 4, '89877683783', 982739488, true,
    87487, 3, false, true, 3, 7388,
    3, 2, 487
),
(
    1, 2, 1, 2, 2, 2, 1, 2006, false, false, false, 2784, '98987987987', 4, '89877683783', 982739488, true,
    84883, 3, false, true, 3, 92222,
    3, 2, 438
),
(
    2, 2, 1, 2, 2, 2, 1, 2022, false, false, false, 2739, '98987987987', 4, '89877683783', 982739488, true,
    388948, 3, false, true, 3, 272727,
    3, 2, 3
),
(
    1, 2, 1, 2, 2, 2, 1, 2021, true, false, false, 283847, '0987654321', 4, '01234567890', 38742973, true,
    15900, 3, false, true, 3, 262662,
    3, 2, 45
),
(
    1, 3, 1, 3, 3, 3, 1, 2022, false, false, false, 2837959, '1122334455', 4, '01234567890', 3485, true,
    2387478, 2, false, true, 1, 773737,
    2, 1, 12
),
(
    2, 3, 1, 3, 3, 3, 1, 2022, false, false, false, 2348859, '1122334455', 4, '01234567890', 1289397, true,
    2377485, 2, false, true, 1, 327478,
    3, 3, 8
),
(
    2, 3, 1, 3, 3, 3, 1, 2022, false, false, false, 234788, '1122334455', 4, '01234567890', 23487, true,
    27272722, 3, false, true, 1, 2348782,
    1, 3, 97
);
-- 1-pending, 2-not sale (my cars), 3-on sale,


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

insert into images (vehicle_id, image) values (2, '/images/cars/2/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (2, '/images/cars/2/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (2, '/images/cars/2/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (2, '/images/cars/2/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (3, '/images/cars/3/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (3, '/images/cars/3/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (3, '/images/cars/3/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (3, '/images/cars/3/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (4, '/images/cars/4/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (4, '/images/cars/4/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (4, '/images/cars/4/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (4, '/images/cars/4/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (5, '/images/cars/5/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (5, '/images/cars/5/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');    
insert into images (vehicle_id, image) values (5, '/images/cars/5/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (5, '/images/cars/5/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (6, '/images/cars/6/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (6, '/images/cars/6/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (6, '/images/cars/6/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (6, '/images/cars/6/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (7, '/images/cars/7/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (7, '/images/cars/7/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (7, '/images/cars/7/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (7, '/images/cars/7/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (8, '/images/cars/8/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (8, '/images/cars/8/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (8, '/images/cars/8/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (8, '/images/cars/8/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (9, '/images/cars/9/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (9, '/images/cars/9/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (9, '/images/cars/9/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (9, '/images/cars/9/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

-- steps for add vehicle
-- 1. vin code, 2. marka, 3. model, 4. year, 5. generations, 6. body types, 7. engine, 8. driveride, 
-- 9. transmission, 10. modification, 11. color, 12. images
