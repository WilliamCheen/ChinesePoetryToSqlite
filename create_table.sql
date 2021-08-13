create table poems
(
    id         integer primary key autoincrement,
    category   varchar(50) not null,
    dynasty    varchar(50) not null,
    title      varchar(200)  default "",
    author     varchar(200)  default "",
    rhythmic   varchar(200)  default "",
    chapter    varchar(200)  default "",
    section    varchar(200)  default "",
    notes      varchar(2000) default "",
    paragraphs varchar(5000) default ""
);