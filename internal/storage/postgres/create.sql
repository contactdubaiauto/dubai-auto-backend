-- create user da with password '1234';
-- grant all privileges on database da to da;
-- grant all privileges on schema public to da;
-- grant all privileges on all tables in schema public to da;
-- grant all privileges on all sequences in schema public to da;
-- alter default privileges in schema public grant all on tables to da;
-- alter default privileges in schema public grant all on sequences to da;



drop table if exists temp_users;
drop table if exists configurations;
drop table if exists images;
drop table if exists videos;
drop table if exists cities;
drop table if exists vehicles;
drop table if exists profiles;
drop table if exists generation_modifications;
drop table if exists colors;
drop table if exists services;
drop table if exists service_types;
drop table if exists regions;
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
    "role_id" int not null default 1,
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
    "role_id" int not null default 1,
    "password" varchar(100) not null,
    "phone" varchar(100),
    "updated_at" timestamp default now(),
    "created_at" timestamp default now(),
    unique("email"),
    unique("phone")
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
insert into body_types (name, image) values ('Universal', '/images/body/universal.png');
insert into body_types (name, image) values ('Cabriolet', '/images/body/cabriolet.png');


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


create table images (
    "vehicle_id" int not null,
    "image" varchar(255) not null,
    "created_at" timestamp not null default now();
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


drop table if exists motorcycle_parameters;
drop table if exists motorcycles;
drop table if exists moto_category_parameters;
drop table if exists moto_parameter_values;
drop table if exists moto_parameters;
drop table if exists moto_models;
drop table if exists moto_brands;
drop table if exists moto_categories;



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
CREATE TYPE price_type_enum AS ENUM ('USD', 'AED', 'RUB');

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
                on update cascade
);


--  comtrans

-- create table comtran_categories (
--     "id" serial primary key,
--     "name" varchar(100) not null,
--     "created_at" timestamp not null default now(),

-- );

-- create table comtran_brands (
--     "id" serial primary key,
--     "name" varchar(100) not null,
--     "comtran_category_id" integer not null,
--     "created_at" timestamp not null default now(),
--     constraint comtran_brands_comtran_category_id_fk
--         foreign key (comtran_category_id)
--             references comtran_categories(id)
--                 on delete cascade
--                 on update cascade
-- );

-- create table comtran_models (
--     "id" serial primary key,
--     "name" varchar(100) not null,
--     "comtran_brand_id" integer not null,
--     "created_at" timestamp not null default now(),
--     constraint comtran_brand_models_comtran_brand_id_fk
--         foreign key (comtran_brand_id)
--             references comtran_brands(id)
--                 on delete cascade
--                 on update cascade
-- );

-- create table comtran_drivetrains(
--     "id" serial primary key,
--     "comtran_category_id" integer not null,
--     "name" varchar(100) not null,
--     "created_at" timestamp default now(),
--     constraint comtran_drivetrains_comtran_category_id_fk
--         foreign key (comtran_category_id)
--             references comtran_categories(id)
--                 on delete cascade
--                 on update cascade,
--     unique("name")
-- );

-- create table comtran_parameters (
--     "id" serial primary key,
--     "comtran_category_id" int not null,
--     "name" varchar(100) not null,
--     "type" varchar(50) not null,
--     "value" json not null,
--     "created_at" timestamp default now(),
--     constraint comtran_parameters_comtran_category_id_fk
--         foreign key (comtran_category_id)
--             references comtran_categories(id)
--                 on delete cascade
--                 on update cascade,
--     unique("name", "comtran_category_id")
-- );

-- create table comtran_category_parameters (
--     "parameter_id" int not null,
--     "comtran_category_id" int not null,
--     "created_at" timestamp not null default now(),
--     constraint comtran_category_parameters_comtran_category_id_fk
--         foreign key (comtran_category_id)
--             references comtran_categories(id)
--                 on delete cascade
--                 on update cascade
-- );

-- create table comtrans (
--     "id" serial primary key,
--     "user_id" int not null,
--     "comtran_category_id" int not null,
--     "comtran_brand_id" int not null,
--     "comtran_model_id" int not null,
--     "fuel_type_id" int not null,
--     "city_id" int not null,
--     "color_id" int not null,
--     "engine" int, -- cm3
--     "power" int, -- hp
--     "year" int not null,
--     "number_of_cycles" int,
--     "odometer" int not null default 0,
--     "crash" boolean,
--     "not_cleared" boolean,
--     "wheel" boolean,
--     "owners" int,
--     "date_of_purchase" varchar (50),
--     "warranty_date" varchar(50),
--     "ptc" boolean,
--     "vin_code" varchar(50) not null,
--     "certificate" varchar(50),
--     "description" text,
--     "can_look_coordinate" varchar(50),
--     "phone_number" varchar(50) not null,
--     "refuse_dealers_calls" boolean,
--     "only_chat" boolean,
--     "protect_spam" boolean,
--     "verified_buyers" boolean,
--     "contact_person" varchar(50), -- email or user_id
--     "email" varchar(50),
--     "price" int not null,
--     constraint comtrans_user_id_fk
--         foreign key (user_id)
--             references users(id)
--                 on delete cascade
--                 on update cascade,
--     constraint comtrans_category_id_fk
--         foreign key (category_id)
--             references comtran_categories(id)
--                 on delete cascade
--                 on update cascade,
--     constraint comtrans_brand_id_fk
--         foreign key (brand_id)
--             references comtran_brands(id)
--                 on delete cascade
--                 on update cascade,
--     constraint comtrans_model_id_fk
--         foreign key (model_id)
--             references comtran_models(id)
--                 on delete cascade
--                 on update cascade,
--     constraint comtrans_fuel_type_id_fk
--         foreign key (fuel_type_id)
--             references fuel_types(id)
--                 on delete cascade
--                 on update cascade,
--     constraint comtrans_color_id_fk
--         foreign key (color_id)
--             references colors(id)
--                 on delete cascade
--                 on update cascade,
--     constraint comtrans_city_id_fk
--         foreign key (city_id)
--             references cities(id)
--                 on delete cascade
--                 on update cascade
-- );





-- create table transmissions (
--     "id" serial primary key,
--     "name" varchar(255) not null,
--     "created_at" timestamp default now(),
--     unique("name")
-- );

-- insert into transmissions (name) values ('Automatic');
-- insert into transmissions (name) values ('Manual');
-- insert into transmissions (name) values ('Semi-Automatic');


-- create table engines (
--     "id" serial primary key,
--     "value" varchar(255) not null,
--     "created_at" timestamp default now(),
--     unique("value")
-- );

-- insert into engines (value) values ('1.0L');
-- insert into engines (value) values ('1.5L');
-- insert into engines (value) values ('2.0L');
-- insert into engines (value) values ('2.5L');
-- insert into engines (value) values ('3.0L');
-- insert into engines (value) values ('4.0L');
-- insert into engines (value) values ('5.0L');
-- insert into engines (value) values ('6.0L');
-- insert into engines (value) values ('7.0L');


-- create table drivetrains (
--     "id" serial primary key,
--     "name" varchar(255) not null,
--     "created_at" timestamp default now(),
--     unique("name")
-- );

-- insert into drivetrains (name) values ('Front-Wheel Drive');
-- insert into drivetrains (name) values ('Rear-Wheel Drive');
-- insert into drivetrains (name) values ('All-Wheel Drive');


-- create table fuel_types (
--     "id" serial primary key,
--     "name" varchar(255) not null,
--     "created_at" timestamp default now(),
--     unique("name")
-- );

-- insert into fuel_types (name) values ('Gasoline');
-- insert into fuel_types (name) values ('Diesel');
-- insert into fuel_types (name) values ('Electric');
-- insert into fuel_types (name) values ('Hybrid');
