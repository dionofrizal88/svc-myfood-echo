# svc-myfood-echo

## Backend local setup

**Create database 'myfood':**
```
CREATE DATABASE myfood;
```

```
-- myfood.baskets definition

CREATE TABLE `baskets` (
  `row_id` int(11) NOT NULL AUTO_INCREMENT,
  `id` bigint(20) unsigned DEFAULT NULL,
  `id_product` bigint(20) unsigned DEFAULT NULL,
  `jumlah` int(32) DEFAULT NULL,
  `keterangan` varchar(200) DEFAULT NULL,
  `flag_order` int(11) DEFAULT 0,
  `flag_aktif` int(11) DEFAULT 1,
  `tgl_input` datetime DEFAULT current_timestamp(),
  `user_input` varchar(255) CHARACTER SET latin1 DEFAULT 'system',
  `tgl_update` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `user_update` varchar(255) CHARACTER SET latin1 DEFAULT 'system',
  `uuid` varchar(128) CHARACTER SET latin1 DEFAULT NULL,
  PRIMARY KEY (`row_id`),
  KEY `id` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;
```

-- myfood.order_pivot definition
```
CREATE TABLE `order_pivot` (
  `row_id` int(11) NOT NULL AUTO_INCREMENT,
  `id` bigint(20) unsigned DEFAULT NULL,
  `nomor_meja` varchar(10) DEFAULT NULL,
  `nama_pemesan` varchar(64) DEFAULT NULL,
  `flag_aktif` int(11) DEFAULT 1,
  `tgl_input` datetime DEFAULT current_timestamp(),
  `user_input` varchar(255) CHARACTER SET latin1 DEFAULT 'system',
  `tgl_update` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `user_update` varchar(255) CHARACTER SET latin1 DEFAULT 'system',
  `uuid` varchar(128) CHARACTER SET latin1 DEFAULT NULL,
  PRIMARY KEY (`row_id`),
  KEY `id` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
```

```
-- myfood.orders definition

CREATE TABLE `orders` (
  `row_id` int(11) NOT NULL AUTO_INCREMENT,
  `id` bigint(20) unsigned DEFAULT NULL,
  `id_order_pivot` bigint(20) unsigned DEFAULT NULL,
  `id_baskets` bigint(20) unsigned DEFAULT NULL,
  `flag_aktif` int(11) DEFAULT 1,
  `tgl_input` datetime DEFAULT current_timestamp(),
  `user_input` varchar(255) CHARACTER SET latin1 DEFAULT 'system',
  `tgl_update` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `user_update` varchar(255) CHARACTER SET latin1 DEFAULT 'system',
  `uuid` varchar(128) CHARACTER SET latin1 DEFAULT NULL,
  PRIMARY KEY (`row_id`),
  KEY `id` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

```

```
-- myfood.products definition

CREATE TABLE `products` (
  `row_id` int(11) NOT NULL AUTO_INCREMENT,
  `id` bigint(20) unsigned DEFAULT NULL,
  `id_status` bigint(20) unsigned DEFAULT NULL,
  `kode` varchar(10) DEFAULT NULL,
  `nama` varchar(64) DEFAULT NULL,
  `harga` varchar(64) DEFAULT NULL,
  `gambar` varchar(200) DEFAULT NULL,
  `flag_ready` int(1) DEFAULT 1,
  `flag_aktif` int(11) DEFAULT 1,
  `tgl_input` datetime DEFAULT current_timestamp(),
  `user_input` varchar(255) CHARACTER SET latin1 DEFAULT 'system',
  `tgl_update` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `user_update` varchar(255) CHARACTER SET latin1 DEFAULT 'system',
  `uuid` varchar(128) CHARACTER SET latin1 DEFAULT NULL,
  `catatan_penghapusan` text CHARACTER SET latin1 DEFAULT NULL,
  PRIMARY KEY (`row_id`),
  KEY `id` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4;
```

**Create triger for all table:**
**This example for baskets table:**

```
CREATE DEFINER=`root`@`localhost` TRIGGER `bi_baskets` BEFORE INSERT ON `baskets` FOR EACH ROW BEGIN
	if new.id is null then
    SET new.id = UUID_SHORT();
	end if;
	if new.uuid is null then
		set new.uuid = uuid();
	end if;
END
```

**Initialize the Go project:**
Initialize the Go project using the following command
```
go mod init svc-myfood-echo
```

**Adding the modules required for the project:**
**Get all module in go.mod:**
```
go get github.com/labstack/echo/v4
go get github.com/go-sql-driver/mysql
go get github.com/gorilla/mux
```
**Run the backend app**
```
go run main.go
```
