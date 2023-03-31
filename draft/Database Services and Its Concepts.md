# Database Services and Its Concepts

## How a Query Work?

Table:

| column_1 | column_2 | column_3 |
|:--------:|:--------:|:--------:|
|  row_1   |  row_1   |  row_1   |
|  row_2   |  row_2   |  row_2   |
|  row_3   |  row_3   |  row_3   |

SQL Query:

```
SELECT * FROM column_3 = row_2
```

Mongo Query:

```
db.table.find({column_3: row_2})
```

Result:

| ________ | ________ | column_3 |
|:--------:|:--------:|:--------:|
| ________ | ________ | ________ |
|  row_2   |  row_2   |  row_2   |
| ________ | ________ | ________ |

## NoSQL

- [ ] ~~(It is) not SQL~~
- [x] Not-Only SQL

### NoSQL Services List

- `MongoDB` - Document-Oriented Database
- `Redis` - Key-Value Database
- `ElasticSearch` - Full-text Search Database
- `InfluxDB` / `TimescaleDB` - Time-Series Database
- `Neo4J` - Graph Database

----------

### 1. RDBMS / SQL Based

| name | job | salary |
|:----:|:---:|:------:|
| Alam | DBA | 50.000 |
| Rio  | DBA | 40.000 |
| Riko | NOC | 45.000 |

- Can relate between a table to another table
- Structured data

### 2. Document-Oriented

```
[
	{"name": "Edi Riyanto", "roles": "DevOps", "resign": false, "project": ["Deploy", "RND"]},
	{"name": "Deri Satria", "roles": "DevOps", "resign": true}
]
```

- Unstructured data

### 3. Key-Value

Get all `PENDING` data.

```
PENDING TRX111222333
        TRX444555666
        TRX777888999
```

- In-memory data store
- Faster because its smaller (data)

### 4. Search Database

Search: `Rio NOC`.

```
[
	{"nama": "Rio Pribumi", "roles": "NOC", "extra_info": null},
	{"nama": "Rio Pendatang", "roles": "DBA", "extra_info": "Ex-NOC"}
]
```

- Can search `values` even in different columns
