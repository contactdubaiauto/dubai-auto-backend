create user da with password '1234';
grant all privileges on database da to da;
grant all privileges on schema public to da;
grant all privileges on all tables in schema public to da;
grant all privileges on all sequences in schema public to da;
alter default privileges in schema public grant all on tables to da;
alter default privileges in schema public grant all on sequences to da;


drop table if exists images;
drop table if exists vehicles;
drop table if exists services;
drop table if exists service_types;
drop table if exists regions;
drop table if exists cities;
drop table if exists fuel_types;
drop table if exists drives;
drop table if exists engines;
drop table if exists transmissions;
drop table if exists body_types;
drop table if exists models;
drop table if exists brands;
drop table if exists users;
drop table if exists admins;


create table users (
    "id" serial primary key,
    "username" varchar(255) not null,
    "email" varchar(255) not null,
    "password" varchar(255) not null,
    "phone" varchar(255) not null,
    "notification" boolean default false,
    "last_active_date" timestamp default now(),
    "created_at" timestamp default now()
);

insert into users (username, email, password, phone, notification, last_active_date, created_at) 
    values ('user', 'user@gmail.com', '$2a$10$Cya9x0xSJSnRknBmJpW.Bu8ukZpVTqzwgrQgAYNPXdrX2HYGRk33W', '01234567890', true, now(), now()); -- password: 12345678


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
    "created_at" timestamp default now()
);

insert into body_types (name) values ('Sedan');
insert into body_types (name) values ('Hatchback');
insert into body_types (name) values ('SUV');
insert into body_types (name) values ('Crossover');
insert into body_types (name) values ('Coupe');
insert into body_types (name) values ('Convertible');
insert into body_types (name) values ('Wagon');
insert into body_types (name) values ('Pickup Truck');
insert into body_types (name) values ('Van');
insert into body_types (name) values ('Minivan');
insert into body_types (name) values ('Roadster');
insert into body_types (name) values ('Sports Car');
insert into body_types (name) values ('Off-Road');
insert into body_types (name) values ('Limousine'); 
insert into body_types (name) values ('Utility');




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


create table drives (
    "id" serial primary key,
    "name" varchar(255) not null,
    "created_at" timestamp default now()
);

insert into drives (name) values ('Front-Wheel Drive');
insert into drives (name) values ('Rear-Wheel Drive');
insert into drives (name) values ('All-Wheel Drive');


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


create table vehicles (
    "id" serial primary key,
    "user_id" int,
    "brand_id" int,
    "region_id" int,
    "city_id" int,
    "model_id" int,
    "transmission_id" int,
    "engine_id" int,
    "drive_id" int,
    "body_type_id" int,
    "fuel_type_id" int,
    "ownership" int not null default 1,
    "announcement_type" int not null default 0,
    "view_count" int not null default 0,
    "year" int not null,
    "exchange" boolean not null default false,
    "credit" boolean not null default false,
    "mileage" int,
    "vin_code" varchar(255),
    "door_count" int,
    "phone_number" varchar(255) not null,
    "price" int not null,
    "new" boolean not null default false,
    "color" varchar(255),
    "credit_price" int,
    "status" int not null default 1,
    "updated_at" timestamp default now(),
    "created_at" timestamp default now(),
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
    constraint vehicles_drive_id_fk
        foreign key (drive_id)
            references drives(id)
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
                on update cascade
);

insert into vehicles (
        user_id, brand_id, city_id, model_id, body_type_id, fuel_type_id, ownership, year, 
        exchange, credit, mileage, vin_code, door_count, phone_number, price, new, color, 
        credit_price, status
    ) 
    values (
        1, 1, 1, 1, 1, 1, 1, 2020, false, false, 100000, '1234567890', 4, '01234567890', 100000, true, '#ffffff', 100000, 1
    );
insert into vehicles (
        user_id, brand_id, city_id, model_id, body_type_id, fuel_type_id, ownership, year, 
        exchange, credit, mileage, vin_code, door_count, phone_number, price, new, color, 
        credit_price, status
    ) 
    values (
        1, 1, 1, 1, 1, 1, 1, 2020, false, false, 100000, '1234567890', 4, '01234567890', 100000, true, '#rfffff', 100000, 1
    );
insert into vehicles (
        user_id, brand_id, city_id, model_id, body_type_id, fuel_type_id, ownership, year, 
        exchange, credit, mileage, vin_code, door_count, phone_number, price, new, color, 
        credit_price, status
    ) 
    values (
        1, 1, 1, 1, 1, 1, 1, 2020, false, false, 100000, '1234567890', 4, '01234567890', 100000, true, '#rfffff', 100000, 1
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

insert into images (vehicle_id, image) values (1, '/images/cars/1/1.jpg');
insert into images (vehicle_id, image) values (1, '/images/cars/1/2.jpg');
insert into images (vehicle_id, image) values (1, '/images/cars/1/3.jpg');
insert into images (vehicle_id, image) values (1, '/images/cars/1/4.jpg');
insert into images (vehicle_id, image) values (1, '/images/cars/1/5.jpg');
insert into images (vehicle_id, image) values (1, '/images/cars/1/6.jpg');
insert into images (vehicle_id, image) values (1, '/images/cars/1/7.jpg');
insert into images (vehicle_id, image) values (1, '/images/cars/1/8.jpg');
insert into images (vehicle_id, image) values (1, '/images/cars/1/9.jpg');
insert into images (vehicle_id, image) values (1, '/images/cars/1/10.jpg');

