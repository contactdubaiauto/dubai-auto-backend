-- Union query for vehicles, comtrans, and motorcycles
with vs as (        
    select 
        vs.id,
        'car' as type,
        bs.name as brand,
        ms.name as model,
        vs.year,
        vs.price,
        vs.credit,
        vs.status,
        vs.created_at,
        images.images,
        vs.view_count,
        true as my_car,
        vs.crash
    from vehicles vs
    left join brands bs on vs.brand_id = bs.id
    left join models ms on vs.model_id = ms.id
    LEFT JOIN LATERAL (
        SELECT json_agg(img.image) AS images
        FROM (
            SELECT image as image
            FROM images
            WHERE vehicle_id = vs.id
            ORDER BY created_at DESC
        ) img
    ) images ON true
    where vs.user_id = 1 and status = 2
    order by vs.id desc
),
cms as (
    select
        cm.id,
        'comtran' as type,
        cbs.name as brand,
        cms.name as model,
        cm.year,
        cm.price,
        cm.credit,
        cm.status,
        cm.created_at,
        images.images,
        cm.view_count,
        true as my_car,
        cm.crash
    from comtrans cm
    left join com_brands cbs on cbs.id = cm.comtran_brand_id
    left join com_models cms on cms.id = cm.comtran_model_id
    LEFT JOIN LATERAL (
        SELECT json_agg(img.image) AS images
        FROM (
            SELECT image
            FROM comtran_images
            WHERE comtran_id = cm.id
            ORDER BY created_at DESC
        ) img
    ) images ON true
    where cm.user_id = 1 and cm.status = 2
),
mts as (
    select
        mt.id,
        'motorcycle' as type,
        mbs.name as brand,
        mms.name as model,
        mt.year,
        mt.price,
        mt.credit,
        mt.status,
        mt.created_at,
        mt.view_count,
        images.images,
        true as my_car,
        mt.crash
    from motorcycles mt
    left join moto_brands mbs on mbs.id = mt.moto_brand_id
    left join moto_models mms on mms.id = mt.moto_model_id
    LEFT JOIN LATERAL (
        SELECT json_agg(img.image) AS images
        FROM (
            SELECT image
            FROM moto_images
            WHERE moto_id = mt.id
            ORDER BY created_at DESC
        ) img
    ) images ON true
    where mt.user_id = 1 and mt.status = 2
)
-- Union all three CTEs
select 
    id, type, brand, model, 
    year, price, credit, 
    status, created_at, 
    view_count, images, my_car, 
    crash 
from vs
union all
select 
    id, type, brand, model, 
    year, price, credit, 
    status, created_at, 
    view_count, images, my_car, 
    crash 
from cms
union all
select 
    id, type, brand, model, 
    year, price, credit, 
    status, created_at, 
    view_count, images, my_car, 
    crash 
from mts
order by created_at desc;




alter table motorcycles add column "view_count" int not null default 0;
alter table comtrans add column "view_count" int not null default 0;
alter table motorcycles add column "credit" boolean not null default false;
alter table comtrans add column "credit" boolean not null default false;



select 
    c.updated_at,
    u.username,
    p.avatar,
    u.id
from conversations c
join users u on u.id = 
    case 
        when c.user_id_1 = 23 then c.user_id_2 
        else c.user_id_1 
    end
join profiles p on p.user_id = u.id
order by c.updated_at desc;
