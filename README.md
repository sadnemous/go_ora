Sources: 
```url
https://blogs.oracle.com/developers/post/connecting-a-go-application-to-oracle-database
https://github.com/lucasjellema/go-oracle-database/tree/main
```

Oracle: 
```bash
    docker pull gvenzl/oracle-xe
    docker run -d -p 1521:1521 -e ORACLE_PASSWORD=oracle gvenzl/oracle-xe
```
sqldeveloper:

Table: https://www.mycompiler.io/view/IXp1OnToZGf
```sql
    create user demo identified by demo
    grant create session to demo;
    grant all privileges to demo;


    CREATE TABLE demo.library (
        bookid varchar2(5)primary key,
        bname varchar2(20),author varchar2(20),
        purchased date,publisher varchar2(20),price number(10)
    );

    INSERT INTO demo.library VALUES ('b101','cost accounting','jainnarang','11-feb-2013','kalyani',800);
    INSERT INTO demo.library VALUES ('b102','business statistics','apaggarwal','22-dec-2011','himalaya',750);
    INSERT INTO demo.library VALUES ('b103','rdbms','cjdate','02-mar-2015','tmh',900);
    INSERT INTO demo.library VALUES ('b104','managment accounting','rksharma','17-apr-2014','kalyani',950);
    INSERT INTO demo.library VALUES ('b105','operating system','galvin','25-nov-2013','phi',750);
    INSERT INTO demo.library VALUES ('b106','advance accounting','cgupta','16-apr-2018','himalaya',600);

    select * from demo.library;
    commit;
```


