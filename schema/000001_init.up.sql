CREATE TABLE restaurants_list (
    id serial NOT NULL unique,
    rest_name varchar (255) NOT NULL,
    av_time TIME NOT NULL,
    av_check int NOT NULL
);

CREATE TABLE user_list (
    id serial NOT NULL unique,
    users_name varchar (255) NOT NULL,
    phone varchar (255) NOT NULL unique
);

CREATE TABLE table_list (
    id serial NOT NULL unique,
    table_num int NOT NULL,
    max_persons int NOT NULL,
    free_tables int NOT NULL
);

CREATE TABLE restaurants_tables (
    id serial NOT NULL unique,
    restaurant_id int NOT NULL,
    table_id int NOT NULL
);

CREATE TABLE booking_list (
    id serial NOT NULL unique,
    table_id int NOT NULL,
    users_id int  NOT NULL,
    table_num int NOT NULL,
    booking_date DATE NOT NULL,
    booking_time TIME NOT NULL
);