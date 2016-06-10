create table price(
  landcode varchara(10) primary key,
  year integer DEFAULT NULL,
  price_per_acreage double(15,3) DEFAULT NULL
 );
