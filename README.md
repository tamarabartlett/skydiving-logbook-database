
### Local Postgres Development

To use a local Postgres instance, install Postgres
(`brew install postgresql`) and follow these steps:

1. Start a Postgres server: `pg_ctl -D /usr/local/var/postgres start`
2. Create db `createdb jumps`
3. Created a table `CREATE TABLE jumps (number int, altitude int, dropzone varchar(255));`
4. Create data in table `INSERT INTO jumps (number, altitude, dropzone) VALUES ('1', '12500', 'Elsinore'),('2', '12500', 'Elsinore'), ('3', '12500', 'Lodi'), ('4', '13000', 'Byron'), ('5', '13000', 'Tracy'),('4', '13000', 'Skydance');
INSERT 0 6`
5. Oops. I added jumps with the same number. Delete one of them `DELETE FROM jumps WHERE dropzone='Skydance';`
5. Make jump-number the primary key `ALTER TABLE jumps ADD PRIMARY KEY (number);`

- Log into db: `psql jumps`
  - List dbs `\l`
  - List tables `\dt`
