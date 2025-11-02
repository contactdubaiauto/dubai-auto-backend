
		select
			about_me,
			whatsapp,
			telegram,
			address,
			coordinates,
			avatar,
			banner,
			company_name,
			message,
			vat_number,
			company_types.name as company_type,
			activity_fields.name as activity_field,
			profiles.created_at,
            destinations.countries as destinations
		from users 
        left join profiles on profiles.user_id = users.id
        left join (
                SELECT json_agg(
                        json_build_object(
                        'from_country', json_build_object(
                            'id', cf.id,
                            'name', cf.name,
                            'flag', cf.flag,
                        ),
                        'to_country', json_build_object(
                            'id', ct.id,
                            'name', ct.name,
                            'flag', ct.flag,
                        )
                        )
                    ) AS countries
                FROM user_destinations ds
                LEFT JOIN countries cf ON cf.id = ds.from_id
                LEFT JOIN countries ct ON ct.id = ds.to_id
                WHERE ds.user_id = 77
        ) destinations on true
		left join company_types on company_types.id = profiles.company_type_id
		left join activity_fields on activity_fields.id = profiles.activity_field_id
        where users.id = 77;



insert into user_destinations (user_id, from_id, to_id) values 
    (76, 1, 2),
    (76, 2, 3),
    (76, 3, 4);




insert into user_destinations (user_id, from_id, to_id) values 
    (80, 1, 2),
    (80, 2, 3),
    (80, 3, 4);



insert into user_destinations (user_id, from_id, to_id) values 
    (84, 1, 2),
    (84, 2, 3),
    (84, 3, 4);


insert into user_destinations (user_id, from_id, to_id) values 
    (88, 1, 2),
    (88, 2, 3),
    (88, 3, 4);




SELECT 
    u.id, 
    p.company_name, 
    ds.licence_issue_date, 
    ds.licence_expiry_date, 
    u.username, 
    u.email, 
    u.phone, 
    u.status, 
    u.created_at,
    $2 || ds.copy_of_id_path as copy_of_id_url,
    $2 || ds.memorandum_path as memorandum_url,
    $2 || ds.licence_path as licence_url,
    p.address
FROM users u
left join profiles p on p.user_id = u.id
left join documents ds on ds.id = p.documents_id
WHERE u.id = $1