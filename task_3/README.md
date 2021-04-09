## Task 3

## Description

Create separate github repo with the following pattern: `golang-training-YOUR-DATABASE-CHEMA-NAME`, Example: golang-training-theater.

Database schema names:
- Theater: @Alexandr59
- Hospital: @DaniilDomaskanov
- Shop: @dimaxdqwerty
- Enterprise: @jeka2708
- Library: @PutskouDzmitry
- University: @tonymontanapaffpaff
- Social network: @VSiniak12
- Restaurant: @miladouski

### First part
1. Create your `PostgresSQL`database consisting of 5 tables and 2 lookup tables (without links to parent tables), if necessary, you can create third. Your database should have a third normal form.
Create a data schema in `draw.io` (you can in access), add the image with your schema to repo with task implementation.
2. Add SQL script that creates tables, and add 5 records to each table. Save the script to the `script.sql` file.
3. Required type of fields in the database:
- `number`
- `date`
- `string`
- `boolean`

### Second part
Write a Golang application for manipulating data from the database created above. Use Standard library or `Gorm` library.
What should be present in the code:
 - Create the necessary entities at your discretion. 
 - At least one JOIN (all other requests at your discretion);
 - Organize CRUD (create, read, update, delete) operations
