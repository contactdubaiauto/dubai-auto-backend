-- 0001	-	add status to users table
alter table users add column status int not null default 1;
alter table temp_users add column status int not null default 1;



alter table temp_users add column "company_name" varchar(100);
alter table temp_users add column "company_type_id" int;
alter table temp_users add column "activity_field_id" int;
alter table temp_users add column "vat_number" varchar(100);
alter table temp_users add column "address" varchar(100);
alter table temp_users add column "license_issue_date" timestamp;
alter table temp_users add column "license_expiry_date" timestamp;



alter table temp_users add constraint fk_temp_users_activity_field_id
        foreign key (activity_field_id)
            references activity_fields(id)
                on delete cascade
                on update cascade;

alter table temp_users add column documents_id int;
alter table temp_users add constraint fk_temp_users_documents_id
        foreign key (documents_id)
            references documents(id)
                on delete cascade
                on update cascade;

alter table documents
    alter column licence_issue_date drop not null,
    alter column licence_expiry_date drop not null;


alter table temp_users rename column license_issue_date to licence_issue_date;
alter table temp_users rename column license_expiry_date to licence_expiry_date;


