
insert into brands (name, logo, car_count, popular) values ('Ford', '/images/logo/ford.png', 38, false);
insert into brands (name, logo, car_count, popular) values ('Chevrolet', '/images/logo/chevrolet.png', 5, false);
insert into brands (name, logo, car_count, popular) values ('Nissan', '/images/logo/nissan.png', 22, true);
insert into brands (name, logo, car_count, popular) values ('Hyundai', '/images/logo/hyundai.png', 74, false);
insert into brands (name, logo, car_count, popular) values ('Kia', '/images/logo/kia.png', 33, true);




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