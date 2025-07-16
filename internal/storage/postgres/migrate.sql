
-- insert into brands (name, logo, car_count, popular) values ('Ford', '/images/logo/ford.png', 38, false);
-- insert into brands (name, logo, car_count, popular) values ('Chevrolet', '/images/logo/chevrolet.png', 5, false);
-- insert into brands (name, logo, car_count, popular) values ('Nissan', '/images/logo/nissan.png', 22, true);
-- insert into brands (name, logo, car_count, popular) values ('Hyundai', '/images/logo/hyundai.png', 74, false);
-- insert into brands (name, logo, car_count, popular) values ('Kia', '/images/logo/kia.png', 33, true);




-- -- ford
-- insert into models (name, brand_id, popular, car_count) values ('F-150', 3, true, 56);
-- insert into models (name, brand_id, popular, car_count) values ('Mustang', 3, true, 29);
-- insert into models (name, brand_id, popular, car_count) values ('Explorer', 3, false, 664);
-- insert into models (name, brand_id, popular, car_count) values ('Bronco', 3, false, 34);

-- -- chevrolet
-- insert into models (name, brand_id, popular, car_count) values ('Camaro', 4, true, 23);
-- insert into models (name, brand_id, popular, car_count) values ('Corvette', 4, true, 645);
-- insert into models (name, brand_id, popular, car_count) values ('Equinox', 4, false, 23);
-- insert into models (name, brand_id, popular, car_count) values ('Silverado', 4, false, 55);

-- -- nissan
-- insert into models (name, brand_id, popular, car_count) values ('Altima', 5, true, 23);
-- insert into models (name, brand_id, popular, car_count) values ('Pathfinder', 5, true, 56);
-- insert into models (name, brand_id, popular, car_count) values ('Rogue', 5, false, 22);
-- insert into models (name, brand_id, popular, car_count) values ('Sentra', 5, false, 53);

-- -- hyundai
-- insert into models (name, brand_id, popular, car_count) values ('Elantra', 6, true, 54);
-- insert into models (name, brand_id, popular, car_count) values ('Sonata', 6, false, 12);
-- insert into models (name, brand_id, popular, car_count) values ('Kona', 6, false, 45);

-- -- kia
-- insert into models (name, brand_id, popular, car_count) values ('K5', 7, true, 50);
-- insert into models (name, brand_id, popular, car_count) values ('K7', 7, false, 98);


insert into vehicles (
    user_id, modification_id, brand_id, region_id, city_id, model_id, ownership_type_id, owners, view_count, year, popular, description, exchange, credit, wheel, crash, odometer, vin_code, phone_numbers, price, new, color_id, trade_in, status
) values
(1, 1, 310, 1, 1, 3232, 1, 1, 10, 2020, 0, 'Mock car 1', false, true, true, false, 10000, 'VIN00001', ARRAY['01234567890'], 20000, false, 1, 1, 3),
(2, 9, 310, 2, 2, 3233, 2, 2, 20, 2019, 1, 'Mock car 2', true, false, false, true, 20000, 'VIN00002', ARRAY['0111222222'], 25000, true, 2, 2, 3),
(1, 13, 310, 3, 3, 3234, 1, 1, 15, 2018, 0, 'Mock car 3', false, false, true, false, 15000, 'VIN00003', ARRAY['01234567890'], 18000, false, 3, 3, 3),
(2, 16, 310, 4, 1, 3232, 2, 2, 5, 2021, 1, 'Mock car 4', true, true, true, false, 5000, 'VIN00004', ARRAY['0111222222'], 30000, true, 4, 1, 3),
(1, 19, 310, 5, 2, 3233, 1, 1, 8, 2022, 0, 'Mock car 5', false, false, false, true, 8000, 'VIN00005', ARRAY['01234567890'], 22000, false, 5, 2, 3),
(2, 20, 310, 6, 3, 3234, 2, 2, 12, 2023, 1, 'Mock car 6', true, true, true, false, 12000, 'VIN00006', ARRAY['0111222222'], 27000, true, 6, 3, 3),
(1, 25, 310, 7, 1, 3232, 1, 1, 18, 2020, 0, 'Mock car 7', false, false, false, false, 18000, 'VIN00007', ARRAY['01234567890'], 21000, false, 7, 1, 3),
(2, 27, 310, 8, 2, 3233, 2, 2, 22, 2019, 1, 'Mock car 8', true, true, true, true, 22000, 'VIN00008', ARRAY['0111222222'], 26000, true, 1, 2, 3),
(1, 35, 310, 9, 3, 3234, 1, 1, 25, 2018, 0, 'Mock car 9', false, false, true, false, 25000, 'VIN00009', ARRAY['01234567890'], 19000, false, 2, 3, 3),
(2, 36, 310, 10, 1, 3232, 2, 2, 30, 2021, 1, 'Mock car 10', true, false, false, false, 30000, 'VIN00010', ARRAY['0111222222'], 32000, true, 3, 1, 3),
(1, 37, 310, 11, 2, 3233, 1, 1, 35, 2022, 0, 'Mock car 11', false, true, true, true, 35000, 'VIN00011', ARRAY['01234567890'], 23000, false, 4, 2, 3),
(2, 39, 310, 12, 3, 3234, 2, 2, 40, 2023, 1, 'Mock car 12', true, false, false, false, 40000, 'VIN00012', ARRAY['0111222222'], 28000, true, 5, 3, 3),
(1, 41, 310, 13, 1, 3232, 1, 1, 45, 2020, 0, 'Mock car 13', false, true, true, false, 45000, 'VIN00013', ARRAY['01234567890'], 24000, false, 6, 1, 3),
(2, 84, 310, 1, 2, 3233, 2, 2, 50, 2019, 1, 'Mock car 14', true, false, false, true, 50000, 'VIN00014', ARRAY['0111222222'], 29000, true, 7, 2, 3),
(1, 89, 310, 2, 3, 3234, 1, 1, 55, 2018, 0, 'Mock car 15', false, true, true, false, 55000, 'VIN00015', ARRAY['01234567890'], 20000, false, 1, 3, 3),
(2, 93, 310, 3, 1, 3232, 2, 2, 60, 2021, 1, 'Mock car 16', true, false, false, false, 60000, 'VIN00016', ARRAY['0111222222'], 33000, true, 2, 1, 3),
(1, 96, 310, 4, 2, 3233, 1, 1, 65, 2022, 0, 'Mock car 17', false, true, true, true, 65000, 'VIN00017', ARRAY['01234567890'], 25000, false, 3, 2, 3),
(2, 97, 310, 5, 3, 3234, 2, 2, 70, 2023, 1, 'Mock car 18', true, false, false, false, 70000, 'VIN00018', ARRAY['0111222222'], 30000, true, 4, 3, 3),
(1, 98, 310, 6, 1, 3232, 1, 1, 75, 2020, 0, 'Mock car 19', false, true, true, false, 75000, 'VIN00019', ARRAY['01234567890'], 26000, false, 5, 1, 3),
(2, 99, 310, 7, 2, 3233, 2, 2, 80, 2019, 1, 'Mock car 20', true, false, false, true, 80000, 'VIN00020', ARRAY['0111222222'], 31000, true, 6, 2, 3),
(1, 107, 310, 8, 3, 3234, 1, 1, 85, 2018, 0, 'Mock car 21', false, true, true, false, 85000, 'VIN00021', ARRAY['01234567890'], 21000, false, 7, 3, 3),
(2, 108, 310, 9, 1, 3232, 2, 2, 90, 2021, 1, 'Mock car 22', true, false, false, false, 90000, 'VIN00022', ARRAY['0111222222'], 34000, true, 1, 1, 3),
(1, 109, 310, 10, 2, 3233, 1, 1, 95, 2022, 0, 'Mock car 23', false, true, true, true, 95000, 'VIN00023', ARRAY['01234567890'], 27000, false, 2, 2, 3),
(2, 111, 310, 11, 3, 3234, 2, 2, 100, 2023, 1, 'Mock car 24', true, false, false, false, 100000, 'VIN00024', ARRAY['0111222222'], 35000, true, 3, 3, 3),
(1, 113, 310, 12, 1, 3232, 1, 1, 105, 2020, 0, 'Mock car 25', false, true, true, false, 105000, 'VIN00025', ARRAY['01234567890'], 28000, false, 4, 1, 3),
(2, 156, 310, 13, 2, 3233, 2, 2, 110, 2019, 1, 'Mock car 26', true, false, false, true, 110000, 'VIN00026', ARRAY['0111222222'], 36000, true, 5, 2, 3),
(1, 160, 310, 1, 3, 3234, 1, 1, 115, 2018, 0, 'Mock car 27', false, true, true, false, 115000, 'VIN00027', ARRAY['01234567890'], 22000, false, 6, 3, 3),
(2, 165, 310, 2, 1, 3232, 2, 2, 120, 2021, 1, 'Mock car 28', true, false, false, false, 120000, 'VIN00028', ARRAY['0111222222'], 37000, true, 7, 1, 3),
(1, 167, 310, 3, 2, 3233, 1, 1, 125, 2022, 0, 'Mock car 29', false, true, true, true, 125000, 'VIN00029', ARRAY['01234567890'], 29000, false, 1, 2, 3),
(2, 168, 310, 4, 3, 3234, 2, 2, 130, 2023, 1, 'Mock car 30', true, false, false, false, 130000, 'VIN00030', ARRAY['0111222222'], 38000, true, 2, 3, 1);





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

insert into images (vehicle_id, image) values (10, '/images/cars/10/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (10, '/images/cars/10/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (10, '/images/cars/10/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (10, '/images/cars/10/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (11, '/images/cars/11/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (11, '/images/cars/11/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (11, '/images/cars/11/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (11, '/images/cars/11/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (12, '/images/cars/12/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (12, '/images/cars/12/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (12, '/images/cars/12/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (12, '/images/cars/12/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (13, '/images/cars/13/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (13, '/images/cars/13/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (13, '/images/cars/13/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (13, '/images/cars/13/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (14, '/images/cars/14/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (14, '/images/cars/14/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (14, '/images/cars/14/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (14, '/images/cars/14/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (15, '/images/cars/15/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (15, '/images/cars/15/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (15, '/images/cars/15/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (15, '/images/cars/15/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (16, '/images/cars/16/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (16, '/images/cars/16/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (16, '/images/cars/16/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (16, '/images/cars/16/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (17, '/images/cars/17/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (17, '/images/cars/17/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (17, '/images/cars/17/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (17, '/images/cars/17/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (18, '/images/cars/18/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (18, '/images/cars/18/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (18, '/images/cars/18/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (18, '/images/cars/18/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (19, '/images/cars/19/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (19, '/images/cars/19/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (19, '/images/cars/19/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (19, '/images/cars/19/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (20, '/images/cars/20/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (20, '/images/cars/20/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (20, '/images/cars/20/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (20, '/images/cars/20/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (21, '/images/cars/21/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (21, '/images/cars/21/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (21, '/images/cars/21/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (21, '/images/cars/21/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (22, '/images/cars/22/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (22, '/images/cars/22/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (22, '/images/cars/22/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (22, '/images/cars/22/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (23, '/images/cars/23/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (23, '/images/cars/23/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (23, '/images/cars/23/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (23, '/images/cars/23/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (24, '/images/cars/24/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (24, '/images/cars/24/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (24, '/images/cars/24/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (24, '/images/cars/24/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (25, '/images/cars/25/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (25, '/images/cars/25/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (25, '/images/cars/25/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (25, '/images/cars/25/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (26, '/images/cars/26/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (26, '/images/cars/26/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (26, '/images/cars/26/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (26, '/images/cars/26/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (27, '/images/cars/27/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (27, '/images/cars/27/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (27, '/images/cars/27/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (27, '/images/cars/27/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (28, '/images/cars/28/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (28, '/images/cars/28/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (28, '/images/cars/28/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (28, '/images/cars/28/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (29, '/images/cars/29/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (29, '/images/cars/29/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (29, '/images/cars/29/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (29, '/images/cars/29/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

insert into images (vehicle_id, image) values (30, '/images/cars/30/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1');
insert into images (vehicle_id, image) values (30, '/images/cars/30/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2');
insert into images (vehicle_id, image) values (30, '/images/cars/30/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3');
insert into images (vehicle_id, image) values (30, '/images/cars/30/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4');

