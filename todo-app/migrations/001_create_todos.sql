create table todos (
    id serial primary key,
    title text not null,
    completed boolean not null default false,
    created_at timestamp not null default now()
);